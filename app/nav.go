package cms

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"github.com/gioapp/cms/pkg/icontextbtn"
	"github.com/gioapp/cms/pkg/theme"
	"github.com/gioapp/gel/helper"
	"github.com/w-ingsolutions/c/pkg/lyt"
)

func (n *Navigation) Nav(th *theme.Theme, gtx layout.Context, width int, noText bool, logo func(gtx layout.Context) layout.Dimensions) layout.Dimensions {
	gtx.Constraints.Min.X = width
	gtx.Constraints.Max.X = width
	gtx.Constraints.Min.Y = gtx.Constraints.Max.Y
	helper.Fill(gtx, helper.HexARGB(n.Bg))
	return lyt.Format(gtx, "vflexs(start,r(_),f(1,_))",
		logo,
		func(gtx layout.Context) layout.Dimensions {
			return navList.Layout(gtx, len(n.Items), func(gtx layout.Context, i int) layout.Dimensions {
				item := n.Items[i]
				btn := icontextbtn.IconTextBtn(th, item.Btn, th.Icons[item.Icon], unit.Dp(85), th.Colors["Info"], item.Title, noText)
				btn.TextSize = unit.Dp(16)
				btn.CornerRadius = unit.Dp(0)
				btn.Background = helper.HexARGB(th.Colors["NavBg"])
				for item.Btn.Clicked() {
					n.currentPage = &item.Title

				}
				return btn.Layout(gtx)
			})
		})
}
