package globals

import (
	"api.go.com/echo/config"
	"api.go.com/echo/country"
	"api.go.com/echo/language"
	"regexp"
)

func GetLanguageById(langId uint) (error, *language.Language) {
	languageInstance := new(language.Language)
	if err := config.DbGlobal.Find(&languageInstance, langId); err != nil {
		return err.Error, languageInstance
	}

	return nil, languageInstance
}

func GetCountryById(countryId uint) (error, *country.Country) {

	countryInstance := new(country.Country)
	if err := config.DbGlobal.Find(&countryInstance, countryId); err != nil {
		return err.Error, countryInstance
	}

	return nil, countryInstance
}

func CheckNumberParamString(parameter string) bool {

	return regexp.MustCompile(`\d+`).MatchString(parameter)
}
