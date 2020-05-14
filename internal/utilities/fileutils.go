package utilities

import (
	"os"
	"strings"
)

func CloseFile(file *os.File, err *error){
	*err = file.Close()
}


func ReturnParentDirPath(path string)(parentPath string){
	pathList := strings.Split(path, "/")
	parentPath = pathList[0]
	parentIndex := len(pathList) - 1
	for _,v := range pathList[1:parentIndex]{
		parentPath += "/" + v
	}
	return parentPath
}