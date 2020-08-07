package exampleApp

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
	cms "github.com/gioapp/cms/app"
	"github.com/gioapp/cms/pkg/container"
	"github.com/gioapp/cms/pkg/theme"
	"github.com/gioapp/gel/helper"
	"github.com/w-ingsolutions/c/pkg/lyt"
)

func Header(th theme.Theme) func(gtx C) D {
	return container.ContainerLayout(th.Colors["Info"], 0, 0, 0, 0, func(gtx C) D {
		gtx.Constraints.Min.X = gtx.Constraints.Max.X
		return lyt.Format(gtx, "hflexb(middle,r(inset(0dp0dp0dp6dp,_)),r(inset(20dp30dp20dp3dp,_)))",
			headerSearch(th),
			headerMenu(th),
		)
	})
}
func headerMenu(th theme.Theme) func(gtx C) D {
	return func(gtx C) D {
		for tourBtn.Clicked() {
			currentPage = "Welcome"
		}
		return lyt.Format(gtx, "hflexb(middle,r(_),r(_),r(_))",
			cms.IconButton(th, tourBtn, func() {}, "StrokeCase", ""),
			helper.DuoUIline(true, 0, 2, 2, th.Colors["DarkGrayI"]),
			//g.pageButton(tourBtn, func() {}, "GlyphPencil", "Welcome"),
			func(gtx C) D {
				btn := material.IconButton(th.T, welcomeBtn, th.Icons["GlyphPencil"])
				btn.Inset = layout.Inset{unit.Dp(2), unit.Dp(2), unit.Dp(2), unit.Dp(2)}
				btn.Size = unit.Dp(21)
				btn.Background = helper.HexARGB(th.Colors["Secondary"])
				for welcomeBtn.Clicked() {
					//f()
					currentPage = "Welcome"
				}
				return btn.Layout(gtx)
			},
		)
	}

}

func headerSearch(th theme.Theme) func(gtx C) D {
	return func(gtx C) D {
		return lyt.Format(gtx, "hflexb(middle,r(inset(20dp0dp20dp30dp,_)),r(_))",
			container.ContainerLayout(th.Colors["White"], 8, 8, 8, 8, func(gtx C) D {
				gtx.Constraints.Min.X = 430
				e := material.Editor(th.T, headerSearchInput, "QmHash")
				return e.Layout(gtx)
			}),
			func(gtx C) D {
				e := material.Button(th.T, browseBtn, "Browse")
				e.Inset = layout.Inset{
					Top:    unit.Dp(4),
					Right:  unit.Dp(4),
					Bottom: unit.Dp(4),
					Left:   unit.Dp(4),
				}
				e.CornerRadius = unit.Dp(4)
				return e.Layout(gtx)
			},
		)
	}
}
