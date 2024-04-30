package bootstrap

import (
	"log"

	"github.com/LXJ0000/go-grpc/user/orm"
	"github.com/LXJ0000/go-grpc/user/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewOrmDatabase(env *Env) orm.Database {

	db, err := gorm.Open(mysql.Open(env.MySQLAddress), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
	}

	if err = db.AutoMigrate(
		&domain.User{},
	); err != nil {
		log.Fatal(err)
	}
	database := orm.NewDatabase(db)

	return database
}
