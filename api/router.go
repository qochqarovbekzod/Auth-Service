package router

import (
	_ "users/api/docs"

	"users/api/handlers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// RouterApi @title Auth Service
// @version 1.0
// @description Auth service
// @host localhost:8081
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func RouterApi(h *handlers.Handler) *gin.Engine {
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	users := router.Group("/api/v1/auth")
	{
		users.POST("/register", h.SignUpHandler)
		users.POST("/login", h.LogInHandler)
		users.GET("/profile", h.ViewProfileHandler)
		users.PUT("/profile", h.EditProfile)
		users.PUT("/type", h.ChangeUserTypeHandler)
		users.GET("/", h.GetAllUsersHandler)
		users.DELETE("/:user_id", h.DeleteUserHandler)
		users.POST("/reset-password", h.PasswordResetHandler)
		users.POST("/refresh", h.RefreshToken)
		// users.POST("/logout", h.TokenCancellationHandler)
	}
	return router
}
