package models

type Request struct {
	Query         string                  `json:"query"`
	OperationName *string                 `json:"operationName,omitempty"`
	Variables     *map[string]interface{} `json:"variables,omitempty"`
}
