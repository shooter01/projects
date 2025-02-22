package req

import (
	response "go/adv-demo2/pkg/res"
	"net/http"
)

func HandleBody[T any](w http.ResponseWriter, res *http.Request) (*T, error) {
	body, err := Decode[T](res.Body)
	if err != nil {
		response.Json(w, err.Error(), 402)
		return nil, err
	}
	err = isValid(body)
	if err != nil {
		response.Json(w, err.Error(), 402)
		return nil, err
	}
	return &body, nil
}
