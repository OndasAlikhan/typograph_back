package main

import (
	"fmt"
	"os"
	_ "typograph_back/docs"
	"typograph_back/src/application"
	"typograph_back/src/application/database"
	"typograph_back/src/application/routes"

	"github.com/labstack/gommon/log"
	echoSwagger "github.com/swaggo/echo-swagger"
	"gorm.io/gorm/logger"
)

// @title Typograph Backend
// @version 1.0
// @description Backend for speedtyping service
// @host localhost:8080
// @BasePath /api/v1
// @SecurityDefinitions.apiKey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// if err := godotenv.Load(); err != nil {
	// 	log.Fatalf("Env exception: %v", err)
	// }

	application.InitializeApp(log.DEBUG)
	application.InitializeDB(logger.Info)

	application.GlobalApp.GET("/swagger/*", echoSwagger.WrapHandler)

	routes.RegisterRoute("/api/v1")
	database.RunFixtures()

	err := application.GlobalApp.Start(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
	if err != nil {
		log.Fatalf("App start exception: %v", err)
	}
}
