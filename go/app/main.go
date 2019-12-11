package main

import (
	_ "database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/json-iterator/go"
	_ "github.com/lib/pq"
)

type User struct {
	User_id   int    `json:"user_id"`
	User_name string `json:"user_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type Recipes struct {
	Recipe_id          int    `json:"recipe_id"`
	Recipe_name        string `json:"recipe_name"`
	Recipe_description string `json:"recipe_description"`
	User_id            int    `json:"user_id"`
}

type Ingredients struct {
	Ingredient_id   int    `json:"ingredient_id"`
	Ingredient_name string `json:"ingredient_name"`
}

type Directions struct {
	Direction_id      int    `json:"direction_id"`
	Direction_details string `json:"direction_details"`
	Direction_order   int    `json:"direction_number"`
	Recipe_id         int    `json:"recipe_id"`
}

type RecipeIngredients struct {
	Recipe_ingredient_id int `json:"recipe_ingredient_id"`
	Ingredient_id        int `json:"ingredient_id"`
	Recipe_id            int `json:"recipe_id"`
}

//Variaveis fixas de postgresBD
const (
	dbHost     = "postgres"
	dbUser     = "docker"
	dbPassword = "docker"
	dbName     = "recipes"
)

// função para abrir connexão com DB
func openConnDB() *sqlx.DB {
	dbinfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable", dbHost, dbUser, dbPassword, dbName)
	db, err := sqlx.Connect("postgres", dbinfo)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

// Função para fechar porta aberta na DB
func closeConnDB(db *sqlx.DB) {
	db.Close()
}

func searchAllUsers() []byte {
	row := []User{}
	db := openConnDB()
	err := db.Select(&row, "SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}
	j, _ := json.Marshal(row)
	closeConnDB(db)
	return j
}

func searchUserBDbyID(id string) []byte {
	row := []User{}
	db := openConnDB()
	err := db.Select(&row, "SELECT * FROM users WHERE user_id ="+id)
	if err != nil {
		log.Fatal(err)
	}
	//var json = jsoniter.ConfigCompatibleWithStandardLibrary
	j, _ := json.Marshal(row)
	closeConnDB(db)
	return j
}

func getUsersByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	rows := searchUserBDbyID(vars["id"])
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(rows)
}

//Função para inserir na tabela, este recebe uma estrutura completa e só junta à query
func insertUser(user User) bool {
	db := openConnDB()
	tx := db.MustBegin()
	tx.NamedExec("INSERT INTO users (user_id, user_name, email, password) VALUES (:user_id, :user_name, :email, :password)", &user)
	err := tx.Commit()
	if err != nil {
		return false
	}
	closeConnDB(db)
	return true
}

func deleteUser(id string) bool {
	db := openConnDB()
	tx := db.MustBegin()
	tx.MustExec("DELETE FROM users WHERE user_id=" + id)
	err := tx.Commit()
	if err != nil {
		return false
	}
	closeConnDB(db)
	return true
}

func editUser(user User) bool {

	db := openConnDB()
	tx := db.MustBegin()
	tx.NamedExec("UPDATE users SET user_name=:user_name, email=:email, password=:password WHERE user_id=:user_id", &user)
	err := tx.Commit()
	if err != nil {
		return false
	}
	closeConnDB(db)
	return true
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	rows := searchAllUsers()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(rows)
}

func insert(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	//var json = jsoniter.ConfigCompatibleWithStandardLibrary

	json.Marshal(&user)
	json.NewDecoder(r.Body).Decode(&user)
	result := insertUser(user)
	j, _ := json.Marshal(result)
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	result := deleteUser(vars["id"])
	j, _ := json.Marshal(result)
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func edit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	//var json = jsoniter.ConfigCompatibleWithStandardLibrary
	json.NewDecoder(r.Body).Decode(&user)
	result := editUser(user)
	j, _ := json.Marshal(result)
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func checkUser(user_name string, pass string) bool {
	var u User
	db := openConnDB()
	err := db.Get(&u, "SELECT user_name, password FROM users WHERE user_name like "+"'"+user_name+"'")
	if err != nil {
		return false
	}
	if u.Password != pass {
		return false
	}
	closeConnDB(db)
	return true
}

func login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	//var json = jsoniter.ConfigCompatibleWithStandardLibrary
	json.NewDecoder(r.Body).Decode(&user)
	result := checkUser(user.User_name, user.Password)
	j, _ := json.Marshal(result)
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func test(w http.ResponseWriter, r *http.Request) {
	//var json = jsoniter.ConfigCompatibleWithStandardLibrary
	j, _ := json.Marshal("test")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func main() {
	//fmt.Println("Test")

	//Init router
	r := mux.NewRouter() // := atribui tipo á variavel
	//CORS
	corsObj := handlers.AllowedOrigins([]string{"*"})
	headersOk := handlers.AllowedHeaders([]string{"Content-Type", "Bearer", "Bearer ", "content-type", "Origin", "Accept"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	// router handlers
	//r.HandleFunc("/api/porcurar/{key}", getbyNome).Methods("GET")
	r.HandleFunc("/api/test", test).Methods("GET")
	r.HandleFunc("/api/insertUser", insert).Methods("POST")
	r.HandleFunc("/api/searchUser/id/{id}", getUsersByID).Methods("GET")
	r.HandleFunc("/api/deleteUser/id/{id}", delete).Methods("DELETE")
	r.HandleFunc("/api/editUser", edit).Methods("POST")
	r.HandleFunc("/api/login", login).Methods("POST")
	r.HandleFunc("/api/allUsers", getAllUsers).Methods("GET")

	// pagina de testes
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(corsObj, headersOk, methodsOk)(r))) // se falhar dá erro !*/
}
