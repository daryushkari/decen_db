package filemgr

import (
	"os"
	"../utilities"
)

// each seperate line is one element in lines
func makeAndWriteFile(filename string, lines []string){
	var allLines string
	for _, i := range lines{
		allLines += i + "\n"
	}
	if _, err := os.Stat(filename); os.IsNotExist(err){
		file, fileError := os.Create(filename)
		utilities.PanicError(fileError)
		defer file.Close()
		file.WriteString(allLines)
	}else{
		utilities.PanicError(err)
	}
}
