package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
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

func main() {
	jsonData, err := os.Open(FOLDER + BIG_TEST_FILE)
	if err != nil {
		log.Fatal(err)
	}
	defer jsonData.Close()

	byteValue, err := ioutil.ReadAll(jsonData)
	if err != nil {
		log.Fatal(err)
	}

	var users []User

	json.Unmarshal(byteValue, &users)

	for i := 0; i < len(users); i++ {
		fmt.Println("#" + strconv.Itoa(i))
		fmt.Println("name: " + users[i].Name)
		fmt.Println("email: " + users[i].Email)
		fmt.Println("company: " + users[i].Company)
		fmt.Println("created: " + users[i].Created + "\n")
	}
}
