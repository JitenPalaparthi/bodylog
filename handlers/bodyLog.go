// Package handlers contains all http restful handlers
// Author : readyGo "JitenP@Outlook.Com"
// This code is generated by readyGo. You are free to make amendments as and where required
package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/jitenpalaparthi/bodylog/global"
	"github.com/jitenpalaparthi/bodylog/interfaces"
	"github.com/jitenpalaparthi/bodylog/messaging"
	"github.com/jitenpalaparthi/bodylog/models"

	"github.com/gin-gonic/gin"
)

// BodyLog type used as a container for db interface and receiver for handler functions
type BodyLog struct {
	IBodyLog  interfaces.BodyLogInterface
	Messaging *messaging.Messaging
}

// CreateCovidData is to post an object provided by json in the body
func (c *BodyLog) CreateBodyLog() func(ctx *gin.Context) {
	var err error
	return func(ctx *gin.Context) {
		if ctx.Request.Method == "POST" {

			bodyLog := &models.BodyLog{}

			err = json.NewDecoder(ctx.Request.Body).Decode(&bodyLog)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"status":  "failed",
					"message": err.Error(),
				})
				ctx.Abort()
				return
			}
			// Validate model
			//err = models.ValidateBodyLog(coviddata)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"status":  "failed",
					"message": err.Error(),
				})
				ctx.Abort()
				return
			}
			bodyLog.Status = global.GetDefaultStr(`active`)
			bodyLog.LastModified = global.GetUnixTimeInStr()

			result, err := c.IBodyLog.CreateBodyLog(bodyLog)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"status":  "failed",
					"message": err.Error(),
				})
				ctx.Abort()
				return
			}
			// If it is publisher
			b, _ := models.ToBytes(bodyLog)
			c.Messaging.ChanMessage <- messaging.Message{Data: b, Subject: "bodyLog_topic"}
			ctx.JSON(http.StatusOK, gin.H{
				"status":  "success",
				"message": result,
			})
			ctx.Abort()
			return
		}
	}
}

// // GetCovidDataByID is to get respective object from database provided by id as a param
// func (c *CovidData) GetCovidDataByID() func(ctx *gin.Context) {
// 	return func(ctx *gin.Context) {
// 		if ctx.Request.Method == "GET" {
// 			id := ctx.Param("id")
// 			if id == "" {
// 				ctx.JSON(http.StatusBadRequest, gin.H{
// 					"status":  "failed",
// 					"message": "id parameter has not been provided",
// 				})
// 				ctx.Abort()
// 				return
// 			}
// 			coviddata, err := c.ICovidData.GetCovidDataByID(id)
// 			if err != nil {
// 				ctx.JSON(http.StatusBadRequest, gin.H{
// 					"status":  "failed",
// 					"message": err.Error(),
// 				})
// 				ctx.Abort()
// 				return
// 			}
// 			ctx.JSON(http.StatusOK, coviddata)
// 		}
// 	}
// }

// // GetAllCovidDatas is to get more number of available objects provided by skip and limit params
// func (c *CovidData) GetAllCovidDatas() func(ctx *gin.Context) {
// 	return func(ctx *gin.Context) {
// 		if ctx.Request.Method == "GET" {
// 			skip := ctx.Param("skip")
// 			limit := ctx.Param("limit")

// 			if skip == "" {
// 				ctx.JSON(http.StatusBadRequest, gin.H{
// 					"status":  "failed",
// 					"message": "skip parameter has not been provided",
// 				})
// 				ctx.Abort()
// 				return
// 			}

// 			if limit == "" {
// 				ctx.JSON(http.StatusBadRequest, gin.H{
// 					"status":  "failed",
// 					"message": "limit parameter has not been provided",
// 				})
// 				ctx.Abort()
// 				return
// 			}

// 			iskip, err := strconv.ParseInt(skip, 10, 64)
// 			if err != nil {
// 				ctx.JSON(http.StatusBadRequest, gin.H{
// 					"status":  "failed",
// 					"message": err,
// 				})
// 				ctx.Abort()
// 				return
// 			}

// 			ilimit, err := strconv.ParseInt(limit, 10, 64)
// 			if err != nil {
// 				ctx.JSON(http.StatusBadRequest, gin.H{
// 					"status":  "failed",
// 					"message": err,
// 				})
// 				ctx.Abort()
// 				return
// 			}
// 			selector := make(map[string]interface{})
// 			jsonMap := ctx.Request.URL.Query()

// 			for key, val := range jsonMap {
// 				selector[key] = val[0]
// 			}

// 			coviddatas, err := c.ICovidData.GetAllCovidDatas(int64(iskip), int64(ilimit), selector)
// 			if err != nil {
// 				ctx.JSON(http.StatusBadRequest, gin.H{
// 					"status":  "failed",
// 					"message": err.Error(),
// 				})
// 				ctx.Abort()
// 				return
// 			}
// 			ctx.JSON(http.StatusOK, coviddatas)
// 		}
// 	}
// }

// // GetAllCovidDatasBy is to get more number of available objects provided by skip and limit params and with a condition as key and value
// func (c *CovidData) GetAllCovidDatasBy() func(ctx *gin.Context) {
// 	return func(ctx *gin.Context) {
// 		if ctx.Request.Method == "GET" {
// 			skip := ctx.Param("skip")
// 			limit := ctx.Param("limit")

// 			if skip == "" {
// 				ctx.JSON(http.StatusBadRequest, gin.H{
// 					"status":  "failed",
// 					"message": "skip parameter has not been provided",
// 				})
// 				ctx.Abort()
// 				return
// 			}

// 			if limit == "" {
// 				ctx.JSON(http.StatusBadRequest, gin.H{
// 					"status":  "failed",
// 					"message": "limit parameter has not been provided",
// 				})
// 				ctx.Abort()
// 				return
// 			}

// 			iskip, err := strconv.ParseInt(skip, 10, 64)
// 			if err != nil {
// 				ctx.JSON(http.StatusBadRequest, gin.H{
// 					"status":  "failed",
// 					"message": err,
// 				})
// 				ctx.Abort()
// 				return
// 			}

// 			ilimit, err := strconv.ParseInt(limit, 10, 64)
// 			if err != nil {
// 				ctx.JSON(http.StatusBadRequest, gin.H{
// 					"status":  "failed",
// 					"message": err,
// 				})
// 				ctx.Abort()
// 				return
// 			}

// 			qkey := ctx.Request.URL.Query().Get("key")
// 			qvalue := ctx.Request.URL.Query().Get("value")

// 			coviddatas, err := c.ICovidData.GetAllCovidDatasBy(qkey, qvalue, int64(iskip), int64(ilimit))
// 			if err != nil {
// 				ctx.JSON(http.StatusBadRequest, gin.H{
// 					"status":  "failed",
// 					"message": err.Error(),
// 				})
// 				ctx.Abort()
// 				return
// 			}
// 			//ctx.BindJSON(&profiles)
// 			ctx.JSON(http.StatusOK, coviddatas)
// 		}
// 	}
// }

// // UpdateCovidDataByID is to update an object(put) provided by id as a param. Update keys and values to be given in the body.
// func (c *CovidData) UpdateCovidDataByID() func(ctx *gin.Context) {
// 	var err error
// 	return func(ctx *gin.Context) {
// 		if ctx.Request.Method == "PUT" {

// 			id := ctx.Param("id")
// 			if id == "" {
// 				ctx.JSON(http.StatusBadRequest, gin.H{
// 					"status":  "failed",
// 					"message": "id parameter has not been provided",
// 				})
// 				ctx.Abort()
// 				return
// 			}

// 			var coviddata map[string]interface{}
// 			coviddata = make(map[string]interface{})

// 			err = json.NewDecoder(ctx.Request.Body).Decode(&coviddata)

// 			if err != nil {
// 				ctx.JSON(http.StatusBadRequest, gin.H{
// 					"status":  "failed",
// 					"message": err.Error(),
// 				})
// 				ctx.Abort()
// 				return
// 			}

// 			result, err := c.ICovidData.UpdateCovidDataByID(id, coviddata)
// 			if err != nil {
// 				ctx.JSON(http.StatusBadRequest, gin.H{
// 					"status":  "failed",
// 					"message": err.Error(),
// 				})
// 				ctx.Abort()
// 				return
// 			}

// 			ctx.JSON(http.StatusOK, gin.H{
// 				"status":  "success",
// 				"message": result,
// 			})
// 			ctx.Abort()
// 			return
// 		}
// 	}
// }

// // DeleteCovidDataByID is to delete an object provided by id
// func (c *CovidData) DeleteCovidDataByID() func(ctx *gin.Context) {
// 	return func(ctx *gin.Context) {
// 		if ctx.Request.Method == "DELETE" {

// 			id := ctx.Param("id")
// 			if id == "" {
// 				ctx.JSON(http.StatusBadRequest, gin.H{
// 					"status":  "failed",
// 					"message": "id parameter has not been provided",
// 				})
// 				ctx.Abort()
// 				return
// 			}

// 			result, err := c.ICovidData.DeleteCovidDataByID(id)
// 			if err != nil {
// 				ctx.JSON(http.StatusBadRequest, gin.H{
// 					"status":  "failed",
// 					"message": err.Error(),
// 				})
// 				ctx.Abort()
// 				return
// 			}

// 			ctx.JSON(http.StatusOK, gin.H{
// 				"status":  "success",
// 				"message": result,
// 			})
// 			ctx.Abort()
// 			return
// 		}
// 	}
// }
