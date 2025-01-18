package main

import (
	"errors"
	"log"
	"net/http"
	"os"
	"petprojectmed/handlers"
)

var err_env error = errors.New("not found port for current application")

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	port, exists := os.LookupEnv("PORT_GOLANG")
	if !exists {
		check(err_env)
	}

	http.HandleFunc("/", handlers.RedirectHandler)
	http.HandleFunc("/main", handlers.MainHandler)
	http.HandleFunc("/doctors/list", handlers.ListHandler)
	http.HandleFunc("/patients/list", handlers.ListHandler)
	err := http.ListenAndServe("localhost:"+port, nil)
	log.Fatal(err)
}
