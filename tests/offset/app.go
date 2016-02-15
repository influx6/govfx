package main

import (
	"fmt"

	"github.com/influx6/faux/vfx"
)

func main() {

	div := vfx.Document().QuerySelector("div.offset")

	top, left := vfx.Offset(div)
	ptop, pleft := vfx.Position(div)

	color := "#cccccc"
	scolor := "#ccc"

	div.SetInnerHTML(fmt.Sprintf(`
    Offset: %.2f %.2f
    <br/>
    Position: %.2f %.2f
    <br/>
    Color: Hex(%s) Rgba(%s)
    <br/>
    Color: Hex(%s) Rgba(%s)
    <br/>
  `, top, left, ptop, pleft, color, vfx.RGBA(color, 50), scolor, vfx.RGBA(scolor, 50)))
}
