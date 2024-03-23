package main

import (
	_ "embed"
	"fmt"
	"log"
	"os"
)

//go:embed "rights/usa.txt"
var usa string

//go:embed "rights/norway.txt"
var norway string

func main() {
	fmt.Println("Hello world")
	if len(os.Args) < 2 {
		log.Fatal("expected one argument: country")
	}
	lang := os.Args[1]
	if lang == "" {
		log.Fatal("Got empty string expected country")
	}

	rights, err := getRights(lang)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("These are the rights of %s\n", lang)
	fmt.Print(rights)
}

func getRights(lang string) (string, error) {
	if lang == "usa" {
		return usa, nil
	} else if lang == "norway" {
		return norway, nil
	} else {
		return "", fmt.Errorf("got unexected contry: %s", lang)
	}
}
