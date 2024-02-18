package api

import (
	"fmt"
	"learn-word/controllers"
	"learn-word/services"
	"net/http"

	"github.com/gorilla/mux"
)

func rootPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Go Api Server")
	fmt.Println("Root endpoint is hooked!")
}

func StartWebServer() error {
	router := mux.NewRouter().StrictSlash(true)

	wordController := controllers.NewWordController(services.NewGetWordService())
	router.HandleFunc("/", rootPage)
	router.HandleFunc("/v1/learn/word/all", wordController.FetchAllWords)

	return http.ListenAndServe(fmt.Sprintf(":%d", 8080), router)
}
