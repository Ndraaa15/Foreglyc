package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // PostgreSQL driver
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func New() *sqlx.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s timezone=%s",
		viper.GetString("database.host"),
		viper.GetString("database.port"),
		viper.GetString("database.user"),
		viper.GetString("database.password"),
		viper.GetString("database.name"),
		viper.GetString("database.sslmode"),
		viper.GetString("database.timezone"),
	)

	db, err := sqlx.Connect(viper.GetString("database.driver"), dsn)
	if err != nil {
		logrus.WithError(err).Fatal("failed to connect to database")
	}

	return db
}
