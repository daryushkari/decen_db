package filemgr

import (
	"os"
	"../utilities"
	// "fmt"
)

// InitDataFolder sets folder which all databse files and logs are stored
func InitDataFolder(folderName string){

	databaseInitPath := "config/database_init.cnf"

	file, fileError := os.Create(databaseInitPath)
	utilities.PanicError(fileError)
	defer file.Close()
	file.WriteString(folderName+"\n")
	
	if _, err := os.Stat(folderName+"/data_config"); os.IsNotExist(err){
		makeDirErr := os.MkdirAll(folderName+"/data_config", 0700)
		utilities.PanicError(makeDirErr)
		makeDataConfig(folderName)
	}else if err == nil{
		makeDataConfig(folderName)
	}else{
		utilities.PanicError(err)
	}

}

// makeDataConfig makes data config files which is needed for managing all databases
func makeDataConfig(foldeName string){
	
	localDatabaseListCnf := foldeName+"/data_config/local_database_list.cnf"
	ledgerDatabaseListCnf := foldeName+"/data_config/ledger_database_list.cnf"

	localDataListLines := []string{"use database:", "", "list of local databases:", ""}
	ledgerDataListLines := []string{"use database:", "", "list of ledger databases:", ""}
	
	makeAndWriteFile(localDatabaseListCnf, localDataListLines)
	makeAndWriteFile(ledgerDatabaseListCnf, ledgerDataListLines)
}

