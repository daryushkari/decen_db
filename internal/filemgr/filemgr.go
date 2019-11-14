package filemgr

import (
	"../utilities"
	"fmt"
	"os"
)

// InitDataDir sets folder which all database files and logs are stored
func InitDataDir(dirName string) {

	// check if directory exists delete everything inside folder
	if _, err := os.Stat(dirName); err == nil {
		deleteInDir(dirName)
	}

	err := os.MkdirAll(dirName+"/data_config", 0700)
	utilities.PanicError(err)

	legErr := os.MkdirAll(dirName+"/ledger_database", 0700)
	utilities.PanicError(legErr)

	locErr := os.MkdirAll(dirName+"/local_database", 0700)
	utilities.PanicError(locErr)

	dBaseInitPath := "config/database_init.cnf"
	dBasePathList := []string{"all : " + dirName, "ledgerdb : " + dirName +
		"/ledger_database", "localdb : " + dirName + "/local_database"}

	makeAndWriteFile(dBaseInitPath, dBasePathList, true)

	makeDataConfig(dirName)
}

// MakeDatabase creates a database with name
func MakeDatabase(dBaseType string, dBaseName string) {
	dBasePathDir := returnDataBaseDir(dBaseType) + "/" + dBaseName

	if checkDataBaseExist(dBaseName, dBasePathDir) {
		fmt.Println("database already exist")
		return
	}

	err := os.MkdirAll(dBasePathDir+"/data/collection", 0700)
	utilities.PanicError(err)

	err = os.MkdirAll(dBasePathDir+"/logs", 0700)
	utilities.PanicError(err)

	err = os.MkdirAll(dBasePathDir+"/config", 0700)
	utilities.PanicError(err)

	collectionListPath := dBasePathDir + "/data" + "collection_list.cnf"
	collectionList := []string{"list of collections :"}
	makeAndWriteFile(collectionListPath, collectionList, false)

	dBaseConfigPath := dBasePathDir + "/config" + "database_config.cnf"
	dBaseConfigLines := []string{"database_path_folder : " + dBasePathDir,
		"database_log_path" ,
		"database_data_path : " + dBasePathDir +"/data",
		"collection_path_folder : " + dBasePathDir +"/data/collection",
		}
	makeAndWriteFile(dBaseConfigPath, dBaseConfigLines, false)
}

