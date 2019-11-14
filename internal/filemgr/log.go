package filemgr

import (
	"../utilities"
	"time"
)


// LogInsert is a function which inserts every log of database
func LogInsert(logPath string, logMassage string){
	date := time.Now()
	lines := []string{logMassage + " at_time : " + date.String()}
	utilities.AppendFile(lines, logPath)
}
