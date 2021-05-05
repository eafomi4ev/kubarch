package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func handlerRoot(w http.ResponseWriter, r *http.Request) {
	env, ok := os.LookupEnv("HOSTNAME")
	if !ok {
		fmt.Println("No env HOSTNAME")
	}

	fmt.Fprintf(w, "Hi on %s, I'm %s!\n", r.URL.Path, env)
}

func handlerHealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "{\"status\": \"OK\"}\n")
}

func handlerVersion(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "{\"version\": \"1.0.0\"}\n")
}

func main() {
	http.HandleFunc("/health", handlerHealthCheck)
	http.HandleFunc("/version", handlerVersion)
	http.HandleFunc("/", handlerRoot)
	time.Sleep(4 * time.Second)
	fmt.Println("Wake up!")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
