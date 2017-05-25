package log

import (
	l "github.com/labstack/gommon/log"
)

func init() {
	l.SetLevel(l.DEBUG)
}

func Debugf(format string, values ...interface{}) {
	l.Debugf(format, values...)
}
