package router

import (
	"login-register/db"
	"login-register/handlers"
	"login-register/middleware"

	"github.com/gin-gonic/gin"
)

// InitRouter initializes the router with the routes and returns it
func InitRouter(collections *db.Collections) *gin.Engine {
	r := gin.Default()

	r.POST("/register", handlers.Register(collections.Users))
	r.POST("/login", handlers.Login(collections.Users))

	// Post routes
	r.POST("/posts", middleware.Authenticate(), handlers.AddPost(collections.Posts))
	r.GET("/posts", middleware.Authenticate(), handlers.GetPosts(collections.Posts))

	// File rotes
	r.POST("/upload", middleware.Authenticate(), handlers.UploadFile(collections.Files))

	// // Comment routes
	// r.POST("/comments", handlers.AddComment(collections.Comments))
	// r.GET("/comments", handlers.GetComments(collections.Comments))

	// Add more routes and pass the appropriate collections
	return r
}
