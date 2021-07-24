package user

import (
	"api.go.com/echo/config"
	"api.go.com/echo/country"
	"api.go.com/echo/globals"
	"api.go.com/echo/user/user_models_dto"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func (user User) MakeUserDTO() *DTOUser {

	userDTO := new(DTOUser)
	userDTO.ID = user.ID
	userDTO.Name = user.Name
	userDTO.Surname = user.Surname
	userDTO.Description = user.Description
	userDTO.Birthday = user.Birthday
	err, gender := getUserGender(user)
	if err != nil {
		fmt.Println(err)
	}
	userDTO.Gender = gender
	userDTO.Birthday = user.Birthday
	err, countryInstance := getUserCountry(user)
	if err != nil {
		fmt.Println(err)
	}
	userDTO.Country = countryInstance

	return userDTO
}

func (user User) MaKeNativeLanguageDTO(nativeLanguage *NativeLanguages) *user_models_dto.NativeLanguagesDTO {
	nativeLangDTO := new(user_models_dto.NativeLanguagesDTO)
	err, lang := globals.GetCountryById(nativeLanguage.LanguageID)
	if err != nil {
		log.Fatalf("Couldn't get language name from GetCountryById. Err: %v", err)
	}
	nativeLangDTO.LanguageName = lang.Name
	nativeLangDTO.UserId = user.ID
	return nativeLangDTO

}

func (user User) MakeNativeLanguagesArrayDTO(nativeLanguages []*NativeLanguages) []*user_models_dto.NativeLanguagesDTO {

	var nativeLanguagesDTO []*user_models_dto.NativeLanguagesDTO

	for _, nativeLang := range nativeLanguages {

		err, lang := globals.GetLanguageById(nativeLang.LanguageID)

		if err != nil {
			log.Fatalf("Couldn't get language name from GetCountryById. Err: %v", err)

		}
		instanceNativeLanguagesDTO := new(user_models_dto.NativeLanguagesDTO)
		instanceNativeLanguagesDTO.LanguageId = lang.ID
		instanceNativeLanguagesDTO.LanguageName = lang.Name
		instanceNativeLanguagesDTO.UserId = user.ID

		nativeLanguagesDTO = append(nativeLanguagesDTO, instanceNativeLanguagesDTO)

	}

	return nativeLanguagesDTO

}

func (user User) MakeLearningLanguagesArrayDTO(learningLanguages []*LearningLanguages) []*user_models_dto.LearningLanguagesDTO {

	var learningLanguagesDTO []*user_models_dto.LearningLanguagesDTO

	for _, nativeLang := range learningLanguages {

		err, lang := globals.GetLanguageById(nativeLang.LanguageID)

		if err != nil {
			log.Fatalf("Couldn't get language name from GetLanguageById. Err: %v", err)
		}

		instanceLearningLanguagesDTO := new(user_models_dto.LearningLanguagesDTO)
		instanceLearningLanguagesDTO.LanguageId = lang.ID
		instanceLearningLanguagesDTO.LanguageName = lang.Name
		instanceLearningLanguagesDTO.UserId = user.ID

		learningLanguagesDTO = append(learningLanguagesDTO, instanceLearningLanguagesDTO)

	}

	return learningLanguagesDTO
	//var nativeLanguagesDTO []user_models_dto.NativeLanguagesDTO
}

func getUserCountry(user User) (error, *country.Country) {
	countryInstance := new(country.Country)
	countryQuery := config.DbGlobal.First(&countryInstance, user.CountryID)

	if countryQuery.Error != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Invalid country id"), countryInstance
	}
	return nil, countryInstance
}

func getUserGender(user User) (error, *Gender) {

	gender := new(Gender)
	if err := config.DbGlobal.First(&gender, user.GenderID); err != nil {
		return err.Error, gender
	}

	return nil, gender
}