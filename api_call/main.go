package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Match struct {
	DateTimeGMT, Status, T1, T2, Series string
}

type Response struct {
	Data []Match
}

func main() {
	apiKey := "***********"
	url := "https://api.cricapi.com/v1/cricScore?apikey=" + apiKey

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var response Response
	if err := json.Unmarshal(body, &response); err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, match := range response.Data {
		fmt.Printf("DateTimeGMT: %s\nStatus: %s\nTeam 1: %s\nTeam 2: %s\n\n",
			match.DateTimeGMT, match.Status, match.T1, match.T2)
		// Uncomment the following lines if you want to filter by series
		// if match.Series == "Big Bash League 2024-25" {
		// 	fmt.Printf("DateTimeGMT: %s\nStatus: %s\nTeam 1: %s\nTeam 2: %s\n\n",
		// 		match.DateTimeGMT, match.Status, match.T1, match.T2)
		// }
	}
}
