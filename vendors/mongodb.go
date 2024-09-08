package vendors

import (
	"context"
	"log"
	"time"

	"github.com/SamPariatIL/rrqd/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client
var Database *mongo.Database

func InitMongoDB() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	conf := config.GetConfig()

	opts := options.Client().ApplyURI(conf.MongoConfig.MongoDBInstance)
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB %s", err)
	}

	defer func() {
		if err = client.Disconnect(context.Background()); err != nil {
			log.Fatalf("Failed to disconnect from MongoDB %s", err)
		}
	}()

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB server %s", err)
	}

	log.Println("Connected to MongoDB")
	MongoClient = client
	Database = client.Database(conf.MongoConfig.MongoDBName)
}
