package dbasemgr

import (
	"decen_db/internal/loadcnf"
	"decen_db/internal/utilities"
	"os"
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

func makeNewDataBase(dBasename string, filePath string)(err error){

	dBaseDirPath := utilities.JoinDirPath([]string{filePath, dBasename})

	err = makeDataBaseFiles(dBaseDirPath)
	if err != nil{
		return err
	}

	return err
}


func makeDataBaseFiles(dBaseDirPath string)(err error){

	fileNamesList := returnNeededFilesForDataBase(dBaseDirPath)

	for _,v := range fileNamesList{
		err = os.MkdirAll(v, 0700)
		if err != nil{
			return err
		}
	}

	return nil

}

func returnNeededFilesForDataBase(dBaseDirPath string)[]string{
	return []string{dBaseDirPath,
		dBaseDirPath + loadcnf.DataBaseLogPath,
		dBaseDirPath + loadcnf.DataBaseLogPath,
		dBaseDirPath + loadcnf.DataBaseLogPath}
}