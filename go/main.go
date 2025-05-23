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
			log.Fatalf("cannot close db: %v", err)
		}
	}()

	warnFile, errorFile := prep.PrepFile()
	logger := setup.SetupEnhanceLogger(warnFile, errorFile)

	router := routing.SetupRouter(logger)
	setup.SetupApplication(router, db)

	go func() {
		grpcServer, netListener := setup.SetupGrpc(logger)
		if err := grpcServer.Serve(netListener); err != nil {
			log.Fatalf("failed to serve: %v", err)
		} else {
			log.Println("grpc serve success")
		}
	}()

	router.Run(fmt.Sprintf(":%s", config.GetServerPort()))
}
