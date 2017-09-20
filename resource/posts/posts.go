package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/andrewcopp/store"
	"github.com/andrewcopp/store/postgres"
)

var db store.Writer

func init() {

	config := postgres.NewConfig()
	db = postgres.NewPostgres(config)

	if err := db.Connect(); err != nil {
		log.Fatalln(err.Error())
	}

	if err := db.Ping(); err != nil {
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

type Post struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func handler(w http.ResponseWriter, r *http.Request) {

	path := r.URL.Path
	cmps := strings.Split(strings.Trim(path, "/"), "/")
	query := r.URL.Query()
	body := r.Body

	// decoder := json.NewDecoder(r.Body)
	// var p *Post
	// if err := decoder.Decode(p); err != nil {
	// 	log.Fatalln(err.Error())
	// }
	// defer r.Body.Close()

	switch r.Method {
	case http.MethodGet:
		switch len(cmps) {
		case 1:
			println(query)
			objs, _ := db.Search(0, 0)
			res, err := json.Marshal(objs)
			if err != nil {
				log.Println(err.Error())
			}

			w.Write(res)
		case 2:
			id, err := strconv.Atoi(cmps[1])
			if err != nil {

			}
			obj, _ := db.Find(id)
			res, err := json.Marshal(obj)
			if err != nil {
				log.Println(err.Error())
			}

			w.Write(res)
		default:
			fmt.Println(400)
		}
	case http.MethodPost:
		print(body)
		db.Create(map[string]interface{}{})
	case http.MethodPut:
		print(body)
		db.Update(1, map[string]interface{}{})
	case http.MethodDelete:
		id, err := strconv.Atoi(cmps[1])
		if err != nil {

		}
		db.Delete(id)
		print("DELETE")
	}
}

func status(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
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
