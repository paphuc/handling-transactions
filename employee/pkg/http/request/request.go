package request

import (
	"bytes"
	"encoding/json"
	"net/http"

	"handling-transactions/employee/pkg/http/response"
)

func Post(url string, body []byte) (response.Base, error) {
	result := response.Base{}
	r, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return result, err
	}
	defer r.Body.Close()
	json.NewDecoder(r.Body).Decode(&result)
	return result, nil
}
