package loadcnf

import (
	"decen_db/internal/filemgr"
	"decen_db/internal/utilities"
	"errors"
	"sync"
	"time"
)

type CollectionInfo struct{
	Name string
}

type DatabaseInfo struct{
	Collections []CollectionInfo `json:"Collections"`
	Name string                  `json:"Name"`
}

//allDataConfig includes database_init.cnf file and config information in it
//including directory path which all databases are and directory of ledger databases
//and local databases and their config file path
type localDbConfig struct {
	UseDataBase  string   `json:"UseDataBase"`
	DataBaseList []string `json:"DataBaseList"`
	// if directory for storing data has not been defined yet HasCnf is false
	HasCnf   bool      `json:"-"`
	LastRead time.Time `json:"-"`
}

func (localCnf *localDbConfig) updateLastRead(){
	localCnf.HasCnf = true
	localCnf.LastRead = time.Now()
}

var LocalDbCnf  = new(localDbConfig)
var LocalDbCnfMu sync.Mutex


//LoadLocalDbConfig reads localdb config from file and loads to localDbConfig struct with singleton pattern
func LoadLocalDbConfig() (locCnf *localDbConfig, err error){
	allCnf, err := LoadAllDataConfig()
	if err != nil{
		return nil, err
	}
	err = loadConfigOnce(cnfMap["localDb"],allCnf.LocalDbCnf, AllDataCnf.LastRead)
	return LocalDbCnf, err
}

// SaveLocalDbConfig gets directory path which all databases should be stored in and
// saves in /config/database_init.cnf file as json
func SaveLocalDbConfig() (localDbConfig *localDbConfig,err error) {
	LocalDbCnfMu.Lock()
	defer LocalDbCnfMu.Unlock()
	setLocalDbConfig()

	dataCnf, err := LoadAllDataConfig()
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

// AddDataBaseToConfig gets database name and adds to config list
func AddDataBaseToConfig(dBasename string)(err error){
	LocalDbCnf, err := LoadLocalDbConfig()
	if err != nil{
		return err
	}

	if utilities.CheckStringInSlice(dBasename, LocalDbCnf.DataBaseList){
		return errors.New("error: database with name " + dBasename + " exist. please drop it before creating new one.")
	}

	LocalDbCnf.DataBaseList = append(LocalDbCnf.DataBaseList, dBasename)
	_, err = SaveLocalDbConfig()
	return err
}

// AddDataBaseToConfig gets database name and adds to config list
func RemoveDataBaseFromConfig(dBasename string)(err error){
	LocalDbCnf, err := LoadLocalDbConfig()
	if err != nil{
		return err
	}

	if !utilities.CheckStringInSlice(dBasename, LocalDbCnf.DataBaseList){
		return errors.New("error:" + dBasename + " does not exist")
	}

	LocalDbCnf.DataBaseList = utilities.RemoveFromSlice(dBasename, LocalDbCnf.DataBaseList)
	_, err = SaveLocalDbConfig()
	return err
}