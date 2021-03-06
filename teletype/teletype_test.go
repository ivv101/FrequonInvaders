package teletype

import (
	"github.com/ArchRobison/Gophetica/nimble"
	"testing"
)

type context struct{}

func (*context) Init(int32, int32) {
}

func (*context) KeyDown(k nimble.Key) {
	DisplayCursor(true)
	if 0x20 <= k && k < 0x7F {
		PrintChar(rune(k))
	} else {
		switch k {
		case nimble.KeyReturn:
			PrintChar('\n')
		case nimble.KeyEscape:
			nimble.Quit()
		case nimble.KeyBackspace, nimble.KeyDelete:
			Backup()
		}
	}
}

var flag bool

func (*context) Render(pm nimble.PixMap) {
	if !flag {
		Print("(0123456789),$/\n")
		Print("Type some text and 'enter'.\n")
		Print("Try backspace and del.\n" +
			"Press Esc to quit.")
		flag = true
	}
	Draw(pm)
}

type windowSpec struct {
	width, height int32
	title         string
}

func (w windowSpec) Size() (width, height int32) {
	return w.width, w.height
}

func (w windowSpec) Title() string {
	return w.title
}

// Requires visual inspection
func TestTeletype(t *testing.T) {
	Init("../Characters.png")
	nimble.AddRenderClient(&context{})
	nimble.AddKeyObserver(&context{})
	nimble.Run(windowSpec{1024, 600, "teletype"})
}
