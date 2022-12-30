package user

import ( 
	"api.go.com/echo/config"
	"github.com/jinzhu/gorm" 
	"strings"
)


type Gender struct {
	gorm.Model
	Name string `gorm:"type:varchar(20)"`
}

var gender = [3]string{"MALE","FEMALE","NONE"};

func MigrateGenderTable() {
	//var Gender := [3]string{}
	config.DbGlobal.AutoMigrate(&Gender{})
}

func CheckGender(value string) bool  {

	var valueToUpper string = strings.ToUpper(value)
	
	for _, a := range gender {
		if a == valueToUpper {
			return true
		}
	}
	return false
}

 
 
//Gender = make([]string, 3)


