package data

import (
	"context"
	"errors"
	"fmt"
	"log"

	"example.com/task_manager/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var DB *mongo.Database
var Collection *mongo.Collection

func StartMongoDB() {
	// Set client options for local MongoDB
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	var err error
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connected to MongoDB!")
	}

	DB = client.Database("taskManager")
	Collection = DB.Collection("tasks")

	if err = client.Ping(context.TODO(), nil); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
	}

	
}

var Data = map[string]models.Task{}

func AddTask(task models.Task) (*mongo.InsertOneResult, error) {
	result, err := Collection.InsertOne(context.TODO(), task)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetAllTask() ([]models.Task, error) {
	findOptions := options.Find()

	var data []models.Task
	var task models.Task

	cursor, err := Collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		return []models.Task{}, err
	}

	for cursor.Next(context.TODO()) {
		err := cursor.Decode(&task)
		if err != nil {
			return []models.Task{}, err
		}
		data = append(data, task)
	}
	return data, nil
}

func GetTask(id string) (models.Task, error) {
	var task models.Task
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Task{}, err
	}
	err = Collection.FindOne(context.TODO(), bson.D{{Key: "_id", Value: ID}}).Decode(&task)
	if err != nil {
		return models.Task{}, err
	}
	return task, nil
}

func DeleteTask(id string) (models.Task, error) {
	var task models.Task
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Task{}, err
	}
	err = Collection.FindOneAndDelete(context.TODO(), bson.D{{Key: "_id", Value: ID}}).Decode(&task)
	if err != nil {
		return models.Task{}, err
	}
	return task, nil
}

func UpdateTask(task models.Task, id string) (models.Task, error) {
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Task{}, err
	}

	updateFields := bson.M{}
	if task.Name != "" {
		updateFields["name"] = task.Name
	}
	if task.Detail != "" {
		updateFields["detail"] = task.Detail
	}
	if task.Start != "" {
		updateFields["start"] = task.Start
	}
	if task.Duration != "" {
		updateFields["duration"] = task.Duration
	}

	// Check if there are fields to update
	if len(updateFields) == 0 {
		return models.Task{}, errors.New("no fields to update")
	}

	// Define the filter and update
	filter := bson.M{"_id": ID}
	update := bson.M{"$set": updateFields}

	// Perform the update operation
	result, err := Collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return models.Task{}, err
	}

	// Send the response
	if result.MatchedCount == 0 {
		return models.Task{}, errors.New("document not found")
	}
	return task, nil
}
