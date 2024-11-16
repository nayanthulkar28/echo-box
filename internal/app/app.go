package app

import (
	"echo-box/config"
	v1 "echo-box/internal/controller/http/v1"
	"echo-box/internal/repo"
	"echo-box/internal/usecase"
	"echo-box/pkg/postgres"
	"fmt"

	"github.com/gin-gonic/gin"
)

// Run is use to setup your database connection, servers, usecase etc...
func Run(cfg *config.Config) {
	// Repository
	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		fmt.Println("app - Run - postgres.New: %w", err)
	}
	defer pg.Close()

	// HTTP Server
	router := gin.Default()
	userRepo := repo.NewUserRepo(pg)
	middleware := v1.NewMiddleware(userRepo)
	authUsecase := usecase.NewAuthUsecase(userRepo)
	userUsecase := usecase.NewUserUsecase(userRepo)
	friendUsecase := usecase.NewFriendUsecase(userRepo)
	wsManager := v1.NewWsManager()
	explorer := v1.NewExplorer(wsManager, friendUsecase)

	go explorer.RunExplorer()
	defer close(explorer.List)

	v1.NewRouter(router, middleware, authUsecase, userUsecase, friendUsecase, wsManager, explorer)
	router.Run(fmt.Sprintf(":%s", cfg.Http.Port))
}
