// Package jdb is a lightweight IPFS peer which runs the minimal setup to
// provide an `ipld.DAGService`, "Add" and "Get" UnixFS files from IPFS.
package jdb

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"strings"
	"sync"

	"github.com/gioapp/cms/pkg/jdb/repo"
	"github.com/ipfs/go-bitswap"
	"github.com/ipfs/go-bitswap/network"
	blockservice "github.com/ipfs/go-blockservice"
	"github.com/ipfs/go-cid"
	"github.com/ipfs/go-datastore"
	syncds "github.com/ipfs/go-datastore/sync"
	blockstore "github.com/ipfs/go-ipfs-blockstore"
	chunker "github.com/ipfs/go-ipfs-chunker"
	provider "github.com/ipfs/go-ipfs-provider"
	"github.com/ipfs/go-ipfs-provider/queue"
	"github.com/ipfs/go-ipfs-provider/simple"
	cbor "github.com/ipfs/go-ipld-cbor"
	ipld "github.com/ipfs/go-ipld-format"
	logging "github.com/ipfs/go-log/v2"
	"github.com/ipfs/go-merkledag"
	"github.com/ipfs/go-unixfs/importer/balanced"
	"github.com/ipfs/go-unixfs/importer/helpers"
	"github.com/ipfs/go-unixfs/importer/trickle"
	ufsio "github.com/ipfs/go-unixfs/io"
	"github.com/libp2p/go-libp2p-core/crypto"
	host "github.com/libp2p/go-libp2p-core/host"
	inet "github.com/libp2p/go-libp2p-core/network"
	peer "github.com/libp2p/go-libp2p-core/peer"
	routing "github.com/libp2p/go-libp2p-core/routing"
	swarm "github.com/libp2p/go-libp2p-swarm"
	"github.com/multiformats/go-multiaddr"
	multihash "github.com/multiformats/go-multihash"
)

func init() {
	ipld.Register(cid.DagProtobuf, merkledag.DecodeProtobufBlock)
	ipld.Register(cid.Raw, merkledag.DecodeRawBlock)
	ipld.Register(cid.DagCBOR, cbor.DecodeBlock) // need to decode CBOR
}

var logger = logging.Logger("ipfslite")

// Peer is an IPFS-Lite peer. It provides a DAG service that can fetch and put
// blocks from/to the IPFS network.
type Peer struct {
	Ctx             context.Context
	Host            host.Host
	Store           datastore.Batching
	Bstore          blockstore.Blockstore
	DHT             routing.Routing
	Bserv           blockservice.BlockService
	Repo            repo.Repo
	Provider        provider.System
	ipld.DAGService // become a DAG service
	bstore          blockstore.Blockstore
	bserv           blockservice.BlockService
	reprovider      provider.System
}

// New creates an IPFS-Lite Peer. It uses the given datastore, libp2p Host and
// Routing (usuall the DHT). The Host and the Routing may be nil if
// config.Offline is set to true, as they are not used in that case. Peer
// implements the ipld.DAGService interface.
func NewPeer(
	ctx context.Context,
	r repo.Repo,
) (*Peer, error) {
	store := syncds.MutexWrap(datastore.NewMapDatastore())
	cfg, err := r.Config()
	if err != nil {
		return nil, err
	}

	privb, _ := base64.StdEncoding.DecodeString(cfg.Identity.PrivKey)
	privKey, _ := crypto.UnmarshalPrivateKey(privb)

	listenAddrs := []multiaddr.Multiaddr{}
	confAddrs := cfg.Addresses.Swarm
	for _, v := range confAddrs {
		listen, _ := multiaddr.NewMultiaddr(v)
		listenAddrs = append(listenAddrs, listen)
	}

	h, dht, err := SetupLibp2p(
		ctx,
		privKey,
		nil,
		listenAddrs,
		store,
		Libp2pOptionsExtra...,
	)

	if err != nil {
		return nil, err
	}
	p := &Peer{
		Ctx:   ctx,
		Host:  h,
		DHT:   dht,
		Store: store,
		Repo:  r,
	}

	err = p.setupBlockstore()
	if err != nil {
		return nil, err
	}
	err = p.setupBlockService()
	if err != nil {
		return nil, err
	}
	err = p.setupDAGService()
	if err != nil {
		p.bserv.Close()
		return nil, err
	}
	err = p.setupReprovider()
	if err != nil {
		p.bserv.Close()
		return nil, err
	}

	go p.autoclose()

	return p, nil
}

func (p *Peer) setupBlockstore() error {
	bs := blockstore.NewBlockstore(p.Store)
	bs = blockstore.NewIdStore(bs)
	cachedbs, err := blockstore.CachedBlockstore(p.Ctx, bs, blockstore.DefaultCacheOpts())
	if err != nil {
		return err
	}
	p.bstore = cachedbs
	return nil
}

func (p *Peer) setupBlockService() error {
	bswapnet := network.NewFromIpfsHost(p.Host, p.DHT)
	bswap := bitswap.New(p.Ctx, bswapnet, p.bstore)
	p.bserv = blockservice.New(p.bstore, bswap)
	return nil
}

func (p *Peer) setupDAGService() error {
	p.DAGService = merkledag.NewDAGService(p.bserv)
	return nil
}

func (p *Peer) setupReprovider() error {
	queue, err := queue.NewQueue(p.Ctx, "repro", p.Store)
	checkError(err)
	prov := simple.NewProvider(
		p.Ctx,
		queue,
		p.DHT,
	)
	cfg, err := p.Repo.Config()
	checkError(err)
	reprov := simple.NewReprovider(
		p.Ctx,
		cfg.ReprovideInterval,
		p.DHT,
		simple.NewBlockstoreProvider(p.bstore),
	)

	p.reprovider = provider.NewSystem(prov, reprov)
	p.reprovider.Run()
	return nil
}

func (p *Peer) autoclose() {
	<-p.Ctx.Done()
	p.reprovider.Close()
	p.bserv.Close()
}

// Bootstrap is an optional helper to connect to the given peers and bootstrap
// the Peer DHT (and Bitswap). This is a best-effort function. Errors are only
// logged and a warning is printed when less than half of the given peers
// could be contacted. It is fine to pass a list where some peers will not be
// reachable.
func (p *Peer) Bootstrap(peers []peer.AddrInfo) {
	connected := make(chan struct{})

	var wg sync.WaitGroup
	for _, pinfo := range peers {
		//h.Peerstore().AddAddrs(pinfo.ID, pinfo.Addrs, peerstore.PermanentAddrTTL)
		wg.Add(1)
		go func(pinfo peer.AddrInfo) {
			defer wg.Done()
			err := p.Host.Connect(p.Ctx, pinfo)
			if err != nil {
				logger.Warn(err)
				return
			}
			logger.Info("Connected to", pinfo.ID)
			connected <- struct{}{}
		}(pinfo)
	}

	go func() {
		wg.Wait()
		close(connected)
	}()

	i := 0
	for range connected {
		i++
	}
	if nPeers := len(peers); i < nPeers/2 {
		logger.Warnf("only connected to %d bootstrap peers out of %d", i, nPeers)
	}

	err := p.DHT.Bootstrap(p.Ctx)
	if err != nil {
		logger.Error(err)
		return
	}
}

// Session returns a session-based NodeGetter.
func (p *Peer) Session(ctx context.Context) ipld.NodeGetter {
	ng := merkledag.NewSession(ctx, p.DAGService)
	if ng == p.DAGService {
		logger.Warn("DAGService does not support sessions")
	}
	return ng
}

// AddParams contains all of the configurable parameters needed to specify the
// importing process of a file.
type AddParams struct {
	Layout    string
	Chunker   string
	RawLeaves bool
	Hidden    bool
	Shard     bool
	NoCopy    bool
	HashFun   string
}

// AddFile chunks and adds content to the DAGService from a reader. The content
// is stored as a UnixFS DAG (default for IPFS). It returns the root
// ipld.Node.
func (p *Peer) AddFile(ctx context.Context, r io.Reader, params *AddParams) (ipld.Node, error) {
	if params == nil {
		params = &AddParams{}
	}
	if params.HashFun == "" {
		params.HashFun = "sha2-256"
	}

	prefix, err := merkledag.PrefixForCidVersion(1)
	if err != nil {
		return nil, fmt.Errorf("bad CID Version: %s", err)
	}

	hashFunCode, ok := multihash.Names[strings.ToLower(params.HashFun)]
	if !ok {
		return nil, fmt.Errorf("unrecognized hash function: %s", params.HashFun)
	}
	prefix.MhType = hashFunCode
	prefix.MhLength = -1

	dbp := helpers.DagBuilderParams{
		Dagserv:    p,
		RawLeaves:  params.RawLeaves,
		Maxlinks:   helpers.DefaultLinksPerBlock,
		NoCopy:     params.NoCopy,
		CidBuilder: &prefix,
	}

	chnk, err := chunker.FromString(r, params.Chunker)
	checkError(err)
	dbh, err := dbp.New(chnk)
	checkError(err)
	var n ipld.Node
	switch params.Layout {
	case "trickle":
		n, err = trickle.Layout(dbh)
	case "balanced", "":
		n, err = balanced.Layout(dbh)
	default:
		return nil, errors.New("invalid Layout")
	}
	return n, err
}

// GetFile returns a reader to a file as identified by its root CID. The file
// must have been added as a UnixFS DAG (default for IPFS).
func (p *Peer) GetFile(ctx context.Context, c cid.Cid) (ufsio.ReadSeekCloser, error) {
	n, err := p.Get(ctx, c)
	if err != nil {
		return nil, err
	}
	return ufsio.NewDagReader(ctx, n, p)
}

// BlockStore offers access to the blockstore underlying the Peer's DAGService.
func (p *Peer) BlockStore() blockstore.Blockstore {
	return p.bstore
}

// HasBlock returns whether a given block is available locally. It is
// a shorthand for .Blockstore().Has().
func (p *Peer) HasBlock(c cid.Cid) (bool, error) {
	return p.BlockStore().Has(c)
}

const connectionManagerTag = "user-connect"
const connectionManagerWeight = 100

// Connect connects host to a given peer
func (p *Peer) Connect(ctx context.Context, pi peer.AddrInfo) error {
	if p.Host == nil {
		return errors.New("peer is offline")
	}

	if swrm, ok := p.Host.Network().(*swarm.Swarm); ok {
		swrm.Backoff().Clear(pi.ID)
	}

	if err := p.Host.Connect(ctx, pi); err != nil {
		return err
	}

	p.Host.ConnManager().TagPeer(pi.ID, connectionManagerTag, connectionManagerWeight)
	return nil
}

// Peers returns a list of connected peers
func (p *Peer) Peers(ctx context.Context) ([]string, error) {
	pIDs := p.Host.Network().Peers()
	peerList := []string{}
	for _, pID := range pIDs {
		peerList = append(peerList, pID.String())
	}
	return peerList, nil
}

// Disconnect host from a given peer
func (p *Peer) Disconnect(ctx context.Context, addr multiaddr.Multiaddr) error {
	if p.Host == nil {
		return errors.New("peer is offline")
	}

	taddr, id := peer.SplitAddr(addr)
	if id == "" {
		return peer.ErrInvalidAddr
	}

	net := p.Host.Network()
	if taddr == nil {
		if net.Connectedness(id) != inet.Connected {
			return errors.New("not connected")
		}
		if err := net.ClosePeer(id); err != nil {
			return err
		}
		return nil
	}
	for _, conn := range net.ConnsToPeer(id) {
		if !conn.RemoteMultiaddr().Equal(taddr) {
			continue
		}
		return conn.Close()
	}
	return nil
}
