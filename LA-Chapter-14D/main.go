package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Definisikan struktur data untuk sebuah buku
type Book struct {
	Title  string `json:"title"`
	Author string `json:"author,omitempty"` // omitempty akan menghilangkan field jika kosong
	ISBN   string `json:"isbn,omitempty"`   // omitempty akan menghilangkan field jika kosong
	Pages  int    `json:"-"`                // '-' akan mengabaikan field ini dalam proses encoding dan decoding
}

func main() {
	// call function encodeJson
	encodeJson()
	// call function decodeJson
	decodeJson()
	// call function ignoreTagJson
	ignoreTagJson()
}

// encoding json with json tag
func encodeJson() {
	p := Person{Name: "John", Age: 30}
	jsonData, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}
	fmt.Println(string(jsonData))

}

// decoding json with json tag
func decodeJson() {
	jsonStr := `{"name":"Jane Smith","age":25}`
	var p Person
	err := json.Unmarshal([]byte(jsonStr), &p)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}
	fmt.Println("Name:", p.Name)
	fmt.Println("Age:", p.Age)
}

// implementation ingore tag json
func ignoreTagJson() {
	// Contoh data buku dengan beberapa field kosong
	book := Book{
		Title: "Golang Programming",
		// Author dan ISBN kosong
		Pages: 256,
	}

	// Encoding (marshalling) struct Book ke JSON
	jsonData, err := json.Marshal(book)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}
	fmt.Println("Encoded JSON:", string(jsonData))
	// Decoding (unmarshalling) JSON ke struct Book
	jsonStr := `{"title":"Golang Programming","isbn":"978-3-16-148410-0"}`
	var decodedBook Book
	err = json.Unmarshal([]byte(jsonStr), &decodedBook)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}
	fmt.Printf("Decoded Book:\nTitle: %s\nAuthor: %s\nISBN: %s\nPages: %d\n",
		decodedBook.Title, decodedBook.Author, decodedBook.ISBN, decodedBook.Pages)

}
