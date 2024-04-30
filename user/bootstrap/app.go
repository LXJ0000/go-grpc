package bootstrap

import (
	"github.com/LXJ0000/go-grpc/user/orm"
	"github.com/LXJ0000/go-grpc/user/util/log"
	"github.com/LXJ0000/go-grpc/user/util/snowflake"
)

type Application struct {
	Env *Env
	Orm orm.Database
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.Orm = NewOrmDatabase(app.Env)
	log.Init(app.Env.AppEnv)
	snowflake.Init(app.Env.SnowflakeStartTime, app.Env.SnowflakeMachineID)

	return *app
}
