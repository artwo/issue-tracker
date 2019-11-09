package utils

import (
	"bytes"
	"encoding/json"
	"log"
)

func ToString(v interface{}) string {
	out, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		log.Printf("Unable to Marshal interface %T", v)
	}
	return string(out)
}

func ToIoReader(v interface{}) *bytes.Reader {
	vBytes, err := json.Marshal(v)
	if err != nil {
		log.Printf("Unable to Marshal interface %T", v)
	}
	return bytes.NewReader(vBytes)
}

func ErrorsToString(errors []error) string {
	var result = "[ "
	for _, err := range errors {
		result = result + "\"" + err.Error() + "\", "
	}
	result = result + "]"
	return result
}

// TODO: Add UUID Generator
