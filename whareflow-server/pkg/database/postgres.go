package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sync"
)

type postgresDb struct {
	Db *gorm.DB
}

var (
	once       sync.Once
	dbInstance *postgresDb
)

func NewPostgres(dsn string) Database {
	once.Do(func() {
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}

		dbInstance = &postgresDb{Db: db}
	})

	return dbInstance

}

func (p *postgresDb) GetDb() *gorm.DB {
	return dbInstance.Db
}
