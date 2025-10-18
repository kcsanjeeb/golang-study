package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name       string   `json:"full_name"`
	Age        int      `json:"years_old,omitempty"`
	Occupation string   `json:"-"`
	Language   []string `json:"spoken_languages"`
}

func main() {
	jsonData := `{"full_name":"Jane Doe" , "years_old":12 , "spoken_languages":["Chinese","English"]}`
	var person Person

	err := json.Unmarshal([]byte(jsonData), &person)
	if err != nil {
		fmt.Println("Error unmarshalling Json ", err)
		panic(err)
	}

	fmt.Println("name", person.Name)
	fmt.Println("age", person.Age)
	fmt.Println("lang", person.Language)
	fmt.Println("occ", person.Occupation)

}
