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
		file, err := os.Create(filename)
		utilities.PanicError(err)
		_, writeErr := file.WriteString(allLines)
		utilities.PanicError(writeErr)
		return
	}

	if _, err := os.Stat(filename); os.IsNotExist(err){
		file, err := os.Create(filename)
		utilities.PanicError(err)
		defer file.Close()
		_, writeErr := file.WriteString(allLines)
		utilities.PanicError(writeErr)
	}else{
		utilities.PanicError(err)
	}
}


// return where should database be stored based on it's type if give all type returns parent folder
func returnDataBaseDir(dBaseType string) string {
	file, err := os.Open("config/database_init.cnf")
	utilities.PanicError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		if strings.Contains(scanner.Text(), dBaseType){
			dBasePath := strings.Fields(scanner.Text())
			return dBasePath[len(dBasePath)-1]
		}
	}
	panic("config file is corrupted please run command: \n localdb init folder name to fix it")
}

// makeDataConfig makes data config files which is needed for managing all databases
func makeDataConfig(dirName string) {

	locCnf := dirName + "/data_config/local_database_list.cnf"
	legCnf := dirName + "/data_config/ledger_database_list.cnf"

	locLines := []string{"use database:", "", "list of local databases:", ""}
	legLines := []string{"use database:", "", "list of ledger databases:", ""}

	makeAndWriteFile(locCnf, locLines, false)
	makeAndWriteFile(legCnf, legLines, false)
}

func checkDataBaseExist(dBaseName string, dBaseDir string) bool {
	dataPath := returnDataBaseDir("all")
	legPath := utilities.ReturnFileLines(dataPath + "/ledger_database_list.cnf")
	locPath := utilities.ReturnFileLines(dataPath + "/local_database_list.cnf")

	if _, err := os.Stat(dBaseDir); os.IsNotExist(err) {
		if !utilities.CheckStringInSlice(dBaseName, append(legPath, locPath...)) {
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
