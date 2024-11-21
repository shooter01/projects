package req

import (
	responses "go/adv-demo/pkg/res"
	"net/http"
)

func HandleBody[T any](w http.ResponseWriter, r *http.Request) (*T, error) {
	body, err := Decode[T](r.Body)
	if err != nil {
		responses.Json(w, err.Error(), 402)
		return nil, err
	}
	err = Validate(body)

	return &body, nil
}
