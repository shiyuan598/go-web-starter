package main

import (
	"go-web-starter/docs"
	"go-web-starter/internal/api"
	"go-web-starter/internal/middleware"
	"go-web-starter/pkg/db"
	"go-web-starter/pkg/logger"

	_ "go-web-starter/docs"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
)

// @title Go Web Starter
// @version 1.0
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	viper.SetConfigFile("config/config.yaml")
	_ = viper.ReadInConfig()

	// 动态设置 swagger basePath
	basePath := viper.GetString("server.swagger_base_path")
	if basePath == "" {
		basePath = "/api"
	}
	docs.SwaggerInfo.BasePath = basePath

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

	if err := r.Run(":" + viper.GetString("server.port")); err != nil {
		logger.Log.Error("server start failed", zap.Error(err))
	}
}
