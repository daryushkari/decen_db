package collectionmgr

import "decen_db/internal/loadcnf"

func MakeNewCollection(cmd []string)(msg string){
	dBaseNameIndex := 2
	colNameIndex := 4

	DbaseInfo, err := loadcnf.ReturnDataBaseBasicInfoByName(cmd[dBaseNameIndex])
	if err != nil{
		return err.Error()
	}



	return ""
}
