package jdb

import (
	"context"
	"fmt"
	"github.com/gioapp/cms/pkg/jdb/cfg"
	"github.com/gioapp/cms/pkg/jdb/repo"
	"github.com/ipfs/go-cid"
	crypto "github.com/libp2p/go-libp2p-core/crypto"
	"os"
)

// JavazacDb Structure
type JavazacDB struct {
	ctx       context.Context
	peer      *Peer
	index     map[string]string
	Datastore string
}

func New(datastore string) *JavazacDB {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	root := "/tmp" + string(os.PathSeparator) + repo.Root
	conf, err := cfg.ConfigInit(2048)
	if err != nil {
		return
	}
	err = repo.Init(root, conf)
	if err != nil {
		return
	}

	r, err := repo.Open(root)
	if err != nil {
		return
	}

	lite, err := NewPeer(ctx, r)
	if err != nil {
		panic(err)
	}

	lite.Bootstrap(DefaultBootstrapPeers())

	c, _ := cid.Decode("QmSv66pvzJfjwLHuQCYhd3cekGWNX6Q2o5Y268SNMw8fd8")
	rsc, err := lite.Get(ctx, c)
	if err != nil {
		panic(err)
	}
	//lite.
	fmt.Println("Reeeeee", rsc.Tree("", 0))
	//defer rsc.Close()
	//content, err := ioutil.ReadAll(rsc)
	//if err != nil {
	//	panic(err)
	//}

	//fmt.Println(string(content))
}
