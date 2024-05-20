package main

import (
	"log"
	"net/http"

	"github.com/fredy-bambang/bank-ina/configs"
	"github.com/fredy-bambang/bank-ina/initializers"
	"github.com/fredy-bambang/bank-ina/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	configs.LoadConfig()
	initializers.InitDB()
	initializers.InitServices()
	initializers.InitControllers()
}

func main() {
	r := gin.Default()
	routes.SetupRoutes(r, &initializers.UserController, &initializers.TaskController)
	log.Fatal(http.ListenAndServe(":8080", r))
}
