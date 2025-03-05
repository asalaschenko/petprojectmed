package common

import (
	"regexp"

	"github.com/go-playground/validator/v10"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func ValidNameFamily(fl validator.FieldLevel) bool {
	return !regexp.MustCompile(`[^а-яА-Я]`).MatchString(fl.Field().String())
}

func ValidDate(fl validator.FieldLevel) bool {
	flag, _ := CheckAndParseDateValue(fl.Field().String())
	return flag
}

func ValidateTime(fl validator.FieldLevel) bool {
	flag, _ := CheckAndParseTimeValue(fl.Field().String())
	return flag
}

func ValidPhoneNumber(fl validator.FieldLevel) bool {
	return regexp.MustCompile(`^7[0-9]{10}$`).MatchString(fl.Field().String())
}

func ValidGender(fl validator.FieldLevel) bool {
	caser := cases.Lower(language.Russian)
	return caser.String(fl.Field().String()) == "мужской" || caser.String(fl.Field().String()) == "женский"
}

func ValidSpecialization(fl validator.FieldLevel) bool {
	return !regexp.MustCompile(`[^а-яА-Я\s]`).MatchString(fl.Field().String())
}
