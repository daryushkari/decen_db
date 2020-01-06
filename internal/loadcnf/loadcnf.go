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
	LedgerDbList []string `json:"LocalDatabaseNameList"`
 	// if directory for storing is has not been defined yet HasCnf is false
	HasCnf bool `json:"-"`
	LastRead time.Time `json:"-"`
}


var allDataCnf = new(allDataConfig)
var mu sync.Mutex

// InitAllDataConfig gets directory path which all databases should be stored in and
// saves in /config/database_init.cnf file as json the information it saves includes
// directory path where all data are stored
// directory path where all data of ledger databases are stored
// directory path where all data of local databases are stored
// path of local databases config file and ledger databases config file
func InitAllDataConfig(allDataDir string) *allDataConfig{
	mu.Lock()
	defer mu.Unlock()
	setAllDataConfig(allDataDir)

	file, err := json.MarshalIndent(allDataCnf, "", " ")
	utilities.PanicError(err)
	err = ioutil.WriteFile(DataBaseInitCnfName, file, 0700)
	utilities.PanicError(err)

	allDataCnf.LastRead = time.Now()
	return allDataCnf
}

func setAllDataConfig(allDataDir string){

	allDataCnf.DataDir = allDataDir
	allDataCnf.DataCnfDir = allDataDir + DataCnfDirNAme
	allDataCnf.LedgerDataDir = allDataDir + LedgerDirName
	allDataCnf.LocalDataDir = allDataDir + LocalDirName
	allDataCnf.LedgerDbCnf = allDataDir + LedgerDbCnfPath
	allDataCnf.LocalDbCnf = allDataDir + LocalDbCnfPath
	allDataCnf.LocalDbList = []string{}
	allDataCnf.LedgerDbList = []string{}
	allDataCnf.HasCnf = true

}


//var once, onceReload *sync.Once

// LoadDatabaseConfig reads information from ./config/database_init.cnf and returns allDatabaseConfig struct
// if refresh is True reload data
//func LoadDatabaseConfig() *allDatabaseConfig {
//	info, err := os.Stat(DataBaseInitCNF)
//
//	if(sync.Tr)
//	once.Do{
//
//	}
//
//}

//func timeReload(){
//	info, err := os.Stat(DataBaseInitCNF)
//	if err == nil &&
//}
//
//func setConstantConfigs(){
//	allDataCnf.EachDBCnf = make(map[string]string)
//	allDataCnf.DatabaseInitPath = "config/database_init.cnf"
//	allDataCnf.EachDBCnf["configDir"] = "/config"
//	allDataCnf.EachDBCnf["dataDir"] = "/data"
//	allDataCnf.EachDBCnf["collectionDir"] = "/data/collection"
//	allDataCnf.EachDBCnf["logDir"] = "/logs"
//	allDataCnf.EachDBCnf["ConfigDatabaseFile"] = "/configdatabase_config.cnf"
//	allDataCnf.EachDBCnf["collectionListFile"] = "/datacollection_list.cnf"
//}
//
//func returnDBLists(DBPath string)(useDB string, DBLists []string){
//	DBLines := utilities.ReturnFileLines(DBPath)
//	useDatabasePlace := 1
//	DataBaseListPlace := 3
//
//	useDB = DBLines[useDatabasePlace]
//	for _, i := range DBLines[DataBaseListPlace:]{
//		if i != ""{
//			DBLists = append(DBLists, i)
//		}
//	}
//	return useDB, DBLists
//}
