package cms

import (
	"gioui.org/layout"
	"gioui.org/widget"
)

var (
	submitBtn = new(widget.Clickable)

	navBtn = new(widget.Clickable)

	navList = &layout.List{
		Axis: layout.Vertical,
	}
	contentList = &layout.List{
		Axis: layout.Vertical,
	}
)
