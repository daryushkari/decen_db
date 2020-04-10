package dbasemgr

import (
	"decen_db/internal/filemgr"
	"decen_db/internal/loadcnf"
	"os"
)


func setPathManage(cmd []string) string{
	dirNameIndex := 2
	if len(cmd) <= dirNameIndex {
		return "please specify folder name"
	}

	return setDataPath(cmd[dirNameIndex])
}


// setDataPath sets folder which all database files and logs are stored
func setDataPath(dirName string) string{

	// check if directory exists delete everything inside folder
	if _, err := os.Stat(dirName); err == nil {
		err = filemgr.DeleteInDir(dirName)
		if err != nil{
			return err.Error()
		}
	}

	err := os.MkdirAll(dirName+loadcnf.LedgerDirName, 0700)
	if err != nil{
		return err.Error()
	}

	err = os.MkdirAll(dirName+loadcnf.LocalDirName, 0700)
	if err != nil{
		return err.Error()
	}

	_, err = loadcnf.InitAllDataConfig(dirName)
	if err != nil{
		return err.Error()
	}

	_, err = loadcnf.InitLocalDbConfig()
	if err != nil{
		return err.Error()
	}

	return "Success"
}