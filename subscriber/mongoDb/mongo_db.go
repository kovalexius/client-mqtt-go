package mongoDb

import (
	"context"
    "fmt"
    "log"
	
	. "../conf"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var connection = struct {
	ctx context.Context
	advancedTelemetry *mongo.Collection
	telemetry *mongo.Collection
}{}

func init() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://"+ 
			Conf.Database.DbUser + ":" + Conf.Database.DbPwd + "@" + 
			Conf.Database.DbURL + ":" + Conf.Database.DbPort + "/" +
			Conf.Database.DbName + "?authMechanism=SCRAM-SHA-1"))
	
	if err != nil {
		panic(err)
	}

	connection.ctx = context.Background()
	err = client.Connect(connection.ctx)
	if err != nil {
		panic(err)
	}
	
	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to MongoDB!")
	
	connection.advancedTelemetry = client.Database(Conf.Database.DbName).Collection("advancedTelemetry")
	connection.telemetry = client.Database(Conf.Database.DbName).Collection("telemetry")
}

type Record struct {
	Topic string
	Payload string
}

func InsertTelemetry(topic string, payload string) {
	//fmt.Println("topic: ", topic, "payload: ", payload)
	if connection.advancedTelemetry != nil {
		res, err := connection.advancedTelemetry.InsertOne(connection.ctx, bson.M{"Topic": topic, "Payload": payload})
		if err != nil {
			panic(err)
		}
		_ = res
	} else {
		log.Fatal("connection.collection is null")
	}
	
	if connection.telemetry != nil {
		record := Record{Topic: topic, Payload: payload}
		res, err := connection.telemetry.InsertOne(connection.ctx, record)
		if err != nil {
			panic(err)
		}
		_ = res
	} else {
		log.Fatal("connection.collection is null")
	}
}


func GetAdvancedTelemetry() []Record {
	
	result := []Record{}
	
	if connection.advancedTelemetry != nil {
		filter := bson.D{}
		cur, err := connection.advancedTelemetry.Find(connection.ctx, filter)
		if err != nil {
			panic(err)
		}
		
		defer cur.Close(connection.ctx)
		
		for cur.Next(connection.ctx) {
			var records bson.M
			var rec Record
			err := cur.Decode(&records)
			if err != nil { 
				panic(err)
			}
			
			for key, value := range records {
				if key == "Topic" {
					rec.Topic = value.(string)
				} else if key == "Payload" {
					rec.Payload = value.(string)
				}
			}

			result = append(result, rec)
		}
	} else {
		log.Fatal("connection.advancedTelemetry is null")
	}
	
	return result
}


func GetTelemetry() []Record {
	result := []Record{}
	
	if connection.telemetry != nil {
		filter := bson.D{}
		cur, err := connection.telemetry.Find(connection.ctx, filter)
		if err != nil {
			panic(err)
		}
		
		defer cur.Close(connection.ctx)
		
		for cur.Next(connection.ctx) {
			var rec Record
			err := cur.Decode(&rec)
			if err != nil { 
				panic(err)
			}
			
			result = append(result, rec)
		}
	} else {
		log.Fatal("connection.telemetry is null")
	}
	
	return result
}
