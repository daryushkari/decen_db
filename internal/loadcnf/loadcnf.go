package loadcnf

import(
	"../utilities"
	"os"
	"strings"
	"encoding/json"
	"io/ioutil"
	"sync"
	"time"

)


// allDatabaseConfig includes database_init file and config information in it
// including directory path which all databases are and directory of ledger databases
// and local databases and their config file path
type allDataConfig struct {
	DataDir string `json:"DataDir"`
	LedgerDataDir string `json:"LedgerDataDir"`
	LocalDataDir string `json:"LocalDataDir"`
	LedgerDbConfig string `json:"LedgerDatabaseConfig"`
	LocalDbConfig string `json:"LocalDatabaseConfig"`
	LocalDbList []string `json:"LocalDatabaseNameList"`
	LedgerDbList []string `json:"LocalDatabaseNameList"`
 	// if directory for storing is has not been defined yet HasCnf is false
	HasCnf bool
	LastRead time.Time
	LastModify time.Time
}


var allDataCnf = new(allDatabaseConfig)
var once, onceReload *sync.Once

// LoadDatabaseConfig reads information from ./config/database_init.cnf and returns allDatabaseConfig struct
// if refresh is True reload data
func LoadDatabaseConfig() *allDatabaseConfig {
	info, err := os.Stat(DataBaseInitCNF)

	if(sync.Tr)
	once.Do{

	}

}

func timeReload(){

}

func setConstantConfigs(){
	allDataCnf.EachDBCnf = make(map[string]string)
	allDataCnf.DatabaseInitPath = "config/database_init.cnf"
	allDataCnf.EachDBCnf["configDir"] = "/config"
	allDataCnf.EachDBCnf["dataDir"] = "/data"
	allDataCnf.EachDBCnf["collectionDir"] = "/data/collection"
	allDataCnf.EachDBCnf["logDir"] = "/logs"
	allDataCnf.EachDBCnf["ConfigDatabaseFile"] = "/configdatabase_config.cnf"
	allDataCnf.EachDBCnf["collectionListFile"] = "/datacollection_list.cnf"
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
