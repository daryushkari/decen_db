package loadcnf

import (
	"decen_db/internal/filemgr"
	"encoding/json"
	"io/ioutil"
)

type CollectionConfig struct{
	Name string 				`json:"Name"`
	MainDir string 				`json:"MainDir"`
	CollectionDataPath string 	`json:"CollectionDataDir"`
}

// Todo: should be thread safe

func MakeCollectionConfig(colName string, colMainDir string) (colCnf *CollectionConfig) {

	colCnf = &CollectionConfig{
		Name : colName,
		MainDir : colMainDir,
		CollectionDataPath : colMainDir + CollectionMainDataPath,
	}

	return colCnf

}

// Todo: should be thread safe

func LoadCollectionConfig(colConfigPath string) (colCnf *CollectionConfig, err error){
	file, err := ioutil.ReadFile(colConfigPath)
	if err != nil{
		return nil, err
	}
	err = json.Unmarshal([]byte(file), colCnf)
	if err != nil{
		return nil, err
	}

	return colCnf, err
}

func UpdateCollectionByConfigPath(cnfPath string, colCnf *CollectionConfig)(err error){
	err = filemgr.WriteAsJson(LocalDbCnf, dataCnf.LocalDbCnf)
	return err
}
