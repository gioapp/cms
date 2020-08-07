package main

import (
	"fmt"
	_ "gioui.org/app/permission/storage"
	"gioui.org/layout"
	"gioui.org/op/paint"
	"gioui.org/widget"
	"github.com/gioapp/cms/app"
	"github.com/gioapp/cms/pkg/theme"
	"image"
	"image/png"
	"os"
)

type (
	D = layout.Dimensions
	C = layout.Context
)

var (
	currentPage string

	ipfsLogoTextImageOp = paint.ImageOp{}
	ipfsLogoTextImage   image.Image
	ipfsLogoImageOp     = paint.ImageOp{}
	ipfsLogoImage       image.Image
	ipldLogoImageOp     = paint.ImageOp{}
	ipldLogoImage       image.Image
	daemonConnected     bool

	logoTextImage widget.Image
	logoImage     widget.Image
	logoIpldImage widget.Image

	headerSearchInput = &widget.Editor{
		SingleLine: true,
		Submit:     true,
	}
	apiAddressInput = &widget.Editor{
		SingleLine: true,
		Submit:     true,
	}
	navList = &layout.List{
		Axis: layout.Vertical,
	}
	contentList = &layout.List{
		Axis: layout.Vertical,
	}
	addressesList = &layout.List{
		Axis: layout.Vertical,
	}
	tourBtn    = new(widget.Clickable)
	welcomeBtn = new(widget.Clickable)
	browseBtn  = new(widget.Clickable)
)

func pages(th *theme.Theme, status *Status, live *statLive) map[string]cms.Page {
	return map[string]cms.Page{
		//"Welcome": cms.Page{
		//	Title:  "Welcome",
		//	Header: g.welcomeHeader(),
		//	Body:   g.welcomeBody(),
		//},
		"Status": cms.Page{
			Title:  "Status",
			Header: statusHeader(th, status),
			Body:   statusBody(th, live),
		},
		//"Files": cms.Page{
		//	Title:  "Files",
		//	Header: g.itemsList(),
		//	Body:   g.filesBody(),
		//},
		//"Explore": cms.Page{
		//	Title:  "Explore",
		//	Header: g.exploreHeader(),
		//	Body:   g.exploreBody(),
		//},
		//"Peers": cms.Page{
		//	Title:  "Peers",
		//	Header: g.peersHeader(),
		//	Body:   g.peersBody(),
		//},
		//"Settings": cms.Page{
		//	Title:  "Settings",
		//	Header: g.settingsHeader(),
		//	Body:   g.settingsBody(),
		//},
	}
}

func menuItems() []cms.Item {
	return []cms.Item{
		cms.Item{
			Title: "Status",
			Icon:  "StrokeMarketing",
			Btn:   new(widget.Clickable),
		},
		//cms.Item{
		//	Title: "Files",
		//	Icon:  g.UI.Theme.Icons["StrokeWeb"],
		//	Btn:   new(widget.Clickable),
		//},
		//cms.Item{
		//	Title: "Explore",
		//	Icon:  g.UI.Theme.Icons["StrokeIpld"],
		//	Btn:   new(widget.Clickable),
		//},
		//cms.Item{
		//	Title: "Peers",
		//	Icon:  g.UI.Theme.Icons["StrokeCube"],
		//	Btn:   new(widget.Clickable),
		//},
		//cms.Item{
		//	Title: "Settings",
		//	Icon:  g.UI.Theme.Icons["StrokeSettings"],
		//	Btn:   new(widget.Clickable),
		//},
	}
}

func getImages() {
	ipfsLogoTextImageFile, err := os.Open("/home/marcetin/go/src/github.com/gioapp/cms/pkg/icon/logo/ipfs-text.png")
	checkError(err)
	defer ipfsLogoTextImageFile.Close()
	ipfsLogoTextImage, err = png.Decode(ipfsLogoTextImageFile)
	checkError(err)

	ipfsLogoImageFile, err := os.Open("/home/marcetin/go/src/github.com/gioapp/cms/pkg/icon/logo/ipfs.png")
	checkError(err)
	defer ipfsLogoImageFile.Close()
	ipfsLogoImage, err = png.Decode(ipfsLogoImageFile)
	checkError(err)

	ipldLogoImageFile, err := os.Open("/home/marcetin/go/src/github.com/gioapp/cms/pkg/icon/logo/ipld.png")
	checkError(err)
	defer ipldLogoImageFile.Close()
	ipldLogoImage, err = png.Decode(ipldLogoImageFile)
	checkError(err)

	logoText := paint.NewImageOp(ipfsLogoTextImage)
	logoTextImage = widget.Image{
		Src:   logoText,
		Scale: 1,
	}
	logo := paint.NewImageOp(ipfsLogoImage)
	logoImage = widget.Image{
		Src:   logo,
		Scale: 1,
	}
	logoIpld := paint.NewImageOp(ipldLogoImage)
	logoIpldImage = widget.Image{
		Src:   logoIpld,
		Scale: 1,
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
