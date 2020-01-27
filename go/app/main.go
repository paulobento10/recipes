package main

import (
	_ "database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

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
	Duration           string `json:"duration"`
	Picture            string `json:"picture"`
	Category           string `json:"category"`
	Kcal               string `json:"kcal"`
	User_id            int    `json:"user_id"`
}

type Ingredients struct {
	Ingredient_id   int    `json:"ingredient_id"`
	Ingredient_name string `json:"ingredient_name"`
	Kcal            string `json:"kcal"`
	User_id         int    `json:"user_id"`
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
* [Model][User] Receives a user as parameter and updates its name in the database
 */
func editUserName(user User) bool {

	db := openConnDB()
	tx := db.MustBegin()
	tx.NamedExec("UPDATE users SET user_name=:user_name WHERE user_id=:user_id", &user)
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
* [Model][User] Returns a bool, it verifies if the email and password received from parameters are the same as the ones in the database
 */
func checkUser(email string, pass string) int {
	var u User
	db := openConnDB()
	err := db.Get(&u, "SELECT * FROM users WHERE email = "+"'"+email+"'")
	if err != nil {
		return -1
	}
	// Comparing the password with the hash
	//var aux string = u.Password
	//hashedPassword := []byte(u.Password)
	/*log.Fatal("hashedPassword: ")
	log.Fatal(hashedPassword)*/
	//password := []byte(pass)
	/*log.Fatal("password: ")
	log.Fatal(password)*/
	//e := bcrypt.CompareHashAndPassword(hashedPassword, password)
	/*log.Fatal("Bycript: ")
	log.Fatal(e)  // nil means it is a match*/
	//if e != nil { //u.Password != pass {
	if u.Password != pass {
		return -1
	}
	closeConnDB(db)
	return u.User_id
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
* [Model][Recipes] Queries the database to get all recipes
 */
func getRecipeAll() []byte {
	row := []Recipes{}
	db := openConnDB()
	err := db.Select(&row, "SELECT * FROM recipes")
	if err != nil {
		log.Fatal(err)
	}
	//var json = jsoniter.ConfigCompatibleWithStandardLibrary
	j, _ := json.Marshal(row)
	closeConnDB(db)
	return j
}

/**
* [Model][Recipes] Queries the database to get all recipes no json
 */
func getRecipeAllNoJson() []Recipes {
	row := []Recipes{}
	db := openConnDB()
	err := db.Select(&row, "SELECT * FROM recipes")
	if err != nil {
		log.Fatal(err)
	}
	closeConnDB(db)
	return row
}

/**
* [Model][Recipes] Receives a recipe as parameter and inserts it in the database
 */
func insertRecipe(r Recipes) bool {
	db := openConnDB()
	tx := db.MustBegin()
	tx.NamedExec("INSERT INTO recipes (recipe_name, recipe_description, duration, picture, category, kcal, user_id) VALUES (:recipe_name, :recipe_description, :duration, :picture, :category, :kcal, :user_id)", &r)
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
	tx.NamedExec("UPDATE recipes SET recipe_name=:recipe_name, recipe_description=:recipe_description, duration=:duration, picture=:picture, category=:category, kcal=:kcal, user_id=:user_id WHERE recipe_id=:recipe_id", &r)
	err := tx.Commit()
	if err != nil {
		return false
	}
	closeConnDB(db)
	return true
}

/**
* [Model][Recipes] Receives a recipe as parameter and updates its name in the database
 */
func editRecipeName(r Recipes) bool {

	db := openConnDB()
	tx := db.MustBegin()
	tx.NamedExec("UPDATE recipes SET recipe_name=:recipe_name WHERE recipe_id=:recipe_id", &r)
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
	tx.MustExec("DELETE FROM recipeingredients WHERE recipe_id=" + id)
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
* [Model][Recipes] Queries the database to get a recipe kcal by its user_id
 */
func getRecipeKcalByRecipeId(recipe_id string) int {
	db := openConnDB()
	recipe := []Recipes{}
	query := "SELECT * FROM recipes WHERE recipe_id =" + recipe_id
	err := db.Select(&recipe, query)
	recipe_kcal := recipe[0].Kcal
	if err != nil {
		log.Fatal(err)
	}
	closeConnDB(db)
	recipe_kcal_int, _ := strconv.Atoi(recipe_kcal)
	return recipe_kcal_int
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

/**
* [Model][Recipes] Queries the database to get a recipe by its recipe_name
 */
func getRecipeByExactName(recipe_name string) []byte {
	row := []Recipes{}
	db := openConnDB()
	recipe_name = "'" + recipe_name + "'"
	querry := "SELECT recipe_id FROM recipes WHERE recipe_name = " + recipe_name + " ORDER BY recipe_id DESC"
	err := db.Select(&row, querry)
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
func getRecipeByCategory(category string) []byte {
	row := []Recipes{}
	db := openConnDB()
	err := db.Select(&row, "SELECT * FROM recipes WHERE category LIKE "+"'%"+category+"%'")
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
* [Model][Ingredients] Queries the database to get all ingredients
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
* [Model][Ingredients] Queries the database to get all ingredients by user_id
 */
func getIngredientByUserId(id string) []byte {
	row := []Ingredients{}
	db := openConnDB()
	err := db.Select(&row, "SELECT * FROM ingredients WHERE user_id ="+id)
	if err != nil {
		log.Fatal(err)
	}
	//var json = jsoniter.ConfigCompatibleWithStandardLibrary
	j, _ := json.Marshal(row)
	closeConnDB(db)
	return j
}

/**
* [Model][Ingredients] Queries the database to get a ingredients by its id
 */
func getIngredientAll() []byte {
	row := []Ingredients{}
	db := openConnDB()
	err := db.Select(&row, "SELECT * FROM ingredients")
	if err != nil {
		log.Fatal(err)
	}
	//var json = jsoniter.ConfigCompatibleWithStandardLibrary
	j, _ := json.Marshal(row)
	closeConnDB(db)
	return j
}

/**
* [Model][Ingredients] Queries the database to get a ingredients by its name and returns an array with them
 */
func getIngredientAllByNameNoJson(names []string) []Ingredients {
	row := []Ingredients{}
	db := openConnDB()

	size := len(names)

	for i := 0; i < size; i++ {
		aux_name := "'" + names[i] + "'"
		err := db.Select(&row, "SELECT * FROM ingredients where ingredient_name = "+aux_name)
		if err != nil {
			log.Fatal(err)
		}
	}
	//var json = jsoniter.ConfigCompatibleWithStandardLibrary
	closeConnDB(db)
	return row
}

/**
* [Model][Ingredients] Queries the database to get a ingredients by its id and returns the kcal
 */
func getIngredientKcalById(id string) int {
	db := openConnDB()
	ingredient := []Ingredients{}
	query := "SELECT * FROM ingredients where ingredient_id =" + id
	err := db.Select(&ingredient, query)
	ingredient_kcal := ingredient[0].Kcal

	if err != nil {
		log.Fatal(err)
	}
	closeConnDB(db)
	ingredient_kcal_int, _ := strconv.Atoi(ingredient_kcal)
	return ingredient_kcal_int
}

/**
* [Model][Ingredients] Receives a ingredient as parameter and inserts it in the database
 */
func insertIngredient(i Ingredients) int {
	db := openConnDB()
	tx := db.MustBegin()
	tx.NamedExec("INSERT INTO ingredients (ingredient_name, kcal, user_id) VALUES (:ingredient_name, :kcal, :user_id)", &i)
	err := tx.Commit()
	if err != nil {
		return false
	}
	closeConnDB(db)
	return true*
}

/**
* [Model][Ingredients] Receives a ingredient as parameter and updates it in the database
 */
func editIngredient(i Ingredients) bool {

	db := openConnDB()
	tx := db.MustBegin()
	tx.NamedExec("UPDATE ingredients SET ingredient_name=:ingredient_name, kcal=:kcal, user_id=:user_id WHERE ingredient_id=:ingredient_id", &i)
	//err := tx.Commit()

	//As an ingredient gets its kcal updated, so does the recipes it belongs to should

	/*row := []RecipeIngredients{}
	query := "SELECT * FROM recipeingredients where ingredient_id = " + strconv.Itoa(i.Ingredient_id)
	err := db.Select(&row, strings.ToLower(query))
	if err != nil {
		log.Fatal(err)
	}

	s := len(row)
	var r Recipes
	for j := 0; j < s; j++ {
		r.Recipe_id = row[j].Recipe_id
		kcal_recipe := getRecipeKcalByRecipeId(strconv.Itoa(row[j].Recipe_id))
		log.Println("kcal_recipe: " + strconv.Itoa(kcal_recipe))
		kcal_ingredient := getIngredientKcalById(strconv.Itoa(row[j].Ingredient_id))
		log.Println("kcal_ingredient: " + strconv.Itoa(kcal_ingredient))
		kcal_total := kcal_recipe - kcal_ingredient
		log.Println("kcal_total: " + strconv.Itoa(kcal_total))
		kcal_ingredient, _ = strconv.Atoi(i.Kcal)
		log.Println("kcal_ingredient 2: " + strconv.Itoa(kcal_ingredient))
		kcal_total = kcal_recipe + kcal_ingredient
		r.Kcal = strconv.Itoa(kcal_total)
		log.Println("kcal_total: " + strconv.Itoa(kcal_total))
		log.Println("Recipe new kcal: " + r.Kcal)
		log.Println("Recipe id: " + strconv.Itoa(r.Recipe_id))
		q := "UPDATE recipes SET kcal=:kcal WHERE recipe_id=:recipe_id"
		tx.NamedExec(q, &r)
		log.Println("querry: " + q)
		/*err = tx.Commit()
		if err != nil {
			return false
		}
	}*/

	//tx.NamedExec()
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
func editIngredientName(i Ingredients) bool {

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

/**
* [Model][Ingredients] Queries the database to get the ingredients of a recipe
 */
func getIngredientsByRecipeId(recipe_id string) []byte {
	row := []Ingredients{}
	db := openConnDB()
	recipe_id = "'" + recipe_id + "'"
	query := "SELECT * FROM ingredients WHERE ingredient_id IN(SELECT ingredient_id FROM recipeingredients WHERE recipe_id = " + recipe_id + ")"
	err := db.Select(&row, query)
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

func getRecipeByIngredientNameTotal(ingredient_name string) []byte {
	row := []Recipes{}
	db := openConnDB()
	//query := "SELECT * FROM recipeingredients WHERE ingredient_id = " + ingredient_id
	ingredient_name = "'" + ingredient_name + "'"
	query := "SELECT * FROM recipes where recipe_id IN (SELECT recipe_id FROM recipeingredients WHERE ingredient_id IN (select ingredient_id from ingredients where ingredient_name = " + ingredient_name + "))"
	err := db.Select(&row, query)
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
	//As an ingredient was added to the recipe its kcal, must be added to the kcal of the recipe
	db := openConnDB()
	tx := db.MustBegin()
	/*recipe_kcal := getRecipeKcalByRecipeId(strconv.Itoa(r.Recipe_id))
	ingredient_kcal := getIngredientKcalById(strconv.Itoa(r.Ingredient_id))

	var re Recipes
	kcal_total := recipe_kcal + ingredient_kcal
	tx := db.MustBegin()
	query := "UPDATE recipes SET kcal=" + strconv.Itoa(kcal_total) + " WHERE recipe_id= " + strconv.Itoa(r.Recipe_id)
	tx.NamedExec(strings.ToLower(query), &re)*/

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
	//As an ingredient was removed from recipeIngredients so should its kcal be removed
	db := openConnDB()
	tx := db.MustBegin()

	/*var r RecipeIngredients
	r.Recipe_id = getRecipeIngredientsByIdRecipeId(id)
	r.Ingredient_id = getRecipeIngredientsByIdIngredientId(id)
	log.Println("recipe_id: " + strconv.Itoa(r.Recipe_id))

	recipe_kcal := getRecipeKcalByRecipeId(strconv.Itoa(r.Recipe_id))
	ingredient_kcal := getIngredientKcalById(strconv.Itoa(r.Ingredient_id))
	kcal_total := recipe_kcal - ingredient_kcal
	log.Println("kcal_total: " + strconv.Itoa(kcal_total))
	var re Recipes
	tx := db.MustBegin()
	query := "UPDATE recipes SET kcal=" + strconv.Itoa(kcal_total) + " WHERE recipe_id= " + strconv.Itoa(r.Recipe_id)
	tx.NamedExec(strings.ToLower(query), &re)*/

	tx.MustExec("DELETE FROM recipeingredients WHERE recipeingredient_id=" + id)
	err := tx.Commit()
	if err != nil {
		return false
	}

	closeConnDB(db)
	return true
}

/**
* [Model][RecipeIngredients] Queries the database to get a RecipeIngredients by its id
 */
func getRecipeIngredientsByIdRecipeId(id string) int {
	db := openConnDB()
	row := []RecipeIngredients{}
	query := "SELECT * FROM recipeingredients WHERE recipeingredient_id = " + id
	err := db.Select(&row, strings.ToLower(query))
	if err != nil {
		log.Fatal(err)
	}
	reciped_id_int := row[0].Recipe_id
	closeConnDB(db)
	return reciped_id_int
}

/**
* [Model][RecipeIngredients] Queries the database to get a RecipeIngredients by its id
 */
func getRecipeIngredientsByIdIngredientId(id string) int {
	db := openConnDB()
	row := []RecipeIngredients{}
	query := "SELECT * FROM recipeingredients WHERE recipeingredient_id = " + id
	err := db.Select(&row, query)
	if err != nil {
		log.Fatal(err)
	}

	ingredient_id_int := row[0].Ingredient_id
	closeConnDB(db)
	return ingredient_id_int
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
* [Controller][User] function to edit a user name
 */
func editUserNameRoute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	//var json = jsoniter.ConfigCompatibleWithStandardLibrary
	json.NewDecoder(r.Body).Decode(&user)
	result := editUserName(user)
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
	result := checkUser(user.Email, user.Password)

	/*var u User
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

	j, _ := json.Marshal(result) //b
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
* [Controller][Recipes] function to get all recipes
 */
func getRecipeAllRoute(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	rows := getRecipeAll()
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
* [Controller][Recipes] function to edit a recipe name
 */
func editRecipeNameRoute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var recipe Recipes
	//var json = jsoniter.ConfigCompatibleWithStandardLibrary
	json.NewDecoder(r.Body).Decode(&recipe)
	result := editRecipeName(recipe)
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
	//w.WriteHeader(http.StatusOK)
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

/**
* [Controller][Recipes] function to get a recipe by recipe_name
 */
func getRecipeByExactNameRoute(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	rows := getRecipeByExactName(vars["name"])
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(rows)
}

/**
* [Controller][Recipes] function to get a recipe by category
 */
func getRecipeByCategoryRoute(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	rows := getRecipeByCategory(vars["category"])
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(rows)
}

/**
* [Controller][Recipes] function to get a recipe by ingredient name
 */
func getRecipeByIngredientNameTotalRoute(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	rows := getRecipeByIngredientNameTotal(vars["name"])
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
* [Controller][Ingredients] function to get all ingredients by user_id
 */
func getIngredientByUserIdRoute(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	rows := getIngredientByUserId(vars["id"])
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(rows)
}

/**
* [Controller][Ingredients] function to get all ingredients
 */
func getIngredientAllRoute(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	rows := getIngredientAll()
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
* [Controller][Ingredients] function to edit an ingredients name
 */
func editIngredientNameRoute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var ingredient Ingredients
	json.NewDecoder(r.Body).Decode(&ingredient)
	result := editIngredientName(ingredient)
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

/**
* [Controller][Ingredients] function to get an ingredient by ingredient_name
 */
func getIngredientsByRecipeIdRoute(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	rows := getIngredientsByRecipeId(vars["id"])
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

/**
* Pesquisar uma receita por ingredientes
 */
func getRecipeByIngredients(names []string) []byte {
	db := openConnDB()

	row := []RecipeIngredients{}
	valid_recipe := true
	s := len(names)
	recipes := getRecipeAllNoJson()
	s_recipes := len(recipes)
	ingredients := getIngredientAllByNameNoJson(names)

	valid_recipes := make([]Recipes, s)

	for i := 0; i < s_recipes; i++ {

		recipe_id := strconv.Itoa(recipes[i].Recipe_id)
		for j := 0; j < s; j++ {
			ingredient_id := strconv.Itoa(ingredients[j].Ingredient_id)
			query := "SELECT * FROM recipeingredients WHERE recipe_id = " + recipe_id + " and ingredient_id = " + ingredient_id
			err := db.Select(&row, strings.ToLower(query))

			if row == nil {
				valid_recipe = false
			}

			if err != nil {
				valid_recipe = false
				log.Fatal(err)
			}
		}
		if valid_recipe == true {
			valid_recipes = append(valid_recipes, recipes[i])
		}
	}

	j, _ := json.Marshal(valid_recipes)
	closeConnDB(db)
	return j
}

/**
* [Controller][RecipeIngredients] function to get RecipeIngredients by id
 */
func getRecipeByIngredientsRoute(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	ingredient_names := []string{"tomate", "alface", "Banana"}
	rows := getRecipeByIngredients(ingredient_names)
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
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})

	// router handlers
	//r.HandleFunc("/api/porcurar/{key}", getbyNome).Methods("GET")

	//User routes
	r.HandleFunc("/api/insertUser", insertUserRoute).Methods("POST")
	r.HandleFunc("/api/searchUser/id/{id}", getUsersByID).Methods("GET")
	r.HandleFunc("/api/deleteUser/id/{id}", deleteUserRoute).Methods("DELETE")
	r.HandleFunc("/api/editUser", editUserRoute).Methods("POST")
	r.HandleFunc("/api/editUserName", editUserNameRoute).Methods("POST")
	r.HandleFunc("/api/login", login).Methods("POST")
	r.HandleFunc("/api/allUsers", getAllUsers).Methods("GET")

	//Recipes routes
	r.HandleFunc("/api/insertRecipe", insertRecipeRoute).Methods("POST")
	r.HandleFunc("/api/searchRecipe/id/{id}", getRecipeByIdRoute).Methods("GET")
	r.HandleFunc("/api/deleteRecipe/id/{id}", deleteRecipeRoute).Methods("DELETE")
	r.HandleFunc("/api/editRecipe", editRecipeRoute).Methods("POST")
	r.HandleFunc("/api/editRecipeName", editRecipeNameRoute).Methods("POST")
	r.HandleFunc("/api/searchUserRecipe/id/{id}", getRecipeByUserIdRoute).Methods("GET")
	r.HandleFunc("/api/searchRecipeName/name/{name}", getRecipeByNameRoute).Methods("GET")
	r.HandleFunc("/api/searchRecipeExactName/name/{name}", getRecipeByExactNameRoute).Methods("GET")
	r.HandleFunc("/api/searchRecipeCategory/category/{category}", getRecipeByCategoryRoute).Methods("GET")
	r.HandleFunc("/api/searchRecipeAll", getRecipeAllRoute).Methods("GET")
	r.HandleFunc("/api/searchRecipeByIngredients", getRecipeByIngredientsRoute).Methods("GET")
	r.HandleFunc("/api/searchRecipeNameTotal/name/{name}", getRecipeByIngredientNameTotalRoute).Methods("GET")

	//Ingredients routes
	r.HandleFunc("/api/insertIngredient", insertIngredientRoute).Methods("POST")
	r.HandleFunc("/api/searchIngredient/id/{id}", getIngredientByIdRoute).Methods("GET")
	r.HandleFunc("/api/deleteIngredient/id/{id}", deleteIngredientRoute).Methods("DELETE")
	r.HandleFunc("/api/editIngredient", editIngredientRoute).Methods("POST")
	r.HandleFunc("/api/editIngredientName", editIngredientNameRoute).Methods("POST")
	r.HandleFunc("/api/searchIngredientName/name/{name}", getIngredientByNameRoute).Methods("GET")
	r.HandleFunc("/api/searchIngredientAll", getIngredientAllRoute).Methods("GET")
	r.HandleFunc("/api/getIngredientsByRecipeId/id/{id}", getIngredientsByRecipeIdRoute).Methods("GET")
	r.HandleFunc("/api/getIngredientByUserIdRoute/id/{id}", getIngredientByUserIdRoute).Methods("GET")

	//RecipeIngredients routes
	r.HandleFunc("/api/insertRecipeIngredients", insertRecipeIngredientsRoute).Methods("POST")
	r.HandleFunc("/api/searchRecipeIngredients/id/{id}", getRecipeIngredientsByIdRoute).Methods("GET")
	r.HandleFunc("/api/deleteRecipeIngredients/id/{id}", deleteRecipeIngredientsRoute).Methods("DELETE")
	r.HandleFunc("/api/editRecipeIngredients", editRecipeIngredientsRoute).Methods("POST")
	r.HandleFunc("/api/searchRecipeIngredientsName/name/{name}", getRecipeIngredientsByIngredientRoute).Methods("GET")
	r.HandleFunc("/api/searchRecipeRecipesName/name/{name}", getRecipeIngredientsByRecipeRoute).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(corsObj, headersOk, methodsOk)(r))) // se falhar dá erro !*/
}
