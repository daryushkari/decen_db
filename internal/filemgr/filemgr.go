package filemgr

import (
	"../utilities"
	"fmt"
	"os"
)

// InitDataFolder sets folder which all database files and logs are stored
func InitDataFolder(folderName string) {
	deleteInDir(folderName)

	makeDirErr := os.MkdirAll(folderName+"/data_config", 0700)
	utilities.PanicError(makeDirErr)

	makeLedgerFolderErr := os.MkdirAll(folderName+"/ledger_database", 0700)
	utilities.PanicError(makeLedgerFolderErr)

	makeLocalFolderErr := os.MkdirAll(folderName+"/local_database", 0700)
	utilities.PanicError(makeLocalFolderErr)

	databaseInitPath := "config/database_init.cnf"
	databasePathList := []string{"all : " + folderName, "ledgerdb : " + folderName +
		"/ledger_database", "localdb : " + folderName + "/local_database"}

	makeAndWriteFile(databaseInitPath, databasePathList, true)

	makeDataConfig(folderName)
}

// MakeNewDatabase creates a database with name 
func MakeNewDatabase(databaseType string, databaseName string) {
	databasePathFolder := returnDatabaseFolder(databaseType) + "/" + databaseName

	if checkDatabaseExist(databaseName, databasePathFolder) {
		fmt.Println("database already exist")
		return
	}

	makeDirErr := os.MkdirAll(databasePathFolder+"/data/collection", 0700)
	utilities.PanicError(makeDirErr)

	makeDirErr = os.MkdirAll(databasePathFolder+"/logs", 0700)
	utilities.PanicError(makeDirErr)

	makeDirErr = os.MkdirAll(databasePathFolder+"/config", 0700)
	utilities.PanicError(makeDirErr)

	collectionListPath := databasePathFolder + "/data" + "collection_list.cnf"
	collectionList := []string{"list of collections :"}
	makeAndWriteFile(collectionListPath, collectionList, false)

	databaseConfigPath := databasePathFolder + "/config" + "database_config.cnf"
	databaseConfigLines := []string{"database_path_folder : " + databasePathFolder,
		"database_log_path" ,
		"database_data_path : " + databasePathFolder+"/data",
		"collection_path_folder : " + databasePathFolder+"/data/collection",
		}
	makeAndWriteFile(databaseConfigPath, databaseConfigLines, false)
}

