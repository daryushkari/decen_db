package collectionmgr

import (
	"decen_db/internal/loadcnf"
	"decen_db/internal/utilities"
	"os"
)
// Todo: refactor ASAP!!
func ManageNewCollection(cmd []string)(msg string){
	dBaseNameIndex := 2
	colNameIndex := 4

	if len(cmd) <= colNameIndex{
		return "invalid input"
	}

	dBaseCnf, err := loadcnf.LoadDataBaseConfigByName(cmd[dBaseNameIndex])
	if err != nil{
		return err.Error()
	}

	if loadcnf.CheckCollectionExist(dBaseCnf, cmd[colNameIndex]){
		return "collection does exist"
	}

	err = createCollectionConfigs(cmd[colNameIndex], dBaseCnf)
	if err != nil{
		return err.Error()
	}

	return "collection created successfully"
}

func createCollectionConfigs(colName string, dBaseCnf *loadcnf.DataBaseConfig)(err error){
	colBasicInfo := loadcnf.ReturnNewCollectionBasicInfo(colName, dBaseCnf)
	err = loadcnf.AddCollectionBasicInfoToConfig(dBaseCnf, colBasicInfo)
	if err != nil{
		return err
	}

	colCnf := loadcnf.MakeCollectionConfig(colName, dBaseCnf.MainDirPath)

	err = makeCollectionDir(colCnf.MainDir)
	if err != nil{
		loadcnf.RemoveCollectionFromDataBaseConfig(dBaseCnf ,colBasicInfo)
		return err
	}

	err = loadcnf.UpdateCollectionByConfigPath(colBasicInfo.ConfigFilePath, colCnf)
	if err != nil{
		loadcnf.RemoveCollectionFromDataBaseConfig(dBaseCnf ,colBasicInfo)
		return err
	}

	err = createCollectionFile(colCnf.CollectionDataPath)
	if err != nil{
		loadcnf.RemoveCollectionFromDataBaseConfig(dBaseCnf ,colBasicInfo)
	}

	return err
}


func createCollectionFile(dataFilePath string)(err error){

	file, err := os.Create(dataFilePath)
	defer utilities.CloseFile(file, &err)
	if err != nil{
		panic(err)
		return err
	}

	_, err = file.WriteString(" ")
	return err
}


// Todo: duplication should be fixed
func makeCollectionDir(colDirPath string)(err error){

	fileNamesList := returnNeededFilesForCollection(colDirPath)

	for _,v := range fileNamesList{
		err = os.MkdirAll(v, 0700)
		if err != nil{
			return err
		}
	}

	return nil

}

func returnNeededFilesForCollection(colDirPath string)[]string{
	return []string{
		colDirPath + loadcnf.CollectionMainDataDirPath,
		colDirPath + loadcnf.CollectionLogDirPath,
		colDirPath + loadcnf.CollectionConfigDirPath,
	}
}