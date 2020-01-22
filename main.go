package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type config struct {
	file string
	wd   string
}

var (
	c config
	fileLock lock
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, `Usage: %s <file to edit>
`, os.Args[0])
		os.Exit(1)
	}

	dir, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't get working directory\n")
		return
	}

	c.file = fmt.Sprintf("%s/%s", dir, os.Args[1])

	server := &http.Server{
		Addr:           ":8080",
		MaxHeaderBytes: 1 << 32,
	}

	router := getRouter()
	http.Handle("/", router)

	fileLock.Init()

	log.Fatal(server.ListenAndServe())

	return
}
