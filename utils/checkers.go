package utils

import (
	"errors"
	"log"
	"path/filepath"
	"runtime"
	"strconv"
	"time"
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

func CheckTimeValue(timeValue string) (bool, time.Time) {
	parseTime, err := time.Parse("15", timeValue)
	if err == nil {
		return true, parseTime
	}
	parseTime, err = time.Parse("15:04", timeValue)
	if err == nil {
		return true, parseTime
	}
	parseTime, err = time.Parse("15:04:05", timeValue)
	if err == nil {
		return true, parseTime
	}
	return false, time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC)
}
