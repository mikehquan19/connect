package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mikehquan19/connect/schema"
	"github.com/mikehquan19/connect/setup"
	"go.mongodb.org/mongo-driver/bson"
)

func GetUsers(c *gin.Context) {
	cursor, err := setup.DB.Collection("users").Find(context.TODO(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var users []schema.User
	if err = cursor.All(context.TODO(), &users); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}
