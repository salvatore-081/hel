package pkg

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/salvatore-081/hel/internal/models"
)

const contentType string = "application/json"

type Client struct {
	Url string
}

func (c *Client) Do(query string, v interface{}) error {
	request := models.Request{
		Query:         query,
		OperationName: nil,
		Variables:     nil,
	}

	jsonEncodedRequest, e := json.Marshal(request)
	if e != nil {
		return e
	}

	buffer := bytes.NewBuffer(jsonEncodedRequest)

	r, e := http.Post(c.Url, contentType, buffer)
	if e != nil {
		return e
	}
	defer r.Body.Close()

	var response models.Response

	e = json.NewDecoder(r.Body).Decode(&response)

	data, e := json.Marshal(*&response.Data)
	if e != nil {
		return e
	}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.UseNumber()
	decoder.DisallowUnknownFields()
	e = decoder.Decode(&v)

	return e
}
