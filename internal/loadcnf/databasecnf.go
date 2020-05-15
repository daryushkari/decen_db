package loadcnf

import (
	"decen_db/internal/filemgr"
	"decen_db/internal/utilities"
	"encoding/json"
	"io/ioutil"
)

type CollectionBasicInfo struct{
	Name string 			`json:"Name"`
	ConfigFilePath string 	`json:"ConfigFilePath"`
}

type DataBaseConfig struct{
	Name string 						`json:"Name"`
	MainDirPath string					`json:"MainDirPath"`
	Collections []CollectionBasicInfo 	`json:"Collections"`
}

// Todo: should be thread safe

func MakeNewDataBaseConfig(dBaseInfo *DataBaseBasicInfo) (dBaseCnf *DataBaseConfig) {

	dBaseCnf = &DataBaseConfig{
		Name:dBaseInfo.Name,
		MainDirPath:dBaseInfo.DataBaseDirPath,
		Collections:[]CollectionBasicInfo{},
	}

	return dBaseCnf

}

// Todo: should be thread safe

func LoadDataBaseConfig(dBaseConfigPath string) (dBaseCnf *DataBaseConfig, err error){
	dBaseCnf = &DataBaseConfig{}
	file, err := ioutil.ReadFile(dBaseConfigPath)
	if err != nil{
		return nil, err
	}
	err = json.Unmarshal([]byte(file), dBaseCnf)
	if err != nil{
		return nil, err
	}

	return dBaseCnf, err
}


func LoadDataBaseConfigByName(dBaseName string)(dBaseCnf *DataBaseConfig, err error){
	dBaseInfo, err := ReturnDataBaseBasicInfoByName(dBaseName)
	if err != nil{
		return nil, err
	}
	dBaseCnf, err = LoadDataBaseConfig(dBaseInfo.ConfigFilePath)
	return dBaseCnf, err

}

func CheckCollectionExist(dBaseCnf *DataBaseConfig, colName string)(colExist bool){
	for _, v := range dBaseCnf.Collections{
		if v.Name == colName{
			return true
		}
	}
	return false
}

func ReturnNewCollectionBasicInfo(colName string, dBaseCnf *DataBaseConfig) *CollectionBasicInfo{
	return &CollectionBasicInfo{
		Name: colName,
		ConfigFilePath: utilities.JoinDirPath([]string{dBaseCnf.MainDirPath, colName, CollectionConfigPath}),
	}
}

func AddCollectionBasicInfoToConfig(dBaseCnf *DataBaseConfig, colInfo *CollectionBasicInfo)(err error){
	colBasicInfoList := &dBaseCnf.Collections
	*colBasicInfoList = append(dBaseCnf.Collections, *colInfo)
	err = UpdateDataBaseConfig(dBaseCnf)
	return err
}


func removeFromCollectionSlice(colName string, dataBaseList []CollectionBasicInfo) []CollectionBasicInfo {
	for i, v := range dataBaseList{
		if v.Name == colName{
			dataBaseList[i] = dataBaseList[len(dataBaseList) - 1]
			dataBaseList = dataBaseList[:len(dataBaseList) - 1]
		}
	}
	return dataBaseList
}

func RemoveCollectionFromDataBaseConfig(dBaseCnf *DataBaseConfig, colInfo *CollectionBasicInfo)(err error){
	dBaseCnf.Collections = removeFromCollectionSlice(colInfo.Name, dBaseCnf.Collections)
	err = UpdateDataBaseConfig(dBaseCnf)
	return err
}

func UpdateDataBaseConfig(dBaseCnf *DataBaseConfig)(err error){
	dBInfo, err := ReturnDataBaseBasicInfoByName(dBaseCnf.Name)
	if err != nil{
		return err
	}


	err = filemgr.WriteAsJson(dBaseCnf, dBInfo.ConfigFilePath)
	return err
}