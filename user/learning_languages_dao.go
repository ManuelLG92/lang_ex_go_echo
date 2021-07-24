package user

import (
	"api.go.com/echo/config"
)

func InsertIntoDbLearningLanguagesFromArray(learningLanguageArray []*LearningLanguages) {
	for _, singleLearningLang := range learningLanguageArray {
		config.DbGlobal.Create(&singleLearningLang)
	}
}
