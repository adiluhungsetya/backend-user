package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"user-go-test/config"
	"user-go-test/controller"
	"user-go-test/repositories"
	"user-go-test/service"
)

func main() {
	viper.AutomaticEnv()
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	db := config.Connect()

	config.MigrateTable(db)
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	userRepository := repositories.CreateUserRepository(db)
	userService := service.CreateUserService(userRepository)
	controller.CreateUserHandler(router,userService)

	err = router.Run(":"+viper.GetString("port"))
	if err != nil {
		log.Fatal(err)
	}
}
