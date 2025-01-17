package DAO

import (
	"bufio"
	"fmt"
	"os"
)

type ID_DB int

func GetFileDatabaseSize(filename string) int {
	var count ID_DB = 0
	recordDB := ""
	for {
		recordDB = count.Select(filename)
		if len(recordDB) <= 2 {
			return int(count)
		}
		count++
	}
}

func SelectHeader(filename string) string {
	file, err := os.Open("../files/" + filename + ".txt")
	if err != nil {
		fmt.Printf("Ошибка открытия файла: %v\n", err)
		return ""
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Printf("Ошибка закрытия файла: %v\n", err)
		}
	}()

	reader := bufio.NewReader(file)

	line, err := reader.ReadString('\n')
	if err != nil {
		if err.Error() != "EOF" {
			fmt.Printf("Ошибка чтения файла: %v\n", err)
		}
	}

	return line
}

func (index ID_DB) Select(filename string) string {
	file, err := os.Open("../files/" + filename + ".txt")
	if err != nil {
		fmt.Printf("Ошибка открытия файла: %v\n", err)
		return ""
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Printf("Ошибка закрытия файла: %v\n", err)
		}
	}()

	reader := bufio.NewReader(file)

	for range index + 1 {
		_, err := reader.ReadString('\n')
		if err != nil {
			if err.Error() != "EOF" {
				fmt.Printf("Ошибка чтения файла: %v\n", err)
			}
			break
		}
	}
	line, err := reader.ReadString('\n')
	if err != nil {
		if err.Error() != "EOF" {
			fmt.Printf("Ошибка чтения файла: %v\n", err)
		}
	}

	return line
}
