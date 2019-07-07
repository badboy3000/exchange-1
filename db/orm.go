package db

import (
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	orm    *gorm.DB
	dbOnce sync.Once
)

func initORM() error {
	var err error
	orm, err = gorm.Open("postgres", "host=127.0.0.1 port=5432 user=kingyang dbname=exchange_development sslmode=disable")
	if err != nil {
		return err
	}
	err = orm.DB().Ping()
	if err != nil {
		return err
	}
	orm.LogMode(true)
	return nil
}

// ORM ...
func ORM() *gorm.DB {
	if orm == nil {
		dbOnce.Do(func() {
			err := initORM()
			if err != nil {
				panic(err)
			}
		})
	}
	return orm
}
