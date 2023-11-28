package utils

import (
	"encoding/json"
	"net/http"
)

func ZwrocJson(w http.ResponseWriter, dane interface{}, nowyStatusCode ...int) error {
	resp, err := json.Marshal(dane)
	if err != nil {
		return err
	}

	statusCode := http.StatusOK
	if len(nowyStatusCode) > 0 {
		statusCode = nowyStatusCode[0]
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(resp)

	return nil
}
