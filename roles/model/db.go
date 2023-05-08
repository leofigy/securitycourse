package model

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Provider interface {
	GetDB() (*gorm.DB, error)
}

type Persistence struct {
	db          *gorm.DB
	location    string
	mysql       bool
	credentials string
}

func NewPersistence(location string, production bool, credentials string) Provider {
	return &Persistence{
		nil, location, production, credentials,
	}
}

func (provider *Persistence) GetDB() (*gorm.DB, error) {
	if provider.db != nil {
		return provider.db, nil
	}
	err := provider.InitModels()
	return provider.db, err
}

func (provider *Persistence) InitModels() error {

	var (
		db  *gorm.DB
		err error
	)

	if provider.mysql {
		log.Println("-------- mysql -------")
		db, err = gorm.Open(
			mysql.Open(provider.credentials), &gorm.Config{},
		)
	} else {
		log.Println("-------- starting in testing mode with slite -------")
		db, err = gorm.Open(sqlite.Open(provider.location), &gorm.Config{})
	}

	if err != nil {
		return err
	}

	err = db.AutoMigrate(
		&User{},
		&Role{},
		&Resource{},
	)

	if provider.db == nil {
		provider.db = db
	}

	return err
}
