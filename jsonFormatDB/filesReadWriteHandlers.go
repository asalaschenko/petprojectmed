package jsonFormatDB

import (
	"encoding/json"
	"log"
	"os"
)

func ReadDoctorsJsonFile() []Doctor {
	file, err := os.Open("./tables/json/doctors.json")
	if err != nil {
		log.Printf("Ошибка открытия файла: %v\n", err)
	}
	defer file.Close()

	var doctors []Doctor
	decoder := json.NewDecoder(file)

	err = decoder.Decode(&doctors)
	if err != nil {
		log.Fatalf("Ошибка декодирования: %v", err)
	}

	return doctors
}

func WriteDoctorsJsonFile(doctors []Doctor) {
	file, _ := os.Create("./tables/json/doctors.json")
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(doctors); err != nil {
		log.Fatalf("Ошибка кодирования: %v", err)
	}
}

func ReadPatientsJsonFile() []Patient {
	file, err := os.Open("./tables/json/patients.json")
	if err != nil {
		log.Printf("Ошибка открытия файла: %v\n", err)
	}
	defer file.Close()

	var patients []Patient
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&patients); err != nil {
		log.Fatalf("Ошибка декодирования: %v", err)
	}

	return patients
}

func WritePatientsJsonFile(doctors []Patient) {
	file, _ := os.Create("./tables/json/patients.json")
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(doctors); err != nil {
		log.Fatalf("Ошибка кодирования: %v", err)
	}
}

func ReadScheduleJsonFile() []Appointment {
	file, err := os.Open("./tables/json/schedule.json")
	if err != nil {
		log.Printf("Ошибка открытия файла: %v\n", err)
	}
	defer file.Close()

	var appointments []Appointment
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&appointments); err != nil {
		log.Fatalf("Ошибка декодирования: %v", err)
	}

	return appointments
}

func WriteScheduleJsonFile(appointments []Appointment) {
	file, _ := os.Create("./tables/json/schedule.json")
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(appointments); err != nil {
		log.Fatalf("Ошибка кодирования: %v", err)
	}
}
