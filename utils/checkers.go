package utils

import (
	"errors"
	"log"
	"path/filepath"
	"runtime"
	"strconv"
)

func CheckErr(err error) {
	if err != nil {
		pc, fileName, line, _ := runtime.Caller(1)
		log.Println(GenLocationError(runtime.FuncForPC(pc).Name(), fileName, strconv.Itoa(line)))
		log.Fatalln(err)
	}
}

func GenLocationError(extFunction string, fileName string, lineOfCallCheckErr string) error {
	outputString := extFunction + ", " + filepath.Base(fileName) + ", " + "line of triggering error handler: " + lineOfCallCheckErr
	outputErr := errors.New(outputString)
	return outputErr
}
