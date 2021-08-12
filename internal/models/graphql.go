package models

type GraphQLRequest struct {
	Query         string                  `json:"query"`
	OperationName *string                 `json:"operationName,omitempty"`
	Variables     *map[string]interface{} `json:"variables,omitempty"`
}

type GraphQLResponse struct {
	Data   *interface{}   `json:"data,omitempty"`
	Errors *[]interface{} `json:"errors,omitempty"`
}
