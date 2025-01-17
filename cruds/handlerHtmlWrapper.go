package cruds

import (
	"strings"
)

func wrapperTable(header string, body string) string {
	replacer := strings.NewReplacer("_", " ")

	outputString := "<table>"

	outputString += "<thead><tr>"
	splitValues := strings.Fields(header)
	for _, str := range splitValues {
		outputString += "<th>" + str + "</th>"
	}
	outputString += "</tr></thead>"

	outputString += "<tbody>"
	splitValues = strings.Split(body, "\n")
	for _, str1 := range splitValues {
		outputString += "<tr>"
		splitValues2 := strings.Fields(str1)
		for _, str2 := range splitValues2 {
			word := replacer.Replace(str2)
			outputString += "<td>" + word + "</td>"
		}
		outputString += "</tr>"
	}
	outputString += "</tbody>"

	outputString += "</table>"

	return outputString
}
