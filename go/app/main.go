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
	"github.com/json-iterator/go"
	_ "github.com/lib/pq"*/)

type Users struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Recipes struct {
	Id      int    `json:"id"`
	User_id int    `json:"user_id"`
	Name    string `json:"name"`
}

type IngredientsList struct {
	Recipe_id      int `json:"recipe_id"`
	Ingredient_id  int `json:"ingredient_id"`
	Ingredient_qty int `json:"recipe_qty"`
}

type Ingredients struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type DirectionsList struct {
	Recipe_id        int    `json:"recipe_id"`
	Direction_text   string `json:"direction_text"`
	Direction_number int    `json:"direction_number"`
}

//Variaveis fixas de postgresBD
const (
	dbHost     = "postgres"
	dbUser     = "docker"
	dbPassword = "docker"
	dbName     = "recipes"
)

func main() {
	var i Ingredients
	i.Id = 1
	i.Name = "Paulo"
	fmt.Println("O nome deste objeto é: "+i.Name+" e o ID: %d", i.Id)
}
