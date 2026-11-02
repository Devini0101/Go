package routes

import (
	"api-go-rest/controllers"
	"api-go-rest/middleware"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	apiUrl := "/api/"
	r := mux.NewRouter()

	r.Use(middleware.ContentTypeMiddleware)

	r.HandleFunc("/", controllers.Home)
	r.HandleFunc(apiUrl+"personalities/{id}", controllers.FindPersonality).Methods("Get")
	r.HandleFunc(apiUrl+"personalities", controllers.Personalities).Methods("Get")
	r.HandleFunc(apiUrl+"personalities", controllers.CreatePersonality).Methods("Post")
	r.HandleFunc(apiUrl+"personalities/{id}", controllers.DeletePersonality).Methods("Delete")
	r.HandleFunc(apiUrl+"personalities/{id}", controllers.EditPersonality).Methods("Put")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"localhost:3000"}))(r)))
}
