package helper

import (
	"encoding/json"
)

func ErrorMessageJson(msg string) []byte {
	ret, _ := json.Marshal(map[string]string{
		"Error": msg,
	})
	return ret
}
