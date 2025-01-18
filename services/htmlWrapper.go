package services

import (
	"strings"
)

func wrapperTable(header string, body string) string {
	replacer := strings.NewReplacer("_", " ")

	outputString := "<table>"

	outputString += "<thead><tr>"
	splitValues := strings.Fields(header)
	for _, columnName := range splitValues {
		outputString += "<th>" + columnName + "</th>"
	}
	outputString += "</tr></thead>"

	outputString += "<tbody>"
	splitValues = strings.Split(body, "\n")
	for _, line := range splitValues {
		outputString += "<tr>"
		splitValues2 := strings.Fields(line)
		for _, word := range splitValues2 {
			word = replacer.Replace(word)
			outputString += "<td>" + word + "</td>"
		}
		outputString += "</tr>"
	}
	outputString += "</tbody>"

	outputString += "</table>"

	return outputString
}
