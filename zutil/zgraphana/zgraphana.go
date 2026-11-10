package zgraphana

import (
	"time"

	"github.com/torlangballe/zed/zutil/zfile"
	"github.com/torlangballe/zed/zutil/zhttp"
	"github.com/torlangballe/zed/zutil/zkeyvalrpc"
	"github.com/torlangballe/zed/zutil/zlog"
)

type Annotation struct {
	DashboardUID string    `json:"dashboardUID"`
	PanelID      int       `json:"panelId"`
	Time         time.Time `json:"-"` // is converted to TimeEpocMS
	TimeEnd      time.Time `json:"-"` // is converted to TimeEndEpocMS
	Tags         []string  `json:"tags"`
	Text         string    `json:"text"`

	TimeEpocMS    int64 `json:"time"`
	TimeEndEpocMS int64 `json:"timeEnd"`
}

type GraphanaResult struct {
	Message string `json:"message"`
	ID      int    `json:"id"`
}

var (
	APIKey       = zkeyvalrpc.NewOption[string]("GraphanaAPIKey", "")
	URLPrefix    = zkeyvalrpc.NewOption[string]("GraphanaURLPrefix", "")
	DashboardUID = zkeyvalrpc.NewOption[string]("GraphanaDashboardUID", "")
)

func SetAnnotation(graphanaURLPrefix string, a Annotation) error {
	if a.DashboardUID == "" {
		a.DashboardUID = DashboardUID.Get()
	}
	if graphanaURLPrefix == "" {
		graphanaURLPrefix = URLPrefix.Get()
	}
	var got GraphanaResult
	if !a.Time.IsZero() {
		a.TimeEpocMS = a.Time.UnixMilli()
	}
	if !a.TimeEnd.IsZero() {
		a.TimeEndEpocMS = a.TimeEnd.UnixMilli()
	}
	// zlog.Info("ANNO:", a.Text, a.Time, a.TimeEnd)
	surl := zfile.JoinPathParts(graphanaURLPrefix, "api/annotations")
	params := zhttp.MakeParameters()
	// params.PrintBody = true
	key := APIKey.Get()
	if a.TimeEnd.IsZero() && a.Time.IsZero() || a.DashboardUID == "" || graphanaURLPrefix == "" || key == "" {
		return zlog.Error("Missing parameters for SetAnnotation", a)
	}
	params.Headers["Authorization"] = "Bearer " + key
	_, err := zhttp.Post(surl, params, a, &got)
	if zlog.OnError(err, "annot") {
		return err
	}
	return nil
}
