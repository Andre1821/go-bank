package initializers

import "bank.com/models"

func SysncDatabase(){
	DB.AutoMigrate(&models.User{})
}