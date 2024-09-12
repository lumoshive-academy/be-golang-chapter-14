package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	// call function encodingJsonMap
	encodingJsonMap()
	// call function decodingJsonMap
	decodingJsonMap()
}

func encodingJsonMap() {
	// Definisikan map dengan data JSON
	data := map[string]interface{}{
		"name": "John Doe",
		"age":  30,
		"address": map[string]interface{}{
			"street": "123 Main St",
			"city":   "Anytown",
		},
		"hobbies": []string{"reading", "traveling", "swimming"},
	}

	// Encoding (marshalling) map ke JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	// Tampilkan JSON yang dihasilkan
	fmt.Println(string(jsonData))
}

func decodingJsonMap() {
	// JSON data yang akan di-decode
	jsonStr := `{
			"name": "Gigi",
			"age": 25,
			"address": {
				"street": "Kedoya",
				"city": "Jakbar"
			},
			"hobbies": ["Renang", "Nonton"]
		}`
	// Variabel untuk menampung hasil decoding
	var data map[string]interface{}

	// Decoding (unmarshalling) JSON ke map
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	// Tampilkan data yang telah di-decode
	fmt.Println("Name:", data["name"])
	fmt.Println("Age:", data["age"])
	fmt.Println("Address:", data["address"].(map[string]interface{}))
	fmt.Println("Hobbies:", data["hobbies"].([]interface{}))
}
