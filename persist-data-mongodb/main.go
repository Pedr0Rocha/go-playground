package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Company string `json:"company"`
	Created string `json:"created"`
}

const (
	FOLDER           = "../test-data/"
	SMALL_TEST_FILE  = "user-test-data-small.json"
	MEDIUM_TEST_FILE = "user-test-data-med.json"
	BIG_TEST_FILE    = "user-test-data-big.json"
)

const (
	MONGODB_HOST = "mongodb://root:admin@localhost:27017"
)

func main() {
	users, err := GetUsersFromJson()
	if err != nil {
		log.Fatal(err)
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(MONGODB_HOST))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	collection := client.Database("tests").Collection("users")

	for i := 0; i < len(users); i++ {
		_, err := collection.InsertOne(context.TODO(), users[i])
		if err != nil {
			log.Fatal(err)
		}
	}
}

func GetUsersFromJson() ([]User, error) {
	jsonData, err := os.Open(FOLDER + SMALL_TEST_FILE)
	if err != nil {
		return nil, err
	}
	defer jsonData.Close()

	byteValue, err := ioutil.ReadAll(jsonData)
	if err != nil {
		return nil, err
	}

	var users []User

	json.Unmarshal(byteValue, &users)

	return users, nil
}
