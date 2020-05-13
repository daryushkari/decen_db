package collectionmgr

import "decen_db/internal/loadcnf"

func ManageNewCollection(cmd []string)(msg string){
	dBaseNameIndex := 2
	colNameIndex := 4

	if len(cmd) <= colNameIndex{
		return "invalid input"
	}

	dBaseInfo, err := loadcnf.ReturnDataBaseBasicInfoByName(cmd[dBaseNameIndex])
	if err != nil{
		return err.Error()
	}

	dBaseCnf, err := loadcnf.LoadDataBaseConfig(dBaseInfo.ConfigFilePath)

	if loadcnf.CheckCollectionExist(dBaseCnf, cmd[dBaseNameIndex]){
		return "collection does exist"
	}



	return "collection created successfully"
}


