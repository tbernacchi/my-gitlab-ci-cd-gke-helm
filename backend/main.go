package main

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	os.Remove("./backend.db")
	db, err := sql.Open("sqlite3", "./backend.db")
	if err != nil {
		log.WithError(err).Error("Sql open database error")
	}
	defer db.Close()

	sqlDdl := `
	create table users (id integer not null primary key, name text);
	delete from users;
	`

	_, err = db.Exec(sqlDdl)
	if err != nil {
		log.WithError(err).Error(sqlDdl)
		return
	}
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	log.Info(r)
	w.WriteHeader(200)
	w.Write([]byte("ok"))
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	log.Info(r)
	log.Info("Getting users")

	db, err := sql.Open("sqlite3", "./backend.db")
	db.Begin()
	if err != nil {
		log.WithError(err).Error("Database begin error")
	}

	statement := "select id, name from users"
	rows, err := db.Query(statement)
	if err != nil {
		log.WithError(err).Error(statement)
	}

	var users []User

	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name); err != nil {
			log.WithError(err).Error("No users found")
		}
		users = append(users, u)
	}

	output, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)

	defer db.Close()
	defer rows.Close()
}

func getUser(w http.ResponseWriter, r *http.Request) {
	log.Info(r)
	log.Info("Getting user")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	db, err := sql.Open("sqlite3", "./backend.db")
	db.Begin()
	if err != nil {
		log.WithError(err).Error("Database begin error")
	}

	statement, err := db.Prepare("select id, name from users where id = ?")
	if err != nil {
		log.WithError(err).Error(statement)
	}

	u := User{}
	err = statement.QueryRow(id).Scan(&u.ID, &u.Name)
	if err != nil {
		log.WithError(err).Error(statement)
	}

	output, err := json.Marshal(u)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)

	defer db.Close()
}

func createUser(w http.ResponseWriter, r *http.Request) {
	log.Info(r)
	log.Info("Creating user")
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Unmarshal
	var u User
	err = json.Unmarshal(b, &u)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	db, err := sql.Open("sqlite3", "./backend.db")
	transaction, err := db.Begin()
	if err != nil {
		log.WithError(err).Error("Database begin error")
	}

	statement := "insert into users (id, name) values ($1, $2)"
	_, err = transaction.Exec(statement, u.ID, u.Name)
	if err != nil {
		log.WithError(err).Error(statement)
		http.Error(w, err.Error(), 500)
	}
	log.Infof("Created user '%d'.", u.ID)

	output, err := json.Marshal(u)
	if err != nil {
		log.WithError(err).Error(output)
		http.Error(w, err.Error(), 500)
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)

	defer db.Close()
	transaction.Commit()
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	log.Info(r)
	log.Info("Deleting user")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	db, err := sql.Open("sqlite3", "./backend.db")
	transaction, err := db.Begin()
	if err != nil {
		log.WithError(err).Error("Database begin error")
	}

	statement := "delete from users where id = $1"
	_, err = transaction.Exec(statement, id)
	if err != nil {
		log.WithError(err).Error(statement)
	}
	log.Infof("Created user '%d'.", id)

	defer db.Close()
	transaction.Commit()
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/healthz", healthzHandler).Methods("GET")
	router.HandleFunc("/users/{id:[0-9]+}", getUser).Methods("GET")
	router.HandleFunc("/users", createUser).Methods("POST")
	router.HandleFunc("/users/{id:[0-9]+}", deleteUser).Methods("DELETE")
	log.SetFormatter(&log.JSONFormatter{})
	log.Info("Starting backend")

	http.ListenAndServe(":8080", router)

}
