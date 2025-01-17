package main

import (
	"errors"
	"log"
	"net/http"
	"os"
	"petprojectmed/cruds"
)

var err_env error = errors.New("not found port for current application")

func mainHandler(writer http.ResponseWriter, request *http.Request) {
	var message string = "<h1 style=text-align:center;><p><samp>Пет-проект Клиника</samp></p></h1>"
	message += "<p style=text-align:center;><span style=\"font-size:24px;background-color:#B8DAF8;\">Здравствуйте! Вы находитесь на главной странице.</span></p>"
	byteMsg := []byte(message)
	_, err := writer.Write(byteMsg)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	port, exists := os.LookupEnv("PORT_GOLANG")
	if !exists {
		log.Fatal(err_env)
	}

	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/doctors/list", cruds.List)
	http.HandleFunc("/patients/list", cruds.List)
	err := http.ListenAndServe("localhost:"+port, nil)
	log.Fatal(err)
}
