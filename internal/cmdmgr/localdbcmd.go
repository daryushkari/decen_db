package cmdmgr

import (
	"../filemgr"
	"fmt"
)


// LocaldbManage manages all commands starting with localdb and checks user input and calls related functions
func LocaldbManage(inputCommands []string){
	switch inputCommands[2] {

	case "init":
		if len(inputCommands) < 4{
			fmt.Println("please specify folder name")
			return
		}
		filemgr.InitDataFolder(inputCommands[3])

	case "new":
		if len(inputCommands) < 4{
			fmt.Println("please enter database name for creating new database")
			return
		}
		filemgr.MakeNewDatabase(inputCommands[1], inputCommands[3])
	}
}
