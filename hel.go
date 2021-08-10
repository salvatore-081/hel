package hel

import "github.com/salvatore-081/hel/pkg"

func NewClient(host string) pkg.Client {
	return pkg.Client{
		Url: host,
	}
}
