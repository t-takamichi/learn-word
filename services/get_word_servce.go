package services

import (
	"context"
	"learn-word/domains"
)

type GetWordService interface {
	GetWordAll(ctx context.Context) []domains.WordDomain
}

type getWordService struct{}

func NewGetWordService() GetWordService {
	return &getWordService{}
}

func (*getWordService) GetWordAll(ctx context.Context) []domains.WordDomain {
	words := []domains.WordDomain{}
	words = append(words,
		domains.WordDomain{
			Id:         "1",
			Vocabulary: "learn",
			Mean:       "稼ぐ",
		},
		domains.WordDomain{
			Id:         "2",
			Vocabulary: "study",
			Mean:       "勉強する",
		},
	)
	return words
}
