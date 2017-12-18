package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

//ParseJSONtoMapInterface - membaca JSON dan mengubah ke map[string]interface{}
func ParseJSONtoMapInterface(file string) map[string]interface{} {
	var data map[string]interface{}

	if raw, err := ioutil.ReadFile(file); err != nil {
		log.Fatalln(err.Error())
	} else {
		_ = json.Unmarshal(raw, &data)
	}
	return data
}

//ConnDB - fungsi untuk membuka koneksi ke database
func ConnDB() (*sql.DB, error) {

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s;",
		paramKoneksi["Server"],
		paramKoneksi["User"],
		paramKoneksi["Password"],
		paramKoneksi["Port"],
		paramKoneksi["Database"])

	db, err := sql.Open("sqlserver", connString)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func angka(txt string) int {
	/*nilai, err := strconv.Atoi(txt)
	if err != nil {
		return 0
	}
	return nilai
	*/
	nilai, _ := strconv.Atoi(txt)

	return nilai
}

func waktu(txt string) time.Time {
	nilai, _ := time.Parse("2006-01-02", txt)

	return nilai
}

//catch - fungsi untuk handle panic, defer di fungsi yang akan memanggil panic
func catch() error {
	if r := recover(); r != nil {
		return errors.New(fmt.Sprint(r))
	}
	return nil
}

//HandleError505 fungsi handle error 500 - StatusInternalServerError
func HandleError505(w http.ResponseWriter, err error) {
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	log.Println(err)
}

//TulisFile fungsi untuk membuat file notif
func TulisFile(namafile string, pesan string) error {
	err := ioutil.WriteFile(fmt.Sprintf("%s.txt", namafile), []byte(pesan), 0644)
	return err
}
