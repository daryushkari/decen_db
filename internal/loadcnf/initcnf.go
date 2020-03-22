package loadcnf

import(
	"time"
	"sync"
	"encoding/json"
	"io/ioutil"
)

var allDataCnf  = new(AllDataConfig)
var mu sync.Mutex

// InitAllDataConfig gets directory path which all databases should be stored in and
// saves in /config/database_init.cnf file as json the information it saves includes
// directory path where all data are stored
// directory path where all data of ledger databases are stored
// directory path where all data of local databases are stored
// path of local databases config file and ledger databases config file
func InitAllDataConfig(allDataDir string) *AllDataConfig{
	mu.Lock()
	defer mu.Unlock()
	setAllDataConfig(allDataDir)

	file, _ := json.MarshalIndent(allDataCnf, "", " ")
	_ = ioutil.WriteFile(DataInitCnfPath, file, 0700)


	allDataCnf.LastRead = time.Now()
	return allDataCnf
}

func setAllDataConfig(allDataDir string){

	allDataCnf.DataDir = allDataDir
	allDataCnf.LedgerDataDir = allDataDir + LedgerDirName
	allDataCnf.LocalDataDir = allDataDir + LocalDirName
	allDataCnf.LedgerDbCnf = allDataDir + LedgerDbCnfPath
	allDataCnf.LocalDbCnf = allDataDir + LocalDbCnfPath
	allDataCnf.HasCnf = true

}
