package dbasemgr

// LocalDbManage manages all commands starting with localdb and calls function related to sub command
func LocalDbManage(cmd []string) string {

	subCmdIndex := 1

	if len(cmd) <= subCmdIndex {
		return "error: too few arguments please enter more"
	}

	switch cmd[subCmdIndex] {

	case "set-path":
		return setPathManage(cmd)

	case "new":
		return manageNewDataBase(cmd)

	case "show":
		return showDataBases()

	//case "drop":
	//	filemgr.DropDatabase()



	}
	return "invalid input please enter help localdb for more information"
}

