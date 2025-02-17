package router

import (
	"net/http"
	"people-matching-api/handler"
	"people-matching-api/helper"
	"people-matching-api/repository"
	"people-matching-api/service"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoute(db *gorm.DB) *gin.Engine {
	userRepo := repository.NewUserRepository(db)

	userService := service.NewUserService(*userRepo)

	userHandler := handler.NewUserHandler(*userService)

	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	url := ginSwagger.URL(helper.GetEnvWithKey("HOST") + ":" + helper.GetEnvWithKey("PORT") + "/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	api := r.Group("/api")
	api.GET("/users", userHandler.GetUser)
	api.POST("/users", userHandler.RegisterUser)
	api.GET("/match_user", userHandler.MatchUser)

	return r
}
