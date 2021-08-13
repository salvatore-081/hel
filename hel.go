package main

import (
	"fmt"

	"github.com/salvatore-081/hel/pkg"
)

func NewClient(host string) pkg.Client {
	return pkg.Client{
		Url: host,
	}
}

func main() {
	// client := NewClient("https://api.spacex.land/graphql/")
	// 	query := `
	// query{
	//   ships{
	//     name
	//   }
	// }
	//   `
	// 	type Ships struct {
	// 		Ships []struct {
	// 			Name string
	// 		}
	// 	}
	// 	var ships interface{}
	var user map[string]interface{}
	var errors []map[string]interface{}
	client := NewClient("http://localhost:8081/graphql")
	query := `
query{
  user(username:"diagnostc"){
    id
    organization{
      id
      name
      eid
    }
  }
}
`

	e := client.Do(query, nil, nil, &user)
	if e != nil {
		fmt.Println(fmt.Sprintf("DO ERROR => %s", e.Error()))
	}
	fmt.Println(fmt.Sprintf("data => %v", user))
	fmt.Println(fmt.Sprintf("errors => %v", errors))

}
