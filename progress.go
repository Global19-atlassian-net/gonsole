package gonsole

import (
	"fmt"
	"strings"

	xs "github.com/huandu/xstrings"
)

type Progress struct {
	BaseControl

	// a value between 0 and 1
	Value float32
}

func NewProgress(win *Window, parent Container, id string) *Progress {
	p := &Progress{}
	p.Init(win, parent, id, "progress")
	parent.AddControl(p)
	return p
}

func (p *Progress) Repaint() {
	if !p.Dirty() {
		return
	}
	p.BaseControl.Repaint()

	cb := p.ContentBox()

	text := strings.Repeat(" ", (cb.Width/2)-1)
	text += fmt.Sprintf("%d%%", int(p.Value*100))
	text += strings.Repeat(" ", (cb.Width/2)-3)

	t := p.Theme()
	percent := int(p.Value * float32(cb.Width))
	DrawTextSimple(xs.Slice(text, 0, percent), false, p.ContentBox(), t.ColorTermbox("filled.fg"), t.ColorTermbox("filled.bg"))
	DrawTextSimple(xs.Slice(text, percent, -1), false, Box{cb.Left + percent, cb.Top, cb.Width - percent, cb.Height}, t.ColorTermbox("empty.fg"), t.ColorTermbox("empty.bg"))
}
