package main

import (
	"PhonesReviewAPI/config"
	"PhonesReviewAPI/models"

	"github.com/julienschmidt/httprouter"

	_ "context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	db, e := config.MySQL()

	if e != nil {
		log.Fatal(e)
	}

	eb := db.Ping()
	if eb != nil {
		panic(eb.Error())
	}

	fmt.Println("Success")

	router := httprouter.New()
	router.GET("/brands", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		readBrands(w, r, p, db)
	})

	router.PUT("/brands/:brandID", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		updateBrand(w, r, p, db)
	})

	router.POST("/brands", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		createBrand(w, r, p, db)
	})

	router.DELETE("/brands/:brandID", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		deleteBrand(w, r, p, db)
	})

	fmt.Println("Server Running at Port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func readBrands(w http.ResponseWriter, r *http.Request, _ httprouter.Params, db *sql.DB) {
	rows, err := db.Query("SELECT brand_id, brand_name FROM brands")
	if err != nil {
		log.Fatal(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var brandList []models.Brand
	for rows.Next() {
		var brand models.Brand
		if err := rows.Scan(&brand.ID, &brand.Name); err != nil {
			log.Fatal(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		brandList = append(brandList, brand)
	}

	// Mengembalikan data dalam format JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(brandList); err != nil {
		log.Fatal(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func updateBrand(w http.ResponseWriter, r *http.Request, p httprouter.Params, db *sql.DB) {
	brandID := p.ByName("brandID")

	// Membaca data yang ingin diupdate dari request
	var updatedBrand models.Brand
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&updatedBrand); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Melakukan update data merek ponsel
	_, err := db.Exec("UPDATE brands SET brand_name = ? WHERE brand_id = ?", updatedBrand.Name, brandID)
	if err != nil {
		log.Fatal(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Mengembalikan respons sukses
	w.WriteHeader(http.StatusNoContent)
}

func createBrand(w http.ResponseWriter, r *http.Request, p httprouter.Params, db *sql.DB) {
	// Membaca data yang ingin ditambahkan dari request
	var newBrand models.Brand
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newBrand); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Menyimpan data merek ponsel baru ke database
	result, err := db.Exec("INSERT INTO brands (brand_name) VALUES (?)", newBrand.Name)
	if err != nil {
		log.Fatal(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Mengambil ID merek yang baru ditambahkan
	brandID, _ := result.LastInsertId()

	// Mengembalikan ID merek yang baru ditambahkan dalam format JSON
	w.Header().Set("Content-Type", "application/json")
	response := map[string]int{"brandID": int(brandID)}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Fatal(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func deleteBrand(w http.ResponseWriter, r *http.Request, p httprouter.Params, db *sql.DB) {
	brandID := p.ByName("brandID")

	// Menghapus merek ponsel berdasarkan brandID
	_, err := db.Exec("DELETE FROM brands WHERE brand_id = ?", brandID)
	if err != nil {
		log.Fatal(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Mengembalikan respons sukses
	w.WriteHeader(http.StatusNoContent)
}
