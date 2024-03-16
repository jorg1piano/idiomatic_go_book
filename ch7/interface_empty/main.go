package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	// any is an alias for  interface{}
	data := map[string]any{}
	contents, err := os.ReadFile("ex.json")

	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(contents, &data)

	fmt.Println(data)
}
