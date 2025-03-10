package common

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

func CheckAndParseTimeValue(timeValue string) (bool, string) {
	_, err := time.Parse("15", timeValue)
	if err == nil {
		return true, "15"
	}
	_, err = time.Parse("15:04", timeValue)
	if err == nil {
		return true, "15:04"
	}
	_, err = time.Parse("15:04:05", timeValue)
	if err == nil {
		return true, "15:04:05"
	}
	return false, ""
}

func CheckAndParseDateValue(dateValue string) (bool, string) {
	_, err := time.Parse("2006-01-02", dateValue)
	if err == nil {
		return true, "2006-01-02"
	}
	_, err = time.Parse("02-01-2006", dateValue)
	if err == nil {
		return true, "02-01-2006"
	}
	return false, ""
}

func CheckAndParseDateValueForFilter(dateValue string) (bool, string) {
	_, err := time.Parse("2006-01-02", dateValue)
	if err == nil {
		return true, "2006-01-02"
	}

	_, err = time.Parse("2006-01", dateValue)
	if err == nil {
		return true, "2006-01"
	}

	_, err = time.Parse("2006", dateValue)
	if err == nil {
		return true, "2006"
	}

	_, err = time.Parse("01-2006", dateValue)
	if err == nil {
		return true, "01-2006"
	}

	_, err = time.Parse("02-01-2006", dateValue)
	if err == nil {
		return true, "02-01-2006"
	}

	_, err = time.Parse("2006-01-02 15", dateValue)
	if err == nil {
		return true, "2006-01-02 15"
	}

	_, err = time.Parse("2006-01-02 15:04", dateValue)
	if err == nil {
		return true, "2006-01-02 15:04"
	}

	return false, ""
}
