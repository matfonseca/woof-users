package main

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func connect() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"<mongo_url>",
	))
	if err != nil { log.Fatal(err) }
	return client
}

func insertOne(user User){
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client := connect()
	test_database := client.Database("test")
	users_collection := test_database.Collection("users")

	_, err := users_collection.InsertOne(ctx, user)

	if err != nil { log.Fatal(err) }
}

func findOne(id string) User {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client := connect()

	testDatabase := client.Database("test")
	usersCollection := testDatabase.Collection("users")

	objId, _ :=  primitive.ObjectIDFromHex(id)
	result := usersCollection.FindOne(ctx, bson.M{"_id": objId})
	user := User{}
	result.Decode(&user)
	return user
}
