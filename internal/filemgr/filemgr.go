package filemgr

import (
	"os"
	"../utilities"
	// "fmt"
)

// InitDataFolder sets folder which all databse files and logs are stored
func InitDataFolder(folderName string){
	
	if _, err := os.Stat(folderName+"/data_config"); os.IsNotExist(err) || err == nil{
		makeDirErr := os.MkdirAll(folderName+"/data_config", 0700)
		utilities.PanicError(makeDirErr)

		makeLedgerFolderErr := os.MkdirAll(folderName+"/ledger_database", 0700)
		utilities.PanicError(makeLedgerFolderErr)

		makeLoclaFolderErr := os.MkdirAll(folderName+"/local_database", 0700)
		utilities.PanicError(makeLoclaFolderErr)
		
		makeDataConfig(folderName)
	}else{
		utilities.PanicError(err)
	}

	databaseInitPath := "config/database_init.cnf"
	databasePathes := []string{folderName, folderName + "/ledger_database", folderName + "/local_database"}

	makeAndWriteFile(databaseInitPath, databasePathes, true)

}

// makeDataConfig makes data config files which is needed for managing all databases
func makeDataConfig(foldeName string){
	
	localDatabaseListCnf := foldeName+"/data_config/local_database_list.cnf"
	ledgerDatabaseListCnf := foldeName+"/data_config/ledger_database_list.cnf"

	localDataListLines := []string{"use database:", "", "list of local databases:", ""}
	ledgerDataListLines := []string{"use database:", "", "list of ledger databases:", ""}
	
	makeAndWriteFile(localDatabaseListCnf, localDataListLines, false)
	makeAndWriteFile(ledgerDatabaseListCnf, ledgerDataListLines, false)
}

