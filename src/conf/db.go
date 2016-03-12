package conf

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

func DB() *gorm.DB {
	sqlConnection := DB_USER_NAME + ":" + DB_PASSWORD + "@tcp(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", sqlConnection)
	if err != nil {
		panic(err)
	}

	return db
}

const (
	DefaultKey  = "conf/db"
	errorFormat = "[gorm] ERROR! %s\n"
)

func DBInit() echo.HandlerFunc {
	return func(c *echo.Context) error {
		sqlConnection := DB_USER_NAME + ":" + DB_PASSWORD + "@tcp(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?charset=utf8mb4&parseTime=True&loc=Local"
		db, err := gorm.Open("mysql", sqlConnection)
		if err != nil {
			panic(err)
		}
		c.Set(DefaultKey, db)

		return nil
	}
}

// shortcut to get DB
func DefaultDB(c *echo.Context) gorm.DB {
	return c.Get(DefaultKey).(gorm.DB)
}
