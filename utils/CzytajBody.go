package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func CzytajBody(r *http.Request, wynik interface{}) ([]byte, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	if err := json.Unmarshal(body, &wynik); err != nil {
		return nil, err
	}

	return body, waliduj(wynik)
}
