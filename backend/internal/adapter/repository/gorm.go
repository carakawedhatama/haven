package repository

import (
	"fmt"
	"haven/pkg/config"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Gorm struct {
	*gorm.DB
	Conf *config.Config `inject:"config"`
}

func (g *Gorm) Startup() error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		g.Conf.Database.Host, g.Conf.Database.User, g.Conf.Database.Password, g.Conf.Database.DBName, g.Conf.Database.Port, g.Conf.Database.SSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	if g.Conf.Env == "development" {
		db.Logger = logger.Default.LogMode(logger.Info)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	sqlDB.SetMaxIdleConns(g.Conf.Database.MaxIdleConn)
	sqlDB.SetConnMaxLifetime(time.Duration(g.Conf.Database.ConnMaxLifetime) * time.Hour)
	sqlDB.SetMaxOpenConns(g.Conf.Database.MaxOpenConn)

	g.DB = db

	return nil
}

func (g *Gorm) Shutdown() error {
	sqlDB, _ := g.DB.DB()

	return sqlDB.Close()
}
