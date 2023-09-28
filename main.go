package main

import (
	"log"
	"net/http"

	"github.com/Mau005/MyExpenses/configuration"
	"github.com/Mau005/MyExpenses/db"
	"github.com/Mau005/MyExpenses/router"
)

func main() {
	configuration.GenerateSecretPassword()
	db.ConnectionSqlite()

	log.Fatal(http.ListenAndServe(":8000", router.NewRouter()))
}
