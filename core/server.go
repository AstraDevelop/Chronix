package core

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
)

func ServePage(w http.ResponseWriter, r *http.Request) {
	pagePath := filepath.Join("pages", "index.html")

	data, err := ioutil.ReadFile(pagePath)
	if err != nil {
		http.Error(w, "Page not found", http.StatusNotFound)
		log.Printf("Error reading %s: %v", pagePath, err)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func StartServer(addr string) {
	http.HandleFunc("/", ServePage)

	fmt.Println("Starting server at", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
