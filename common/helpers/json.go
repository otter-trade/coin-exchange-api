package helpers

import (
	"encoding/json"
	jsoniter "github.com/json-iterator/go"
)

// JsonEncode - Returns the JSON representation of a value
func JsonEncode(v interface{}) ([]byte, error) {

	return json.Marshal(v)
}

// JsonDecode - Decodes a JSON string
func JsonDecode(data []byte, v interface{}) error {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Unmarshal(data, v)
}
