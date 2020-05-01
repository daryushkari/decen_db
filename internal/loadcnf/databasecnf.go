package loadcnf

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

type CollectionBasicInfo struct{
	Name string 			`json:"Name"`
	ConfigFilePath string 	`json:"ConfigFilePath"`
}

type DataBaseConfig struct{
	Name string 						`json:"Name"`
	Collections []CollectionBasicInfo 	`json:"Collections"`
}

// Todo: should be thread safe

func MakeNewDataBaseConfig(dBaseName string) (dBaseCnf *DataBaseConfig, err error) {
	doesExist, err := CheckDataBaseExist(dBaseName)
	if err != nil {
		return nil, err
	}

	if doesExist{
		return nil, errors.New("database already does exist")
	}

	dBaseCnf = &DataBaseConfig{Name:dBaseName}

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

