package dbasemgr

import (
	"decen_db/internal/filemgr"
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

	err = loadcnf.AddDataBaseToConfig(loadcnf.CreateNewDataBaseInfo(cmd[dBaseNameIndex],allDataCnf.LocalDataDir))
	if err != nil{
		return err.Error()
	}

	err = makeNewDataBase(allDataCnf.LocalDataDir, cmd[dBaseNameIndex])

	if err != nil{
		_ = loadcnf.RemoveDataBaseFromConfig(cmd[dBaseNameIndex])
		return err.Error()
	}

	return cmd[dBaseNameIndex] + "created successfully"
}

func createDataBaseConfigFile(dBaseName string)(err error){
	dBaseInfo, err := loadcnf.ReturnDataBaseBasicInfoByName(dBaseName)
	if err != nil{
		return err
	}

	dBaseCnf := loadcnf.MakeNewDataBaseConfig(dBaseInfo)

	err = filemgr.WriteAsJson(dBaseCnf, dBaseInfo.ConfigFilePath)
	return err
}

func makeNewDataBase(filePath string, dBaseName string)(err error){

	dBaseDirPath := utilities.JoinDirPath([]string{filePath, dBaseName})

	err = makeDataBaseDirs(dBaseDirPath)
	if err != nil{
		return err
	}

	err = createDataBaseConfigFile(dBaseName)
	return err
}

func makeDataBaseDirs(dBaseDirPath string)(err error){

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
		dBaseDirPath + loadcnf.DataBaseCollectionPath}
}

