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

func (c *Client) Do(query string, d interface{}, e interface{}) error {
	gqlRequest := models.GraphQLRequest{
		Query:         query,
		OperationName: nil,
		Variables:     nil,
	}

	jsonEncodedRequest, err := json.Marshal(gqlRequest)
	if err != nil {
		return err
	}

	buffer := bytes.NewBuffer(jsonEncodedRequest)

	response, err := http.Post(c.Url, contentType, buffer)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	var gqlResponse models.GraphQLResponse

	err = json.NewDecoder(response.Body).Decode(&gqlResponse)
	if err != nil {
		return err
	}

	if d != nil {
		err = graphqlDecoder(gqlResponse.Data, d)
		if err != nil {
			return err
		}
	}

	if e != nil {
		err = graphqlDecoder(gqlResponse.Errors, e)
		if err != nil {
			return err
		}
	}

	return nil
}
