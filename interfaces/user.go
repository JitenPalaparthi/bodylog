// Package interfaces is to define all interface methods for each model type
// Author : readyGo "JitenP@Outlook.Com"
// This code is generated by readyGo. You are free to make amendments as and where required
package interfaces

import (
	"github.com/jitenpalaparthi/bodylog/models"
)

// UserInterface interfaces
type UserInterface interface {
	Register(user *models.User) error
	Signin(userLogin *models.User) bool
	ResetPassword(ResetPasseord *models.UserPasswodReset) error
	GetUserBy(email *string) (*models.User, error)
	GetUsers(skip, limit int64, selector interface{}) ([]models.User, error)
	UpdateById(id string, data interface{}) error
	//GetSummary() (map[string]interface{}, error)
}
