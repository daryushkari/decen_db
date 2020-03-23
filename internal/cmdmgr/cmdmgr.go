package cmdmgr

import (
	"decen_db/internal/utilities"
)

// CommandManager Main command manager
func CommandManager(cmd []string) string {

	mainCmdIndex := 0
	commandList, err := utilities.ReturnFileLines("config/commands/main_commands.cnf")

	if err != nil {
		return err.Error()
	}

	// check if user hasn't entered valid command
	if !utilities.CheckStringInSlice(cmd[mainCmdIndex], commandList) {
		return "error: " + cmd[mainCmdIndex] + " is an invalid command please enter help command for more information"
	}

	switch cmd[mainCmdIndex] {
	case "help":
		response := HelpCommand(cmd)
		return response
		//case "localdb":
		//	LocaldbManage(cmd)
	}
	return " "

}
