//go:build zui && !js

package zbutton

import (
	"github.com/torlangballe/zed/zui/zkeyboard"
	"github.com/torlangballe/zed/zutil/zgeo"
)

func New(text string) *Button {
	return nil
}

func (v *Button) MakeReturnKeyDefault() {
}

func (v *Button) SetMargin(m zgeo.Rect) {
	v.margin = m
}

func (v *Button) MakeEscapeCanceler() {}
