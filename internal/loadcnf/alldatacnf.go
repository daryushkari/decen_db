package loadcnf


import (
	"decen_db/internal/filemgr"
	"sync"
	"time"
)


//allDataConfig includes database_init.cnf file and config information in it
//including directory path which all databases are and directory of ledger databases
//and local databases and their config file path
type allDataConfig struct {
	DataDir string 		 `json:"DataDir"`
	LedgerDataDir string `json:"LedgerDataDir"`
	LocalDataDir string  `json:"LocalDataDir"`
	LedgerDbCnf string   `json:"LedgerDatabaseConfig"`
	LocalDbCnf string    `json:"LocalDatabaseConfig"`
	// if directory for storing data has not been defined yet HasCnf is false
	HasCnf bool          `json:"-"`
	LastRead time.Time   `json:"-"`
}

func (localCnf *allDataConfig) updateLastRead(){
	localCnf.HasCnf = true
	localCnf.LastRead = time.Now()
}

var AllDataCnf = new(allDataConfig)
var allDataCnfMu sync.Mutex

// LoadAllDataConfig reads all of data config from database_init.cnf file
//and loads to allDataConfig with singleton pattern
func LoadAllDataConfig() (allCnf *allDataConfig,err error){
	err = loadConfigOnce(cnfMap["allData"],DataInitCnfPath, AllDataCnf.LastRead)
	return AllDataCnf, err
}


// InitAllDataConfig gets directory path which all databases should be stored in and
// saves in /config/database_init.cnf file as json
func InitAllDataConfig(allDataDir string) (allDataConfig *allDataConfig, err error) {
	allDataCnfMu.Lock()
	defer allDataCnfMu.Unlock()
	setAllDataConfig(allDataDir)

	err = filemgr.WriteAsJson(AllDataCnf, DataInitCnfPath)
	if err != nil{
		return nil, err
	}

	AllDataCnf.LastRead = time.Now()
	return AllDataCnf, nil
}


func setAllDataConfig(allDataDir string){
	AllDataCnf.DataDir = allDataDir
	AllDataCnf.LedgerDataDir = allDataDir + LedgerDbDirName
	AllDataCnf.LocalDataDir = allDataDir + LocalDbDirName
	AllDataCnf.LedgerDbCnf = allDataDir + LedgerDbCnfPath
	AllDataCnf.LocalDbCnf = allDataDir + LocalDbCnfPath
	AllDataCnf.HasCnf = true
}
