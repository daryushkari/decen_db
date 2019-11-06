package cmdmgr

import (
	"fmt"
	//"../utilities"
	"../filemgr"
)


// LocaldbManage manages all commands starting with localdb
func LocaldbManage(inputCommands []string){
	fmt.Println(inputCommands[1])
	switch(inputCommands[2]){
	case "init":
		if len(inputCommands) < 4{
			fmt.Println("please specify folder name")
			return
		}
		fmt.Println(inputCommands[3])
		filemgr.InitDataFolder(inputCommands[3])
	}
}
