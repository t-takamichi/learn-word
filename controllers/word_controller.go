package controllers

import (
	"encoding/json"
	"net/http"
)

type Word struct {
	Id         string `json:"id"`
	Vocabulary string `json:"vocabulary"`
	Mean       string `json:"mean"`
}

func getData() []*Word {
	words := []*Word{}
	words = append(words, &Word{
		Id:         "1",
		Vocabulary: "learn",
		Mean:       "稼ぐ",
	})
	return words
}

type WordController struct {
}

func NewWordController() *WordController {
	return &WordController{}
}

func (c *WordController) FetchAllWords(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(getData())
}
