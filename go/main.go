package main

import (
	"context"
	"fmt"
	"frascati/config"
	"frascati/routing"
	"frascati/setup"
	"log"
	"net/http"
	"os/signal"
	"syscall"
)

func main() {
	app, err := setup.SetupApp()
	if err != nil {
		log.Fatalf("app start error: %v", err.ToMap())
	}
	defer func() {
		if closeErr := app.CloseComp(); closeErr != nil {
			log.Fatalf("cannot close app: %v", closeErr.ToMap())
		}
	}()

	app.Open()
	router := routing.SetupRouter(app)

	grpcServer, netListener := setup.SetupGrpc(app)
	go func() {
		if err := grpcServer.Serve(netListener); err != nil {
			log.Fatalf("failed to serve: %v", err)
		} else {
			log.Println("grpc serve success")
		}
	}()

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", config.GetServerPort()),
		Handler: router,
	}

	go func() {
		// router.Run(fmt.Sprintf(":%s", config.GetServerPort()))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("serve run error: %v", err)
		}
	}()

	shutdown, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-shutdown.Done()
	closeApp(app, srv)
}

func closeApp(app setup.App, server *http.Server) {
	log.Println("shutdown signal received. initiating shutdown")

	serverClosedSig := make(chan struct{})
	defer close(serverClosedSig)

	appClosedSig := make(chan struct{})
	defer close(appClosedSig)

	gateClosedSig := make(chan struct{})
	defer close(gateClosedSig)

	log.Println("waiting for last operations to be complete")

	go func() {
		if err := app.Close(appClosedSig, serverClosedSig, gateClosedSig); err != nil {
			log.Fatalf("app close error: %v", err.ToMap())
		}
	}()

	<-gateClosedSig

	if err := server.Shutdown(context.Background()); err != nil {
		serverClosedSig <- struct{}{}
		log.Fatalf("shutdown fail: %v", err)
	}

	serverClosedSig <- struct{}{}
	log.Println("server shutdown complete")

	<-appClosedSig
	log.Println("app shutdown complete")
}
