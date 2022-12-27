package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	// 21.7 Customizing JSON with struct tags
	type location struct {
		Lat  float64 `json:"latitude"`
		Long float64 `json:"longitude"`
	}
	curiosity := location{-4.5895, 137.4417}
	// 21.6. Encoding structures to JSON
	bytes, err := json.Marshal(curiosity)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(bytes))

}
