package dbasemgr

import "decen_db/internal/loadcnf"

func dropDataBase(dBaseName string)(msg string){
	localDbCnf, err := loadcnf.LoadAllDataConfig()
	if err != nil{
		return err.Error()
	}

	err = loadcnf.RemoveDataBaseFromConfig(dBaseName)
	if err != nil{
		return err.Error()
	}

}
