package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_DBNAME")

	info := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, pass, dbname)

	var err error
	db, err = sql.Open("postgres", info)
	if err != nil {
		log.Fatalln(err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln(err.Error())
	}

	http.HandleFunc("/posts/", handler)
	http.HandleFunc("/posts/status/", status)
}

func main() {
	defer db.Close()
	log.Println("Deploying on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Println(err.Error())
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	_ = r.URL.Query()
	cmpts := strings.Split(strings.Trim(path, "/"), "/")

	if len(cmpts) == 2 {
		id, err := strconv.Atoi(cmpts[1])
		if err != nil {

		}

		row := db.QueryRow("SELECT title FROM posts WHERE id = $1;", id)

		var title string
		row.Scan(
			&title,
		)

		res, err := json.Marshal(title)
		if err != nil {

		}

		w.Write(res)

	} else {
		rows, err := db.Query("SELECT title FROM posts;")
		if err != nil {
			log.Fatalln(err.Error())
		}

		posts := []string{}
		for rows.Next() {
			var title string
			rows.Scan(
				&title,
			)
			posts = append(posts, title)
		}

		res, err := json.Marshal(posts)
		if err != nil {
			log.Println(err.Error())
		}

		w.Write(res)
	}
}

func status(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

type ReadPostRequest struct {
	Identifier int
}

type ReadPostRequests struct {
	Requests []*ReadPostRequest
}

// Reader
type Reader interface {
}

// obviously can read posts

// Writer
type Writer interface {
}

// has the authority to create, update, and delete
// those should probably also be all their own interfaces

// Creater
type Create interface {
}

// Updater
type Updater interface {
}

// Deleter
type Deleter interface {
}

// Authenticater
type Authenticater interface {
}

// one type of authenticater
// checks that you have credentials
// not required of all endpoints

// Authorizer
type Authorizer interface {
}

// four types of accounts
// user
// author who can modify original
// moderator
// administrator who can do anything
