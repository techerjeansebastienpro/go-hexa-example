package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/techerjeansebastienpro/go-hexa-example/internal/application"
	"github.com/techerjeansebastienpro/go-hexa-example/internal/domain"
	"github.com/techerjeansebastienpro/go-hexa-example/internal/infrastructure"
	db "github.com/techerjeansebastienpro/go-hexa-example/pkg/models"
)

func main() {
	envConfig()
	fmt.Println(viper.GetString("DATABASE_URL"))
	prismaClient := db.NewClient(
		db.WithDatasourceURL(viper.GetString("DATABASE_URL")),
	)
	if err := prismaClient.Prisma.Connect(); err != nil {
		panic(err)
	}

	defer func() {
		if err := prismaClient.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()
	userService := domain.NewUserService(infrastructure.NewUserRepository((prismaClient)))
	api := infrastructure.NewUserApi(*userService)

	app := gin.New()
	application.NewUserHttpHandler(app, api).RegisterRoutes()

	app.Run(":8080")

}

func envConfig() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
}
