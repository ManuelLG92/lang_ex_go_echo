package user

import "api.go.com/echo/user/user_models_dto"

type DTOUserResponse struct {
	UserDTO              *DTOUser
	NativeLanguagesDTO   []*user_models_dto.NativeLanguagesDTO
	LearningLanguagesDTO []*user_models_dto.LearningLanguagesDTO
}
