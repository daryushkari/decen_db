package cmdmgr

import (
	"fmt"
	"os"
	"../utilities"
)

// HelpCmd shows help
func HelpCmd(){
	if len(os.Args) < 3{
		helpList := utilities.ReturnFileLines("config/help/main_help.cnf")
		fmt.Print("here is list of available commands:\n")
		for _, i := range helpList{
			fmt.Println(i)
		}
	}
}