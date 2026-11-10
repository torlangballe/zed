//go:build zui && test

package main

import (
	"github.com/torlangballe/zed/zui/zapp"
	"github.com/torlangballe/zed/zui/zcontainer"
	"github.com/torlangballe/zed/zui/zlabel"
	"github.com/torlangballe/zed/zui/zpresent"
	"github.com/torlangballe/zed/zutil/zgeo"
	"github.com/torlangballe/zed/zutil/zgrapher"
)

func main() {
	zapp.New()
	zapp.SetUIDefaults(true)
	stack := zcontainer.StackViewVert("stack")
	att := zpresent.AttributesDefault()
	att.MakeFull = true
	label := zlabel.New("Graph Test")
	stack.Add(label, zgeo.TopLeft)

	var job zgrapher.Job
	job.ID = "test1"
	job.PixelHeight = 20
	job.SecondsPerPixel = 5
	job.WindowMinutes = 20
	graph := zgrapher.NewGraphView(job, "grapherTest")
	graph.MinWidth = 500
	graph.ImagePathPrefix = "http://localhost:7776/"
	stack.Add(graph, zgeo.TopLeft)

	zpresent.PresentView(stack, att)
	select {}
}
