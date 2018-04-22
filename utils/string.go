package utils

import (
	"fmt"
)

func Fstring(format string, v ...interface{}) string {
	return fmt.Sprintf(format, v...)
}
