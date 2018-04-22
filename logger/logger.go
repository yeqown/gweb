/*
 * logger package selfmade
 * output to stdout or stderr also output to file
 * colorful to enhance you foucsing point
 */

package logger

import (
	// "fmt"
	// "time"
	L "github.com/yeqown/log"
)

var (
	ReqL = L.NewLogger()
	AppL = L.NewLogger()
)

func init() {
	// TODO: init a global logger instance, so that can log anywhere
	// and cant output to console and file both
	// file can be splited day by day format like "app.20180412.log"
	ReqL.SetFileOutput("../logs", "app")
	ReqL.SetFileOutput("../logs", "request")
}
