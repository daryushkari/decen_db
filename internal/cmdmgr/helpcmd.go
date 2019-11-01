package cmdmgr

import (
	"fmt"
	"../utilities"
)

// HelpCmd shows help
func HelpCmd(inputCommands []string){
	if len(inputCommands) < 3{
		helpList := utilities.ReturnFileLines("config/help/main_help.cnf")
		fmt.Print("here is list of available commands:\n")
		for _, i := range helpList{
			fmt.Println(i)
		}
	}
}