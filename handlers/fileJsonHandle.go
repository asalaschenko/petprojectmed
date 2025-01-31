package handlers

import (
	"encoding/json"
	"log"
	"os"
	"petprojectmed/models"
)

func ReadDoctorsJsonFile() []models.Doctor {
	file, err := os.Open("./tables/json/doctors.json")
	if err != nil {
		log.Printf("Ошибка открытия файла: %v\n", err)
	}
	defer file.Close()

	var doctors []models.Doctor
	decoder := json.NewDecoder(file)

	err = decoder.Decode(&doctors)
	if err != nil {
		log.Fatalf("Ошибка декодирования: %v", err)
	}

	return doctors
}

func WriteDoctorsJsonFile(doctors []models.Doctor) {
	file, _ := os.Create("./tables/json/doctors.json")
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(doctors); err != nil {
		log.Fatalf("Ошибка кодирования: %v", err)
	}
}

func ReadPatientsJsonFile() []models.Patient {
	file, err := os.Open("./tables/json/patients.json")
	if err != nil {
		log.Printf("Ошибка открытия файла: %v\n", err)
	}
	defer file.Close()

	var patients []models.Patient
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&patients); err != nil {
		log.Fatalf("Ошибка декодирования: %v", err)
	}

	return patients
}

func WritePatientsJsonFile(doctors []models.Patient) {
	file, _ := os.Create("./tables/json/patients.json")
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(doctors); err != nil {
		log.Fatalf("Ошибка кодирования: %v", err)
	}
}
