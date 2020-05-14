package collectionmgr

import (
	"decen_db/internal/loadcnf"
	"decen_db/internal/utilities"
	"os"
)

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

	colBasicInfo := loadcnf.ReturnNewCollectionBasicInfo(cmd[colNameIndex], dBaseCnf)
	err = loadcnf.AddCollectionBasicInfoToConfig(dBaseCnf, colBasicInfo)
	if err != nil{
		return err.Error()
	}

	colCnf := loadcnf.MakeCollectionConfig(dBaseCnf.Name, dBaseCnf.MainDirPath)
	err = createCollectionFile(colCnf.CollectionDataPath)
	if err != nil{
		loadcnf.RemoveCollectionFromDataBaseConfig(dBaseCnf ,colBasicInfo)
		return err.Error()
	}

	return "collection created successfully"
}


func createCollectionFile(dataFilePath string)(err error){

	file, err := os.Create(dataFilePath)
	defer utilities.CloseFile(file, &err)
	if err != nil{
		return err
	}

	_, err = file.WriteString("")
	return err
}
