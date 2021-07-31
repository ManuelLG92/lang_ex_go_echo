package user

import (
	"api.go.com/echo/config"
	"api.go.com/echo/globals"
	"api.go.com/echo/globals/password"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
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
			return echo.NewHTTPError(http.StatusConflict, "User already registered with email %V", userInstance.Email)
		}
	}

	if err := config.DbGlobal.Create(&userInstance); err != nil {
		if err.Error != nil {
			return err.Error
		}
	}

	userDTO := userInstance.MakeUserDTO()

	return c.JSON(http.StatusCreated, userDTO)

}

func Update(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	userSearch := &User{}
	//var userSearch *User

	if err := config.DbGlobal.First(&userSearch, id); err.Error != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User not found")
	}

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

	userUpdatedDTO := userSearch.MakeUserDTO()

	return c.JSON(http.StatusAccepted, userUpdatedDTO)
}

func Remove(c echo.Context) error {
	return c.JSON(http.StatusAccepted, "Delete User!")
}
