package log

import (
	l "log"
)

func Print(format string, values ...interface{}) {
	l.Printf("[DEBUG] "+format, values...)
}

func PrinError(err error) {
	if err != nil {
		Print("[ERROR] %v\n", err)
	}
}
