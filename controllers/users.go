package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mikehquan19/connect/schema"
	"github.com/mikehquan19/connect/setup"
	"go.mongodb.org/mongo-driver/bson"
)

func GetUsers(c *gin.Context) {
	cursor, err := setup.DB.Collection("Users").Find(context.TODO(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var users []schema.User
	// FIX THIS
	if err = cursor.All(context.TODO(), &users); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

func CreateUser(c *gin.Context) {
	var user schema.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.CreatedAt = time.Now().Unix()

	res, err := setup.DB.Collection("Users").InsertOne(context.TODO(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, res)
}

/*
	--HINTS
	Line 21 combines the unmarshal step with error handling from GO. Seperate these steps how line 14 & 15 do.



	ANSWER:
	err = cursor.All(context.TODO(), &users)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

*/
