//go:build zui

package zwidgets

import (
	"github.com/torlangballe/zed/zui/zpresent"
	"github.com/torlangballe/zed/zui/zview"
)

func init() {
	zpresent.DocumentationIconViewNewFunc = func(path string) zview.View {
		return DocumentationIconViewNew(path)
	}
}
