package router

import (
	"time"

	"github.com/LXJ0000/go-grpc/user/orm"
	"github.com/LXJ0000/go-grpc/user/api/handler"
	"github.com/LXJ0000/go-grpc/user/api/middleware"
	"github.com/LXJ0000/go-grpc/user/bootstrap"
	"github.com/LXJ0000/go-grpc/user/repository"
	"github.com/LXJ0000/go-grpc/user/service"
	"github.com/gin-gonic/gin"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db orm.Database, gin *gin.Engine) {
	userRepo := repository.NewUserRepository(db)
	userSvc := service.NewUserService(userRepo, timeout)
	userHandler := &handler.UserHandler{
		UserSvc: userSvc,
		Env:     env,
	}
	publicRouter := gin.Group("/api")
	publicRouter.POST("/login", userHandler.Login)
	publicRouter.POST("/refresh", userHandler.RefreshToken)
	publicRouter.POST("/signup", userHandler.Signup)

	protectedRouter := gin.Group("/api/user")
	// Middleware to verify AccessToken
	protectedRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))
	protectedRouter.GET("/profile", userHandler.Fetch)

}
