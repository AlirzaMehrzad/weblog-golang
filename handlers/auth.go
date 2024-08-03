package handlers

import (
	"context"
	"fmt"
	"log"
	"login-register/middleware"
	"login-register/models"
	"net/http"
	"net/smtp"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Register(usersCollection *mongo.Collection) gin.HandlerFunc {
	return func(c *gin.Context) {
		var creds Credentials

		err := c.ShouldBindJSON(&creds)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ورودی نامعتبر"})
			return
		}

		// go SendWelcomeEmail(creds.Username)

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}

		user := models.User{
			Username: creds.Username,
			Password: string(hashedPassword),
		}

		_, err = usersCollection.InsertOne(context.TODO(), user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "کاربر ایجاد نشد"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
	}
}

func Login(usersCollection *mongo.Collection) gin.HandlerFunc {
	return func(c *gin.Context) {

		// Check for incoming data from body using ShouldBindJSON method
		var creds Credentials
		err := c.ShouldBindJSON(&creds)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ورودی نامعتبر"})
			return
		}

		// Query to database to find user by username property
		var user models.User
		err = usersCollection.FindOne(context.TODO(), bson.M{"username": creds.Username}).Decode(&user)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "کاربر پیدا نشد"})
			return
		}

		// Compare user claimed password with real password in database using CompareHashAndPassword method
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password))
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "اطلاعات وارد شده اشتباه است"})
			return
		}

		// At last stage if password true, then generate token for logged in user using jwt middleware
		token, err := middleware.GenerateJWT(creds.Username)
		if err != nil {
			c.JSON(500, gin.H{"error": "Could not generate token"})
			return
		}

		// Return token and a custom message as response request
		c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})
	}
}

func SendWelcomeEmail(to string) {
	from := "alireza.me@hotmail.com"
	password := "mypassword"

	smtpHost := "smtp.office355.com"
	smtpPort := "465"

	msg := []byte(fmt.Sprintf("To: %s\r\n"+
		"Subject: به وبسایت ما خوش امدید\r\n"+
		"\r\n"+
		"Welcome, %s!\r\n", to, to))

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, msg)
	if err != nil {
		log.Printf("ارسال ایمیل ناموفق به %s: %v", to, err)
		return
	}

	log.Printf("ایمیل خوش آمد ارسال شد به %s", to)
}
