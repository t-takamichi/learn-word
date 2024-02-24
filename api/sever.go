package api

import (
	"context"
	"fmt"
	"learn-word/configs"
	"learn-word/controllers"
	"learn-word/repositories"
	"learn-word/services"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func rootPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Go Api Server")
	fmt.Println("Root endpoint is hooked!")
}

type Trainer struct {
	Name string
	Age  int
	City string
}

func StartWebServer() error {
	router := mux.NewRouter().StrictSlash(true)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := configs.InitMongoDB(&ctx)
	if err != nil {
		log.Fatal(err)
		panic("データベースでの接続でエラーとなりました。")
	}

	defer cancel()

	// collection := client.Database("learn-world").Collection("words")

	// maxWord := collections.Word{
	// 	ID:         primitive.NewObjectID().Hex(),
	// 	Vocabulary: "max",
	// 	Meaning:    "最大値",
	// 	CreatedAt:  time.Now(),
	// 	UpdatedAt:  time.Now(),
	// }

	// collection.InsertOne(ctx, maxWord)

	repository := repositories.NewWordRepository(client)
	fmt.Println(repository.FetchAll(ctx))
	// ash := Trainer{Name: "Ash", Age: 10, City: "Pallet Town"}
	// misty := Trainer{Name: "Misty", Age: 10, City: "Cerulean City"}
	// brock := Trainer{Name: "Brock", Age: 15, City: "Pewter City"}
	// insertManyResult, err := collection.InsertMany(context.TODO(), []interface{}{ash, misty, brock})
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// fmt.Println(insertManyResult)

	// // 1つのドキュメントをupdate
	// filter := bson.D{{"name", "Ash"}}
	// update := bson.D{{"$inc", bson.D{{"age", 1}}}}
	// updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// log.Printf("Matched %d documents and updated %d documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	// cur, err := collection.Find(ctx, bson.D{{"age", 15}, {"city", "Pewter City"}})
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// for cur.Next(context.TODO()) {
	// 	fmt.Println("------------------------")
	// 	var result Trainer
	// 	if err = cur.Decode(&result); err != nil {
	// 		log.Fatalln(err)
	// 	}
	// 	fmt.Println("------------------------")
	// 	log.Println(result)
	// }
	// _ = cur.Close(context.TODO())

	wordController := controllers.NewWordController(services.NewGetWordService())
	router.HandleFunc("/", rootPage)
	router.HandleFunc("/v1/learn/word/all", wordController.FetchAllWords)

	return http.ListenAndServe(fmt.Sprintf(":%d", 8080), router)
}
