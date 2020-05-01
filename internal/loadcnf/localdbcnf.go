package loadcnf

import (
	"decen_db/internal/filemgr"
	"decen_db/internal/utilities"
	"errors"
	"sync"
	"time"
)

type DatabaseBasicInfo struct{
	Name string          	`json:"Name"`
	ConfigFilePath string 	`json:"ConfigFilePath"`
}

//allDataConfig includes database_init.cnf file and config information in it
//including directory path which all databases are and directory of ledger databases
//and local databases and their config file path
type localDbConfig struct {
	UseDataBase  string              `json:"UseDataBase"`
	DataBaseList []DatabaseBasicInfo `json:"DataBaseList"`
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
	if LocalDbCnf.HasCnf{
		return
	}
	LocalDbCnf.UseDataBase = ""
	LocalDbCnf.DataBaseList = []DatabaseBasicInfo{}
	LocalDbCnf.HasCnf = true
}


func getDataBaseList(dBaseInfo []DatabaseBasicInfo)(dBaseList []string){
	for _, v := range dBaseInfo{
		dBaseList = append(dBaseList, v.Name)
	}
	return dBaseList
}


func CheckDataBaseExist(dBasename string)(doesExist bool,err error){
	LocalDbCnf, err := LoadLocalDbConfig()
	if err != nil{
		return false, err
	}
	if utilities.CheckStringInSlice(dBasename, getDataBaseList(LocalDbCnf.DataBaseList)){
		return true, nil
	}
	return false, nil
}


func CreateNewDataBaseInfo(dBasename string, localDbDirPath string)(newDBaseInfo *DatabaseBasicInfo){
	newDBaseInfo = &DatabaseBasicInfo{
		Name: dBasename,
		ConfigFilePath:utilities.JoinDirPath([]string{localDbDirPath,dBasename,DataBaseConfigPath}),
	}
	return newDBaseInfo
}

// AddDataBaseToConfig gets database name and adds to config list
func AddDataBaseToConfig(newDBaseInfo *DatabaseBasicInfo)(err error){
	LocalDbCnf, err := LoadLocalDbConfig()
	if err != nil{
		return err
	}

	// Todo: take string to somewhere else
	doesExist, err := CheckDataBaseExist(newDBaseInfo.Name)
	if err != nil{
		return err
	}
	if doesExist{
		return errors.New("error: database with name " + newDBaseInfo.Name + " exist. please drop it before creating new one.")
	}

	dataBaseListPtr := &LocalDbCnf.DataBaseList


	*dataBaseListPtr = append(LocalDbCnf.DataBaseList, *newDBaseInfo)
	_, err = SaveLocalDbConfig()

	return err
}

// AddDataBaseToConfig gets database name and adds to config list
func RemoveDataBaseFromConfig(dBasename string)(err error){
	LocalDbCnf, err := LoadLocalDbConfig()
	if err != nil{
		return err
	}

	if !utilities.CheckStringInSlice(dBasename, getDataBaseList(LocalDbCnf.DataBaseList)){
		return errors.New("error:" + dBasename + " does not exist")
	}

	LocalDbCnf.DataBaseList = removeFromDataBaseSlice(dBasename, LocalDbCnf.DataBaseList)
	_, err = SaveLocalDbConfig()
	return err
}


func removeFromDataBaseSlice(name string, dataBaseList []DatabaseBasicInfo) []DatabaseBasicInfo {
	for i, v := range dataBaseList{
		if v.Name == name{
			dataBaseList[i] = dataBaseList[len(dataBaseList) - 1]
			dataBaseList = dataBaseList[:len(dataBaseList) - 1]
		}
	}
	return dataBaseList
}

func ReturnDataBaseBasicInfoByName(dBaseName string)(newDBaseInfo *DatabaseBasicInfo, err error){
	LocalDbCnf, err := LoadLocalDbConfig()
	if err != nil{
		return nil, err
	}

	for _,v := range LocalDbCnf.DataBaseList{
		if v.Name == dBaseName{
			return &v, nil
		}
	}
	return nil,  errors.New("error:" + dBaseName + " does not exist")
}