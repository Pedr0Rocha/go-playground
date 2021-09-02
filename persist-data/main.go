package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/go-sql-driver/mysql"
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
	DATABASE_HOST     = "127.0.0.1:3308"
	DATABASE_USER     = "root"
	DATABASE_PASSWORD = "admin"
	DATABASE_DATABASE = "users_test_data"
)

var db *sql.DB

func main() {
	_, err := GetUsersFromJson()
	if err != nil {
		log.Fatal(err)
	}

	ConnectToDatabase()
}

func ConnectToDatabase() {
	config := mysql.Config{
		User:                 DATABASE_USER,
		Passwd:               DATABASE_PASSWORD,
		Net:                  "tcp",
		Addr:                 DATABASE_HOST,
		DBName:               DATABASE_DATABASE,
		AllowNativePasswords: true,
	}

	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected!")
}

func SaveUsers() {

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

func PrettyPrintUsers(users []User) {
	for i := 0; i < len(users); i++ {
		fmt.Println("#" + strconv.Itoa(i))
		fmt.Println("name: " + users[i].Name)
		fmt.Println("email: " + users[i].Email)
		fmt.Println("company: " + users[i].Company)
		fmt.Println("created: " + users[i].Created + "\n")
	}
}
