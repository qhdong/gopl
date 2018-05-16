package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

func main() {
	var movies = []Movie{
		{Title: "让子弹飞", Year: 2012, Color: true, Actors: []string{"姜文", "葛优", "胡军"}},
		{Title: "无人区", Year: 2013, Color: true, Actors: []string{"徐峥", "黄渤", "于南"}},
	}
	data, err := json.MarshalIndent(movies, "", "  ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}

	var titles []struct{ Title string }
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)
	fmt.Printf("titles = %+v\n", titles)
}
