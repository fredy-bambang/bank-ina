package routes

import (
	"fmt"
	"net/http"

	"github.com/fredy-bambang/bank-ina/configs"
	"github.com/fredy-bambang/bank-ina/controllers"
	"github.com/gin-gonic/gin"
	ginserver "github.com/go-oauth2/gin-server"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
)

func SetupRoutes(r *gin.Engine, userController *controllers.UserController, taskController *controllers.TaskController) {
	// oauth2
	manager := manage.NewDefaultManager()
	manager.MustTokenStorage(store.NewFileTokenStore("data.db"))
	clientStore := store.NewClientStore()
	clientStore.Set(configs.GetConfig().Oauth.ClientID, &models.Client{
		ID:     configs.GetConfig().Oauth.ClientID,
		Secret: configs.GetConfig().Oauth.ClientSecret,
		Domain: configs.GetConfig().Oauth.Domain,
	})
	manager.MapClientStorage(clientStore)

	ginserver.InitServer(manager)
	ginserver.SetAllowGetAccessRequest(true)
	ginserver.SetClientInfoHandler(server.ClientFormHandler)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	auth := r.Group("/oauth2")
	{
		auth.GET("/token", ginserver.HandleTokenRequest)
	}

	userRoutes := r.Group("/users")
	{
		userRoutes.POST("", userController.CreateUser)
		userRoutes.GET("", userController.FindAllUsers)
		userRoutes.GET("/:id", userController.FindUserByID)
		userRoutes.PUT("/:id", userController.UpdateUserByID)
		userRoutes.DELETE("/:id", userController.DeleteUserByID)
	}

	taskRoutes := r.Group("/tasks")
	{
		taskRoutes.Use(handleOAuthErrors(ginserver.HandleTokenVerify()))
		taskRoutes.POST("", taskController.CreateTask)
		taskRoutes.GET("", taskController.FindAllTasks)
		taskRoutes.GET("/:id", taskController.FindTaskByID)
		taskRoutes.PUT("/:id", taskController.UpdateTaskByID)
		taskRoutes.DELETE("/:id", taskController.DeleteTaskByID)
	}
}

// middleware to handle return invalid token
func handleOAuthErrors(handler gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler(c)
		fmt.Println("status:", c.Writer.Status())
		if c.Writer.Status() == http.StatusInternalServerError {
			c.JSON(http.StatusUnauthorized, gin.H{"error": c.Errors.Last().Error()})
			c.Abort()
			return
		}
	}
}
