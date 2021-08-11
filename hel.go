package main

import (
	"github.com/salvatore-081/hel/pkg"
)

func NewClient(host string) pkg.Client {
	return pkg.Client{
		Url: host,
	}
}

// func main() {
// 	client := NewClient("https://api.spacex.land/graphql/")
// 	query := `
// query{
//   ships{
//     name
//   }
// }
//   `
// 	type Ships struct {
// 		Ships []struct {
// 			Names string
// 		}
// 	}
// 	var ships Ships
// 	e := client.Do(query, &ships)
// 	if e != nil {
// 		fmt.Println(fmt.Sprintf("error => %s", e.Error()))
// 	}
// 	fmt.Println(fmt.Sprintf("ships => %v", ships))

// }
