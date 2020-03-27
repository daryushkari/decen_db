package cmdmgr

// CommandManager Main command manager
func CommandManager(cmd []string) string {

	mainCmdIndex := 0

	switch cmd[mainCmdIndex] {
		case "help":
			response := HelpCommand(cmd)
			return response
		//case "localdb":
			//LocaldbManage(cmd)
		default:
			return "error: " + cmd[mainCmdIndex] + " is an invalid command please enter help command for more information"
	}

}
