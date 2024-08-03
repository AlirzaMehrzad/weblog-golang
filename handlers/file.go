package handlers

import (
	"context"
	"fmt"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

const DST = "files"

func UploadFile(filesCollection *mongo.Collection) gin.HandlerFunc {
	return func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ورودی نامعتبر"})
			return
		}

		// Get the current time and format it
		formattedTime := time.Now().Format("20060102_150405") // Format: YYYYMMDD_HHMMSS

		// Create a new filename by appending the formatted time
		newFilename := fmt.Sprintf("%s_%s", formattedTime, file.Filename)

		// Create the destination file path
		dst := filepath.Join(DST, newFilename)

		// Save the uploaded file to the specified destination
		if err := c.SaveUploadedFile(file, dst); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "فایل آپلود نشد"})
			return
		}
		
		// Save the file path to the database
		fileRecord := bson.M{"fileName": newFilename, "path": dst}
		_, err = filesCollection.InsertOne(context.TODO(), fileRecord)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "مسیر فایل ذخیره نشد"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": fmt.Sprintf("'%s' uploaded and path saved!", newFilename)})
	}
}
