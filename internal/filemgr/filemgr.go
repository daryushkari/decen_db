package filemgr

import (
	"../utilities"
	"fmt"
	"os"
	"../loadcnf"
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
	dBasePathList := []string{"all : " + dirName,
		                      "ledgerdb : " + dirName + "/ledger_database",
		                      "localdb : " + dirName + "/local_database",
		                      "cnf : " + dirName + "/data_config",
							  "loc_cnf : " + dirName + "/data_config/local_database_list.cnf",
		                      "leg_cnf : " + dirName + "/data_config/ledger_database_list.cnf"}
	makeAndWriteFile(dBaseInitPath, dBasePathList, true)

	makeDataConfig(dirName)
}


// MakeDatabase creates a new database with name
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
		"database_log_path : " + dBasePathDir + "/logs",
		"database_data_path : " + dBasePathDir +"/data",
		"collection_path_folder : " + dBasePathDir +"/data/collection",
		}
	makeAndWriteFile(dBaseConfigPath, dBaseConfigLines, false)

	// Add database name in existing databases list
	addDatabaseNameToList(dBaseType, dBaseName)
}


// Show list of databases
func ShowDatabase(){

	allDataCnf := loadcnf.LoadDatabaseConfig(true)

	dBaseTypes := map[string][]string{"local":allDataCnf.LocalDBList, "ledger": allDataCnf.LedgerDBList}

	for k, v := range dBaseTypes{
		fmt.Println("list of " + k + " databases:")
		for _, i := range v{
			fmt.Println(i)
		}
	}

}

// DropDatabase deletes selected database
func DropDatabase(){

}

func removeLine(line string, filename string){

}