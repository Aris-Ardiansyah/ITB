package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gorilla/mux"
)

func main() {
	//log.Println(TulisFile("notif1", "pesan dari notif 1"))
	paramKoneksi = ParseJSONtoMapInterface("koneksi.json") //isi parameter dari file JSON

	router := mux.NewRouter()

	//tabel task
	router.HandleFunc("/task", listTask).Methods("GET")
	router.HandleFunc("/task/{seksi}", listTaskSeksi).Methods("GET")
	router.HandleFunc("/task", insertTask).Methods("POST")
	router.HandleFunc("/task/{id}", updateTask).Methods("PUT")
	router.HandleFunc("/task/{id}", deleteTask).Methods("DELETE")

	//tabel taskdesc
	router.HandleFunc("/taskdesc/{field}", listTaskDesc).Methods("GET")
	router.HandleFunc("/taskdesc", insertTaskDesc).Methods("POST")
	router.HandleFunc("/taskdesc/{id}", updateTaskDesc).Methods("PUT")
	router.HandleFunc("/taskdesc/{id}", deleteTaskDesc).Methods("DELETE")

	//tabel tasklist
	router.HandleFunc("/tasklist", listAllTaskList).Methods("GET")
	router.HandleFunc("/tasklist/{field}", listTaskList).Methods("GET")
	router.HandleFunc("/tasklist", insertTaskList).Methods("POST")
	router.HandleFunc("/tasklist/{id}", updateTaskList).Methods("PUT")
	router.HandleFunc("/tasklist/{id}", deleteTaskList).Methods("DELETE")

	router.HandleFunc("/bebankerja", listBebanKerja).Methods("GET")
	router.HandleFunc("/bebankerja/{pegawai}", listBebanKerjaPegawai).Methods("GET")

	router.HandleFunc("/pekerjaanbyseksi", taskDelegatebySeksi).Methods("GET")

	//tabel pegawai
	router.HandleFunc("/pegawai", listPegawai).Methods("GET")
	router.HandleFunc("/pegawai/{seksi}", listPegawaiSeksi).Methods("GET")
	router.HandleFunc("/pegawai", insertPegawai).Methods("POST")
	router.HandleFunc("/pegawai/{nip}", updatePegawai).Methods("PUT")
	router.HandleFunc("/pegawai/{nip}", deletePegawai).Methods("DELETE")

	//tabel kegiatan
	router.HandleFunc("/listkegiatan", listKegiatan).Methods("GET")

	//tabel lainnya
	router.HandleFunc("/list/{tabel}", listTabel).Methods("GET")

	//log.Printf("Server starting on port %.f\n", paramKoneksi["WebPort"].(float64))
	//log.Fatal(http.ListenAndServe(fmt.Sprintf(":%.f", paramKoneksi["WebPort"].(float64)), router))
	log.Printf("Server starting on port %s\n", paramKoneksi["WebPort"])
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", paramKoneksi["WebPort"]), router))
}
