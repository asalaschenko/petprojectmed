package handlers

import (
	"log"
	"net/http"
	"net/url"
	"petprojectmed/services"
	"strings"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ListHandler(writer http.ResponseWriter, request *http.Request) {
	u, err := url.Parse(request.URL.String())
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Internal server error"))
		return
	}

	log.Println("url.Parse(request.URL.String()): ", u)
	params := u.Query()
	log.Println("u.Query():", params)
	url := u.RequestURI()
	log.Println(url)
	array := strings.Split(url, "/")
	urlIdArg := params["id"]
	log.Println(urlIdArg)
	if len(params) != 0 && len(urlIdArg) == 0 {
		http.Error(writer, http.StatusText(http.StatusNotAcceptable), 400)
		check(err)
	} else {
		outputString := services.GetDataOfTables(array[1], urlIdArg)
		_, err = writer.Write([]byte(outputString))
		check(err)
	}
}

func MainHandler(writer http.ResponseWriter, request *http.Request) {
	var message string = "<h1 style=text-align:center;><p><samp>Пет-проект Клиника</samp></p></h1>"
	message += "<p style=text-align:center;><span style=\"font-size:24px;background-color:#B8DAF8;\">Здравствуйте! Вы находитесь на главной странице.</span></p>"
	byteMsg := []byte(message)
	_, err := writer.Write(byteMsg)
	check(err)
}

func RedirectHandler(writer http.ResponseWriter, request *http.Request) {
	http.Redirect(writer, request, "/main", http.StatusFound)
}
