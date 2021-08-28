package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/salvatore-081/hel/models"
)

type Client struct {
	Url  string
	Opts models.Opts
}

func (c *Client) Do(query string, operationName *string, variables *map[string]interface{}, v interface{}) error {
	gqlRequest := models.GraphQLRequest{
		Query:         query,
		OperationName: operationName,
		Variables:     variables,
	}

	jsonEncodedRequest, e := json.Marshal(gqlRequest)
	if e != nil {
		return e
	}

	buffer := bytes.NewBuffer(jsonEncodedRequest)

	response, e := http.Post(c.Url, "application/json", buffer)
	if e != nil {
		return e
	}
	defer response.Body.Close()

	var gqlResponse models.GraphQLResponse

	e = json.NewDecoder(response.Body).Decode(&gqlResponse)
	if e != nil {
		return e
	}

	if gqlResponse.Data != nil && (c.Opts.ErrorPolicy != 0 || gqlResponse.Errors == nil) {
		data, e := json.Marshal(gqlResponse.Data)
		if e != nil {
			return e
		}

		decoder := json.NewDecoder(bytes.NewReader(data))
		decoder.UseNumber()
		decoder.DisallowUnknownFields()

		e = decoder.Decode(&v)
		if e != nil {
			return e
		}
	}

	if gqlResponse.Errors != nil && c.Opts.ErrorPolicy != 1 {
		gqlErrors, e := json.MarshalIndent(gqlResponse.Errors, "", " ")
		if e != nil {
			return e
		}
		return errors.New(fmt.Sprintf("%s", string(gqlErrors)))
	}

	return nil
}

func NewClient(host string, opts models.Opts) Client {
	return Client{
		Url:  host,
		Opts: opts,
	}
}
