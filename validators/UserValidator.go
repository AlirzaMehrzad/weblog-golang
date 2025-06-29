package validators

import (
	"context"
	"login-register/models"
	"regexp"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

// ValidationError represents an error during validation
type ValidationError struct {
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}

// ValidateUser validates the user data
func ValidateUser(user models.User, userCollection *mongo.Collection) error {
	if user.Username == "" {
		return &ValidationError{"یوزرنیم الزامی است"}
	}
	if user.Password == "" {
		return &ValidationError{"رمز عبور الزامی است"}
	}
	if !isValidEmail(user.Email) {
		return &ValidationError{"ایمیل نامعتبر است"}
	}
	if len(user.Password) < 8 || !containsSpecialChar(user.Password) {
		return &ValidationError{"رمز عبور باید حداقل 8 کاراکتر و شامل یک کاراکتر ویژه باشد"}
	}

	if err := CheckRepeatedField(userCollection, "username", user.Username); err != nil {
		return err
	}

	if err := CheckRepeatedField(userCollection, "email", user.Email); err != nil {
		return err
	}
	
	return nil
}


// isValidEmail checks if the email format is valid
func isValidEmail(email string) bool {
	// Simple regex for email validation
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

// containsSpecialChar checks if the password contains at least one special character
func containsSpecialChar(password string) bool {
	specialChars := "!@#$%^&*()-_=+[]{}|;:',.<>?/"
	for _, char := range password {
		if strings.ContainsRune(specialChars, char) {
			return true
		}
	}
	return false
}

// HashPassword hashes the user's password
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func CheckRepeatedField(collections *mongo.Collection, field string, value string) error {
	filter := bson.M{field: value}
	count, err := collections.CountDocuments(context.TODO(), filter)
	if err != nil {
		return err
	}
	if count > 0 {
		return &ValidationError{Message: "یوزرنیم یا ایمیل تکراری است"}
	}
	return nil
}
	