//Funsi-fungsi CRUD untuk tabel Task
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func listTabel(w http.ResponseWriter, r *http.Request) {

	var nilaiBalik struct {
		id        int
		deskripsi string
	}

	db, err := ConnDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Panic(err)
	}
	defer db.Close()

	rows, err := db.Query(fmt.Sprintf("SELECT * FROM %s", mux.Vars(r)["tabel"]))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&nilaiBalik.id,
			&nilaiBalik.deskripsi,
		)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Panic(err)
		}

		json.NewEncoder(w).Encode(fmt.Sprint(nilaiBalik))
	}

	err = rows.Err()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Panic(err)
	}

}

func listKegiatan(w http.ResponseWriter, r *http.Request) {

	var nilaiBalik struct {
		id        int
		deskripsi string
	}

	db, err := ConnDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, deskripsi FROM Kegiatan")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&nilaiBalik.id,
			&nilaiBalik.deskripsi,
		)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Panic(err)
		}

		json.NewEncoder(w).Encode(fmt.Sprint(nilaiBalik))
	}

	err = rows.Err()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Panic(err)
	}

}
