package doctors

import (
	"context"
	"log"
	"petprojectmed/common"
	"petprojectmed/storage"
	"slices"
	"sort"

	//"github.com/gofiber/fiber/v2/log"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const (
	OK              = "OK"
	NOT_FOUND       = "NOT_FOUND"
	LIST_EMPTY      = "LIST_EMPTY"
	INVALID_REQUEST = "INVALID_REQUEST"
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
		return nil, NOT_FOUND
	} else {
		return storage.GetDoctorsByID(conn, []int(*val)), OK
	}
}

func (val *QueryDoctorsListFilter) GetList() (*[]storage.Doctor, string) {
	conn := storage.GetConnectionDB()
	defer conn.Close(context.Background())

	switch val.List {
	case "all":
		return storage.GetAllDoctors(conn), OK
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

			m := storage.GetIDandSpecializations(conn)

			arrayIndex := []int{}
			for _, value := range val.Specializations {
				arrayIndex = append(arrayIndex, common.ReturnIndexOfTargetFilterValue(value, m)...)
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

			log.Println(val.DatesOfBirth)

			m := storage.GetIDandDateofBirth(conn)

			arrayIndex := []int{}
			for _, value := range val.DatesOfBirth {
				if b, layout := common.CheckDateValueForFilter(value); b == true {
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
			return nil, LIST_EMPTY
		}

		outputArray := common.FindIntersectionOfSetsValues(resultArray)

		if len(outputArray) == 0 {
			return nil, NOT_FOUND
		}

		return storage.GetDoctorsByID(conn, outputArray), OK
	default:
		return nil, INVALID_REQUEST
	}
}

func (val *Doctor) Create() string {
	caserT := cases.Title(language.Russian)
	caserL := cases.Lower(language.Russian)
	_, layout := common.CheckDateValue(val.DateOfBirth)
	date := common.ReturnDateFormat(val.DateOfBirth, layout)
	doctor := storage.NewDoctor(caserT.String(val.Name), caserT.String(val.Family), caserL.String(common.TrimSpaces(val.Specialization)), val.Cabinet, date)

	conn := storage.GetConnectionDB()
	defer conn.Close(context.Background())

	storage.InsertDoctor(conn, doctor)
	return OK
}

func (val *Doctor) Update(ID int) string {
	conn := storage.GetConnectionDB()
	defer conn.Close(context.Background())

	intID := []int{ID}
	doctor := storage.GetDoctorsByID(conn, intID)
	updateDoctors := *doctor
	updateEntryDoctor := updateDoctors[0]

	caser := cases.Title(language.Russian)

	if val.Name != "" {
		updateEntryDoctor.Name = caser.String(val.Name)
	}

	if val.Family != "" {
		updateEntryDoctor.Family = caser.String(val.Family)
	}

	if val.Specialization != "" {
		val.Specialization = common.TrimSpaces(val.Specialization)
		updateEntryDoctor.Specialization = val.Specialization
	}

	if val.DateOfBirth != "" {
		_, layout := common.CheckDateValue(val.DateOfBirth)
		date := common.ReturnDateFormat(val.DateOfBirth, layout)
		updateEntryDoctor.DateOfBirth = date
	}

	if val.Cabinet != 0 {
		updateEntryDoctor.Cabinet = val.Cabinet
	}

	storage.UpdateDoctorByID(conn, ID, &updateEntryDoctor)
	return OK
}
