package util

import (
	"encoding/json"
)

// JSONEncode 序列化
func JSONEncode(v interface{}) (string, error) {
	bytes, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
