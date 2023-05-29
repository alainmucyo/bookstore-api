package postgres

import (
	"bookstore/core/environment"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

type Database struct {
	env *environment.Environment
	DB  *gorm.DB
}

func New(env *environment.Environment) *Database {
	var err error
	db, err := gorm.Open(postgres.Open(env.DBURL), &gorm.Config{})
	if err != nil {
		os.Exit(2)
	}
	return &Database{DB: db, env: env}
}

func (db Database) Migrate() {
	err := db.DB.AutoMigrate()
	if err != nil {
		fmt.Println("Failed to migrate tables")
		fmt.Println(err.Error())
		os.Exit(2)
	}
}
