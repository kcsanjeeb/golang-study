package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// GET https://bored-api.appbrewery.com/random
// {
//   "activity": "Learn Express.js",
//   "availability": 0.25,
//   "type": "education",
//   "participants": 1,
//   "price": 0.1,
//   "accessibility": "Few to no challenges",
//   "duration": "hours",
//   "kidFriendly": true,
//   "link": "https://expressjs.com/",
//   "key": "3943506"
// }

type boringResponse struct {
	Activity      string  `json:"activity"`
	Type          string  `json:"type"`
	Participants  int     `json:"participants"`
	Price         float64 `json:"price"`
	Link          string  `json:"link"`
	Key           string  `json:"key"`
	Accessibility string  `json:"accessibility"`
}

func main() {
	ctx := context.Background()
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://bored-api.appbrewery.com/random", nil)
	if err != nil {
		fmt.Println("error creating request: ", err)
		return
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error doing request: ", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error invalid status code: ", resp.Status)
	}

	var boringRes boringResponse
	if err := json.NewDecoder(resp.Body).Decode(&boringRes); err != nil {
		fmt.Println("error reading json response body", err)
		return
	}

	fmt.Printf("%+v\n", boringRes)
}
