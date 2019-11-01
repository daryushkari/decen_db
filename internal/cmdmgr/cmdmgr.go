package cmdmgr

import (
	"fmt"
	"os"
	"../utilities"
)

// CommandManager Main command manager
func CommandManager(){

	commandList := utilities.ReturnFileLines("config/commands/main_commands.cnf")
	inputCommands = commandParser()
	
	// check if user hasn't entered any command
	if len(os.Args) < 2{
		fmt.Println("Welcome to DecenDB")
		return
	}

	// check if user hasn't entered valid command
	if !utilities.CheckStringInSlice(os.Args[1],commandList){
		fmt.Println("error: ",inputCommands[1],"is an invlid command please enter help command for more information")
		return
	}

	switch(os.Args[1]){
	case "help":
		HelpCmd(inputCommands)
	// case "localdb":

	}
}

func commandParser()(inputCommands []string){
	for _, i := range os.Args{
		inputCommands = append(inputCommands, i)
	}
	return inputCommands
}