package cmdmgr

import (
	"fmt"
	"os"
	"../utilities"
)

// CommandManager Main command manager
func CommandManager(){
	availableCommands := []string{"help", "runfile", "database", " listen", "getledger", "postledger", "ledgerdatabase"}
	
	if !utilities.CheckStringInSlice(os.Args[1],availableCommands){
		fmt.Println("error: ",os.Args[1],"is an invlid command please enter help command for more information")
	}

	switch(os.Args[1]){
	case "help":
		HelpCmd()

	}
}
