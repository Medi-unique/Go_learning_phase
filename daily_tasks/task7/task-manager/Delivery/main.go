package main

import (
    "context"
    "fmt"
    "log"
    "time"
    "task-manager/Delivery/routers"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
    client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
    if err != nil {
        log.Fatal(err)
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    err = client.Connect(ctx)
    if err != nil {
        log.Fatal(err)
    }

    defer client.Disconnect(ctx)

    router := routers.SetupRouter()
    fmt.Println("Starting server on port 8080...")
    log.Fatal(router.Run(":8080"))
}