package bootstrap

import (
	"abduselam-arabianmejlis/mongo"
	"context"
	"fmt"
	"log"
	"time"
)

func NewMongoDatabase(env *Env) mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dbHost := env.DBHost
	dbPort := env.DBPort
	dbUser := env.DBUser
	dbPass := env.DBPass

	mongodbURI := fmt.Sprintf("mongodb://%s:%s@%s:%s", dbUser, dbPass, dbHost, dbPort)

	if dbUser == "" || dbPass == "" {
		mongodbURI = fmt.Sprintf("mongodb://%s:%s", dbHost, dbPort)
	}

	client, err := mongo.NewClient(mongodbURI)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func CloseMongoDBConnection(client mongo.Client) {
	if client == nil {
		return
	}

	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connection to MongoDB closed.")
}


func CreateTextIndex(db mongo.Database, colName string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := db.Collection(colName)

	// List all indexes on the collection
	cursor, err := collection.Indexes().List(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)

	indexExists := false

	for cursor.Next(ctx) {
		var index bson.M
		if err := cursor.Decode(&index); err != nil {
			log.Fatal(err)
		}
		// Check if the index name matches the one we want to create
		if index["name"] == "author_text_title_text_content_text" {
			indexExists = true
			break
		}
	}

	if !indexExists {
		indexModel := mongo.IndexModel{Keys: bson.D{{Key: "author", Value: "text"}, {Key: "title", Value: "text"}, {Key: "content", Value: "text"}}}

		name, err := collection.Indexes().CreateOne(ctx, indexModel)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Text index created successfully with name:", name)
	} else {
		fmt.Println("Index already exists, skipping creation.")
	}
}