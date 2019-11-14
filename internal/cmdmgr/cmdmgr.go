package cmdmgr

import (
	"../utilities"
	"fmt"
	"os"
)

// CommandManager Main command manager
func CommandManager(){

	commandList := utilities.ReturnFileLines("config/commands/main_commands.cnf")
	inputCommands := cmdParser()
	
	// check if user hasn't entered any command
	if len(os.Args) < 2{
		fmt.Println("Welcome to DecenDB")
		return
	}

	// check if user hasn't entered valid command
	if !utilities.CheckStringInSlice(os.Args[1],commandList){
		fmt.Println("error: ",inputCommands[1],"is an invalid command please enter help command for more information")
		return
	}

	switch os.Args[1] {
	case "help":
		HelpCmd(inputCommands)
	case "localdb":
		LocaldbManage(inputCommands)
	}
}

func cmdParser()(inputCmd []string){
	for _, i := range os.Args{
		inputCmd = append(inputCmd, i)
	}
	return inputCmd
}
