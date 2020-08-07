package cms

import (
	"github.com/w-ingsolutions/c/pkg/lyt"
)

func (g *CMS) AppMain() {
	lyt.Format(g.UI.Context, "hflexb(start,r(_),f(1,_))",
		func(gtx C) D {
			gtx.Constraints.Min.Y = gtx.Constraints.Max.Y
			n := Navigation{
				Name:  "Navigacion",
				Bg:    g.UI.Theme.Colors["NavBg"],
				Items: g.menuItems,
			}
			width := 252
			if g.UI.mob {
				width = 128
			}
			return n.Nav(g.UI.Theme, gtx, width, g.UI.mob, g.logo())
		},
		func(gtx C) D {
			return lyt.Format(gtx, "vflexb(start,f(1,_))",
				//g.UI.Header,
				g.page(g.UI.pages[g.currentPage]))
			//func(gtx C) D {
			//	return theme.H3(g.UI.Theme, "dfdffdf").Layout(gtx)
			//})
		})
}

func (g *CMS) BeforeMain() {

}

func (g *CMS) AfterMain() {

}

func (g *CMS) Tik() func() {
	return func() {

	}
}
