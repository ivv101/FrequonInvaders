package menu

import (
	"fmt"
	"github.com/ArchRobison/Gophetica/nimble"
	"os"
	"testing"
)

type context struct{}

var theMenu Menu

func (*Menu) ObserveMouse(event nimble.MouseEvent, x, y int32) {
	_ = theMenu.TrackMouse(event, x, y)
}

func (*context) Init(int32, int32) {
}

func (*context) Render(pm nimble.PixMap) {
	pm.Fill(nimble.Gray(0.1))
	theMenu.Draw(pm, 50, 100)
}

type FruitItem struct {
	Item
}

func (f *FruitItem) OnSelect() {
	fmt.Fprintf(os.Stderr, "%v\n", f.Label)
}

var BananaCherry = RadioState{Value: 0, OnSelect: func(value int) {
	fmt.Printf("new state = %v\n", value)
}}

// Requires visual inspection
func TestMenu(t *testing.T) {
	i0 := FruitItem{Item{Label: "Apple"}}
	i1 := Add(MakeRadioItem("Banana", &BananaCherry, 0), Separator)
	i2 := MakeRadioItem("Cherry", &BananaCherry, 1)
	i3 := FruitItem{Item{Label: "Date", Flags: Disabled | Separator}}
	i4 := MakeCheckItem("Elderberry", true, func(val bool) { fmt.Printf("Elderberry=%v\n", val) })
	i5 := MakeSimpleItem("Fig", func() { fmt.Printf("Fig!\n") })
	i6 := MakeSimpleItem("Quit", func() { nimble.Quit() })
	theMenu = Menu{Label: "Fruits",
		Items: []ItemInterface{&i0, i1, i2, &i3, i4, i5, i6}}
	nimble.AddRenderClient(&context{})
	nimble.AddMouseObserver(&theMenu)
	nimble.Run(nil)
}
