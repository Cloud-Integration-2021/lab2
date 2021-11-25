package main

import (
	"lab2/config"
	"lab2/controllers"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	var logger *zap.Logger
	gin.SetMode(gin.ReleaseMode)

	// Load configuration
	cfg, err := config.LoadCfg()
	if err != nil {
		log.Printf("Error to parsing env %v", err)
		os.Exit(0)
	}

	// Set log level
	if cfg.IsDev() {
		logger, err = zap.NewDevelopment()
	} else {
		logger, err = zap.NewProduction()
	}

	if err != nil {
		zap.S().Error("Error to initialize logger")
	}

	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			zap.Error(err)
		}
	}(logger)
	zap.ReplaceGlobals(logger)

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "*"
		},
		MaxAge: 12 * time.Hour,
	}))

	// Route
	r.GET("/actors/:id", controllers.FindActorsByMovieId)

	err = r.Run(":8081")
	if err != nil {
		log.Println("Unable to start web server")
	}
}
