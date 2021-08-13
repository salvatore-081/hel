package pkg

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/salvatore-081/hel/internal/models"
)

const contentType string = "application/json"

type Client struct {
	Url string
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

	response, e := http.Post(c.Url, contentType, buffer)
	if e != nil {
		return e
	}
	defer response.Body.Close()

	var gqlResponse models.GraphQLResponse

	e = json.NewDecoder(response.Body).Decode(&gqlResponse)
	if e != nil {
		return e
	}

	// errors, _ := json.MarshalIndent(gqlResponse.Errors, "", " ")
	fmt.Println(fmt.Sprintf("gqlResponse => %v - data => %v", gqlResponse, *gqlResponse.Data))

	// if gqlResponse.Data != nil {
	// 	data, e := json.Marshal(gqlResponse.Data)
	// 	if e != nil {
	// 		return e
	// 	}

	// 	decoder := json.NewDecoder(bytes.NewReader(data))
	// 	decoder.UseNumber()
	// 	decoder.DisallowUnknownFields()

	// 	e = decoder.Buffered().Read()(&v)
	// 	if e != nil {
	// 		return e
	// 	}
	// }

	// if gqlResponse.Errors != nil {
	// 	errors, e := json.Marshal(gqlResponse.Errors)
	// 	if e != nil {
	// 		return e
	// 	}

	// 	decoder := json.NewDecoder(bytes.NewReader(errors))
	// 	decoder.UseNumber()
	// 	decoder.DisallowUnknownFields()

	//   e =

	// }

	if gqlResponse.Errors != nil {
		gqlErrors, e := json.MarshalIndent(gqlResponse.Errors, "", " ")
		if e != nil {
			return e
		}

		return errors.New(fmt.Sprintf("%s", string(gqlErrors)))
	}

	return nil

}
