package Handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type book struct {
	Id         int       `json:"Id"`
	Title      string    `json:"Title"`
	Author     string    `json:"Author"`
	Page       int       `json:"page"`
	Created_at time.Time `json:"Date and Time"`
}

type err error

// func respondWithError(response http.ResponseWriter, statusCode int, msg string) {
// 	respondWithJSON(response, statusCode, map[string]string{"error": msg})
// }
func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func dbConn() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:Dsp123@tcp(127.0.0.1:3306)/bookdatabase?charset=utf8")
	return db, err
}

func CreateBook(w http.ResponseWriter, r *http.Request) {

	db, err := sql.Open("mysql", "root:Dsp123@tcp(127.0.0.1:3306)/")
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS bookdatabase;")
	check(err)

	_, err = db.Exec("USE bookdatabase;")
	check(err)

	_, err = db.Exec(`CREATE Table IF NOT EXISTS books(
		id int NOT NULL AUTO_INCREMENT,
		title varchar(100) NOT NULL,
		author varchar(100) NOT NULL,
		page int NOT NULL, 
		created_at DATETIME, 
		PRIMARY KEY (id));`)
	check(err)

	db, err = dbConn()
	check(err)
	err = db.Ping()
	check(err)

	b := book{
		Id:         1,
		Title:      "You Can Win",
		Author:     "Shiv khera",
		Page:       250,
		Created_at: time.Now(),
	}
	b1 := book{
		Id:         2,
		Title:      "The Alchemist",
		Author:     "Paulo Cohelo",
		Page:       250,
		Created_at: time.Now(),
	}

	_, err = db.Exec("INSERT INTO books(id, title, author, page, created_at) VALUES (?,?,?,?,?)", b.Id, b.Title, b.Author, b.Page, b.Created_at)
	fmt.Fprintln(w, "Inserted Successfully")
	_, err = db.Exec("INSERT INTO books(id, title, author, page, created_at) VALUES (?,?,?,?,?)", b1.Id, b1.Title, b1.Author, b1.Page, b1.Created_at)
	fmt.Fprintln(w, "Inserted Successfully")
	defer db.Close()
}

func ReadBook(w http.ResponseWriter, r *http.Request) {
	db, err := dbConn()
	check(err)

	row, err := db.Query("SELECT * FROM books")
	defer row.Close()
	check(err)
	var bInfo []book
	for row.Next() {
		var b book
		row.Scan(&b.Id, &b.Title, &b.Author, &b.Page, &b.Created_at)
		bInfo = append(bInfo, b)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	Bbytes, _ := json.Marshal(&bInfo)
	w.Write(Bbytes)
	return
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	db, err := dbConn()
	err = db.Ping()
	check(err)
	id := r.URL.Query().Get(`id`)

	_, err = db.Exec("DELETE FROM books WHERE id" + id + ";")
	check(err)
	fmt.Fprintln(w, "Book Deleted")

}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var b book
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
	} else {
		fmt.Println("Product Info - Updated")
		fmt.Println("Id: ", b.Id)
		fmt.Println("Title: ", b.Title)
		fmt.Println("Author: ", b.Author)
		fmt.Println("Page: ", b.Page)
		fmt.Println("CreCreated_at: ", b.Created_at)
		respondJSON(w, http.StatusOK, b)
	}
}

func respondJSON(r http.ResponseWriter, statusCode int, data interface{}) {
	result, _ := json.Marshal(data)
	r.Header().Set("Content-Type", "application/json")
	r.WriteHeader(statusCode)
	r.Write(result)
}

func respondError(r http.ResponseWriter, statusCode int, msg string) {
	respondJSON(r, statusCode, map[string]string{"error": msg})
}
