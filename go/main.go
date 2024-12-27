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
		log.Fatalln("cannot start db, err: ", err.Error())
	}
	defer func() {
		err := db.Close()
		if err != nil {
			log.Fatalln("cannot close db")
		}
	}()

	warnFile, errorFile := prep.PrepFile()
	logger := setup.SetupEnhanceLogger(warnFile, errorFile)

	router := routing.SetupRouter(logger)
	setup.SetupApplication(router, db)

	router.Run(fmt.Sprintf(":%s", config.GetServerPort()))
}
