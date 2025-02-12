package handlers

import (
	"encoding/json"
	"log"
	"os"
	"petprojectmed/dto"
)

func ReadDoctorsJsonFile() []dto.Doctor {
	file, err := os.Open("./tables/json/doctors.json")
	if err != nil {
		log.Printf("Ошибка открытия файла: %v\n", err)
	}
	defer file.Close()

	var doctors []dto.Doctor
	decoder := json.NewDecoder(file)

	err = decoder.Decode(&doctors)
	if err != nil {
		log.Fatalf("Ошибка декодирования: %v", err)
	}

	return doctors
}

func WriteDoctorsJsonFile(doctors []dto.Doctor) {
	file, _ := os.Create("./tables/json/doctors.json")
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(doctors); err != nil {
		log.Fatalf("Ошибка кодирования: %v", err)
	}
}

func ReadPatientsJsonFile() []dto.Patient {
	file, err := os.Open("./tables/json/patients.json")
	if err != nil {
		log.Printf("Ошибка открытия файла: %v\n", err)
	}
	defer file.Close()

	var patients []dto.Patient
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&patients); err != nil {
		log.Fatalf("Ошибка декодирования: %v", err)
	}

	return patients
}

func WritePatientsJsonFile(doctors []dto.Patient) {
	file, _ := os.Create("./tables/json/patients.json")
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(doctors); err != nil {
		log.Fatalf("Ошибка кодирования: %v", err)
	}
}

func ReadScheduleJsonFile() []dto.Appointment {
	file, err := os.Open("./tables/json/schedule.json")
	if err != nil {
		log.Printf("Ошибка открытия файла: %v\n", err)
	}
	defer file.Close()

	var appointments []dto.Appointment
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&appointments); err != nil {
		log.Fatalf("Ошибка декодирования: %v", err)
	}

	return appointments
}

func WriteScheduleJsonFile(appointments []dto.Appointment) {
	file, _ := os.Create("./tables/json/schedule.json")
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(appointments); err != nil {
		log.Fatalf("Ошибка кодирования: %v", err)
	}
}
