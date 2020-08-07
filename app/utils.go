package cms

import (
	"github.com/gioapp/cms/pkg/container"
	"github.com/w-ingsolutions/c/pkg/lyt"
)

func (g *CMS) twoPanels(l, r int, left, right func(gtx C) D) func(gtx C) D {
	return func(gtx C) D {
		return lyt.Format(gtx, "hflexb(start,f(0.60,inset(30dp10dp30dp10dp,_)),f(0.30,inset(30dp10dp30dp10dp,_)))",
			container.ContainerLayout(g.UI.Theme.Colors["PanelBg"], l, l, l, l, left),
			container.ContainerLayout(g.UI.Theme.Colors["PanelBg"], r, r, r, r, right))
	}
}

//func (g *CMS) GetItems() func(gtx C) D {
//	return items.ItemsList(g.jdb, g.UI.Theme, g.ItemsList)
//}
