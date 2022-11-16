package main

import (
	"fmt"
    "log"
    "os"
	"math/rand"
	"database/sql"
    "github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type user_data struct{
	Username string
	Uuid int 	
}

func generate_uid(low, hi int) int {
    return low + rand.Intn(hi-low)
}

func main() {
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }
  host     := os.Getenv("POSTGRES_HOST")
  user     := os.Getenv("POSTGRES_USER")
  password := os.Getenv("POSTGRES_PASSWORD")
  dbname   := os.Getenv("POSTGRES_DB")
  port     := 5432
  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)
  db, err := sql.Open("postgres", psqlInfo)
  if err != nil {
    panic(err)
  }
  defer db.Close()

  err = db.Ping()
  if err != nil {
    panic(err)
  }
  var users []user_data

  for i:=1; i<10; i++ {
	user_data := user_data{
		Username: "test",
		Uuid: generate_uid(100, 10000),
	}
	users = append(users, user_data)
  }
  for _, user := range users {
	result, err := db.Exec("INSERT INTO bot_data (username, uuid) VALUES ($1, $2)", user.Username, user.Uuid)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.RowsAffected())
  }
  fmt.Println("Successfully connected!")
}