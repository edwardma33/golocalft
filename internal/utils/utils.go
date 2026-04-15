package utils

import (
	"encoding/json"
	"net/http"
)

type JsonMap map[string]any

func WriteJson(w http.ResponseWriter, status int, body JsonMap) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	data, err := json.Marshal(body)
	if err != nil {
		return err
	}
	_, err = w.Write(data)
	return err
}
