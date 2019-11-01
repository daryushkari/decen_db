package filemgr

import (
	"os"
	"bufio"
	"io/ioutil"
)

// InitDataFolder sets folder which databse files and logs are stored
func InitDataFolder(folderName string){

	

	if _, err := os.Stat("config/database_init.cnf"); os.IsNotExist(err){
		
	}
}
