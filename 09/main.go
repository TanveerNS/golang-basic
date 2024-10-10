package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
}

func encodeToJSON(p Person) (string, error) {
	jsonData, err := json.Marshal(p)
	if err != nil {
		return "", fmt.Errorf("error encoding to JSON: %v", err)
	}
	return string(jsonData), nil
}

func decodeFromJSON(jsonStr string) (Person, error) {
	var p Person
	err := json.Unmarshal([]byte(jsonStr), &p)
	if err != nil {
		return Person{}, fmt.Errorf("error decoding from JSON: %v", err)
	}
	return p, nil
}

func main() {
	person := Person{
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
	}

	jsonStr, err := encodeToJSON(person)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Encoded JSON:")
	fmt.Println(jsonStr)

	decodedPerson, err := decodeFromJSON(jsonStr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\nDecoded Struct:")
	fmt.Printf("%+v\n", decodedPerson)
}
