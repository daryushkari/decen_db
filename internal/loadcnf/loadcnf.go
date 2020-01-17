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
	DataCnfDir string `json:"DataConfigDir"`
	LedgerDbCnf string `json:"LedgerDatabaseConfig"`
	LocalDbCnf string `json:"LocalDatabaseConfig"`
 	// if directory for storing is has not been defined yet HasCnf is false
	HasCnf bool `json:"-"`
	LastRead time.Time `json:"-"`
}


var once, onceReload *sync.Once

//LoadDatabaseConfig reads information from ./config/database_init.cnf and returns allDatabaseConfig struct
func LoadDatabaseConfig() *AllDataConfig {

	onceReload.Do(func() {
		if timeReload(){
			refreshOnce(once)
		}
	})
	
	defer refreshOnce(onceReload)

	once.Do(func() {
			file, _ := ioutil.ReadFile(DataBaseInitCnfPath)
			_ = json.Unmarshal([]byte(file), allDataCnf)
			allDataCnf.LastRead = time.Now()
	})

	return allDataCnf
}

func refreshOnce(refOnce *sync.Once){
	refOnce = new(sync.Once)
}

func timeReload()bool{
	info, err := os.Stat(DataBaseInitCnfPath)
	utilities.PanicError(err)
	lastMod := info.ModTime()
	timeDiff := lastMod.Sub(allDataCnf.LastRead)
	if timeDiff > 0{
		return true
	}
	return false
}
	