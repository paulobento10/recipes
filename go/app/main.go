package main

import (
	_ "database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"golang.org/x/crypto/bcrypt"

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
	RecipeIngredient_id int `json:"recipeingredient_id" db:"recipeingredient_id"`
	Ingredient_id       int `json:"ingredient_id" db:"ingredient_id"`
	Recipe_id           int `json:"recipe_id" db:"recipe_id"`
}

//postgres vars
const (
	dbHost     = "postgres"
	dbUser     = "docker"
	dbPassword = "docker"
	dbName     = "recipes"
)

//DB Functions

/**
* Opens up db connection
 */
func openConnDB() *sqlx.DB {
	dbinfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable", dbHost, dbUser, dbPassword, dbName)
	db, err := sqlx.Connect("postgres", dbinfo)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

/**
* Closes the port of the db connection
 */
func closeConnDB(db *sqlx.DB) {
	db.Close()
}

//-----Model-----
//-----User Functions - Model-----

/**
* [Model][User] Queries the database to get a user by its id
 */
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

/**
* [Model][User] Receives a user as parameter and inserts it in the database
 */
func insertUser(user User) bool {
	db := openConnDB()
	tx := db.MustBegin()
	tx.NamedExec("INSERT INTO users (user_name, email, password) VALUES (:user_name, :email, :password)", &user)
	err := tx.Commit()
	if err != nil {
		return false
	}
	closeConnDB(db)
	return true
}

/**
* [Model][User] Receives a user as parameter and updates it in the database
 */
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

/**
* [Model][User] Receives an id as parameter and deletes the user of the received id in the database
 */
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

/**
* [Model][User] Returns a bool, it verifies if the user_name and password received from parameters are the same as the ones in the database
 */
func checkUser(user_name string, pass string) bool {
	var u User
	db := openConnDB()
	err := db.Get(&u, "SELECT user_name, password FROM users WHERE user_name = "+"'"+user_name+"'")
	if err != nil {
		return false
	}
	// Comparing the password with the hash
	//var aux string = u.Password
	hashedPassword := []byte(u.Password)
	/*log.Fatal("hashedPassword: ")
	log.Fatal(hashedPassword)*/
	password := []byte(pass)
	/*log.Fatal("password: ")
	log.Fatal(password)*/
	e := bcrypt.CompareHashAndPassword(hashedPassword, password)
	/*log.Fatal("Bycript: ")
	log.Fatal(e)  // nil means it is a match*/
	if e != nil { //u.Password != pass {
		return false
	}
	closeConnDB(db)
	return true
}

/**
* [Model][User] Returns a json with all the users in the database
 */
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

//-----Recipes Functions - Model-----

/**
* [Model][Recipes] Queries the database to get a recipe by its id
 */
func getRecipeById(id string) []byte {
	row := []Recipes{}
	db := openConnDB()
	err := db.Select(&row, "SELECT * FROM recipes WHERE recipe_id ="+id)
	if err != nil {
		log.Fatal(err)
	}
	//var json = jsoniter.ConfigCompatibleWithStandardLibrary
	j, _ := json.Marshal(row)
	closeConnDB(db)
	return j
}

/**
* [Model][Recipes] Receives a recipe as parameter and inserts it in the database
 */
func insertRecipe(r Recipes) bool {
	db := openConnDB()
	tx := db.MustBegin()
	tx.NamedExec("INSERT INTO recipes (recipe_name, recipe_description, user_id) VALUES (:recipe_name, :recipe_description, :user_id)", &r)
	err := tx.Commit()
	if err != nil {
		return false
	}
	closeConnDB(db)
	return true
}

/**
* [Model][Recipes] Receives a recipe as parameter and updates it in the database
 */
func editRecipe(r Recipes) bool {

	db := openConnDB()
	tx := db.MustBegin()
	tx.NamedExec("UPDATE recipes SET recipe_name=:recipe_name, recipe_description=:recipe_description, user_id=:user_id WHERE recipe_id=:recipe_id", &r)
	err := tx.Commit()
	if err != nil {
		return false
	}
	closeConnDB(db)
	return true
}

/**
* [Model][Recipes] Receives an id as parameter and deletes the recipe of the received id in the database
 */
func deleteRecipe(id string) bool {
	db := openConnDB()
	tx := db.MustBegin()
	tx.MustExec("DELETE FROM recipes WHERE recipe_id=" + id)
	err := tx.Commit()
	if err != nil {
		return false
	}
	closeConnDB(db)
	return true
}

/**
* [Model][Recipes] Queries the database to get a recipe by its user_id
 */
func getRecipeByUserId(user_id string) []byte {
	row := []Recipes{}
	db := openConnDB()
	err := db.Select(&row, "SELECT * FROM recipes WHERE user_id ="+user_id) //+" ORDER BY recipe_id ASC"
	if err != nil {
		log.Fatal(err)
	}
	//var json = jsoniter.ConfigCompatibleWithStandardLibrary
	j, _ := json.Marshal(row)
	closeConnDB(db)
	return j
}

/**
* [Model][Recipes] Queries the database to get a recipe by its recipe_name
 */
func getRecipeByName(recipe_name string) []byte {
	row := []Recipes{}
	db := openConnDB()
	err := db.Select(&row, "SELECT * FROM recipes WHERE recipe_name LIKE "+"'%"+recipe_name+"%'")
	if err != nil {
		log.Fatal(err)
	}
	//var json = jsoniter.ConfigCompatibleWithStandardLibrary
	j, _ := json.Marshal(row)
	closeConnDB(db)
	return j
}

//-----Ingredients Functions - Model-----

/**
* [Model][Ingredients] Queries the database to get a ingredients by its id
 */
func getIngredientById(id string) []byte {
	row := []Ingredients{}
	db := openConnDB()
	err := db.Select(&row, "SELECT * FROM ingredients WHERE ingredient_id ="+id)
	if err != nil {
		log.Fatal(err)
	}
	//var json = jsoniter.ConfigCompatibleWithStandardLibrary
	j, _ := json.Marshal(row)
	closeConnDB(db)
	return j
}

/**
* [Model][Ingredients] Receives a ingredient as parameter and inserts it in the database
 */
func insertIngredient(i Ingredients) bool {
	db := openConnDB()
	tx := db.MustBegin()
	tx.NamedExec("INSERT INTO ingredients (ingredient_name) VALUES (:ingredient_name)", &i)
	err := tx.Commit()
	if err != nil {
		return false
	}
	closeConnDB(db)
	return true
}

/**
* [Model][Ingredients] Receives a ingredient as parameter and updates it in the database
 */
func editIngredient(i Ingredients) bool {

	db := openConnDB()
	tx := db.MustBegin()
	tx.NamedExec("UPDATE ingredients SET ingredient_name=:ingredient_name WHERE ingredient_id=:ingredient_id", &i)
	err := tx.Commit()
	if err != nil {
		return false
	}
	closeConnDB(db)
	return true
}

/**
* [Model][Ingredients] Receives an id as parameter and deletes the ingredient of the received id in the database
 */
func deleteIngredient(id string) bool {
	db := openConnDB()
	tx := db.MustBegin()
	tx.MustExec("DELETE FROM ingredients WHERE ingredient_id=" + id)
	err := tx.Commit()
	if err != nil {
		return false
	}
	closeConnDB(db)
	return true
}

/**
* [Model][Ingredients] Queries the database to get an ingredient by its ingredient_name
 */
func getIngredientByName(ingredient_name string) []byte {
	row := []Ingredients{}
	db := openConnDB()
	err := db.Select(&row, "SELECT * FROM ingredients WHERE ingredient_name LIKE "+"'%"+ingredient_name+"%'")
	if err != nil {
		log.Fatal(err)
	}
	//var json = jsoniter.ConfigCompatibleWithStandardLibrary
	j, _ := json.Marshal(row)
	closeConnDB(db)
	return j
}

//-----RecipeIngredients Functions - Model------

/**
* [Model][RecipeIngredients] Queries the database to get a RecipeIngredients by its id
 */
func getRecipeIngredientsById(id string) []byte {
	row := []RecipeIngredients{}
	db := openConnDB()
	query := "SELECT recipeingredient_id, ingredient_id, recipe_id FROM recipeingredients WHERE recipeingredient_id = " + id
	err := db.Select(&row, strings.ToLower(query))
	if err != nil {
		log.Fatal(err)
	}
	//var json = jsoniter.ConfigCompatibleWithStandardLibrary
	j, _ := json.Marshal(row)
	closeConnDB(db)
	return j
}

/**
* [Model][RecipeIngredients] Receives RecipeIngredients as parameter and inserts it in the database
 */
func insertRecipeIngredients(r RecipeIngredients) bool {
	db := openConnDB()
	tx := db.MustBegin()
	tx.NamedExec("INSERT INTO recipeingredients (ingredient_id, recipe_id) VALUES (:ingredient_id, :recipe_id)", &r)
	err := tx.Commit()
	if err != nil {
		return false
	}
	closeConnDB(db)
	return true
}

/**
* [Model][RecipeIngredients] Receives RecipeIngredients as parameter and updates it in the database
 */
func editRecipeIngredients(r RecipeIngredients) bool {

	db := openConnDB()
	tx := db.MustBegin()
	query := "UPDATE recipeingredients SET ingredient_id=:ingredient_id, recipe_id=:recipe_id WHERE recipeingredient_id=:recipeingredient_id"
	tx.NamedExec(strings.ToLower(query), &r)
	err := tx.Commit()
	if err != nil {
		return false
	}
	closeConnDB(db)
	return true
}

/**
* [Model][RecipeIngredients] Receives an id as parameter and deletes the RecipeIngredients of the received id in the database
 */
func deleteRecipeIngredients(id string) bool {
	db := openConnDB()
	tx := db.MustBegin()
	tx.MustExec("DELETE FROM recipeingredients WHERE recipeingredient_id=" + id)
	err := tx.Commit()
	if err != nil {
		return false
	}
	closeConnDB(db)
	return true
}

/**
* [Model][RecipeIngredients] Queries the database to get RecipeIngredients by its ingredient_name from table ingredients
 */
func getRecipeIngredientsByIngredient(ingredient_name string) []byte {
	row := []RecipeIngredients{}
	db := openConnDB()
	//query := "SELECT * FROM recipeingredients WHERE ingredient_id = " + ingredient_id
	query := "SELECT * FROM recipeingredients WHERE ingredient_id IN (select ingredient_id from ingredients where ingredient_name LIKE " + "'%" + ingredient_name + "%')"
	err := db.Select(&row, strings.ToLower(query))
	if err != nil {
		log.Fatal(err)
	}
	//var json = jsoniter.ConfigCompatibleWithStandardLibrary
	j, _ := json.Marshal(row)
	closeConnDB(db)
	return j
}

/**
* [Model][RecipeIngredients] Queries the database to get RecipeIngredients by its recipe_name from table recipes
 */
func getRecipeIngredientsByRecipe(recipe_name string) []byte {
	row := []RecipeIngredients{}
	db := openConnDB()
	//query := "SELECT * FROM recipeingredients WHERE ingredient_id = " + ingredient_id
	query := "SELECT * FROM recipeingredients WHERE recipe_id IN (select recipe_id from recipes where recipe_name LIKE " + "'%" + recipe_name + "%')"
	err := db.Select(&row, strings.ToLower(query))
	if err != nil {
		log.Fatal(err)
	}
	//var json = jsoniter.ConfigCompatibleWithStandardLibrary
	j, _ := json.Marshal(row)
	closeConnDB(db)
	return j
}

//-----Controller-----
//-----User Functions - Controller-----

/**
* [Controller][User] function to get a user by id
 */
func getUsersByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	rows := searchUserBDbyID(vars["id"])
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(rows)
}

/**
* [Controller][User] function to insert a user
 */
func insertUserRoute(w http.ResponseWriter, r *http.Request) {
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

/**
* [Controller][User] function to edit a user
 */
func editUserRoute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	//var json = jsoniter.ConfigCompatibleWithStandardLibrary
	json.NewDecoder(r.Body).Decode(&user)
	result := editUser(user)
	j, _ := json.Marshal(result)
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

/**
* [Controller][User] function to delete a user
 */
func deleteUserRoute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	result := deleteUser(vars["id"])
	j, _ := json.Marshal(result)
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

/**
* [Controller][User] function to login a user, calls checkUser function
 */
func login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	//var json = jsoniter.ConfigCompatibleWithStandardLibrary
	json.NewDecoder(r.Body).Decode(&user)
	//result := checkUser(user.User_name, user.Password)

	var u User
	var b bool
	db := openConnDB()
	err := db.Get(&u, "SELECT user_name, password FROM users WHERE user_name = "+"'"+user.User_name+"'")
	if err != nil {
		b = false
	}
	// Comparing the password with the hash
	//var aux string = u.Password
	//hashedPassword := []byte(u.Password)
	fmt.Println("hashedPassword: ")
	fmt.Println(u.Password) //hashedPassword
	//password := []byte(user.Password)
	fmt.Println("password: ")
	fmt.Println(user.Password) //password
	/*e := bcrypt.CompareHashAndPassword(hashedPassword, password)
	fmt.Println("Bycript: ")
	fmt.Println(e) // nil means it is a match
	if e != nil {  //u.Password != pass {
		b = false
	}
	closeConnDB(db)
	b = true*/

	j, _ := json.Marshal(b) //result
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

/**
* [Controller][User] function to get all users, returns a json with all the users in the database, calls searchAllUsers function
 */
func getAllUsers(w http.ResponseWriter, r *http.Request) {
	rows := searchAllUsers()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(rows)
}

//-----Recipes Functions - Controller------
/**
* [Controller][Recipes] function to get a recipe by id
 */
func getRecipeByIdRoute(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	rows := getRecipeById(vars["id"])
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(rows)
}

/**
* [Controller][Recipes] function to insert a recipe
 */
func insertRecipeRoute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var recipe Recipes
	//var json = jsoniter.ConfigCompatibleWithStandardLibrary

	json.Marshal(&r)
	json.NewDecoder(r.Body).Decode(&recipe)
	result := insertRecipe(recipe)
	j, _ := json.Marshal(result)
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

/**
* [Controller][Recipes] function to edit a recipe
 */
func editRecipeRoute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var recipe Recipes
	//var json = jsoniter.ConfigCompatibleWithStandardLibrary
	json.NewDecoder(r.Body).Decode(&recipe)
	result := editRecipe(recipe)
	j, _ := json.Marshal(result)
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

/**
* [Controller][Recipes] function to delete a recipe
 */
func deleteRecipeRoute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	result := deleteRecipe(vars["id"])
	j, _ := json.Marshal(result)
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

/**
* [Controller][Recipes] function to get a recipe by user_id
 */
func getRecipeByUserIdRoute(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	rows := getRecipeByUserId(vars["id"])
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(rows)
}

/**
* [Controller][Recipes] function to get a recipe by recipe_name
 */
func getRecipeByNameRoute(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	rows := getRecipeByName(vars["name"])
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(rows)
}

//-----Ingredients Functions - Controller------
/**
* [Controller][Ingredients] function to get an ingredient by id
 */
func getIngredientByIdRoute(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	rows := getIngredientById(vars["id"])
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(rows)
}

/**
* [Controller][Ingredients] function to insert an ingredient
 */
func insertIngredientRoute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var ingredient Ingredients
	//var json = jsoniter.ConfigCompatibleWithStandardLibrary

	json.Marshal(&r)
	json.NewDecoder(r.Body).Decode(&ingredient)
	result := insertIngredient(ingredient)
	j, _ := json.Marshal(result)
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

/**
* [Controller][Ingredients] function to edit an ingredient
 */
func editIngredientRoute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var ingredient Ingredients
	//var json = jsoniter.ConfigCompatibleWithStandardLibrary
	json.NewDecoder(r.Body).Decode(&ingredient)
	result := editIngredient(ingredient)
	j, _ := json.Marshal(result)
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

/**
* [Controller][Ingredients] function to delete an ingredient
 */
func deleteIngredientRoute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	result := deleteIngredient(vars["id"])
	j, _ := json.Marshal(result)
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

/**
* [Controller][Ingredients] function to get an ingredient by ingredient_name
 */
func getIngredientByNameRoute(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	rows := getIngredientByName(vars["name"])
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(rows)
}

//-----RecipeIngredients Functions - Controller------
/**
* [Controller][RecipeIngredients] function to get RecipeIngredients by id
 */
func getRecipeIngredientsByIdRoute(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	rows := getRecipeIngredientsById(vars["id"])
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(rows)
}

/**
* [Controller][RecipeIngredients] function to insert RecipeIngredients
 */
func insertRecipeIngredientsRoute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var recipeIngredients RecipeIngredients
	//var json = jsoniter.ConfigCompatibleWithStandardLibrary

	json.Marshal(&r)
	json.NewDecoder(r.Body).Decode(&recipeIngredients)
	result := insertRecipeIngredients(recipeIngredients)
	j, _ := json.Marshal(result)
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

/**
* [Controller][RecipeIngredients] function to edit RecipeIngredients
 */
func editRecipeIngredientsRoute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var recipeIngredients RecipeIngredients
	//var json = jsoniter.ConfigCompatibleWithStandardLibrary
	json.NewDecoder(r.Body).Decode(&recipeIngredients)
	result := editRecipeIngredients(recipeIngredients)
	j, _ := json.Marshal(result)
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

/**
* [Controller][RecipeIngredients] function to delete RecipeIngredients
 */
func deleteRecipeIngredientsRoute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	result := deleteRecipeIngredients(vars["id"])
	j, _ := json.Marshal(result)
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

/**
* [Controller][RecipeIngredients] function to get RecipeIngredients by ingredient_name from ingredients table
 */
func getRecipeIngredientsByIngredientRoute(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	rows := getRecipeIngredientsByIngredient(vars["name"])
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(rows)
}

/**
* [Controller][RecipeIngredients] function to get RecipeIngredients by recipe_name from recipes table
 */
func getRecipeIngredientsByRecipeRoute(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	rows := getRecipeIngredientsByRecipe(vars["name"])
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(rows)
}

func main() {
	//Init router
	r := mux.NewRouter() // := atribui tipo á variavel
	//CORS
	corsObj := handlers.AllowedOrigins([]string{"*"})
	headersOk := handlers.AllowedHeaders([]string{"Content-Type", "Bearer", "Bearer ", "content-type", "Origin", "Accept"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	// router handlers
	//r.HandleFunc("/api/porcurar/{key}", getbyNome).Methods("GET")

	//User routes
	r.HandleFunc("/api/insertUser", insertUserRoute).Methods("POST")
	r.HandleFunc("/api/searchUser/id/{id}", getUsersByID).Methods("GET")
	r.HandleFunc("/api/deleteUser/id/{id}", deleteUserRoute).Methods("DELETE")
	r.HandleFunc("/api/editUser", editUserRoute).Methods("POST")
	r.HandleFunc("/api/login", login).Methods("POST")
	r.HandleFunc("/api/allUsers", getAllUsers).Methods("GET")

	//Recipes routes
	r.HandleFunc("/api/insertRecipe", insertRecipeRoute).Methods("POST")
	r.HandleFunc("/api/searchRecipe/id/{id}", getRecipeByIdRoute).Methods("GET")
	r.HandleFunc("/api/deleteRecipe/id/{id}", deleteRecipeRoute).Methods("DELETE")
	r.HandleFunc("/api/editRecipe", editRecipeRoute).Methods("POST")
	r.HandleFunc("/api/searchUserRecipe/id/{id}", getRecipeByUserIdRoute).Methods("GET")
	r.HandleFunc("/api/searchRecipeName/name/{name}", getRecipeByNameRoute).Methods("GET")

	//Ingredients routes
	r.HandleFunc("/api/insertIngredient", insertIngredientRoute).Methods("POST")
	r.HandleFunc("/api/searchIngredient/id/{id}", getIngredientByIdRoute).Methods("GET")
	r.HandleFunc("/api/deleteIngredient/id/{id}", deleteIngredientRoute).Methods("DELETE")
	r.HandleFunc("/api/editIngredient", editIngredientRoute).Methods("POST")
	r.HandleFunc("/api/searchIngredientName/name/{name}", getIngredientByNameRoute).Methods("GET")

	//RecipeIngredients routes
	r.HandleFunc("/api/insertRecipeIngredients", insertRecipeIngredientsRoute).Methods("POST")
	r.HandleFunc("/api/searchRecipeIngredients/id/{id}", getRecipeIngredientsByIdRoute).Methods("GET")
	r.HandleFunc("/api/deleteRecipeIngredients/id/{id}", deleteRecipeIngredientsRoute).Methods("DELETE")
	r.HandleFunc("/api/editRecipeIngredients", editRecipeIngredientsRoute).Methods("POST")
	r.HandleFunc("/api/searchRecipeIngredientsName/name/{name}", getRecipeIngredientsByIngredientRoute).Methods("GET")
	r.HandleFunc("/api/searchRecipeRecipesName/name/{name}", getRecipeIngredientsByRecipeRoute).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(corsObj, headersOk, methodsOk)(r))) // se falhar dá erro !*/
}
