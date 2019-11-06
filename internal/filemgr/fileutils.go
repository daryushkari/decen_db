package filemgr

import (
	"os"
	"../utilities"
)

// each seperate line is one element in lines
func makeAndWriteFile(filename string, lines []string, ovewrite bool){
	var allLines string
	for _, i := range lines{
		allLines += i + "\n"
	}

	// if overwrite is true deletes file and overwrites if file does exist
	if ovewrite{
		file, fileError := os.Create(filename)
		utilities.PanicError(fileError)
		file.WriteString(allLines)
		return
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
