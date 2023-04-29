package main

import (
	"encoding/json"
	"fmt"
)

type Human struct {
	Name string
	Age int
}

// Unmarshal: convert string to struct
func unmarshal2Struct(humanStr string) Human {
	h := Human{}
	err := json.Unmarshal([]byte(humanStr), &h)
	if err != nil {
		println(err)
	}
	return h
}

// Marshal: convert struct to string
func marshal2JsonString(h Human) string {
	h.Age = 30                // change age internally
	updatedBytes, err := json.Marshal(&h)
	if err != nil {
		println(err)
	}
	return string(updatedBytes)
}

func main() {
	fmt.Println("---json marshal and unmarshal---")

	human := new(Human)
	human.Name = "Andy Chen"
	human.Age = 28

	jsonStr := marshal2JsonString(*human)
	fmt.Println(jsonStr)       // {"Name":"Andy Chen","Age":30}

	hStruct := unmarshal2Struct(jsonStr)
	fmt.Println(hStruct)       // {Andy Chen 30}
	fmt.Println(hStruct.Name)  // Andy Chen
	fmt.Println(hStruct.Age)   // 30
}