package filemgr

import (
	"../utilities"
	"bufio"
	"os"
	"strings"
	"path/filepath"
)

// each separate line is one element in lines
func makeAndWriteFile(filename string, lines []string, overwrite bool){
	var allLines string
	for _, i := range lines{
		allLines += i + "\n"
	}

	// if overwrite is true deletes file and overwrites if file does exist
	if overwrite {
		file, fileError := os.Create(filename)
		utilities.PanicError(fileError)
		_, writeErr := file.WriteString(allLines)
		utilities.PanicError(writeErr)
		return
	}

	if _, err := os.Stat(filename); os.IsNotExist(err){
		file, fileError := os.Create(filename)
		utilities.PanicError(fileError)
		defer file.Close()
		_, writeErr := file.WriteString(allLines)
		utilities.PanicError(writeErr)
	}else{
		utilities.PanicError(err)
	}
}


// return where should database be stored based on it's type if give all type returns parent folder
func returnDatabaseFolder(databaseType string)string{
	file, err := os.Open("config/database_init.cnf")
	utilities.PanicError(err)
	defer file.Close()

	lineReader := bufio.NewScanner(file)
	for lineReader.Scan(){
		if strings.Contains(lineReader.Text(), databaseType){
			databasePath := strings.Fields(lineReader.Text())
			return databasePath[len(databasePath)-1]
		}
	}
	panic("config file is corrupted please run command: \n localdb init folder name to fix it")
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


func deleteInDir(dirPath string){
	dir, err := os.Open(dirPath)
	utilities.PanicError(err)
	defer dir.Close()
	
	names, err := dir.Readdirnames(-1)
	utilities.PanicError(err)
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dirPath, name))
		utilities.PanicError(err)
	}
}
