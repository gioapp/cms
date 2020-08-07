package items

import (
	"fmt"
	_ "gioui.org/app/permission/storage"
	"gioui.org/layout"
	"gioui.org/widget"
	"github.com/gioapp/cms/pkg/itembtn"
	"github.com/gioapp/cms/pkg/theme"
	"github.com/gioapp/gel/helper"
	"github.com/ipfs/go-cid"
	"github.com/w-ingsolutions/c/pkg/lyt"
)

var (
	l = &layout.List{
		Axis: layout.Vertical,
	}
)

type I []*FolderListItem

type FolderListItem struct {
	Name  string
	Cid   cid.Cid
	Size  uint64
	Type  uint8
	Btn   *widget.Clickable
	Check *widget.Bool
}

func ItemsList(th *theme.Theme, itms I) func(gtx layout.Context) layout.Dimensions {
	return func(gtx layout.Context) layout.Dimensions {
		return l.Layout(gtx, len(itms), func(gtx layout.Context, i int) layout.Dimensions {
			itm := itms[i]
			return lyt.Format(gtx, "vflexb(middle,r(_),r(_))",
				func(gtx layout.Context) layout.Dimensions {
					b := itembtn.ItemBtn(th, itm.Btn, itm.Check, th.Icons["GlyphFolder"], th.Icons["GlyphDots"], itm.Name, itm.Cid.String(), itm.Size).Layout(gtx)
					for itm.Btn.Clicked() {
						fmt.Println("Name", "/"+itm.Name)

						//g.ItemsList = listFolder(g.ctx, g.sh, "/"+itm.Name)
						//itms = g.jdb.ReadList( "QmSv66pvzJfjwLHuQCYhd3cekGWNX6Q2o5Y268SNMw8fd8")
					}
					return b
				},
				helper.DuoUIline(false, 0, 0, 1, th.Colors["Gray"]),
			)
		})
	}
}

//
//func listFolder(ctx context.Context, path string) I {
//	//f, err := sh.FilesLs(ctx, path, shell.FilesLs.Stat(true))
//	//checkError(err)
//	var folder []*FolderListItem
//	for _, item := range f {
//		fmt.Println("item", item)
//		folder = append(folder, &FolderListItem{
//			Name:  item.Name,
//			//Cid:  item.Hash,
//			Size:  item.Size,
//			Type:  item.Type,
//			Btn:   new(widget.Clickable),
//			Check: new(widget.Bool),
//		})
//	}
//	return folder
//}
func filesBody() []func(gtx layout.Context) layout.Dimensions {
	return []func(gtx layout.Context) layout.Dimensions{
		func(gtx layout.Context) layout.Dimensions {
			//return lyt.Format(gtx, "vflexb(middle,r(inset(5dp0dp5dp0dp,_)),r(inset(5dp0dp30dp0dp,_)),r(inset(5dp0dp5dp0dp,_),r(inset(5dp0dp5dp0dp,_)))",
			//	statusRow(g.UI.Theme, "RateIn: ", row(g.UI.Theme, fmt.Sprint(g.Status.Live.RateIn))),
			//	statusRow(g.UI.Theme, "RateOut: ", row(g.UI.Theme, fmt.Sprint(g.Status.Live.RateOut))),
			//	statusRow(g.UI.Theme, "TotalIn: ", row(g.UI.Theme, fmt.Sprint(g.Status.Live.TotalIn))),
			//	statusRow(g.UI.Theme, "TotalOut: ", row(g.UI.Theme, fmt.Sprint(g.Status.Live.TotalOut))),
			//)
			return layout.Dimensions{}
		},
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}
}
