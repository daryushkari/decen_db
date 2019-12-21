package cmdmgr

import (
	"../filemgr"
	"fmt"
)


// LocaldbManage manages all commands starting with localdb and checks user input and calls related functions
func LocaldbManage(inputCommands []string){

	if len(inputCommands) < 3{
		fmt.Println("too few arguments please enter more")
		return
	}

	switch inputCommands[2] {

	case "init":
		if len(inputCommands) < 4{
			fmt.Println("please specify folder name")
			return
		}
		filemgr.InitDataDir(inputCommands[3])

	case "new":
		if len(inputCommands) < 4{
			fmt.Println("please enter database name for creating new database")
			return
		}
		filemgr.MakeDatabase(inputCommands[1], inputCommands[3])

	case "show":
		showCommand(inputCommands)

	//case "drop":
	//	filemgr.DropDatabase()

	default:
		fmt.Println("invalid input please enter help localdb for more information")
	}

}

func showCommand(inputCommands []string){
	if len(inputCommands) < 4{
		filemgr.ShowDatabase()
	}
}