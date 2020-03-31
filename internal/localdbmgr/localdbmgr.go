package localdbmgr

// LocaldbManage manages all commands starting with localdb and calls function related to sub command
func LocaldbManage(cmd []string) string {

	subCmdIndex := 1

	if len(cmd) <= subCmdIndex {
		return "error: too few arguments please enter more"
	}

	switch cmd[subCmdIndex] {

	case "set-path":
		return setPathManage(cmd)


	//case "new":
	//	if len(inputCommands) < 4 {
	//		fmt.Println("please enter database name for creating new database")
	//		return
	//	}
	//	filemgr.MakeDatabase(inputCommands[1], inputCommands[3])

	//case "show":
	//	showCommand(inputCommands)

	//case "drop":
	//	filemgr.DropDatabase()

	default:
		return "invalid input please enter help localdb for more information"
	}

}

//func showCommand(inputCommands []string){
//	if len(inputCommands) < 4{
//		filemgr.ShowDatabase()
//	}
//}
