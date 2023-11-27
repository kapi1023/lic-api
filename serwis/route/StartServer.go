package route

import "net/http"

type Route struct {
	Method string
	Func   http.HandlerFunc
}

func StartServer(port string, routes []Route) {
	for _, route := range routes {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			if r.Method == route.Method {
				route.Func(w, r)
			}
		})
	}

	http.ListenAndServe(":"+port, nil)
}
