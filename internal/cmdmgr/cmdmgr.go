package cmdmgr

import(
	"decen_db/internal/localdbmgr"
)

// CommandManager Main command manager
func CommandManager(cmd []string) string {

	mainCmdIndex := 0

	switch cmd[mainCmdIndex] {
		case "help":
			response := HelpCommand(cmd)
			return response
		case "localdb":
			response := localdbmgr.LocaldbManage(cmd)
			return response
		default:
			return "error: " + cmd[mainCmdIndex] + " is an invalid command please enter help command for more information"
	}

}
