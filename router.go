package main

import (
	"crypto/rand"
	"crypto/sha1"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func logger(next http.Handler) (result http.Handler) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		randomBytes := make([]byte, 4)
		rand.Read(randomBytes)

		hasher := sha1.New()
		hasher.Write(randomBytes)
		requestId := fmt.Sprintf("%x", hasher.Sum(nil)[hasher.Size()-4:])

		log.Printf("START [%s] %s %s\n", requestId, r.Method, r.URL)
		next.ServeHTTP(w, r)
		log.Printf("END [%s]\n", requestId)
	})
}

func getRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/edit", handleEdit).Methods("GET")
	router.HandleFunc("/", handleGet).Methods("GET")
	router.HandleFunc("/", handlePost).Methods("POST")

	router.Use(logger)

	return router
}
