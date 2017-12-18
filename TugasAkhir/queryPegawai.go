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

func insertPegawai(w http.ResponseWriter, r *http.Request) {
	u := r.URL.Query()

	db, err := ConnDB()
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	tsql := fmt.Sprintf("INSERT INTO Pegawai (NIP, Nama, Jabatan, Seksi, Telepon, Email) VALUES (@nip, @nama, @jabatan, @seksi, @telepon, @email);")

	result, err := db.Exec(tsql,
		sql.Named("nip", u["nip"][0]),
		sql.Named("nama", u["nama"][0]),
		sql.Named("jabatan", angka(u["jabatan"][0])),
		sql.Named("seksi", angka(u["seksi"][0])),
		sql.Named("telepon", u["telepon"][0]),
		sql.Named("email", u["email"][0]))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	id, _ := result.LastInsertId()
	msgR := msgResponse{Status: "Success", Message: fmt.Sprintf("%d", id)}

	json.NewEncoder(w).Encode(msgR)
}

func listPegawai(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db, err := ConnDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM Pegawai")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Panic(err)
	}
	defer rows.Close()

	var tabel []PegawaiRow
	for rows.Next() {
		var baris = PegawaiRow{}
		err := rows.Scan(
			&baris.NIP,
			&baris.Nama,
			&baris.Jabatan,
			&baris.Seksi,
			&baris.Telepon,
			&baris.Email)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Panic(err)
		}

		tabel = append(tabel, baris)
	}
	json.NewEncoder(w).Encode(tabel)

	err = rows.Err()
	if err != nil {
		log.Panic(err)
	}

}

func listPegawaiSeksi(w http.ResponseWriter, r *http.Request) {

	db, err := ConnDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM pegawai WHERE seksi = @seksi", sql.Named("seksi", angka(mux.Vars(r)["seksi"])))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Panic(err)
	}
	defer rows.Close()

	var tabel []PegawaiRow
	for rows.Next() {
		var baris = PegawaiRow{}
		err := rows.Scan(
			&baris.NIP,
			&baris.Nama,
			&baris.Jabatan,
			&baris.Seksi,
			&baris.Telepon,
			&baris.Email)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Panic(err)
		}

		tabel = append(tabel, baris)
	}
	json.NewEncoder(w).Encode(tabel)

	err = rows.Err()
	if err != nil {
		log.Panic(err)
	}

}

func updatePegawai(w http.ResponseWriter, r *http.Request) {
	u := r.URL.Query()

	var msgR msgResponse

	db, err := ConnDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Panic(err)
	}
	defer db.Close()

	tsql := fmt.Sprintf("UPDATE Pegawai SET Nama=@nama, Jabatan=@jabatan, Seksi=@seksi, Telepon=@telepon, Email=@email WHERE NIP=@nip;")

	result, err := db.Exec(tsql,
		sql.Named("nama", u["nama"][0]),
		sql.Named("jabatan", angka(u["jabatan"][0])),
		sql.Named("seksi", angka(u["seksimulai"][0])),
		sql.Named("telepon", u["telepon"][0]),
		sql.Named("email", u["email"][0]),
		sql.Named("nip", mux.Vars(r)["id"]))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Panic(err)
	} else {
		jumlahBaris, _ := result.RowsAffected()
		msgR = msgResponse{Status: "Updated Successfully", Message: fmt.Sprintf("%d row", jumlahBaris)}
	}

	json.NewEncoder(w).Encode(msgR)
}

func deletePegawai(w http.ResponseWriter, r *http.Request) {

	var msgR msgResponse

	db, err := ConnDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Panic(err)
	}
	defer db.Close()

	id := mux.Vars(r)["id"]

	//sebelum menghapus, cek dulu di tasklist dan tabel lain yang masih pake NIP yang mau dihapus
	result, err := db.Exec("DELETE FROM Pegawai WHERE NIP=@nip;", sql.Named("nip", id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Panic(err)
	}

	//Hitung baris kehapus
	rowCount, err := result.RowsAffected()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Panic(err)
	}

	//encode to json format and send status as response
	if rowCount > 0 {
		msgR = msgResponse{Status: "Success", Message: fmt.Sprintf("Berhasil menghapus NIP = %s, dengan jumlah = %d", id, rowCount)}
	} else {
		msgR = msgResponse{Status: "Error", Message: "Data tidak ditemukan dengan ID = " + id}
	}

	json.NewEncoder(w).Encode(msgR)
}
