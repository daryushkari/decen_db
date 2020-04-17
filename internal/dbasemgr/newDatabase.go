package dbasemgr

import (
	"decen_db/internal/loadcnf"
)


func manageNewDataBase(cmd []string)(msg string){

	dBaseNameIndex := 2
	if len(cmd) <= dBaseNameIndex {
		return "error: please enter database name for creating new database"
	}

	allDataCnf, err := loadcnf.LoadAllDataConfig()
	if err != nil{
		return err.Error()
	}

	err = loadcnf.AddDataBaseToConfig(cmd[dBaseNameIndex])
	if err != nil{
		return err.Error()
	}

	err = makeNewDataBase(cmd[dBaseNameIndex], allDataCnf.LocalDataDir)
	if err != nil{
		_ = loadcnf.RemoveDataBaseFromConfig(cmd[dBaseNameIndex])
		return err.Error()
	}

	return cmd[dBaseNameIndex] + "created successfully"
}

func makeNewDataBase(Dbname string, filePath string)(err error){
	return nil
}

