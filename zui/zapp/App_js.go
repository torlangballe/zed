package zapp

import (
	"net/url"
	"strconv"
	"strings"

	"github.com/torlangballe/zed/zui"
	"github.com/torlangballe/zed/zui/zdom"
	"github.com/torlangballe/zed/zui/zwidgets"
	"github.com/torlangballe/zed/zutil/xrpc"
	"github.com/torlangballe/zed/zutil/zbool"
	"github.com/torlangballe/zed/zutil/zdebug"
	"github.com/torlangballe/zed/zutil/zdict"
	"github.com/torlangballe/zed/zutil/zerrors"
	"github.com/torlangballe/zed/zutil/zlog"
	"github.com/torlangballe/zed/zutil/zrest"
	"github.com/torlangballe/zed/zutil/zstr"
)

var cookies map[string]string

// URL returns the url that invoked this app
func URL() *url.URL {
	u, err := url.Parse(URLString())
	zlog.AssertNotError(err)
	return u
}

func URLString() string {
	return zdom.WindowJS.Get("location").Get("href").String()
}

// MainArgs returns the path of browser and url parameters as args map of max one parameter of each key
func MainArgs() (path string, args map[string]string) {
	args = map[string]string{}
	u := URL()
	path = strings.TrimRight(u.Path, "/")
	for k, v := range u.Query() {
		args[k] = v[0]
	}
	return
}

func Cookies() map[string]string {
	if cookies == nil {
		params := zdom.DocumentJS.Get("cookie")
		if params.IsUndefined() {
			return map[string]string{}
		}
		cookies = zstr.GetParametersFromURLArgString(params.String())
		zlog.Info("Got Cookies:", cookies)
	}
	return cookies
}

func URLStub() string {
	u := URL()
	u.Path = ""
	u.RawQuery = ""
	return u.String()
}

func guiRestartHandler(err error) {
	if err == nil {
		return
	}
	dict := zdict.Dict{}
	ce := zerrors.MakeContextError(dict, "GUI Restart", err)
	callErr := xrpc.MainCaller().Call("AppCalls.SetGUIError", ce, nil)
	zlog.OnError(callErr)
}

// SetUIDefaults sets up an app, including some sensible defaults for rpc communicated with server counterpart
// if rpcPort is 0, it will be set to the port of the URL, or 80/443 depending on scheme.
// if -1, no rpc client will be set up.
func SetUIDefaults(rpcPort int) (path string, args map[string]string) {
	url := URL()
	ssl := (url.Scheme == "https")
	if rpcPort == 0 {
		if url.Port() != "" {
			rpcPort, _ = strconv.Atoi(url.Port())
		} else {
			rpcPort = 80
			if ssl {
				rpcPort = 443
			}
		}
	}
	args = map[string]string{}
	for k, v := range url.Query() {
		if k == "zdev" && v[0] == "1" {
			zui.DebugOwnerMode = true
		}
		args[k] = v[0]
	}
	url.RawQuery = ""
	DownloadPathPrefix = url.String()
	zwidgets.DocumentationPathPrefix = DownloadPathPrefix + "doc/"
	path, args = MainArgs()
	if zbool.FromString(args["ztest"], false) {
		zdebug.IsInTests = true // for testing gui
	}
	if rpcPort > 0 {
		serverURL := zrest.AppURLPrefix
		if !strings.HasPrefix(serverURL, "/") {
			serverURL = "/" + serverURL
		}
		xrpc.MainClientID = "zui-webclient-" + zstr.GenerateRandomHexBytes(14)
		xrpc.SetupSimpleClient(rpcPort, ssl, serverURL, xrpc.MainClientID)
		zdebug.HandleRestartFunc = guiRestartHandler
		xrpc.MainRPC.Start()
	}
	return
}
