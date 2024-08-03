package handlers

import (
	"context"
	"login-register/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddPost(postsCollection *mongo.Collection) gin.HandlerFunc {
	return func(c *gin.Context) {
		var post models.Post
        
		if err := c.ShouldBindJSON(&post); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ورودی نامعتبر"})
			return
		}

		_, err := postsCollection.InsertOne(context.TODO(), post)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "پست اضافه نشد"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "Post added successfully"})
	}
}

func GetPosts(postsCollection *mongo.Collection) gin.HandlerFunc {
	return func(c *gin.Context) {
		cursor, err := postsCollection.Find(context.TODO(), bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "خبری پیدا نشد"})
			return
		}
		defer cursor.Close(context.TODO())

		var posts []models.Post
		if err = cursor.All(context.TODO(), &posts); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "ارور سرور"})
			return
		}

		c.JSON(http.StatusOK, posts)
	}
}
