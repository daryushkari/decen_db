package filemgr

import (
	"../utilities"
	"fmt"
	"os"
)

// InitDataFolder sets folder which all database files and logs are stored
func InitDataFolder(folderName string) {

	if _, err := os.Stat(folderName + "/data_config"); os.IsNotExist(err) || err == nil {
		makeDirErr := os.MkdirAll(folderName+"/data_config", 0700)
		utilities.PanicError(makeDirErr)

		makeLedgerFolderErr := os.MkdirAll(folderName+"/ledger_database", 0700)
		utilities.PanicError(makeLedgerFolderErr)

		makeLocalFolderErr := os.MkdirAll(folderName+"/local_database", 0700)
		utilities.PanicError(makeLocalFolderErr)

		makeDataConfig(folderName)
	} else {
		utilities.PanicError(err)
	}

	databaseInitPath := "config/database_init.cnf"
	databasePathList := []string{"all : " + folderName, "ledgerdb : " + folderName +
		"/ledger_database", "localdb : " + folderName + "/local_database"}

	makeAndWriteFile(databaseInitPath, databasePathList, true)

}

// makeDataConfig makes data config files which is needed for managing all databases
func makeDataConfig(folderName string) {

	localDatabaseListCnf := folderName + "/data_config/local_database_list.cnf"
	ledgerDatabaseListCnf := folderName + "/data_config/ledger_database_list.cnf"

	localDataListLines := []string{"use database:", "", "list of local databases:", ""}
	ledgerDataListLines := []string{"use database:", "", "list of ledger databases:", ""}

	makeAndWriteFile(localDatabaseListCnf, localDataListLines, false)
	makeAndWriteFile(ledgerDatabaseListCnf, ledgerDataListLines, false)
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
		"database_log_path"
		"database_data_path : " + databasePathFolder+"/data",
		"collection_path_folder : " + databasePathFolder+"/data/collection",
		}
	makeAndWriteFile(databaseConfigPath, databaseConfigLines, false)
}

func checkDatabaseExist(databaseName string, databaseFolder string) bool {
	allDataFolder := returnDatabaseFolder("all")
	allLedgerDatabase := utilities.ReturnFileLines(allDataFolder + "/ledger_database_list.cnf")
	allLocalDatabase := utilities.ReturnFileLines(allDataFolder + "/local_database_list.cnf")

	if _, err := os.Stat(databaseFolder); os.IsNotExist(err) {
		if !utilities.CheckStringInSlice(databaseName, append(allLedgerDatabase, allLocalDatabase...)) {
			return false
		}
	}
	return true
}
