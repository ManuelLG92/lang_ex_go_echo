package user

import (
	"api.go.com/echo/config"
	"api.go.com/echo/globals"
	"api.go.com/echo/globals/password"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

var (
	db *gorm.DB
)

func Show(c echo.Context) error {

	paramId := c.Param("id")
	if err := globals.CheckNumberParamString(paramId); !err {
		return echo.NewHTTPError(http.StatusBadRequest, "Only are allowed number as user id")
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	result := GetUser(id)

	if result.Error != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Invalid User id.")
	}

	user := result.Value.(*User)

	userDTO := user.MakeUserDTO()

	err, nativeLanguages := user.GetAllNativeLanguages()
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Couldn't retrieve native languages for user %v", strconv.Itoa(int(user.ID)))
	}

	nativeLanguagesDTO := user.MakeNativeLanguagesArrayDTO(nativeLanguages)

	err, learningLanguages := user.GetAllLearningLanguages()

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Couldn't retrieve learning languages for user %v", strconv.Itoa(int(user.ID)))

	}

	learningLanguagesDTO := user.MakeLearningLanguagesArrayDTO(learningLanguages)

	return c.JSON(http.StatusAccepted, DTOUserResponse{UserDTO: userDTO, LearningLanguagesDTO: learningLanguagesDTO, NativeLanguagesDTO: nativeLanguagesDTO})

}

func Store(c echo.Context) error {
	userInstance := new(User)

	if err := c.Bind(&userInstance); err != nil {
		return err
	}

	userInstance.Password = password.Generate(userInstance.Password)

	errors := ValidateStruct(*userInstance)

	if errors != nil {
		return c.JSON(http.StatusBadRequest, errors)
	}

	counter := int64(0)
	if err := config.DbGlobal.Where("email = ?", &userInstance.Email).Find(&User{}).Count(&counter); err != nil {
		if counter > 0 {
			return echo.NewHTTPError(http.StatusBadRequest, "User already registered with email %V", userInstance.Email)
		}
	}

	if err := config.DbGlobal.Create(&userInstance); err != nil {
		if err.Error != nil {
			return err.Error
		}
	}

	userDTO := userInstance.MakeUserDTO()
	fmt.Println(userDTO)

	return c.JSON(http.StatusCreated, userDTO)
	/*	var content map[string]interface{}

			err := c.Bind(&content)
			fmt.Println(content)
			fmt.Println(err)
			if err != nil {
				return err
			}
			fmt.Println("aftert bind content err")
			fmt.Println(content)
			fmt.Println(userInstance)
			if reflect.TypeOf(content["native_languages"]) == nil {
				return echo.NewHTTPError(http.StatusBadRequest, "Native languages should be at least one")
			}
			if reflect.TypeOf(content["learning_languages"]) == nil {
				return echo.NewHTTPError(http.StatusBadRequest, "Learning languages should be at least one")
			}

			nativeLanguagesRequest := content["native_languages"].([]interface{})
			learningLanguagesRequest := content["learning_languages"].([]interface{})

		nativeLangArray, err := CreateNativeLanguages(nativeLanguagesRequest, userInstance.ID)
		if err != nil {
			return err
		}

		learningLanguagesArray, err := CreateLearningLanguages(learningLanguagesRequest, userInstance.ID)
		if err != nil {
			return err
		}

		InsertIntoDbLearningLanguagesArray(learningLanguagesArray)
		InsertIntoDbNativeLanguagesArray(nativeLangArray)*/

}

func Update(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	fmt.Println(id)
	userSearch := &User{}
	config.DbGlobal.First(&userSearch, id)

	if err := c.Bind(&userSearch); err != nil {
		return err
	}

	errors := ValidateStruct(*userSearch)

	if errors != nil {
		return c.JSON(http.StatusBadRequest, errors)
	}

	if errorQuery := config.DbGlobal.Save(&userSearch); errorQuery.Error != nil {
		return c.JSON(http.StatusBadRequest, errorQuery.Error.Error())
	}
	fmt.Println("after Save")

	userUpdatedDTO := userSearch.MakeUserDTO()
	fmt.Println(userUpdatedDTO)
	return c.JSON(http.StatusAccepted, userUpdatedDTO)
}

func Remove(c echo.Context) error {
	return c.JSON(http.StatusAccepted, "Delete User!")
}

/*
func CreateNativeLanguages(nativeLanguagesRequest []interface{}, userId uint) ([]*NativeLanguages, error) {

	var nativeLanguagesArray []*NativeLanguages

	if len(nativeLanguagesRequest) > 1 {

		for _, value := range nativeLanguagesRequest {

			native := new(NativeLanguages)
			if len(value.(map[string]interface{})) > 1 {
				for key, val := range value.(map[string]interface{}) {

					if key == "level_language" {
						native.Level = LanguageLevel(val.(string))
					}
					if key == "language_id" {
						if !globals.CheckNumberParamString(val.(string)) {
							return nativeLanguagesArray, echo.NewHTTPError(http.StatusBadRequest, "Language ID should be a number.")
						}
						id, err := strconv.ParseUint(val.(string), 10, 64)
						if err != nil {
							native.LanguageID = uint(1)
						} else {
							native.LanguageID = uint(id)
						}
					}
				}
				native.UserID = userId
				nativeLanguagesArray = append(nativeLanguagesArray, native)

			} else {
				return nativeLanguagesArray, echo.NewHTTPError(http.StatusBadRequest, "Some missing value to add a into native language struct.")
			}

		}
		return nativeLanguagesArray, nil
	} else {
		return nativeLanguagesArray, echo.NewHTTPError(http.StatusBadRequest, "It's mandatory at least one native language.")

	}

}

func CreateLearningLanguages(learningLanguagesRequest []interface{}, userId uint) ([]*LearningLanguages, error) {

	var learningLanguagesArray []*LearningLanguages

	if len(learningLanguagesRequest) > 0 {
		for _, value := range learningLanguagesRequest {
			learning := new(LearningLanguages)
			if len(value.(map[string]interface{})) > 1 {
				for key, val := range value.(map[string]interface{}) {

					if key == "level_language" {
						learning.Level = LanguageLevel(val.(string))
					}
					if key == "language_id" {
						id, err := strconv.ParseUint(val.(string), 10, 64)
						if err != nil {
							learning.LanguageID = uint(1)
						} else {
							learning.LanguageID = uint(id)
						}

					}


				}
				learning.UserID = userId
				learningLanguagesArray = append(learningLanguagesArray, learning)

			} else {
				return learningLanguagesArray, echo.NewHTTPError(http.StatusBadRequest, "Some missing value to add a into learning language struct.")
			}

			//		db.Create(&learning)
		}
		return learningLanguagesArray, nil
	} else {
		return learningLanguagesArray, echo.NewHTTPError(http.StatusBadRequest, "It's mandatory at least one learning language.")

	}
}*/

/*func SetUpDbInstance(dbInstance *gorm.DB) {
	db = dbInstance
}*/
