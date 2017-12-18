//Funsi-fungsi CRUD untuk tabel TaskList
package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func insertTaskList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	u := r.URL.Query()

	db, err := ConnDB()
	if err != nil {
		HandleError505(w, err)
		return
	}

	//cek dulu apakah tugas telah didelegate
	tsql := fmt.Sprintf("INSERT INTO TaskList (TaskID, Pegawai, Parent, Mulai, Selesai, Status) VALUES (@taskid, @pegawai, @parent, @mulai, @selesai, @status);")

	result, err := db.Exec(tsql,
		sql.Named("taskid", angka(u["taskid"][0])),
		sql.Named("pegawai", u["pegawai"][0]),
		sql.Named("parent", angka(u["parent"][0])),
		sql.Named("mulai", waktu(u["mulai"][0])),
		sql.Named("selesai", waktu(u["selesai"][0])),
		sql.Named("status", angka(u["status"][0])))

	if err != nil {
		HandleError505(w, err)
		return
	}
	id, _ := result.LastInsertId()

	//update taskdesc field isdelegated=1 berarti sudah didelegasikan
	result, err = db.Exec("UPDATE TaskDesc SET isDelegated=1 WHERE id=@id;", sql.Named("id", angka(u["taskid"][0])))

	if err != nil {
		HandleError505(w, err)
		return
	}

	//select dari database identitas sesuai dengan nip ini
	var Nama, Telepon, Email string

	err = db.QueryRow(fmt.Sprintf("SELECT Nama, Telepon, Email FROM pegawai WHERE NIP = %s", u["pegawai"][0])).Scan(&Nama, &Telepon, &Email)

	if err != nil {
		HandleError505(w, err)
		return
	}

	//buat file notif
	err = TulisFile(u["pegawai"][0], fmt.Sprintf("%s anda mendapat delegasi tugas baru # %s # %s", Nama, Telepon, Email))

	//jumlahBaris, _ := result.RowsAffected()
	msgR := msgResponse{Status: "Success", Message: fmt.Sprintf("%d", id)}

	json.NewEncoder(w).Encode(msgR)
}

func listAllTaskList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db, err := ConnDB()
	if err != nil {
		HandleError505(w, err)
		return
	}

	rows, err := db.Query("SELECT * FROM TaskList")

	if err != nil {
		HandleError505(w, err)
		return
	}
	defer rows.Close()

	var tabel []TaskListRow
	for rows.Next() {
		var baris = TaskListRow{}
		err := rows.Scan(
			&baris.ID,
			&baris.TaskID,
			&baris.Pegawai,
			&baris.Parent,
			&baris.Mulai,
			&baris.Selesai,
			&baris.Status)

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

func listTaskList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db, err := ConnDB()
	if err != nil {
		HandleError505(w, err)
		return
	}

	perintah := fmt.Sprintf("SELECT * FROM TaskList WHERE %s = @nilai", mux.Vars(r)["field"])
	rows, err := db.Query(perintah, sql.Named("nilai", angka(r.URL.Query()["nilai"][0])))

	if err != nil {
		HandleError505(w, err)
		return
	}
	defer rows.Close()

	var tabel []TaskListRow
	for rows.Next() {
		var baris = TaskListRow{}
		err := rows.Scan(
			&baris.ID,
			&baris.TaskID,
			&baris.Pegawai,
			&baris.Parent,
			&baris.Mulai,
			&baris.Selesai,
			&baris.Status)

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

func updateTaskList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	u := r.URL.Query()

	var msgR msgResponse

	db, err := ConnDB()
	if err != nil {
		HandleError505(w, err)
		return
	}

	tsql := fmt.Sprintf("UPDATE TaskList SET taskid=@taskid, pegawai=@pegawai, parent=@parent, Mulai=@mulai, Selesai=@selesai, status=@status WHERE id=@id;")

	result, err := db.Exec(tsql,
		sql.Named("taskid", angka(u["taskid"][0])),
		sql.Named("pegawai", u["pegawai"][0]),
		sql.Named("parent", angka(u["parent"][0])),
		sql.Named("mulai", waktu(u["mulai"][0])),
		sql.Named("selesai", waktu(u["selesai"][0])),
		sql.Named("status", angka(u["status"][0])),
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

func deleteTaskList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var msgR msgResponse

	db, err := ConnDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Panic(err)
	}

	id := mux.Vars(r)["id"]
	//select taskid from tasklist where id=@id -> query row simpan ke variabel deletedtaskid

	var DeletedTaskID string

	err = db.QueryRow(fmt.Sprintf("SELECT TaskID FROM Tasklist WHERE id = %s", id)).Scan(&DeletedTaskID)

	if err != nil {
		HandleError505(w, err)
		return
	}

	//update taskdesc field isdelegated=0 berarti belum didelegasikan
	result, _ := db.Exec("UPDATE TaskDesc SET isDelegated=0 WHERE id=@id;", sql.Named("id", DeletedTaskID))

	//if err != nil {
	//	HandleError505(w, err)
	//	return
	//}

	tsql := fmt.Sprintf("DELETE FROM TaskList WHERE ID=@id;")

	result, err = db.Exec(tsql, sql.Named("id", id))
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
		msgR = msgResponse{Status: "Deleted Successfully", Message: fmt.Sprintf("Berhasil menghapus ID = %s, dengan jumlah = %d", id, rowCount)}
		json.NewEncoder(w).Encode(msgR)
		return
	}

	http.Error(w, fmt.Sprintf("Data dengan id = %s tidak ditemukan", id), 404)

}

func listBebanKerja(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	type Baris struct {
		Pegawai string `json:"NIP"`
		Nama    string `json:"Nama"`
		Jumlah  int    `json:"Jumlah"`
	}

	db, err := ConnDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT TaskList.Pegawai, pegawai.Nama, count(TaskList.Pegawai) as jumlah	FROM dbo.TaskList ,dbo.Pegawai	WHERE TaskList.Pegawai=Pegawai.NIP 	GROUP BY TaskList.Pegawai, Pegawai.Nama  ")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Panic(err)
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
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Panic(err)
		}
		tabel = append(tabel, baris)
	}
	json.NewEncoder(w).Encode(tabel)

	err = rows.Err()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Panic(err)
	}
}

// list pekerjaan apa saja yang didelegasikan ke pegawai tertentu
//queri belum diperbaiki
func listBebanKerjaPegawai(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db, err := ConnDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM TaskList WHERE Pegawai=@pegawai", sql.Named("pegawai", angka(mux.Vars(r)["pegawai"])))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Panic(err)
	}
	defer rows.Close()

	var tabel []TaskListRow
	for rows.Next() {
		var baris = TaskListRow{}
		err := rows.Scan(
			&baris.ID,
			&baris.TaskID,
			&baris.Pegawai,
			&baris.Parent,
			&baris.Mulai,
			&baris.Selesai,
			&baris.Status)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Panic(err)
		}
		tabel = append(tabel, baris)
	}
	json.NewEncoder(w).Encode(tabel)

	err = rows.Err()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Panic(err)
	}
}
