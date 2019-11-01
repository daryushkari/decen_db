package cmdmgr

import (
	"fmt"
	"../utilities"
)


// LocaldbManage manages all commands starting with localdb
func LocaldbManage(inputCommands){
	switch(inputCommands[1]){
	case "init":
		initDatabase()

	}
}
