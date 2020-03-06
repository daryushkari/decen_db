package filemgr

import (
	"../loadcnf"
	"../utilities"
	"os"
	"fmt"
)


// InitDataDir sets folder which all database files and logs are stored
func InitDataDir(dirName string) {

	// check if directory exists delete everything inside folder
	if _, err := os.Stat(dirName); err == nil {
		deleteInDir(dirName)
	}

	err := os.MkdirAll(dirName+loadcnf.LedgerDirName, 0700)
	utilities.PanicError(err)

	err = os.MkdirAll(dirName+loadcnf.LocalDirName, 0700)
	utilities.PanicError(err)

	loadcnf.InitAllDataConfig(dirName)
}



// MakeDatabase creates a new database with name
func MakeDatabase(dBaseType string, dBaseName string) error{

	databasePath := loadcnf.LoadDataConfig()
	if databasePath == nil{
		return fmt.Errorf("database_init.cnf file corrupted please run localdb init to fix")
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
	return nil
}
//
//
//// Show list of databases
//func ShowDatabase(){
//
//	allDataCnf := loadcnf.LoadDatabaseConfig(true)
//
//	dBaseTypes := map[string][]string{"local":allDataCnf.LocalDBList, "ledger": allDataCnf.LedgerDBList}
//
//	for k, v := range dBaseTypes{
//		fmt.Println("list of " + k + " databases:")
//		for _, i := range v{
//			fmt.Println(i)
//		}
//	}
//
//}
//
//// DropDatabase deletes selected database
//func DropDatabase(dbType string, dbName string){
//
//	allDataCnf := loadcnf.LoadDatabaseConfig(true)
//
//	dbTypesList := map[string][]string{"localdb":allDataCnf.LocalDBList, "ledgerdb": allDataCnf.LedgerDBList}
//	dbCnfFile := map[string]string{"localdb": "loc_cnf", "ledgerdb": "leg_cnf"}
//
//	if !utilities.CheckStringInSlice(dbName, dbTypesList[dbType]){
//		panic("no such database named" + dbName)
//	}
//
//	deleteInDir(allDataCnf.AllDatabaseInfo[dbType] + "/" + dbName)
//	deleteInDir(allDataCnf.AllDatabaseInfo[dbCnfFile[dbType]])
//
//}
//
