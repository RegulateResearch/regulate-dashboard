package main

import (
	"fmt"
	"frascati/config"
	"frascati/prep"
	"frascati/routing"
	"frascati/setup"
	"log"
)

func main() {
	config.InitEnv()
	db, err := prep.ConnectDB()
	if err != nil {
		log.Fatalln("cannot start db")
	}
	defer func() {
		err := db.Close()
		if err != nil {
			log.Fatalln("cannot close db")
		}
	}()

	router := routing.SetupRouter()
	setup.SetupApplication(router, db)

	router.Run(fmt.Sprintf(":%s", config.GetServerPort()))
}
