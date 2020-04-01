package loadcnf


import (
	"decen_db/internal/filemgr"
	"sync"
	"time"
)


//AllDataConfig includes database_init.cnf file and config information in it
//including directory path which all databases are and directory of ledger databases
//and local databases and their config file path
type AllDataConfig struct {
	DataDir       string `json:"DataDir"`
	LedgerDataDir string `json:"LedgerDataDir"`
	LocalDataDir  string `json:"LocalDataDir"`
	LedgerDbCnf   string `json:"LedgerDatabaseConfig"`
	LocalDbCnf    string `json:"LocalDatabaseConfig"`
	// if directory for storing data has not been defined yet HasCnf is false
	HasCnf   bool      `json:"-"`
	LastRead time.Time `json:"-"`
}


var AllDataCnf = new(AllDataConfig)
var allDataCnfMu sync.Mutex


// InitAllDataConfig gets directory path which all databases should be stored in and
// saves in /config/database_init.cnf file as json
func InitAllDataConfig(allDataDir string) (allDataConfig *AllDataConfig, err error) {
	allDataCnfMu.Lock()
	defer allDataCnfMu.Unlock()
	setAllDataConfig(allDataDir)

	err = filemgr.WriteAsJson(allDataDir, DataInitCnfPath)
	if err != nil{
		return nil, err
	}

	AllDataCnf.LastRead = time.Now()
	return AllDataCnf, nil
}


func setAllDataConfig(allDataDir string){
	AllDataCnf.DataDir = allDataDir
	AllDataCnf.LedgerDataDir = allDataDir + LedgerDirName
	AllDataCnf.LocalDataDir = allDataDir + LocalDirName
	AllDataCnf.LedgerDbCnf = allDataDir + LedgerDbCnfPath
	AllDataCnf.LocalDbCnf = allDataDir + LocalDbCnfPath
	AllDataCnf.HasCnf = true
}
