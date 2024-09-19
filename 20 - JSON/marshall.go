package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

var dogJson = `{"name": "Rex", "breed": "Dalmatian", "Age": 4}`

var dogMap = map[string]interface{}{
	"name":  "Toby",
	"breed": "Poodle",
	"age":   3,
}

type Dog struct {
	Name  string
	Breed string
	Age   int
}

func (d Dog) String() string {
	return fmt.Sprintf("%s is a %s and is %d years old\n", d.Name, d.Breed, d.Age)
}

func main() {
	d1 := Dog{"Toby", "Poodle", 3}
	d1Json, _ := json.Marshal(d1)
	println(bytes.NewBuffer(d1Json).String())

	d2Json, _ := json.Marshal(dogMap)
	println(bytes.NewBuffer(d2Json).String())

	var d2 Dog
	if err := json.Unmarshal([]byte(dogJson), &d2); err != nil {
		log.Fatal(err)
	}
	fmt.Print(d2)

	var d3 map[string]interface{}
	if err := json.Unmarshal(d2Json, &d3); err != nil {
		log.Fatal(err)
	}
	fmt.Print(d3)
}
