package handler

import "encoding/json"

func jsonError(msg string) []byte {
	rawError := struct {
		Message string `json:"message"`
	}{
		Message: msg,
	}
	r, err := json.Marshal(rawError)
	if err != nil {
		return []byte(err.Error())
	}
	return r
}
