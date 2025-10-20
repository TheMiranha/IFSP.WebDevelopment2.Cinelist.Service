package infrastructure_http

import (
	"cinelist/application/usecases"
	"cinelist/infrastructure/database/repositories"
	"cinelist/infrastructure/http/controllers"
	"cinelist/infrastructure/http/middlewares"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func InitializeServer(database *sql.DB) {
	server := gin.Default()

	healthController := controllers.NewHealthController()

	userRepository := repositories.NewUserRepository(database)
	movieRepository := repositories.NewMovieRepository(database)

	authenticationUseCase := usecases.NewAuthenticationUseCase(userRepository)
	authenticationController := controllers.NewAuthenticationController(authenticationUseCase)

	userUseCase := usecases.NewUserUseCase(userRepository)
	userController := controllers.NewUserController(userUseCase)

	movieUseCase := usecases.NewMovieUseCase(movieRepository)
	movieController := controllers.NewMovieController(movieUseCase)

	apiV1 := server.Group("/api/v1")
	{
		healthController.RegisterRoutes(apiV1)
		authenticationController.RegisterRoutes(apiV1)
		movieController.RegisterRoutes(apiV1)

		protected := apiV1.Group("")
		protected.Use(middlewares.AuthenticationMiddleware())
		{
			userController.RegisterRoutes(protected)
		}
	}

	server.Run(":8000")
}
