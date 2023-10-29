package models

import (
	_ "github.com/go-sql-driver/mysql"
)

// Struktur data untuk merek ponsel
type Brand struct {
	ID   int
	Name string
}

// Struktur data untuk informasi ponsel
type Phone struct {
	ID       int
	BrandID  int
	Model    string
	Category string
}

// Struktur data untuk pelanggan
type Customer struct {
	ID   int
	Name string
}

// Struktur data untuk ulasan pelanggan
type Review struct {
	ID      int
	CustID  int
	PhoneID int
	Text    string
	Rating  float64
}
