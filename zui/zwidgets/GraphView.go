//go:build zui

package zwidgets

import (
	"github.com/torlangballe/zed/zutil/zgeo"
)

type GraphView struct {
}

func GraphViewNew(minSize zgeo.Size) *GraphView {
	v := &GraphView{}
	return v
}
