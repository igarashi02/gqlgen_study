package main

import (
  "fmt"
	"app/graph"
	"app/graph/generated"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

  "app/internal/models"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)

const defaultPort = "8080"

const (
  Dialect = "mysql"
  DBUser = "root"
  DBPass = "password"
  // DBProtocol = "tcp(127.0.0.1:3306)"
  DBProtocol = "tcp(mysql:3306)"
  DBName = "gql_database"
)

func DBConnect() *gorm.DB {
  // connectTemplate := "%s:%s@%s/%s"
  // connect := fmt.Sprintf(connectTemplate, DBUser, DBPass, DBProtocol, DBName)
  connect := DBUser + ":" + DBPass + "@" + DBProtocol + "/" + DBName + "?parseTime=true"
  db, err := gorm.Open(Dialect, connect)

  if err != nil {
    fmt.Println(err)
  }
  return db
}

func DBMigrate(db *gorm.DB) *gorm.DB {
  db.AutoMigrate(&models.Todo{})
  // db.AutoMigrate(&User{})
  return db
}

func main() {
  db := DBConnect()
  defer db.Close()
  DBMigrate(db)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

  srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{ DB: db }}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
