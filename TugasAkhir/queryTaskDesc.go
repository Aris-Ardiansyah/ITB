//Funsi-fungsi CRUD untuk tabel TaskDesc
package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func insertTaskDesc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	u := r.URL.Query()

	db, err := ConnDB()
	if err != nil {
		HandleError505(w, err)
		return
	}
	defer db.Close()

	tsql := fmt.Sprintf("INSERT INTO TaskDesc (TaskID, Author, Judul, Deskripsi, Jumlah, isDelegated) VALUES (@taskid, @author, @judul, @deskripsi, @jumlah, 0);")

	result, err := db.Exec(tsql,
		sql.Named("taskid", angka(u["taskid"][0])),
		sql.Named("author", angka(u["author"][0])),
		sql.Named("judul", u["judul"][0]),
		sql.Named("deskripsi", u["deskripsi"][0]),
		sql.Named("jumlah", u["jumlah"][0]))

	if err != nil {
		HandleError505(w, err)
		return
	}

	id, _ := result.LastInsertId()
	msgR := msgResponse{Status: "Success", Message: fmt.Sprintf("%d", id)}

	json.NewEncoder(w).Encode(msgR)
}

func listTaskDesc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db, err := ConnDB()
	if err != nil {
		HandleError505(w, err)
		return
	}
	defer db.Close()

	perintah := fmt.Sprintf("SELECT * FROM TaskDesc WHERE %s = @nilai", mux.Vars(r)["field"])
	//nilaiku := angka(r.URL.Query()["nilai"][0])
	//fmt.Println(angka(r.URL.Query()["nilai"][0]))
	rows, err := db.Query(perintah, sql.Named("nilai", angka(r.URL.Query()["nilai"][0])))

	if err != nil {
		HandleError505(w, err)
		return
	}
	defer rows.Close()

	var tabel []TaskDescRow
	for rows.Next() {
		var baris = TaskDescRow{}
		err := rows.Scan(
			&baris.ID,
			&baris.TaskID,
			&baris.Author,
			&baris.Judul,
			&baris.Deskripsi,
			&baris.Jumlah,
			&baris.IsDelegated)

		if err != nil {
			HandleError505(w, err)
			return
		}

		tabel = append(tabel, baris)
	}
	json.NewEncoder(w).Encode(tabel)

}

func updateTaskDesc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	u := r.URL.Query()

	var msgR msgResponse

	db, err := ConnDB()
	if err != nil {
		HandleError505(w, err)
		return
	}
	defer db.Close()

	tsql := fmt.Sprintf("UPDATE TaskDesc SET TaskID=@taskid, Author=@author, Judul=@judul, Deskripsi=@deskripsi, Jumlah=@jumlah WHERE id=@id;")

	result, err := db.Exec(tsql,
		sql.Named("taskid", angka(u["taskid"][0])),
		sql.Named("author", angka(u["author"][0])),
		sql.Named("judul", u["judul"][0]),
		sql.Named("deskripsi", u["deskripsi"][0]),
		sql.Named("jumlah", u["jumlah"][0]),
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

func deleteTaskDesc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var msgR msgResponse

	db, err := ConnDB()
	if err != nil {
		HandleError505(w, err)
		return
	}
	defer db.Close()

	id := mux.Vars(r)["id"]

	tsql := fmt.Sprintf("DELETE FROM TaskDesc WHERE ID=@id;")

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

func taskDelegatebySeksi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	type Baris struct {
		Pegawai string `json:"NIP"`
		Nama    string `json:"Nama"`
		Jumlah  int    `json:"Jumlah"`
	}

	db, err := ConnDB()
	if err != nil {
		HandleError505(w, err)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT Seksi.ID, Seksi.Deskripsi, Count(taskdesc.Author) as Jumlah FROM Seksi, TaskDesc WHERE Seksi.ID=TaskDesc.Author GROUP by Seksi.ID, Seksi.Deskripsi")

	if err != nil {
		HandleError505(w, err)
		return
	}
	defer rows.Close()

	var tabel []Baris
	for rows.Next() {
		var baris = Baris{}
		err := rows.Scan(
			&baris.Pegawai,
			&baris.Nama,
			&baris.Jumlah,
		)

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
