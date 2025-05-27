package req

import (
	"app/test/pkg/res"
	"net/http"
)

func HandleBody[T any](w http.ResponseWriter, r *http.Request) (*T, error) {
	body, err := Decode[T](r)
	if err != nil {
		res.Json(w, map[string]string{"error": err.Error()}, http.StatusBadRequest)
		return nil, err
	}

	err = IsValid(body)
	if err != nil {
		res.Json(w, map[string]string{"error": err.Error()}, http.StatusBadRequest)
		return nil, err
	}

	return body, nil
}
