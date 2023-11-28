package route

import "net/http"

type Route struct {
	MetodaHttp string
	Path       string
	Func       http.HandlerFunc
	logPrefix  string
	Attr       []RouteAttr
	SkipAttr   []RouteAttr
}
