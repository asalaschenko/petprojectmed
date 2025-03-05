package patients

import (
	"context"
	"petprojectmed/common"
	"petprojectmed/storage"
	"slices"
	"sort"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func (val *patientsID) GetList() (*[]storage.Patient, string) {
	conn := storage.GetConnectionDB()
	defer conn.Close(context.Background())

	sort.Ints(*val)
	*val = common.RemoveDuplicateInt(*val)

	arrayIndex := []int{}
	arrayIDs := storage.GetIDofPatients(conn)

	for _, value := range *val {
		if slices.Contains(*arrayIDs, value) {
			arrayIndex = append(arrayIndex, value)
		}
	}

	if len(arrayIndex) == 0 {
		return nil, common.NOT_FOUND
	} else {
		return storage.GetPatientsByID(conn, []int(*val)), common.OK
	}
}

func (val *QueryPatientsListFilter) GetList() (*[]storage.Patient, string) {
	conn := storage.GetConnectionDB()
	defer conn.Close(context.Background())

	switch val.List {
	case "all":
		return storage.GetAllPatients(conn), common.OK
	case "filter":
		resultArray := [][]int{}
		flag := false

		if len(val.PhoneNumbers) != 0 && val.PhoneNumbers[0] != "" {
			flag = true

			for index, _ := range val.PhoneNumbers {
				common.TransformCharsForPhoneNumber(&val.PhoneNumbers[index])
			}

			m := storage.GetPatientIDandPhoneNumbers(conn)

			arrayIndex := []int{}
			for _, value := range val.PhoneNumbers {
				arrayIndex = append(arrayIndex, common.ReturnIndexOfTargetFilterValueString(value, m)...)
			}
			sort.Ints(arrayIndex)
			arrayIndex = common.RemoveDuplicateInt(arrayIndex)
			resultArray = append(resultArray, arrayIndex)
		}

		if len(val.DatesOfBirth) != 0 && val.DatesOfBirth[0] != "" {
			flag = true

			for index, _ := range val.DatesOfBirth {
				common.TransformCharsForDateofBirth(&val.DatesOfBirth[index])
			}

			m := storage.GetPatientIDandDateOfBirth(conn)

			arrayIndex := []int{}
			for _, value := range val.DatesOfBirth {
				if b, layout := common.CheckAndParseDateValueForFilter(value); b {
					date := common.ReturnDateFormat(value, layout)
					arrayIndex = append(arrayIndex, common.ReturnIndexOfTargetDateOfBirth(date, m, layout)...)
				}
			}

			sort.Ints(arrayIndex)
			arrayIndex = common.RemoveDuplicateInt(arrayIndex)
			resultArray = append(resultArray, arrayIndex)
		}

		if !flag {
			return nil, common.FILTER_EMPTY
		}

		outputArray := common.FindIntersectionOfSetsValues(resultArray)

		if len(outputArray) == 0 {
			return nil, common.NOT_FOUND
		}

		return storage.GetPatientsByID(conn, outputArray), common.OK
	default:
		return nil, common.INVALID_REQUEST
	}
}

func (val *Patient) Create() string {
	conn := storage.GetConnectionDB()
	defer conn.Close(context.Background())

	entries := storage.GetPhoneNumberOfPatients(conn)
	if slices.Contains(*entries, val.PhoneNumber) {
		return common.INVALID_PHONE_NUMBER
	}

	caserT := cases.Title(language.Russian)
	caserL := cases.Lower(language.Russian)
	_, layout := common.CheckAndParseDateValue(val.DateOfBirth)
	date := common.ReturnDateFormat(val.DateOfBirth, layout)
	patient := storage.NewPatient(caserT.String(val.Name), caserT.String(val.Family), val.PhoneNumber, caserL.String(val.Gender), date)

	storage.InsertPatient(conn, patient)
	return common.OK
}

func (val *PatientU) Update(ID int) string {
	conn := storage.GetConnectionDB()
	defer conn.Close(context.Background())

	caserT := cases.Title(language.Russian)
	caserL := cases.Lower(language.Russian)

	if val.Name != "" {
		storage.UpdatePatientNameByID(conn, ID, caserT.String(val.Name))
	}
	if val.Family != "" {
		storage.UpdatePatientNameByID(conn, ID, caserT.String(val.Family))
	}
	if val.Gender != "" {
		storage.UpdatePatientGenderByID(conn, ID, caserL.String(val.Gender))
	}
	if val.DateOfBirth != "" {
		_, layout := common.CheckAndParseDateValue(val.DateOfBirth)
		date := common.ReturnDateFormat(val.DateOfBirth, layout)
		storage.UpdatePatientDateOfBirthByID(conn, ID, date)
	}
	if val.PhoneNumber != "" {
		entries := storage.GetPhoneNumberOfPatients(conn)
		if slices.Contains(*entries, val.PhoneNumber) {
			return common.INVALID_PHONE_NUMBER
		}

		storage.UpdatePatientPhoneNumberByID(conn, ID, val.PhoneNumber)
	}
	return common.OK
}

func (ID *patientID) Delete() (string, *storage.Patient) {
	conn := storage.GetConnectionDB()
	defer conn.Close(context.Background())

	output := *storage.GetPatientsByID(conn, []int{int(*ID)})
	storage.DeletePatientByID(conn, int(*ID))

	return common.OK, &output[0]
}
