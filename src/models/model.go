package models

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"

	"conf"
)

const (
	DefaultKey  = "models/model"
	errorFormat = "[models] ERROR! %s\n"
)

type model struct {
	db *gorm.DB
}

func (m model) DB() *gorm.DB {
	return m.db
}

func Model() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			db := conf.DB()
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
