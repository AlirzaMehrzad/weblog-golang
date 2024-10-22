package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"login-register/db"
	"login-register/router"
)

func main() {
    // Initialize MongoDB connection
    client, collections := db.InitDB()

    // Ensure the connection is closed when the application stops
    defer func() {
        if err := client.Disconnect(context.TODO()); err != nil {
            log.Fatal(err)
        }
    }()

    // Initialize the router
    r := router.InitRouter(collections)

    // Start server
    srv := &http.Server{
        Handler:      r,
        Addr:         "127.0.0.1:8000",
        WriteTimeout: 15 * time.Second,
        ReadTimeout:  15 * time.Second,
    }
    log.Fatal((*srv).ListenAndServe()) // or ->   // log.Fatal(srv.ListenAndServe())
}
