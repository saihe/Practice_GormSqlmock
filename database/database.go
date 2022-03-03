package database

import (
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/labstack/gommon/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *gorm.DB

func init() {
	cfg, err := pgx.ParseConfig(os.Getenv("ELEPHANTSQL_URL"))
	if err != nil {
		log.Error(err)
	}
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: stdlib.OpenDB(*cfg),
	}), &gorm.Config{
		Logger: logger.Discard,
	})
	if err != nil {
		log.Error(err)
	}
	Db = db
	defer func() {
		if sqlDb, err := db.DB(); err == nil {
			sqlDb.Close()
		}
	}()
}
