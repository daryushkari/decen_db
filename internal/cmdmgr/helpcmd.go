package cmdmgr

import (
	"fmt"
	"os"
)

// HelpCmd shows help
func HelpCmd(){
	fmt.Println(len(os.Args))
	if len(os.Args) < 3{
		fmt.Print("here is list of available commands: \n help \n runfile \n localdatabase \n listen \n getledger")
	}
}