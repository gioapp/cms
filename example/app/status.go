package exampleApp

import (
	"gioui.org/text"
	"github.com/gioapp/cms/pkg/items"
	"github.com/gioapp/cms/pkg/theme"
	"github.com/w-ingsolutions/c/pkg/lyt"
	"math"
	"strconv"
)

var (
	suffixes = [7]string{
		0: "B",
		1: "KB",
		2: "MB",
		3: "GB",
		4: "TB",
		5: "PB",
		6: "EB",
	}
	//status     = new(Status)
	liveStatus = new(statLive)
)

type Status struct {
	Title       string
	HostingSize uint
	PeerId      string
	Version     string
	Gateway     string
	Api         string
	Addresses   []string
	Pub         string
	Live        statLive
}
type statLive struct {
	RateOut  float64
	RateIn   float64
	TotalIn  int64
	TotalOut int64
}

//func GetStatus(sh shell.Shell) *Status {
//	f, err := sh.ID()
//	checkError(err)
//	fv, ft, err := sh.Version()
//	checkError(err)
//	return &Status{
//		Title: "string",
//		//HostingSize: "uint",
//		PeerId:  f.ID,
//		Version: fv + " " + ft,
//		//Gateway:   ,
//		//Api:       "string",
//		Addresses: f.Addresses,
//		Pub:       f.PublicKey,
//	}
//}

//
//func GetLiveStat() {
//
//	fbw, err := g.sh.StatsBW(g.ctx)
//	checkError(err)
//	live = &statLive{
//		//HostingSize: "uint",
//		RateOut:  fbw.RateOut,
//		RateIn:   fbw.RateIn,
//		TotalIn:  fbw.TotalIn,
//		TotalOut: fbw.TotalOut,
//	}
//}

func statusHeader(th *theme.Theme, s *Status) func(gtx C) D {
	//return container.ContainerLayout(th.Colors["PanelBg"], 10, 10, 10, 10, func(gtx C) D {
	//	gtx.Constraints.Min.X = gtx.Constraints.Max.X
	//	helper.Fill(gtx, helper.HexARGB(th.Colors["PanelBg"]))
	//	return lyt.Format(gtx, "vflexb(middle,r(inset(5dp0dp5dp0dp,_)),r(inset(5dp0dp30dp0dp,_)),r(inset(5dp0dp5dp0dp,_)),r(inset(5dp0dp5dp0dp,_)),r(inset(5dp0dp5dp0dp,_)),r(inset(5dp0dp5dp0dp,_)),r(inset(5dp0dp5dp0dp,_)),r(inset(5dp0dp5dp0dp,_)))",
	//		func(gtx C) D {
	//			gtx.Constraints.Min.X = gtx.Constraints.Max.X
	//			title := theme.H1(th, "Connected to IPFS")
	//			title.Alignment = text.Start
	//			return title.Layout(gtx)
	//		},
	//		func(gtx C) D {
	//			gtx.Constraints.Min.X = gtx.Constraints.Max.X
	//			title := theme.H6(th, "Hosting 54.1 MB of files — Discovered 149 peers")
	//			title.Alignment = text.Start
	//			return title.Layout(gtx)
	//		},
	//		statusRow(th, "PEER ID", row(th, s.PeerId)),
	//		statusRow(th, "VERSION", row(th, s.Version)),
	//		statusRow(th, "GATEWAY", row(th, s.Gateway)),
	//		statusRow(th, "API", row(th, s.Api)),
	//		statusRow(th, "ADDRESSES", container.ContainerLayout(th.Colors["White"], 10, 10, 10, 10, func(gtx C) D {
	//			return addressesList.Layout(gtx, len(s.Addresses), func(gtx C, i int) D {
	//				title := theme.Body(th, s.Addresses[i])
	//				title.Alignment = text.Start
	//				return title.Layout(gtx)
	//			})
	//		})),
	//		statusRow(th, "PUBLIC KEY", container.ContainerLayout(th.Colors["White"], 10, 10, 10, 10, func(gtx C) D {
	//			title := theme.Body(th, s.Pub)
	//			title.Alignment = text.Start
	//			return title.Layout(gtx)
	//		})))
	//})
	return func(gtx C) D {
		return D{}

	}
}

func row(th *theme.Theme, label string) func(gtx C) D {
	return func(gtx C) D {
		t := theme.Body(th, label)
		t.Alignment = text.Start
		return t.Layout(gtx)
	}
}

func statusRow(th *theme.Theme, label string, content func(gtx C) D) func(gtx C) D {
	return func(gtx C) D {
		return lyt.Format(gtx, "hflexb(start,r(inset(0dp16dp0dp0dp,_)),f(1,_))",
			//gtx.Constraints.Min.X = gtx.Constraints.Max.X
			func(gtx C) D {
				gtx.Constraints.Min.X = 100
				title := theme.Body(th, label)
				title.Alignment = text.Start
				return title.Layout(gtx)
			},
			content,
		)
	}
}

//
//func statusBody(th *theme.Theme, l *statLive) []func(gtx C) D {
//	return []func(gtx C) D{
//		func(gtx C) D {
//			var (
//				rateIn   string = "0"
//				rateOut  string = "0"
//				totalIn  string = "0"
//				totalOut string = "0"
//			)
//			//if live != nil {
//			if l.RateIn != 0 {
//				rateIn = formatByteSize(l.RateIn)
//			}
//			if l.RateOut != 0 {
//				rateOut = formatByteSize(l.RateOut)
//			}
//			if l.TotalIn != 0 {
//				totalIn = formatByteSize(float64(l.TotalIn))
//			}
//			if l.TotalOut != 0 {
//				totalOut = formatByteSize(float64(l.TotalOut))
//			}
//
//			//fmt.Println("gore", formatByteSize(0))
//			//}
//			return lyt.Format(gtx, "vflexb(middle,r(_),r(_),r(_),r(_))",
//				statusRow(th, "RateIn: ", row(th, rateIn)),
//				statusRow(th, "RateOut: ", row(th, rateOut)),
//				statusRow(th, "TotalIn: ", row(th, totalIn)),
//				statusRow(th, "TotalOut: ", row(th, totalOut)),
//			)
//			//return D{}
//		},
//	}
//}

func statusBody(th *theme.Theme, itms items.I) []func(gtx C) D {
	return []func(gtx C) D{
		//fmt.Println("gore", formatByteSize(0))
		items.ItemsList(th, itms),
	}
}

func Round(val float64, roundOn float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}

func formatByteSize(size float64) string {
	//size := sizeInMB * 1024 * 1024
	base := math.Log(size) / math.Log(1024)
	getSize := Round(math.Pow(1024, base-math.Floor(base)), .5, 2)
	getSuffix := suffixes[int(math.Floor(base))]
	//fmt.Println("dole", strconv.FormatFloat(getSize, 'f', -1, 64)+" "+string(getSuffix))
	return strconv.FormatFloat(getSize, 'f', -1, 64) + " " + string(getSuffix)
}
