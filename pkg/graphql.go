package pkg

import (
	"bytes"
	"encoding/json"
)

func graphqlDecoder(response interface{}, v interface{}) error {
	jsonEncoding, err := json.Marshal(response)
	if err != nil {
		return err
	}

	decoder := json.NewDecoder(bytes.NewReader(jsonEncoding))
	decoder.UseNumber()
	decoder.DisallowUnknownFields()

	err = decoder.Decode(&v)
	return err
}
