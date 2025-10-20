package infrastructure_http

import (
	"cinelist/application/usecases"
	"cinelist/infrastructure/database/repositories"
	"cinelist/infrastructure/http/controllers"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func InitializeServer(database *sql.DB) {
	server := gin.Default()

	healthController := controllers.NewHealthController()

	userRepository := repositories.NewUserRepository(database)

	authenticationUseCase := usecases.NewAuthenticationUseCase(userRepository)
	authenticationController := controllers.NewAuthenticationController(authenticationUseCase)

	apiV1 := server.Group("/api/v1")
	{
		healthController.RegisterRoutes(apiV1)
		authenticationController.RegisterRoutes(apiV1)
	}

	server.Run(":8000")
}
