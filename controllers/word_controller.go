package controllers

import (
	"encoding/json"
	"learn-word/services"
	"net/http"
)

type WordController struct {
	getWordService services.GetWordService
}

func NewWordController(service services.GetWordService) *WordController {
	return &WordController{service}
}

func (c *WordController) FetchAllWords(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	getWorkServce := services.NewGetWordService()
	words := getWorkServce.GetWordAll(r.Context())
	json.NewEncoder(w).Encode(words)
}
