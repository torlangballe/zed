<!-- File created by Claude Opus 5.5 (claude-opus-5-5) on behalf of Tor Langballe -->
# zed

Go utility and GUI libraries, in a single module: `github.com/torlangballe/zed`.

- **[zutil](zutil)** — general utilities: strings, geometry, logging, files, HTTP, RPC (`xrpc`), SQL, scheduling, users and more.
- **[zui](zui)** — the Zetrus GUI framework, mainly for the browser via WebAssembly, with some native (macOS) support.

```sh
go get github.com/torlangballe/zed
```

```go
import (
	"github.com/torlangballe/zed/zutil/zstr"
	"github.com/torlangballe/zed/zui/zlabel"
)
```

## Build tags

| Tag      | Use                                     |
|----------|-----------------------------------------|
| `zui`    | GUI code (e.g. `GOOS=js GOARCH=wasm go build -tags zui`) |
| `server` | Server-side code (`go build -tags server`) |

Run tests with `go test -tags server ./...`.

## License

[Mozilla Public License 2.0](LICENSE)
