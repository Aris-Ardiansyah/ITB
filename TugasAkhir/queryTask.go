//Funsi-fungsi CRUD untuk tabel Task
package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func insertTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	u := r.URL.Query()

	db, err := ConnDB()
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	tsql := fmt.Sprintf("INSERT INTO Task (Kegiatan, Jumlah, Mulai, Selesai, Seksi, Deskripsi) VALUES (@kegiatan, @jumlah, @mulai, @selesai, @seksi, @deskripsi);")

	result, err := db.Exec(tsql,
		sql.Named("kegiatan", angka(u["kegiatan"][0])),
		sql.Named("jumlah", angka(u["jumlah"][0])),
		sql.Named("mulai", waktu(u["mulai"][0])),
		sql.Named("selesai", waktu(u["selesai"][0])),
		sql.Named("seksi", angka(u["seksi"][0])),
		sql.Named("deskripsi", u["deskripsi"][0]))

	if err != nil {
		HandleError505(w, err)
		return
	}
	id, _ := result.LastInsertId()
	msgR := msgResponse{Status: "Success", Message: fmt.Sprintf("%d", id)}

	json.NewEncoder(w).Encode(msgR)
}

func listTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db, err := ConnDB()
	if err != nil {
		HandleError505(w, err)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT Task.ID, Kegiatan, Kegiatan.Deskripsi as DeskKegiatan, Jumlah, Mulai, Selesai, Seksi, Task.Deskripsi FROM Task, Kegiatan WHERE Task.Kegiatan=Kegiatan.ID ")

	if err != nil {
		HandleError505(w, err)
		return
	}
	defer rows.Close()

	var tabel []TaskRow
	for rows.Next() {
		var baris = TaskRow{}
		err := rows.Scan(
			&baris.ID,
			&baris.Kegiatan,
			&baris.DeskripsiKegiatan,
			&baris.Jumlah,
			&baris.Mulai,
			&baris.Selesai,
			&baris.Seksi,
			&baris.Deskripsi)

		if err != nil {
			HandleError505(w, err)
			return
		}

		tabel = append(tabel, baris)
	}
	json.NewEncoder(w).Encode(tabel)

	err = rows.Err()
	if err != nil {
		HandleError505(w, err)
		return
	}

}

func listTaskSeksi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db, err := ConnDB()
	if err != nil {
		HandleError505(w, err)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM Task WHERE seksi = @seksi", sql.Named("seksi", angka(mux.Vars(r)["seksi"])))

	if err != nil {
		HandleError505(w, err)
		return
	}
	defer rows.Close()

	var tabel []TaskRow
	for rows.Next() {
		var baris = TaskRow{}
		err := rows.Scan(
			&baris.ID,
			&baris.Kegiatan,
			&baris.Jumlah,
			&baris.Mulai,
			&baris.Selesai,
			&baris.Seksi,
			&baris.Deskripsi)

		if err != nil {
			HandleError505(w, err)
			return
		}

		tabel = append(tabel, baris)

	}
	json.NewEncoder(w).Encode(tabel)

	err = rows.Err()
	if err != nil {
		HandleError505(w, err)
		return
	}

}

func updateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	u := r.URL.Query()

	var msgR msgResponse

	db, err := ConnDB()
	if err != nil {
		HandleError505(w, err)
		return
	}
	defer db.Close()

	tsql := fmt.Sprintf("UPDATE Task SET Kegiatan=@kegiatan, Jumlah=@jumlah, Mulai=@mulai, Selesai=@selesai, Seksi=@seksi, Deskripsi=@deskripsi WHERE id=@id;")

	result, err := db.Exec(tsql,
		sql.Named("kegiatan", angka(u["kegiatan"][0])),
		sql.Named("jumlah", angka(u["jumlah"][0])),
		sql.Named("mulai", waktu(u["mulai"][0])),
		sql.Named("selesai", waktu(u["selesai"][0])),
		sql.Named("seksi", angka(u["seksi"][0])),
		sql.Named("deskripsi", u["deskripsi"][0]),
		sql.Named("id", angka(mux.Vars(r)["id"])))

	if err != nil {
		HandleError505(w, err)
		return
	}

	jumlahBaris, _ := result.RowsAffected()

	if jumlahBaris > 0 {
		msgR = msgResponse{Status: "Updated Successfully", Message: fmt.Sprintf("%d row", jumlahBaris)}

		json.NewEncoder(w).Encode(msgR)
	} else {
		http.Error(w, fmt.Sprintf("Data dengan id = %s tidak ditemukan", mux.Vars(r)["id"]), 404)
	}

}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var msgR msgResponse

	db, err := ConnDB()
	if err != nil {
		HandleError505(w, err)
		return
	}
	defer db.Close()

	id := mux.Vars(r)["id"]

	tsql := fmt.Sprintf("DELETE FROM Task WHERE ID=@id;")

	result, err := db.Exec(tsql, sql.Named("id", id))
	if err != nil {
		HandleError505(w, err)
		return
	}

	//Hitung baris kehapus
	rowCount, err := result.RowsAffected()
	if err != nil {
		HandleError505(w, err)
		return
	}

	//encode to json format and send status as response
	if rowCount > 0 {
		msgR = msgResponse{Status: "Deleted Successfully", Message: fmt.Sprintf("Berhasil menghapus ID = %s, dengan jumlah = %d", id, rowCount)}
		json.NewEncoder(w).Encode(msgR)
		return
	}

	http.Error(w, fmt.Sprintf("Data dengan id = %s tidak ditemukan", id), 404)
}
