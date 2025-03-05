package doctors

import (
	"context"
	"log"
	"petprojectmed/common"
	"petprojectmed/storage"
	"slices"
	"sort"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func (val *doctorsID) GetList() (*[]storage.Doctor, string) {
	conn := storage.GetConnectionDB()
	defer conn.Close(context.Background())

	sort.Ints(*val)
	*val = common.RemoveDuplicateInt(*val)

	arrayIndex := []int{}
	arrayIDs := storage.GetIDofDoctors(conn)

	for _, value := range *val {
		if slices.Contains(*arrayIDs, value) {
			arrayIndex = append(arrayIndex, value)
		}
	}

	if len(arrayIndex) == 0 {
		return nil, common.NOT_FOUND
	} else {
		return storage.GetDoctorsByID(conn, []int(*val)), common.OK
	}
}

func (val *QueryDoctorsListFilter) GetList() (*[]storage.Doctor, string) {
	conn := storage.GetConnectionDB()
	defer conn.Close(context.Background())

	switch val.List {
	case "all":
		return storage.GetAllDoctors(conn), common.OK
	case "filter":
		resultArray := [][]int{}
		flag := false

		if len(val.Specializations) != 0 && val.Specializations[0] != "" {
			flag = true

			caser := cases.Lower(language.Russian)
			for index, value := range val.Specializations {
				value = common.TrimSpaces(value)
				val.Specializations[index] = caser.String(value)
			}

			m := storage.GetDoctorsIDandSpecializations(conn)

			arrayIndex := []int{}
			for _, value := range val.Specializations {
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

			m := storage.GetDoctorsIDandDateofBirth(conn)

			arrayIndex := []int{}
			for _, value := range val.DatesOfBirth {
				if b, layout := common.CheckAndParseDateValueForFilter(value); b {
					date := common.ReturnDateFormat(value, layout)
					arrayIndex = append(arrayIndex, common.ReturnIndexOfTargetDateOfBirth(date, m, layout)...)
				}
			}

			log.Println(arrayIndex)

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

		return storage.GetDoctorsByID(conn, outputArray), common.OK
	default:
		return nil, common.INVALID_REQUEST
	}
}

func (val *Doctor) Create() string {
	caserT := cases.Title(language.Russian)
	caserL := cases.Lower(language.Russian)
	_, layout := common.CheckAndParseDateValue(val.DateOfBirth)
	date := common.ReturnDateFormat(val.DateOfBirth, layout)
	doctor := storage.NewDoctor(caserT.String(val.Name), caserT.String(val.Family), caserL.String(common.TrimSpaces(val.Specialization)), val.Cabinet, date)

	conn := storage.GetConnectionDB()
	defer conn.Close(context.Background())

	storage.InsertDoctor(conn, doctor)
	return common.OK
}

func (val *DoctorU) Update(ID int) string {
	conn := storage.GetConnectionDB()
	defer conn.Close(context.Background())

	caserT := cases.Title(language.Russian)
	caserL := cases.Lower(language.Russian)

	if val.Name != "" {
		storage.UpdateDoctorNameByID(conn, ID, caserT.String(val.Name))
	}
	if val.Family != "" {
		storage.UpdateDoctorNameByID(conn, ID, caserT.String(val.Family))
	}
	if val.Specialization != "" {
		storage.UpdateDoctorNameByID(conn, ID, caserL.String(common.TrimSpaces(val.Specialization)))
	}
	if val.DateOfBirth != "" {
		_, layout := common.CheckAndParseDateValue(val.DateOfBirth)
		date := common.ReturnDateFormat(val.DateOfBirth, layout)
		storage.UpdateDoctorDateOfBirthByID(conn, ID, date)
	}
	if val.Cabinet != 0 {
		storage.UpdateDoctorCabinetByID(conn, ID, val.Cabinet)
	}
	return common.OK
}

func (ID *doctorID) Delete() (string, *storage.Doctor) {
	conn := storage.GetConnectionDB()
	defer conn.Close(context.Background())

	output := *storage.GetDoctorsByID(conn, []int{int(*ID)})
	storage.DeleteDoctorByID(conn, int(*ID))

	return common.OK, &output[0]
}
