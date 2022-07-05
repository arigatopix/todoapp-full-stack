package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	config "server/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Todo struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	UserId    int    `json:"userId"`
	Completed bool   `json:"completed"`
}

type User struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
}

func main() {
	dbConn := openConnection()

	// load data
	todos := LoadTodosFromJson()
	users := LoadUsersFromJson()

	arg := os.Args[1]

	if arg == "-i" {
		fmt.Println("Create Data")
		dbConn.Create(&todos)
		dbConn.Create(&users)
		return
	} else if arg == "-d" {
		fmt.Println("Destroy database")
		dbConn.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&todos)
		dbConn.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&users)
		return
	}
}

func openConnection() *gorm.DB {
	env := config.LoadENV()

	dsn := "host=localhost user=" + env.POSTGRES_USER + " password=" + env.POSTGRES_PASSWORD + " dbname=" + env.POSTGRES_DB + " port=54321 sslmode=disable TimeZone=Asia/Bangkok"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Couldn't establish database connection: %s", err)
	}

	db.AutoMigrate(&Todo{}, &User{})

	return db
}

func ParseJSON(filename string) []byte {
	jsonFile, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened ", filename)

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	return byteValue
}

func LoadTodosFromJson() []Todo {
	byteValue := ParseJSON("mock/todos.json")
	var todos []Todo
	json.Unmarshal([]byte(byteValue), &todos)
	return todos
}

func LoadUsersFromJson() []User {
	byteValue := ParseJSON("mock/users.json")
	var users []User
	json.Unmarshal([]byte(byteValue), &users)
	return users
}
