package main

import (
	"encoding/json"
	"fmt"
)

// definisikan struktur data untuk person
type Person struct {
	Name string
	Age  int
}

// Definisikan struktur data untuk Produk
type Product struct {
	Nama  string
	Merk  string
	Harga int
}

// Definisikan struktur data untuk Objek JSON utama
type Catalog struct {
	ID        int
	Nama      string
	Deskripsi string
	Produk    []Product
}

func main() {
	jsonData1 := []byte(`{"name":"Jerry","age":25}`)
	var p Person
	err1 := json.Unmarshal(jsonData1, &p)
	if err1 != nil {
		fmt.Println("Error unmarshalling JSON:", err1)
		return
	}
	fmt.Println("Name:", p.Name)
	fmt.Println("Age:", p.Age)

	// Data JSON yang akan di-decode
	jsonData2 := []byte(`{
			"id": 1,
			"nama": "Produk Elektronik",
			"deskripsi": "Daftar produk elektronik terbaru",
			"produk": [
            {"nama": "Smartphone", "merk": "Samsung", "harga": 3000000},
            {"nama": "Laptop", "merk": "Apple", "harga": 15000000},
            {"nama": "Smartwatch", "merk": "Fitbit", "harga": 2000000}
        ]
    }`)
	// Variabel untuk menampung hasil decoding
	var catalog Catalog

	// Decoding (unmarshalling) JSON ke dalam objek JSON utama (Catalog)
	err2 := json.Unmarshal(jsonData2, &catalog)
	if err2 != nil {
		fmt.Println("Error unmarshalling JSON:", err2)
		return
	}

	// Tampilkan data yang telah di-decode
	fmt.Println("ID:", catalog.ID)
	fmt.Println("Nama:", catalog.Nama)
	fmt.Println("Deskripsi:", catalog.Deskripsi)
	fmt.Println("Produk:")

	for _, product := range catalog.Produk {
		fmt.Printf("  - Nama: %s, Merk: %s, Harga: %d\n", product.Nama, product.Merk, product.Harga)
	}
}
