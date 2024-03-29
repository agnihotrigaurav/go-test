package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title  string `json:"title"`
	Author Author `json:"author"`
}

type Author struct {
	Sales     int  `json:"book_sales"`
	Age       int  `json:"age"`
	Developer bool `json:"is_developer"`
}

type SensorReading struct {
	Name     string `json:"name"`
	Capacity int    `json:"capacity"`
	Time     string `json:"time"`
}

func main() {
	marshel()
	fmt.Println()
	unmarshal()
	fmt.Println()
	unmarshalUnstructred()
}

func marshel() {
	author := Author{Sales: 3, Age: 25, Developer: true}
	book := Book{Title: "Learning Concurrency in Python", Author: author}

	byteArray, err := json.Marshal(book)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(byteArray))
}

func unmarshal() {
	jsonString := `{"name": "battery sensor", "capacity": 40, "time": "2019-01-21T19:07:28Z"}`
	var sensorReading SensorReading
	err := json.Unmarshal([]byte(jsonString), &sensorReading)
	if err != nil {
		fmt.Println("Error while marshling")
		return
	}
	fmt.Println(sensorReading)
}

func unmarshalUnstructred() {
	jsonString := `{"name": "battery sensor", "capacity": 40, "time": "2019-01-21T19:07:28Z"}`
	var reading map[string]interface{}
	err := json.Unmarshal([]byte(jsonString), &reading)
	if err != nil {
		fmt.Println("Error while marshling")
		return
	}
	fmt.Println(reading)
}
