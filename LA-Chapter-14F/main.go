package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Product struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func main() {

	// call function streamEncodingJson
	streamEncodingJson()
	// call function streamDecodeJson
	streamDecodeJson()
}

func streamEncodingJson() {
	// Buka file untuk menulis
	file, err := os.Create("output.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Buat encoder baru
	encoder := json.NewEncoder(file)
	// Data yang akan diencode
	products := []Product{
		{Name: "Laptop", Price: 1500},
		{Name: "Smartphone", Price: 800},
		{Name: "Tablet", Price: 400},
	}

	// Encode dan tulis setiap produk ke file
	for _, product := range products {
		if err := encoder.Encode(&product); err != nil {
			fmt.Println("Error encoding JSON:", err)
			return
		}
	}

}

func streamDecodeJson() {
	// Buka file JSON
	file, err := os.Open("products.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Buat decoder baru
	decoder := json.NewDecoder(file)
	// Decode array JSON
	var products []Product
	if err := decoder.Decode(&products); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	// Tampilkan produk yang telah di-decode
	for _, product := range products {
		fmt.Printf("Product: %+v\n", product)
	}
}
