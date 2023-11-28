package route

import "net/http"

type Ruter struct {
	middlewares map[RouteAttr]Middleware
	apiPrefix   string
	globalAttr  []RouteAttr
	kolejnosc   []RouteAttr
}

type LogPanic func(error, map[string]interface{})
type Middleware func(http.Handler, string) http.Handler
