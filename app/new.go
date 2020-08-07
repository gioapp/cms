package cms

import (
	"context"
	"fmt"
	"gioui.org/app"
	"gioui.org/unit"
	"github.com/gioapp/cms/pkg/icon/icons"
	"github.com/gioapp/cms/pkg/jdb"
	"github.com/gioapp/cms/pkg/theme"
	"github.com/gioapp/gel/helper"
)

func NewCMS(theme *theme.Theme, header func(gtx C) D) *CMS {
	g := &CMS{
		//sh:  shell.NewShell("/ip4/127.0.0.1/tcp/5001"),
		ctx: context.Background(),
	}

	g.jdb = jdb.New(g.ctx, "datastore")
	//fmt.Println("pagespagespagespages", pages)
	g.ItemsList = g.jdb.ReadList("QmSv66pvzJfjwLHuQCYhd3cekGWNX6Q2o5Y268SNMw8fd8")

	g.UI = cmsUI{
		Theme: theme,
		//mob:   make(chan bool),

	}
	g.currentPage = "Status"
	g.UI.Theme.Icons = icons.NewIPFSicons()
	//g.UI.pages = pages

	g.UI.header = header
	g.UI.Theme.T.Color.Primary = helper.HexARGB(g.UI.Theme.Colors["Primary"])
	g.UI.Theme.T.Color.Text = helper.HexARGB(g.UI.Theme.Colors["Charcoal"])
	g.UI.Theme.T.Color.Hint = helper.HexARGB(g.UI.Theme.Colors["Silver"])
	g.UI.Window = app.NewWindow(
		app.Size(unit.Dp(1024), unit.Dp(800)),
		app.Title("IPFS"),
	)
	//g.menuItems = menuItems

	//
	//getImages()
	//g.GetStatus()

	return g
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}
}
