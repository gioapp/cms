package cms

import (
	"context"
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget"
	"github.com/gioapp/cms/pkg/theme"
	shell "github.com/ipfs/go-ipfs-api"
)

type (
	D = layout.Dimensions
	C = layout.Context
)

var (
	selected int
)

type CMS struct {
	sh          *shell.Shell
	ctx         context.Context
	UI          cmsUI
	menuItems   []Item
	Settings    gipfsSettings
	ItemsList   []*folderListItem
	currentPage string
}

type folderListItem struct {
	Name  string
	Hash  string
	Size  uint64
	Type  uint8
	btn   *widget.Clickable
	check *widget.Bool
}

type cmsUI struct {
	Device  string
	Window  *app.Window
	Theme   *theme.Theme
	Context layout.Context
	//Ekran   func(gtx layout.Context) layout.Dimensions
	FontSize float32

	mob    bool
	pages  pages
	header func(gtx C) D
	Images map[string]widget.Image
	Ops    op.Ops
}

type gipfsSettings struct {
	Dir  string
	File string
}

type Page struct {
	Title  string
	Header func(gtx C) D
	Body   []func(gtx C) D
}

type pages map[string]Page

type Navigation struct {
	Name        string
	Bg          string
	Logo        Logo
	Items       []Item
	currentPage *string
}
type Item struct {
	Title string
	Bg    string
	//Icon  *widget.Icon
	Icon string
	Btn  *widget.Clickable
	//Page *Page
}

type Logo struct {
	Title string
	Logo  string
}
