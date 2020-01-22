package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func handleEdit(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	w.Write(template)
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	_, err := os.Stat(c.file)
	if err != nil {
		http.Error(w, "file not found", http.StatusNotFound)
	}

	http.ServeFile(w, r, c.file)
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	ok := fileLock.TryAcquire()
	if !ok {
		http.Error(w, "file is locked", http.StatusConflict)
		return
	}
	defer fileLock.Release()

	_, err := os.Stat(c.file)
	if err != nil {
		http.Error(w, "file not found", http.StatusNotFound)
		return
	}

	f, err := os.Create(c.file)
	if err != nil {
		log.Panic(err)
	}
	defer f.Close()

	_, err = io.Copy(f, r.Body)
	if err != nil {
		log.Panic(err)
	}
	f.Sync()

	http.ServeFile(w, r, c.file)
}
