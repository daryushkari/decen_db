package dbasemgr

import (
	"decen_db/internal/filemgr"
	"decen_db/internal/loadcnf"
	"os"
)


func setPathManage(cmd []string) (msg string){
	dirNameIndex := 2
	if len(cmd) <= dirNameIndex {
		return "please specify folder name"
	}

	return setDataPath(cmd[dirNameIndex])
}


// setDataPath sets folder which all database files and logs are stored
func setDataPath(dirName string) (msg string){

	// check if directory exists delete everything inside folder
	if _, err := os.Stat(dirName); err == nil {
		err = filemgr.DeleteInDir(dirName)
		if err != nil{
			return err.Error()
		}
	}

	err := os.MkdirAll(dirName+loadcnf.LedgerDbDirName, 0700)
	if err != nil{
		return err.Error()
	}

	err = os.MkdirAll(dirName+loadcnf.LocalDbDirName, 0700)
	if err != nil{
		return err.Error()
	}

	_, err = loadcnf.InitAllDataConfig(dirName)
	if err != nil{
		return err.Error()
	}

	_, err = loadcnf.SaveLocalDbConfig()
	if err != nil{
		return err.Error()
	}

	return "Success"
}