package dbasemgr

import (
	"decen_db/internal/loadcnf"
	"decen_db/internal/utilities"
)

func showDataBases()string{
	localDbCnf, err := loadcnf.LoadLocalDbConfig()
	if err != nil{
		return err.Error()
	}

	DBaseList := loadcnf.GetNamesFromDataBaseList(localDbCnf.DataBaseList)
	DBaseListLines := utilities.ConvertSliceToLines(DBaseList)
	return DBaseListLines
}
