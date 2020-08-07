package cms

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/gioapp/cms/pkg/container"
	"github.com/gioapp/cms/pkg/theme"
	"github.com/gioapp/gel/helper"
	"github.com/w-ingsolutions/c/pkg/lyt"
)

func (g *CMS) page(page Page) func(gtx C) D {
	return container.ContainerLayout(g.UI.Theme.Colors["White"], 0, 0, 0, 0, func(gtx C) D {
		return lyt.Format(gtx, "vflexs(start,r(inset(0dp30dp20dp30dp,_)),f(1,inset(0dp0dp0dp16dp,_)))",
			page.Header,
			func(gtx C) D {
				gtx.Constraints.Min.X = gtx.Constraints.Max.X
				return contentList.Layout(gtx, len(page.Body), func(gtx C, i int) D {
					return lyt.Format(gtx, "vflexb(middle,r(_),r(_))",
						page.Body[i],
						helper.DuoUIline(false, 0, 0, 1, g.UI.Theme.Colors["Gray"]),
					)
				})
			})
	})
}

func IconButton(th theme.Theme, b *widget.Clickable, f func(), icon, page string) func(gtx C) D {
	return func(gtx C) D {
		btn := material.IconButton(th.T, b, th.Icons[icon])
		btn.Inset = layout.Inset{unit.Dp(2), unit.Dp(2), unit.Dp(2), unit.Dp(2)}
		btn.Size = unit.Dp(21)
		btn.Background = helper.HexARGB(th.Colors["Secondary"])
		for b.Clicked() {
			f()
			//g.currentPage = page
		}
		return btn.Layout(gtx)
	}
}
