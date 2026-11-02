package controllers

import (
	"api-go-rest/database"
	"api-go-rest/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "home page")
}

func Personalities(w http.ResponseWriter, r *http.Request) {
	var p []models.Personality
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func FindPersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personality models.Personality
	database.DB.First(&personality, id)
	json.NewEncoder(w).Encode(personality)
}

func CreatePersonality(w http.ResponseWriter, r *http.Request) {
	var personality models.Personality

	json.NewDecoder(r.Body).Decode(&personality)

	database.DB.Create(&personality)

	json.NewEncoder(w).Encode(personality)
}

func DeletePersonality(w http.ResponseWriter, r *http.Request) {
	var personality models.Personality
	vars := mux.Vars(r)
	id := vars["id"]

	database.DB.Delete(&personality, id)

	json.NewEncoder(w).Encode(personality)
}

func EditPersonality(w http.ResponseWriter, r *http.Request) {
	var personality models.Personality
	vars := mux.Vars(r)
	id := vars["id"]

	database.DB.First(&personality, id)

	json.NewDecoder(r.Body).Decode(&personality)
	database.DB.Save(&personality)
	json.NewEncoder(w).Encode(personality)
}
