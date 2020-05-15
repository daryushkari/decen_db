package dbasemgr

import (
	"decen_db/internal/loadcnf"
	"os"
)

func dropDataBase(cmd []string)(msg string){
	dBaseNameIndex := 2

	dBaseBasicInfo, err := loadcnf.ReturnDataBaseBasicInfoByName(cmd[dBaseNameIndex])
	if err != nil{
		return err.Error()
	}

	err = loadcnf.RemoveDataBaseFromConfig(cmd[dBaseNameIndex])
	if err != nil{
		return err.Error()
	}

	err = os.RemoveAll(dBaseBasicInfo.DataBaseDirPath)
	if err != nil{
		loadcnf.AddDataBaseToConfig(dBaseBasicInfo)
		return err.Error()
	}
	return "dropped successfully"

}
