package utils

import "encoding/json"

func ToJson(obj interface{}) ([]byte, error) {
	return json.Marshal(obj)
}

func FromJson(data []byte, t interface{}) error {
	return json.Unmarshal(data, t)
}
