package models

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	ErrMandatoryField = "field is mandatory"
)

// User model
type User struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty" mapstructure:"_id"`
	Name        *string            `json:"name" bson:"name"`
	Email       *string            `json:"email" bson:"email"`
	Mobile      *string            `json:"mobile" bson:"mobile"`
	Role        *string            `json:"role" bson:"role"`
	Password    *string            `json:"password" bson:"password"`
	Status      string             `json:"status" bson:"status"`
	LastUpdated string             `json:"lastUpdated" bson:"lastUpdated"`
}

// UserDetails model
type UserDetails struct {
	Email *string `json:"email" bson:"email"`
}

type UserAccess struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty" mapstructure:"_id"`
	UserID    primitive.ObjectID `json:"userId" bson:"_userId,omitempty" mapstructure:"_userId"`
	ProjectID primitive.ObjectID `json:"projectId" bson:"_projectId,omitempty" mapstructure:"_projectId"`
}

// UserPasswodReset model
type UserPasswodReset struct {
	Email       *string `json:"email" bson:"email"`
	VerifyCode  *string `json:"verifyCode" bson:"verifyCode"`
	Password    *string `json:"password" bson:"password"`
	Status      string  `json:"status" bson:"status"`
	LastUpdated string  `json:"lastUpdated" bson:"lastUpdated"`
}

// ValidateUser validates the User type
func ValidateUser(u *User) error {
	if u == nil {
		return fmt.Errorf("User details provided are empty or null")
	}
	if *u.Mobile == "" || u.Mobile == nil {
		return fmt.Errorf("Mobile " + ErrMandatoryField)
	}
	if *u.Email == "" || u.Email == nil {
		return fmt.Errorf("Email " + ErrMandatoryField)
	}

	if *u.Password == "" || u.Password == nil {
		return fmt.Errorf("Password " + ErrMandatoryField)
	}

	return nil
}

func ValidateUserForSignin(u *User) error {

	if u == nil {
		return fmt.Errorf("User details provided are empty or null")
	}
	if *u.Email == "" || u.Email == nil {
		return fmt.Errorf("Email " + ErrMandatoryField)
	}
	if *u.Password == "" || u.Password == nil {
		return fmt.Errorf("Password " + ErrMandatoryField)
	}

	return nil
}

func ValidateUserForPasswordReset(u *UserPasswodReset) error {
	if u == nil {
		return fmt.Errorf("User details provided are empty or null")
	}
	if *u.Email == "" || u.Email == nil {
		return fmt.Errorf("Email " + ErrMandatoryField)
	}
	if *u.Password == "" || u.Password == nil {
		return fmt.Errorf("Password " + ErrMandatoryField)
	}
	if *u.VerifyCode == "" || u.VerifyCode == nil {
		return fmt.Errorf("Verification Code " + ErrMandatoryField)
	}
	return nil
}
