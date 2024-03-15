package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	foo, err := makeFoo()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(foo.name)

	jsonIntro()
	jsonUnmarsalWithoutAnonymousStruct()
}

type Foo struct {
	name string
}

func makeFoo() (Foo, error) {
	f := Foo{
		name: "foo",
	}
	return f, nil
}

func jsonIntro() {
	v := struct {
		Name string `json:"Name"`
		Age  int    `josn:"age"`
	}{}
	err := json.Unmarshal([]byte(`{"Name":"John", "age": 30}`), &v)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(v)
}

type person struct {
	Name string `json:"Name"`
	Age  int    `json:"Age"`
}

func jsonUnmarsalWithoutAnonymousStruct() {
	out := person{}
	err := json.Unmarshal([]byte(`{"Name":"Jorgen", "Age": 69}`), &out)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(out)
}
