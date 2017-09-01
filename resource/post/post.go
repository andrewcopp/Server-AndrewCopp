package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var db *sql.DB
var mock map[int]string

func init() {
	mock = map[int]string{
		1: "Hello, World!",
		2: "It's working. It's working!",
		3: "This one starts like an actual blog post. Too bad it is static.",
	}

	db = &sql.DB{}
	http.HandleFunc("/posts/", handler)
	http.HandleFunc("/posts/status/", status)
}

func main() {
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

		if post, ok := mock[id]; ok {
			res, err := json.Marshal(post)
			if err != nil {

			}

			w.Write(res)
		}

	} else {
		posts := []string{}

		for _, post := range mock {
			posts = append(posts, post)
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
