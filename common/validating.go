package common

import (
	"regexp"

	"github.com/go-playground/validator/v10"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func ValidNameName(fl validator.FieldLevel) bool {
	return !regexp.MustCompile(`[^а-яА-Я]`).MatchString(fl.Field().String())
}

func ValidNameFamily(fl validator.FieldLevel) bool {
	return !regexp.MustCompile(`[^а-яА-Я]`).MatchString(fl.Field().String())
}

func ValidDateofBirth(fl validator.FieldLevel) bool {
	flag, _ := CheckDateValue(fl.Field().String())
	return flag
}

func ValidPhoneNumber(pN string) bool {
	return regexp.MustCompile(`^7[0-9]{10}$`).MatchString(pN)
}

func ValidGender(fl validator.FieldLevel) bool {
	caser := cases.Lower(language.Russian)
	return caser.String(fl.Field().String()) != "мужской" && caser.String(fl.Field().String()) != "женский"
}

func ValidSpecialization(fl validator.FieldLevel) bool {
	return !regexp.MustCompile(`[^а-яА-Я\s]`).MatchString(fl.Field().String())
}

func ValidCabinet(cab int) bool {
	return cab > 0 && cab < 86
}
