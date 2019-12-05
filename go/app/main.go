package main

import (
	//"encoding/json"
	"fmt"
	/*"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/json-iterator/go"*/
	_ "github.com/lib/pq")

	/*------------SQL DA BD------------------
	CREATE TABLE Users
	(
	  user_id INT NOT NULL,
	  user_name VARCHAR(50) NOT NULL,
	  password VARCHAR NOT NULL,
	  PRIMARY KEY (user_id),
	  UNIQUE (user_name)
	);

	CREATE TABLE Recipes
	(
	  recipe_id INT NOT NULL,
	  recipe_name VARCHAR(50) NOT NULL,
	  recipe_description VARCHAR(500) NOT NULL,
	  PRIMARY KEY (recipe_id)
	);

	CREATE TABLE Ingredients
	(
	  ingredient_id INT NOT NULL,
	  ingredient_name VARCHAR(50) NOT NULL,
	  PRIMARY KEY (ingredient_id),
	  UNIQUE (ingredient_name)
	);

	CREATE TABLE Directions
	(
	  direction_id INT NOT NULL,
	  direction_details VARCHAR(500) NOT NULL,
	  direction_order INT NOT NULL,
	  recipe_id INT NOT NULL,
	  PRIMARY KEY (direction_id),
	  FOREIGN KEY (recipe_id) REFERENCES Recipes(recipe_id)
	);

	CREATE TABLE UsersRecipes
	(
	  user_recipe_id INT NOT NULL,
	  user_id INT NOT NULL,
	  recipe_id INT NOT NULL,
	  PRIMARY KEY (user_recipe_id),
	  FOREIGN KEY (user_id) REFERENCES Users(user_id),
	  FOREIGN KEY (recipe_id) REFERENCES Recipes(recipe_id)
	);

	CREATE TABLE RecipeIngredients
	(
	  recipeIngredient_id INT NOT NULL,
	  ingredient_id INT NOT NULL,
	  recipe_id INT NOT NULL,
	  PRIMARY KEY (recipeIngredient_id),
	  FOREIGN KEY (ingredient_id) REFERENCES Ingredients(ingredient_id),
	  FOREIGN KEY (recipe_id) REFERENCES Recipes(recipe_id)
	);
	------------SQL DA BD------------------*/

type Users struct {
	User_id       int    `json:"user_id"`
	User_name     string `json:"user_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Recipes struct {
	Recipe_id      int    `json:"recipe_id"`
	Recipe_name    string `json:"recipe_name"`
	Recipe_description    string `json:"recipe_description"`
}

type Ingredients struct {
	Ingredient_id   int    `json:"ingredient_id"`
	Ingredient_name string `json:"ingredient_name"`
}

type Directions struct {
	Direction_id        int    `json:"direction_id"`
	Direction_details   string `json:"direction_details"`
	Direction_order int    `json:"direction_number"`
	Recipe_id int `json:"recipe_id"`
}

type UsersRecipes struct {
	User_recipe_id      int    `json:"user_recipe_id"`
	User_id      int    `json:"user_id"`
	Recipe_id    int `json:"recipe_id"`
}

type RecipeIngredients struct {
	Recipe_ingredient_id      int `json:"recipe_ingredient_id"`
	Ingredient_id  int `json:"ingredient_id"`
	Recipe_id int `json:"recipe_id"`
}

//Variaveis fixas de postgresBD
const (
	dbHost     = "postgres"
	dbUser     = "docker"
	dbPassword = "docker"
	dbName     = "recipes"
)

func main() {
	fmt.Println("Test")
}
