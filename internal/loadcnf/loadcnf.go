package loadcnf

import (
	"time"
	"sync"
	"encoding/json"
	"io/ioutil"
	"../utilities"
)


// allDatabaseConfig includes database_init file and config information in it
// including directory path which all databases are and directory of ledger databases
// and local databases and their config file path
type allDataConfig struct {
	DataDir string `json:"DataDir"`
	LedgerDataDir string `json:"LedgerDataDir"`
	LocalDataDir string `json:"LocalDataDir"`
	DataCnfDir string `json:"DataConfigDir"`
	LedgerDbCnf string `json:"LedgerDatabaseConfig"`
	LocalDbCnf string `json:"LocalDatabaseConfig"`
	LocalDbList []string `json:"LocalDatabaseNameList"`
	LedgerDbList []string `json:"LedgerDatabaseNameList"`
 	// if directory for storing is has not been defined yet HasCnf is false
	HasCnf bool `json:"-"`
	LastRead time.Time `json:"-"`
}


var once, onceReload *sync.Once

// LoadDatabaseConfig reads information from ./config/database_init.cnf and returns allDatabaseConfig struct
//if refresh is True reload data
func LoadDatabaseConfig() *allDatabaseConfig {
	info, err := os.Stat(DataBaseInitCNF)

	if(sync.Tr)
	once.Do{

	}

}

func timeReload(){
	info, err := os.Stat(DataBaseInitCNF)
	if err == nil &&
}

func returnDBLists(DBPath string)(useDB string, DBLists []string){
	DBLines := utilities.ReturnFileLines(DBPath)
	useDatabasePlace := 1
	DataBaseListPlace := 3

	useDB = DBLines[useDatabasePlace]
	for _, i := range DBLines[DataBaseListPlace:]{
		if i != ""{
			DBLists = append(DBLists, i)
		}
	}
	return useDB, DBLists
}
