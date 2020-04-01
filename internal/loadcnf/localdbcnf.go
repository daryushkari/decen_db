package loadcnf

import (
	"decen_db/internal/filemgr"
	"sync"
	"time"
)

//AllDataConfig includes database_init.cnf file and config information in it
//including directory path which all databases are and directory of ledger databases
//and local databases and their config file path
type localDbConfig struct {
	UseDataBase  string   `json:"UseDataBase"`
	DataBaseList []string `json:"DataBaseList"`
	// if directory for storing data has not been defined yet HasCnf is false
	HasCnf   bool      `json:"-"`
	LastRead time.Time `json:"-"`
}


var LocalDbCnf  = new(localDbConfig)
var LocalDbCnfMu sync.Mutex


// InitLocalDbConfig gets directory path which all databases should be stored in and
// saves in /config/database_init.cnf file as json
func InitLocalDbConfig() (localDbConfig *localDbConfig,err error) {
	LocalDbCnfMu.Lock()
	defer LocalDbCnfMu.Unlock()
	setLocalDbConfig()

	dataCnf, err := LoadDataConfig()
	if err != nil{
		return nil,err
	}
	err = filemgr.WriteAsJson(LocalDbCnf, dataCnf.LocalDbCnf)
	if err != nil{
		return nil, err
	}

	LocalDbCnf.LastRead = time.Now()
	return LocalDbCnf, nil
}


func setLocalDbConfig(){
	LocalDbCnf.UseDataBase = ""
	LocalDbCnf.DataBaseList = []string{""}
	LocalDbCnf.HasCnf = true
}
