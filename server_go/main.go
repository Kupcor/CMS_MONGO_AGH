package main

import (
    "fmt"
    "log"
    "net/http"

    "historycznymonolog/router"
    
    "github.com/gorilla/handlers"
    "go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client

func main() {
    router := router.Router()

    //  Added to allow multiple request types from the same endpoint
    headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
    methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "DELETE", "OPTIONS"})
    origins := handlers.AllowedOrigins([]string{"*"})

    log.Fatal(http.ListenAndServe(":8000", handlers.CORS(headers, methods, origins)(router)))    
    fmt.Println("Server started on port 8000...")
}
