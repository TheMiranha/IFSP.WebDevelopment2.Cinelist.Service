package infrastructure_http

import (
	"cinelist/application/usecases"
	"cinelist/infrastructure/database/repositories"
	infrastructure_services "cinelist/infrastructure/services"
	"cinelist/infrastructure/http/controllers"
	"cinelist/infrastructure/http/middlewares"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func InitializeServer(database *sql.DB) {
	server := gin.Default()

	server.Use(middlewares.CORSMiddleware())

	healthController := controllers.NewHealthController()

	userRepository := repositories.NewUserRepository(database)
	movieRepository := repositories.NewMovieRepository(database)
	actorRepository := repositories.NewActorRepository(database)
	movieInteractionRepository := repositories.NewMovieInteractionRepository(database)

	authService := infrastructure_services.NewAuthService()
	authenticationUseCase := usecases.NewAuthenticationUseCase(userRepository, authService)
	authenticationController := controllers.NewAuthenticationController(authenticationUseCase)

	userUseCase := usecases.NewUserUseCase(userRepository, movieInteractionRepository)
	userController := controllers.NewUserController(userUseCase)

	movieUseCase := usecases.NewMovieUseCase(movieRepository)
	movieController := controllers.NewMovieController(movieUseCase)

	actorUseCase := usecases.NewActorUseCase(actorRepository)
	actorController := controllers.NewActorController(actorUseCase)

	movieInteractionUseCase := usecases.NewMovieInteractionUseCase(movieInteractionRepository)
	movieInteractionController := controllers.NewMovieInteractionController(movieInteractionUseCase)

	apiV1 := server.Group("/api/v1")
	{
		healthController.RegisterRoutes(apiV1)
		authenticationController.RegisterRoutes(apiV1)
		movieController.RegisterRoutes(apiV1)
		actorController.RegisterRoutes(apiV1)

		protected := apiV1.Group("")
		protected.Use(middlewares.AuthenticationMiddleware())
		{
			userController.RegisterRoutes(protected)
			movieInteractionController.RegisterRoutes(protected)
		}
	}

	server.Run(":8000")
}
