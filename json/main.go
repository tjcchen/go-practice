package main

import (
	"encoding/json"
	"fmt"
)

type Human struct {
	Name string
	Age  int
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
	h.Age = 30 // change age internally
	updatedBytes, err := json.Marshal(&h)
	if err != nil {
		println(err)
	}
	return string(updatedBytes)
}

func main() {
	fmt.Println("---json marshal and unmarshal starts---")

	human := new(Human)
	human.Name = "Andy Chen"
	human.Age = 28

	jsonStr := marshal2JsonString(*human)
	fmt.Println(jsonStr) // {"Name":"Andy Chen","Age":30}

	hStruct := unmarshal2Struct(jsonStr)
	fmt.Println(hStruct)      // {Andy Chen 30}
	fmt.Println(hStruct.Name) // Andy Chen
	fmt.Println(hStruct.Age)  // 30

	fmt.Println("---json marshal and unmarshal ends---")

	// json decode example
	err, _ := jsonDecode(jsonStr)
	if err != nil {
		fmt.Println("failed")
	} else {
		fmt.Println("succeed")
	}
}

func jsonDecode(jsonStr string) (error, bool) {
	fmt.Println("\n---json decode starts---")
	var obj interface{}                          // interface{} is a base type
	err := json.Unmarshal([]byte(jsonStr), &obj) // after this, jsonStr becomes a map

	if err != nil {
		return fmt.Errorf("an error occured when unmarshaled json string"), false
	}

	fmt.Println("sss", obj) // sss map[Age:30 Name:Andy Chen]

	objMap, ok := obj.(map[string]interface{})

	if !ok {
		return fmt.Errorf("an error occured when unmarshaled json string"), false
	}

	for k, v := range objMap {
		switch value := v.(type) {
		case string:
			fmt.Printf("type of %s is string, value is %v\n", k, value) // type of Name is string, value is Andy Chen
		case interface{}:
			fmt.Printf("type of %s is interface{}, value is %v\n", k, value) // type of Age is interface{}, value is 30
		default:
			fmt.Printf("type of %s is wrong, value is %v\n", k, value)
		}
	}
	fmt.Println("---json decode ends---")
	return nil, true
}
