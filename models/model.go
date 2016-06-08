package models

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"

	"github.com/hobo-go/echo-web/modules/log"
)

// @TODO Drivers
var db *gorm.DB

func DB() *gorm.DB {
	if db == nil {
		log.DebugPrint("Model NewDB")
		newDb, err := NewDB()
		if err != nil {
			panic(err)
		}
		db = newDb
	}

	return db
}

const (
	DefaultKey  = "github.com/hobo-go/echo-web/models/model"
	errorFormat = "[models] ERROR! %s\n"
)

type model struct {
	db *gorm.DB
}

func (m model) DB() *gorm.DB {
	log.DebugPrint("Model model DB")
	return m.db
}

func Model() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			db, err := NewDB()
			if err != nil {
				panic(err)
			}
			model := model{db}
			c.Set(DefaultKey, model)

			return next(c)
		}
	}
}

// shortcut to get model
func Default(c echo.Context) model {
	// return c.MustGet(DefaultKey).(model)
	return c.Get(DefaultKey).(model)
}
