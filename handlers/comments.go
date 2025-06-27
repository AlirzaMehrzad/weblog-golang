package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"

	"login-register/models"

	"go.mongodb.org/mongo-driver/mongo"
)

func AddComment(commentsCollection *mongo.Collection) gin.HandlerFunc {
	return func(c *gin.Context) {
		var comment models.Comment
		err := c.BindJSON(&comment)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		_, err = commentsCollection.InsertOne(context.TODO(), comment)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusCreated, gin.H{"comment": comment})
	}
}

func GetComments(commentsCollection *mongo.Collection) gin.HandlerFunc {
	return func(c *gin.Context) {
		cursor, err := commentsCollection.Find(context.TODO(), bson.D{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		var comments []models.Comment
		if err = cursor.All(context.TODO(), &comments); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"comments": comments})
	}
}
