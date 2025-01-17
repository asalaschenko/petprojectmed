package cruds

import (
	"log"
	"net/http"
	"net/url"
	"petprojectmed/DAO"
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
	log.Println("u.Query():", params)

	urlArg := params["id"]
	if len(urlArg) == 0 {
		_, err := writer.Write([]byte("все доктора"))
		if err != nil {
			log.Fatal(err)
		}
	} else {
		idArray := ConvertToIdArray(urlArg)
		header := DAO.SelectHeader("doctors")
		body := ""
		for _, value := range idArray {
			body += value.Select("doctors")
		}
		outputString := wrapperTable(header, body)

		_, err := writer.Write([]byte(outputString))
		if err != nil {
			log.Fatal(err)
		}
	}
}
