package repo

import (
	"context"
	"log"
	"os"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var i *Imp

func TestMain(m *testing.M) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("Không thể kết nối với MongoDB:", err)
	}
	i = &Imp{
		imp: client.Database("saleswebsite"),
	}
	os.Exit(m.Run())
}
