package log

import (
	l "log"

	"echo-web/conf"
)

func DebugPrint(format string, values ...interface{}) {
	if conf.RELEASE_MODE == false {
		l.Printf("[DEBUG] "+format, values...)
	}
}

func DebugPrinError(err error) {
	if err != nil {
		DebugPrint("[ERROR] %v\n", err)
	}
}
