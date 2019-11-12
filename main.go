package main

/* import package
"json agar bisa kirim data berupa json"
"log untuk menampilkan log di terminal pada saat debug"
"net/http membuat server golang"
"mysql agar kita bisa menggunakan database mysql"
"mux untuk membuat routing url"
*/
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {

	// routing url "GET, POST, PUT, DELETE"
	router := mux.NewRouter()
	router.HandleFunc("/getstudents", returnAllUsers).Methods("GET")
	router.HandleFunc("/students", insertUsersMultipart).Methods("POST")
	router.HandleFunc("/students", updateUsersMultipart).Methods("PUT")
	router.HandleFunc("/students", deleteUsersMultipart).Methods("DELETE")
	http.Handle("/", router)
	fmt.Println("Connected to port 1234")
	// membuat server dengan port 1234
	log.Fatal(http.ListenAndServe(":1234", router))

}

// fungsi Get all data user
func returnAllUsers(w http.ResponseWriter, r *http.Request) {
	var users Users
	var arrUser []Users
	var response Response

	// konek ke db
	db := connect()
	defer db.Close()

	// query untuk get data
	rows, err := db.Query("Select id, nama, npm, jk from students")
	// jika error
	if err != nil {
		log.Print(err)
	}

	// looping data
	for rows.Next() {
		if err := rows.Scan(&users.Id, &users.Npm, &users.Nama, &users.Jk); err != nil {
			log.Fatal(err.Error())

		} else {
			arrUser = append(arrUser, users)
		}
	}

	// tampilkan responnya
	response.Status = 1
	response.Message = "Success"
	response.Data = arrUser

	// tampilkan sebagai json
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

// fungsi insert data
func insertUsersMultipart(w http.ResponseWriter, r *http.Request) {
	var response Response

	// konek database
	db := connect()
	defer db.Close()

	// kita akan menggunakan tipepost "form-data" dengan fungsi "ParseMultipartForm"
	// jika ingin menggunakan "x-www-form-urlencoded" silahkan ganti dengan "ParseForm"
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	// buat variable untuk menampung request dari form
	npm := r.FormValue("npm")
	nama := r.FormValue("nama")
	jk := r.FormValue("jk")

	// query insert
	_, err = db.Exec("INSERT INTO students (npm, nama, jk) values (?,?,?)",
		npm,
		nama,
		jk,
	)

	// jika error
	if err != nil {
		log.Print(err)
	}

	// tampilkan responnya
	response.Status = 1
	response.Message = "Success Add"
	log.Print("Insert data to database")

	// tampilkan sebagai json
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

// fungsi update data
func updateUsersMultipart(w http.ResponseWriter, r *http.Request) {
	var response Response

	// konek database
	db := connect()
	defer db.Close()

	// kita akan menggunakan tipepost "form-data" dengan fungsi "ParseMultipartForm"
	// jika ingin menggunakan "x-www-form-urlencoded" silahkan ganti dengan "ParseForm"
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	// buat variable untuk menampung request dari form
	id := r.FormValue("id")
	npm := r.FormValue("npm")
	nama := r.FormValue("nama")
	jk := r.FormValue("jk")

	// query update
	_, err = db.Exec("UPDATE students set npm = ?, nama = ?, jk = ? where id = ?",
		npm,
		nama,
		jk,
		id,
	)

	// jika error
	if err != nil {
		log.Print(err)
	}

	// tampilkan responnya
	response.Status = 1
	response.Message = "Success Update Data"
	log.Print("Update data to database")

	// tampilkan sebagai json
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

// delete data
func deleteUsersMultipart(w http.ResponseWriter, r *http.Request) {
	var response Response

	// konek database
	db := connect()
	defer db.Close()

	// kita akan menggunakan tipepost "form-data" dengan fungsi "ParseMultipartForm"
	// jika ingin menggunakan "x-www-form-urlencoded" silahkan ganti dengan "ParseForm"
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	// buat variable untuk menampung request dari form
	id := r.FormValue("id")

	// query delete
	_, err = db.Exec("DELETE from students where id = ?",
		id,
	)

	// jika error
	if err != nil {
		log.Print(err)
	}

	// tampilkan responnya
	response.Status = 1
	response.Message = "Success Delete Data"
	log.Print("Delete data to database")

	// tampilkan sebagai json
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
