package builder

import (
	"github.com/elingsuryo/movie-app/config"
	"github.com/elingsuryo/movie-app/internal/http/handler"
	"github.com/elingsuryo/movie-app/internal/http/router"
	"github.com/elingsuryo/movie-app/internal/repository"
	"github.com/elingsuryo/movie-app/internal/service"
	"github.com/elingsuryo/movie-app/pkg/route"
	"gorm.io/gorm"
)

func BuildPublicRoutes(cfg *config.Config, db *gorm.DB) []route.Route {
	//repository
	userRepository := repository.NewUserRepository(db)
	movieRepository := repository.NewMoviesRepository(db)
	//service
	userService := service.NewUserService(cfg,userRepository)
	tokenService := service.NewTokenService(cfg.JWTConfig.SecretKey)
	movieService := service.NewMovieService(movieRepository)
	//handler
	movieHandler := handler.NewMovieHandler(movieService)
	userHandler := handler.NewUserHandler(tokenService, userService)
	//router

	//end
	return router.PublicRoutes(movieHandler, userHandler)
}

func BuildPrivateRoutes(cfg *config.Config, db *gorm.DB) []route.Route {	
	movieRepository := repository.NewMoviesRepository(db)
	userRepository := repository.NewUserRepository(db)

	movieService := service.NewMovieService(movieRepository)
	userService := service.NewUserService(cfg, userRepository)
	tokenService := service.NewTokenService(cfg.JWTConfig.SecretKey)

	movieHandler := handler.NewMovieHandler(movieService)
	userHandler := handler.NewUserHandler(tokenService, userService)
	return router.PrivateRoutes(movieHandler,userHandler)
}
