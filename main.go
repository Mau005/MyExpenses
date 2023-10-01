package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Mau005/MyExpenses/configuration"
	"github.com/Mau005/MyExpenses/db"
	"github.com/Mau005/MyExpenses/router"
)

func main() {
	err := configuration.LoadConfiguration("config.yml")
	if err != nil {
		log.Panic(err)
	}
	err = db.ConnectionDataBase()
	if err != nil {
		log.Panic(err)
	}

	//Check generate attribute in DataBase
	if len(os.Args) > 1 {
		command := os.Args[1]
		switch command {
		case "generate":
			db.GenerateCategory()
		}
	}

	log.Println("Listening Server Run My Expenses")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", configuration.Config.Server.Ip, configuration.Config.Server.Port), router.NewRouter()))
}
