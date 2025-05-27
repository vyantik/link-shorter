package req

import (
	"encoding/json"
	"log"
	"net/http"
)

func Decode[T any](r *http.Request) (*T, error) {
	var payload T
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		log.Printf("[Req] - [Decode] - [ERROR] : %s", err)
		return nil, err
	}

	return &payload, nil
}
