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
	port string
}

var (
	c        config
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

	if os.Args[1][0] == '/' {
		c.file = os.Args[1]
	} else {
		c.file = fmt.Sprintf("%s/%s", dir, os.Args[1])
	}

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "80"
	}
	c.port = port

	server := &http.Server{
		Addr:           ":" + c.port,
		MaxHeaderBytes: 1 << 32,
	}

	router := getRouter()
	http.Handle("/", router)

	fileLock.Init()

	log.Fatal(server.ListenAndServe())

	return
}
