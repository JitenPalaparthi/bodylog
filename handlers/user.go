package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/jitenpalaparthi/bodylog/interfaces"
	"github.com/jitenpalaparthi/bodylog/models"

	jwt_lib "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var (
	mysupersecretpassword = "TheAimIsToUseThisJWT"
)

type User struct {
	IUser interfaces.UserInterface
}

func (u *User) Register() func(c *gin.Context) {
	var err error
	return func(c *gin.Context) {
		if c.Request.Method == "POST" {
			user := &models.User{}
			user.Status = "inactive"
			user.LastUpdated = time.Now().UTC().String()
			//*user.Role = ""
			err = json.NewDecoder(c.Request.Body).Decode(user)

			fmt.Println(user)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  "failed",
					"message": err.Error(),
				})
				c.Abort()
				return
			}
			err = models.ValidateUser(user)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  "failed",
					"message": err.Error(),
				})
				c.Abort()
				return
			}
			err = u.IUser.Register(user)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  "failed",
					"message": err.Error(),
				})
				c.Abort()
				return
			}
			//c.BindJSON(&u)
			c.JSON(http.StatusOK, gin.H{
				"status":  "success",
				"message": "User successfully registered",
			})
			c.Abort()
			return
		}
	}
}

func (u *User) SignIn() func(c *gin.Context) {
	var err error
	return func(c *gin.Context) {
		if c.Request.Method == "POST" {
			user := &models.User{}
			err = json.NewDecoder(c.Request.Body).Decode(&user)
			//err = c.BindJSON(&user)
			fmt.Println(user)

			if err != nil {
				//	fmt.Println("what is this error --1")

				c.JSON(http.StatusBadRequest, gin.H{
					"status":  "failed",
					"message": err.Error(),
				})
				c.Abort()
				return
			}

			err = models.ValidateUserForSignin(user)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  "failed",
					"message": err.Error(),
				})
				c.Abort()
				return
			}
			if !u.IUser.Signin(user) {
				fmt.Println("what is this error --3")

				fmt.Println(err)

				c.JSON(http.StatusBadRequest, gin.H{
					"status":  "failed",
					"message": "user could not sign in due to invalid email | password , user is inactive or does not exist ",
				})
				c.Abort()
				return
			}

			// Create the token
			token := jwt_lib.New(jwt_lib.GetSigningMethod("HS256"))
			// Set some claims
			token.Claims = jwt_lib.MapClaims{
				"Id":  user.Email,
				"exp": time.Now().Add(time.Hour * 1).Unix(),
			}
			// Sign and get the complete encoded token as a string
			tokenString, err := token.SignedString([]byte(mysupersecretpassword))
			if err != nil {
				c.JSON(500, gin.H{"message": "Could not generate token"})
			}
			//	bytes, _ := json.Marshal(user)
			//	chanMessage <- Message{Data: bytes, Subject: messaging.SIGNIN}

			c.JSON(200, gin.H{"token": tokenString, "status": "success"})
			c.Abort()
			return
		}
	}
}
func (u *User) MobileSignin() func(c *gin.Context) {
	var err error
	return func(c *gin.Context) {
		if c.Request.Method == "POST" {
			user := &models.User{}
			err = json.NewDecoder(c.Request.Body).Decode(&user)
			//err = c.BindJSON(&user)
			fmt.Println(user)

			if err != nil {
				fmt.Println("what is this error --1")

				c.JSON(http.StatusBadRequest, gin.H{
					"status":  "failed",
					"message": err.Error(),
				})
				c.Abort()
				return
			}

			err = models.ValidateUserForSignin(user)
			if err != nil {
				fmt.Println("what is this error --2")

				fmt.Println(err)

				c.JSON(http.StatusBadRequest, gin.H{
					"status":  "failed",
					"message": err.Error(),
				})
				c.Abort()
				return
			}
			if !u.IUser.Signin(user) {
				fmt.Println("what is this error --3")

				fmt.Println(err)

				c.JSON(http.StatusBadRequest, gin.H{
					"status":  "failed",
					"message": "user could not sign in due to invalid email | password or user does not exist",
				})
				c.Abort()
				return
			}

			// Create the token
			token := jwt_lib.New(jwt_lib.GetSigningMethod("HS256"))
			// Set some claims
			token.Claims = jwt_lib.MapClaims{
				"Id":  user.Email,
				"exp": time.Now().Add(time.Hour * 8760).Unix(),
			}
			// Sign and get the complete encoded token as a string
			tokenString, err := token.SignedString([]byte(mysupersecretpassword))
			if err != nil {
				c.JSON(500, gin.H{"message": "Could not generate token"})
			}
			//bytes, _ := json.Marshal(user)
			//chanMessage <- Message{Data: bytes, Subject: messaging.SIGNIN}

			c.JSON(200, gin.H{"token": tokenString, "status": "success"})
			c.Abort()
			return
		}
	}
}

func (u *User) ResetPassword() func(c *gin.Context) {
	var err error
	return func(c *gin.Context) {
		if c.Request.Method == "POST" {
			user := &models.UserPasswodReset{}
			user.Status = "active"
			user.LastUpdated = time.Now().UTC().String()
			err = json.NewDecoder(c.Request.Body).Decode(&user)
			//err = c.BindJSON(&user)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  "failed",
					"message": err.Error(),
				})
				c.Abort()
				return
			}

			err = models.ValidateUserForPasswordReset(user)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  "failed",
					"message": err.Error(),
				})
				c.Abort()
				return
			}
			err = u.IUser.ResetPassword(user)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  "failed",
					"message": err.Error(),
				})
				c.Abort()
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"status":  "success",
				"message": "User Password has been reset",
			})
			c.Abort()
			return
		}
	}
}

func (u *User) GetUserBy() func(c *gin.Context) {
	return func(c *gin.Context) {
		if c.Request.Method == "GET" {
			email := c.Param("email")
			if email == "" {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  "failed",
					"message": "email parameter has not been provieded",
				})
				c.Abort()
				return
			}
			user, err := u.IUser.GetUserBy(&email)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  "failed",
					"message": err.Error(),
				})
				c.Abort()
				return
			}
			///c.BindJSON(&profile)
			c.JSON(http.StatusOK, user)
		}
	}
}

func (u *User) GetUsers() func(c *gin.Context) {
	return func(c *gin.Context) {
		if c.Request.Method == "GET" {
			skip := c.Param("skip")
			limit := c.Param("limit")

			if skip == "" {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  "failed",
					"message": "skip parameter has not been provieded",
				})
				c.Abort()
				return
			}

			if limit == "" {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  "failed",
					"message": "limit parameter has not been provieded",
				})
				c.Abort()
				return
			}

			iskip, err := strconv.ParseInt(skip, 10, 32)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  "failed",
					"message": err,
				})
				c.Abort()
				return
			}

			ilimit, err := strconv.ParseInt(limit, 10, 32)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  "failed",
					"message": err,
				})
				c.Abort()
				return
			}
			selector := make(map[string]interface{})
			jsonMap := c.Request.URL.Query()

			for key, val := range jsonMap {
				selector[key] = val[0]
			}

			profiles, err := u.IUser.GetUsers(int64(iskip), int64(ilimit), selector)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  "failed",
					"message": err.Error(),
				})
				c.Abort()
				return
			}
			//c.BindJSON(&profiles)
			c.JSON(http.StatusOK, profiles)
		}
	}
}

// UpdateById is to create a profile
func (u *User) UpdateById() func(c *gin.Context) {
	var err error
	return func(c *gin.Context) {
		if c.Request.Method == "PUT" {
			id := c.Param("id")
			if id == "" {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  "failed",
					"message": "User  id parameter has not been provieded",
				})
				c.Abort()
				return
			}

			data := make(map[string]interface{})
			err = json.NewDecoder(c.Request.Body).Decode(&data)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  "body seems to be wrong json format",
					"message": err.Error(),
				})
				c.Abort()
				return
			}
			err = u.IUser.UpdateById(id, data)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  "failed",
					"message": err.Error(),
				})
				c.Abort()
				return
			}
			//c.BindJSON(&u)
			c.JSON(http.StatusOK, gin.H{
				"status":  "success",
				"message": "profile successfully updated",
			})
			c.Abort()
			return
		}
	}
}

// func (u *User) GetSummary() func(c *gin.Context) {
// 	return func(c *gin.Context) {
// 		if c.Request.Method == "GET" {

// 			data, err := u.IUser.GetSummary()
// 			if err != nil {
// 				c.JSON(http.StatusBadRequest, gin.H{
// 					"status":  "failed",
// 					"message": err.Error(),
// 				})
// 				c.Abort()
// 				return
// 			}
// 			///c.BindJSON(&profile)
// 			c.JSON(http.StatusOK, data)
// 		}
// 	}
// }
