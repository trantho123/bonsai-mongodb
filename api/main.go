package main

import (
	"context"
	"fmt"
	"log"

	"github.com/trantho123/saleswebsite/repo"
	"github.com/trantho123/saleswebsite/server.go"
	"github.com/trantho123/saleswebsite/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	fmt.Println("Starting the server on port 8080...")
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatalf("cannot load config: %v", err)
	}
	client := ConnectDB(config.DBSource)
	repo := repo.NewRepo(client.Database(config.DBName))
	server := server.NewServer(repo, &config)
	server.Run(config.HTTPServerPort)
}

func ConnectDB(uri string) *mongo.Client {
	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("Không thể kết nối với MongoDB:", err)
	}
	fmt.Println("Kết nối MongoDB thành công!")
	return client
}
