package serwis

import (
	"net/http"

	"github.com/kapi1023/licencjat/api/utils"
)

func pobierzTest() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.ZwrocJson(w, "essa", http.StatusOK)
	}
}
