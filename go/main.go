package main

import (
	"fmt"
	"frascati/config"
	"frascati/routing"
	"frascati/setup"
	"log"
)

func main() {
	app, err := setup.SetupApp()
	if err != nil {
		log.Fatalln(err.ToMap())
	}
	defer closeApp(app)

	router := routing.SetupRouter(app)

	go func() {
		grpcServer, netListener := setup.SetupGrpc(app)
		if err := grpcServer.Serve(netListener); err != nil {
			log.Fatalf("failed to serve: %v", err)
		} else {
			log.Println("grpc serve success")
		}
	}()

	router.Run(fmt.Sprintf(":%s", config.GetServerPort()))
}

func closeApp(app setup.App) {
	err := app.Close()
	if err != nil {
		log.Fatalln(err)
	}
}
