package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID          int      `json:"user_id"`
	Name        string   `json:"name,omitempty"`
	Age         int      `json:"age"`
	Password    string   `json:"-"`
	Permissions []string `json:"roles"`
}

func main() {
	u := User{
		ID:          1,
		Name:        "John Doe",
		Age:         20,
		Password:    "my-pass",
		Permissions: []string{"admin", "group-member"},
	}
	b, err := json.Marshal(u)
	if err != nil {
		fmt.Println("Error marshalling Json", err)
		panic(err)
	}
	fmt.Println(string(b))
}
