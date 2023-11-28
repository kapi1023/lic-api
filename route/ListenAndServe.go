package route

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type RouteAttr string

const (
	ATTR_API_PREFIX RouteAttr = "PREFIX"
)

func (r *Ruter) ListenAndServe(routes []Route, port string) error {
	ruter := mux.NewRouter().StrictSlash(true)
	max := getMaxLen(routes)

	for i := range routes {
		route := routes[i]
		route.setLogPrefix(max)

		r.dodajApiPrefix(&route)
		ruter.
			Methods(route.MetodaHttp).
			Path(route.Path).
			Handler(r.dodajMiddlewares(route))
	}

	log.Printf("http.ListenAndServe \":%s\"", port)
	return http.ListenAndServe(":"+port, ruter)
}

func (r *Ruter) dodajApiPrefix(route *Route) {
	if !r.sprawdzWarunki(ATTR_API_PREFIX, route.Attr, route.SkipAttr) {
		return
	}

	route.Path = r.apiPrefix + route.Path
}

func (r *Ruter) dodajMiddlewares(route Route) http.Handler {
	if len(r.middlewares) == 0 {
		return route.Func
	}

	var handler http.Handler = route.Func

	for _, attr := range r.kolejnosc {
		if !r.sprawdzWarunki(attr, route.Attr, route.SkipAttr) {
			continue
		}
		mid, ok := r.middlewares[attr]
		if !ok {
			continue
		}

		handler = mid(handler, route.logPrefix)
	}

	return handler
}

func (r *Ruter) sprawdzWarunki(attr RouteAttr, attrs, skip []RouteAttr) bool {
	if contains(attr, skip) {
		return false
	}

	if contains(attr, r.globalAttr) {
		return true
	}

	return contains(attr, attrs)
}

func (r *Route) setLogPrefix(max int) {
	r.logPrefix = r.Path + tablicaSpacji(max-len(r.Path))
}

func tablicaSpacji(len int) string {
	wynik := ""
	for i := 0; i < len; i++ {
		wynik += " "
	}

	return wynik
}

func getMaxLen(routes []Route) int {
	max := 0
	for _, route := range routes {
		if max < len(route.Path) {
			max = len(route.Path)
		}
	}

	return max
}

func contains(elem RouteAttr, arr []RouteAttr) bool {
	for _, elem2 := range arr {
		if elem2 == elem {
			return true
		}
	}

	return false
}
