package filemgr

import (
	"../utilities"
	"time"
)

func LogInsert(logFilePath string, logMassage string){
	date := time.Now()
	appendLines := []string{logMassage + " at_time : " + date.String()}
	utilities.AppendFile(appendLines, logFilePath)
}
