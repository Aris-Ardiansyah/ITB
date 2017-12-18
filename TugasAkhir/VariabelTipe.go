package main

//map buat nyimpan parameter koneksi
//dibuat global di package main agar tiap pembukaan koneksi ke database tidak perlu parse file koneksi.json lagi
var paramKoneksi map[string]interface{}

//Struct user
type userData struct {
	ID           int
	NamaDepan    string
	NamaBelakang string
	Alamat       string
}

type msgResponse struct {
	Status  string
	Message string
}

//TaskRow - struct buat nyimpan hasil dari tabel Task
type TaskRow struct {
	ID                int    `json:"ID"`
	Kegiatan          int    `json:"Kegiatan"`
	DeskripsiKegiatan string `json:"DeskKegiatan"`
	Jumlah            int    `json:"Jumlah"`
	Mulai             string `json:"Mulai"`
	Selesai           string `json:"Selesai"`
	Seksi             int    `json:"Seksi"`
	Deskripsi         string `json:"Deskripsi"`
}

//TaskDescRow - struct buat nyimpan hasil dari tabel TaskDesc
type TaskDescRow struct {
	ID          int    `json:"ID"`
	TaskID      int    `json:"TaskID"`
	Author      int    `json:"Author"`
	Judul       string `json:"Judul"`
	Deskripsi   string `json:"Deskripsi"`
	Jumlah      int    `json:"Jumlah"`
	IsDelegated int    `json:"IsDelegated"`
}

//TaskListRow - struct buat nyimpan hasil dari tabel TaskList
type TaskListRow struct {
	ID      int    `json:"ID"`
	TaskID  int    `json:"TaskID"`
	Pegawai string `json:"Pegawai"`
	Parent  int    `json:"Parent"`
	Mulai   string `json:"Mulai"`
	Selesai string `json:"Selesai"`
	Status  int    `json:"Status"`
}

//PegawaiRow - struct buat nyimpan hasil dari tabel pegawai
type PegawaiRow struct {
	NIP     string `json:"NIP"`
	Nama    string `json:"Nama"`
	Jabatan int    `json:"Jabatan"`
	Seksi   int    `json:"Seksi"`
	Telepon string `json:"Telepon"`
	Email   string `json:"Email"`
}
