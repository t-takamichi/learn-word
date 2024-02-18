package services

import "context"

type GetWordService interface {
	GetWordAll(ctx context.Context) []Word
}

type getWordService struct{}

func NewGetWordService() GetWordService {
	return &getWordService{}
}

func (*getWordService) GetWordAll(ctx context.Context) []Word {
	words := []Word{}
	words = append(words,
		Word{
			Id:         "1",
			Vocabulary: "learn",
			Mean:       "稼ぐ",
		},
		Word{
			Id:         "2",
			Vocabulary: "study",
			Mean:       "勉強する",
		},
	)
	return words
}

type Word struct {
	Id         string `json:"id"`
	Vocabulary string `json:"vocabulary"`
	Mean       string `json:"mean"`
}
