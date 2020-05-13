package loadcnf

import (
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

func MakeNewDataBaseConfig(dBaseName string) (dBaseCnf *DataBaseConfig, err error) {

	dBaseCnf = &DataBaseConfig{Name:dBaseName, Collections:[]CollectionBasicInfo{}}

	return dBaseCnf, nil

}

// Todo: should be thread safe

func LoadDataBaseConfig(dBaseConfigPath string) (dBaseCnf *DataBaseConfig, err error){
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


func CheckCollectionExist(dBaseCnf *DataBaseConfig, colName string)(colExist bool){
	for _, v := range dBaseCnf.Collections{
		if v.Name == colName{
			return true
		}
	}
	return false
}

func ReturnNewCollectionBasicInfo(colName string, colCnfPath string) *CollectionBasicInfo{
	return &CollectionBasicInfo{Name: colName, ConfigFilePath: colCnfPath}
}

func AddCollectionBasicInfoToConfig(dBaseCnf *DataBaseConfig, colInfo *CollectionBasicInfo){
	colBasicInfoList := &dBaseCnf.Collections
	*colBasicInfoList = append(dBaseCnf.Collections, *colInfo)
}

func RemoveCollectionFromConfig(dBaseCnf *DataBaseConfig, colInfo *CollectionBasicInfo){

}