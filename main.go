package main


import (
	"go-web-starter/internal/api"
	"go-web-starter/internal/middleware"
	"go-web-starter/pkg/db"
	"go-web-starter/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	_ "go-web-starter/docs"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"
)

// @title Go Web Starter
// @version 1.0
// @host localhost:8080
// @BasePath /api

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	viper.SetConfigFile("config/config.yaml")
	_ = viper.ReadInConfig()

	logger.Init()
	db.Init()

	r := gin.New()
	r.Use(middleware.Logger(), gin.Recovery())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiGroup := r.Group("/api")
	{
		apiGroup.POST("/login", api.Login)

		auth := apiGroup.Group("/users")
		auth.Use(middleware.JWT())
		{
			auth.GET("", api.ListUsers)
			auth.POST("", api.CreateUser)
		}
	}

	r.Run(":8080")
	println("Server started on port 8080")
}
