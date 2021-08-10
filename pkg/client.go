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

func (c *Client) Do(query string) (decodedResponse models.Response, e error) {
	request := models.Request{
		Query:         query,
		OperationName: nil,
		Variables:     nil,
	}

	jsonEncodedRequest, e := json.Marshal(request)
	if e != nil {
		return decodedResponse, e
	}

	buffer := bytes.NewBuffer(jsonEncodedRequest)

	response, e := http.Post(c.Url, contentType, buffer)
	if e != nil {
		return decodedResponse, e
	}
	defer response.Body.Close()

	e = json.NewDecoder(response.Body).Decode(&response)

	return decodedResponse, e
}
