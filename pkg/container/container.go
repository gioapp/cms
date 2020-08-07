package container

import (
	_ "gioui.org/app/permission/storage"
	"gioui.org/layout"
	"gioui.org/unit"
	"github.com/gioapp/gel/helper"
)

func ContainerLayout(bg string, paddingTop, paddingRight, paddingBottom, paddingLeft int, itemContent func(gtx layout.Context) layout.Dimensions) func(gtx layout.Context) layout.Dimensions {
	//vmin := gtx.Constraints.Min.Y
	//if d.FullWidth {
	//	hmin = gtx.Constraints.Max.Y
	//}
	return func(gtx layout.Context) layout.Dimensions {
		return layout.Stack{Alignment: layout.W}.Layout(gtx,
			layout.Expanded(func(gtx layout.Context) layout.Dimensions {
				return helper.Fill(gtx, helper.HexARGB(bg))
			}),
			layout.Stacked(func(gtx layout.Context) layout.Dimensions {
				//gtx.Constraints.Min = hmin
				//gtx.Constraints.Min = vmin
				return layout.Inset{
					Top:    unit.Dp(float32(paddingTop)),
					Right:  unit.Dp(float32(paddingRight)),
					Bottom: unit.Dp(float32(paddingBottom)),
					Left:   unit.Dp(float32(paddingLeft)),
				}.Layout(gtx, itemContent)
			}),
		)
	}
}
