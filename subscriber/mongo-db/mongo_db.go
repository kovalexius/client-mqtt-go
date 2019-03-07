package mongoDb

import (
	"context"
    "fmt"
    "log"
	"time"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)



func init() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://mqtt-owner:Test*123@localhost:27017/mqtt_database?authMechanism=SCRAM-SHA-1"))
	
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	
	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	
	collection := client.Database("mqtt_database").Collection("trainers")
	res, err := collection.InsertOne(ctx, bson.M{"name": "pi", "value": 3.14159})
	if err != nil {
		log.Fatal(err)
	}
	
	_ = res
	_ = collection
}
