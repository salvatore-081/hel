package models

type ErrorPolicy int

const (
	None ErrorPolicy = iota
	Ignore
	All
)

type Opts struct {
	ErrorPolicy ErrorPolicy
}
