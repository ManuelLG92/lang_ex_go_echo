package user

import (
	"api.go.com/echo/config"
	"api.go.com/echo/country"
	"api.go.com/echo/globals"
	"api.go.com/echo/language"
	"api.go.com/echo/user/user_models_dto"
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
	userDTO.Gender = user.Gender
	userDTO.Birthday = user.Birthday
	userDTO.Country = user.Country
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
}

func getUserCountry(user User) (error, *country.Country) {
	countryInstance := new(country.Country)
	countryQuery := config.DbGlobal.First(&countryInstance, user.Country)

	if countryQuery.Error != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Invalid country id"), countryInstance
	}
	return nil, countryInstance
}

func GetUserGender(user User) (error, *Gender) {

	gender := new(Gender)
	if err := config.DbGlobal.First(&gender, user.Gender); err != nil {
		return err.Error, gender
	}

	return nil, gender
}

func CheckIfUserExistById(id string) (error, User) {
	var user User
	counter := int64(0)
	if err := config.DbGlobal.Where("id = ?", id).Find(&user).Count(&counter); err != nil {
		if counter < 1 {
			return echo.NewHTTPError(http.StatusBadRequest, "User Not found"), user
		}
	}
	return nil, user
}

func CheckIfLanguageIdExists(languageId uint) error {
	counter := int64(0)

	if err := config.DbGlobal.Where("id = ?", languageId).First(&language.Language{}).Count(&counter); err != nil {
		if counter < 1 {
			return echo.NewHTTPError(http.StatusBadRequest, "Language ID not found")
		}
	}
	return nil
}
