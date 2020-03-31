package utilities

import (
	"os"
)

func CloseFile(file *os.File, err *error){
	*err = file.Close()
}