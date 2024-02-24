package repositories

import (
	"context"
	"learn-word/collections"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type WordRepository interface {
	FetchAll(ctx context.Context) []collections.Word
	FetchById(ctx context.Context, id string) collections.Word
}

type wordRepository struct {
	client *mongo.Client
}

func NewWordRepository(client *mongo.Client) WordRepository {
	return &wordRepository{client}
}

func (repository *wordRepository) FetchAll(ctx context.Context) []collections.Word {
	collection := repository.client.Database("learn-world").Collection("words")
	if collection == nil {
		return nil
	}
	filter := bson.M{}

	// オフセット設定
	skip := int64(0)    // 最初のx件をスキップ
	limit := int64(100) // 件数X件取得する

	// オプション設定
	opts := options.FindOptions{
		Skip:  &skip,
		Limit: &limit,
		Sort:  bson.D{{Key: "_id", Value: 1}},
	}

	cur, err := collection.Find(ctx, filter, &opts)
	if err != nil {
		log.Fatalln(err)
	}

	var resultList []collections.Word
	for cur.Next(ctx) {
		var result collections.Word
		if err = cur.Decode(&result); err != nil {
			log.Fatalln(err)
			return nil
		}
		resultList = append(resultList, result)

		log.Println(result)
	}

	return resultList
}

func (repository *wordRepository) FetchById(ctx context.Context, id string) collections.Word {

	return collections.Word{}
}
