package filemgr

import (
	"../utilities"
	"time"
)


// LogInsert is a function which inserts every log of database
func LogInsert(logFilePath string, logMassage string){
	date := time.Now()
	appendLines := []string{logMassage + " at_time : " + date.String()}
	utilities.AppendFile(appendLines, logFilePath)
}
