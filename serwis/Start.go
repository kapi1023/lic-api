package serwis

import (
	"log"
	"net/http"

	"github.com/kapi1023/licencjat/api/route"
)

func Start(port string) {
	r := route.Start()
	log.Fatal(r.ListenAndServe([]route.Route{
		{
			MetodaHttp: http.MethodPost,
			Path:       "/essa",
			Func: func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("Hello, GET!"))
			},
		},
		{
			MetodaHttp: http.MethodPost,
			Path:       "/pobierz-test",
			Func:       pobierzTest(),
		},
	}, port))
}
