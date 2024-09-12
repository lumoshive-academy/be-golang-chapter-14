package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

// Definisikan struktur data untuk Produk
type Product struct {
	Nama  string `json:"nama"`
	Merk  string `json:"merk"`
	Harga int    `json:"harga"`
}

// Definisikan struktur data untuk Objek JSON utama
type Catalog struct {
	ID        int       `json:"id"`
	Nama      string    `json:"nama"`
	Deskripsi string    `json:"deskripsi"`
	Produk    []Product `json:"produk"`
}

func main() {
	// encoding basic
	p := Person{Name: "Febry", Age: 30}
	jsonData1, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}
	fmt.Println(string(jsonData1))

	// Data untuk Objek JSON utama (Catalog)
	catalog := Catalog{
		ID:        1,
		Nama:      "Produk Elektronik",
		Deskripsi: "Daftar produk elektronik terbaru",
		Produk: []Product{
			{Nama: "Smartphone", Merk: "Samsung", Harga: 3000000},
			{Nama: "Laptop", Merk: "Apple", Harga: 15000000},
			{Nama: "Smartwatch", Merk: "Fitbit", Harga: 2000000},
		},
	}

	// Encoding (marshalling) objek JSON utama ke dalam format JSON
	jsonData2, err := json.MarshalIndent(catalog, "", " ")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	// Tampilkan JSON yang dihasilkan
	fmt.Println(string(jsonData2))

}
