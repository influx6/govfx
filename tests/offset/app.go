package main

import (
	"fmt"

	"github.com/influx6/govfx"
)

func main() {

	div := govfx.Document().QuerySelector("div.offset")

	top, left := govfx.Offset(div)
	ptop, pleft := govfx.Position(div)

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
  `, top, left, ptop, pleft, color, govfx.RGBA(color, 50), scolor, govfx.RGBA(scolor, 50)))
}
