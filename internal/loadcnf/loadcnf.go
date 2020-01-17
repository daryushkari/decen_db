package loadcnf

import (
	"time"
	"encoding/json"
	"io/ioutil"
	"sync"
	"os"
	"../utilities"
)


//AllDataConfig includes database_init.cnf file and config information in it
//including directory path which all databases are and directory of ledger databases
//and local databases and their config file path
type AllDataConfig struct {
	DataDir string `json:"DataDir"`
	LedgerDataDir string `json:"LedgerDataDir"`
	LocalDataDir string `json:"LocalDataDir"`
	LedgerDbCnf string `json:"LedgerDatabaseConfig"`
	LocalDbCnf string `json:"LocalDatabaseConfig"`
 	// if directory for storing is has not been defined yet HasCnf is false
	HasCnf bool `json:"-"`
	LastRead time.Time `json:"-"`
}


var once, onceReload *sync.Once

//LoadDataConfig reads information from ./config/database_init.cnf and returns allDataConfig struct
func LoadDataConfig() *AllDataConfig {

	onceReload.Do(func() {
		if timeReload(){
			refreshOnce(once)
		}
	})
	
	defer refreshOnce(onceReload)

	once.Do(readDataConfig)

	return allDataCnf
}

func refreshOnce(refOnce *sync.Once){
	refOnce = new(sync.Once)
}

func readDataConfig(){
	file, err := ioutil.ReadFile(DataInitCnfPath)
	utilities.PanicError(err)
	err = json.Unmarshal([]byte(file), allDataCnf)
	
	allDataCnf.LastRead = time.Now()
	allDataCnf.HasCnf = true

	if err != nil{
		allDataCnf = nil
	}
}

func timeReload()bool{
	info, err := os.Stat(DataInitCnfPath)
	utilities.PanicError(err)
	lastMod := info.ModTime()
	timeDiff := lastMod.Sub(allDataCnf.LastRead)
	if timeDiff > 0{
		return true
	}
	return false
}
	