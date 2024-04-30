package main

import (
	"time"

	"github.com/LXJ0000/go-grpc/user/api/middleware"
	"github.com/LXJ0000/go-grpc/user/api/router"
	"github.com/LXJ0000/go-grpc/user/bootstrap"
	"github.com/gin-gonic/gin"
)

func main() {

	app := bootstrap.App()

	env := app.Env

	db := app.Orm

	timeout := time.Duration(env.ContextTimeout) * time.Second // 接口超时时间

	server := gin.Default()
	server.Use(middleware.CORSMiddleware())

	router.Setup(env, timeout, db, server)
	
	_ = server.Run(env.ServerAddress)

}
