package serwis

import "net/http"

func Start(port string) {
	routes := []route.Route{
		{
			Method: "GET",
			Func: func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("Hello, GET!"))
			},
		},
		{
			Method: http.MethodPost,
			Func: func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("Hello, POST!"))
			},
		},
	}
	route.StartServer(port, routes)
}
