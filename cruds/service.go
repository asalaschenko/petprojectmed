package cruds

import (
	"log"
	"net/http"
	"net/url"
	"petprojectmed/DAO"
	"strings"
)

func List(writer http.ResponseWriter, request *http.Request) {
	u, err := url.Parse(request.URL.String())
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Internal server error"))
		return
	}

	log.Println("url.Parse(request.URL.String()): ", u)
	params := u.Query()
	url := u.RequestURI()
	log.Println(url)
	array := strings.Split(url, "/")
	databaseFileName := array[1]
	log.Println(databaseFileName)
	log.Println("u.Query():", params)

	urlArg := params["id"]
	if len(urlArg) == 0 {
		count := DAO.GetFileDatabaseSize(databaseFileName)
		header := DAO.SelectHeader(databaseFileName)
		body := ""
		for i := 0; i < count; i++ {
			index := DAO.ID_DB(i)
			body += index.Select(databaseFileName)
		}
		outputString := wrapperTable(header, body)
		_, err := writer.Write([]byte(outputString))
		if err != nil {
			log.Fatal(err)
		}
	} else {
		idArray := ConvertToIdArray(urlArg)
		header := DAO.SelectHeader(databaseFileName)
		body := ""
		for _, value := range idArray {
			body += value.Select(databaseFileName)
		}
		outputString := wrapperTable(header, body)

		_, err := writer.Write([]byte(outputString))
		if err != nil {
			log.Fatal(err)
		}
	}
}
