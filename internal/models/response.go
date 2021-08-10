package models

type Response struct {
	Data   *interface{}   `json:"data,omitempty"`
	Errors *[]interface{} `json:"errors,omitempty"`
}
