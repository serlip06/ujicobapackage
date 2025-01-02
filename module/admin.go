package module

import (
	"github.com/serlip06/ujicobapackage/model"
)

var AdminCredentials = model.Admin{
	Username : "admin",
	Password : "admin123", 
}

func ValidateAdmin(input model.Admin) bool {
	return input.Username == AdminCredentials.Username && input.Password == AdminCredentials.Password
}