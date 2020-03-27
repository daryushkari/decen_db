package cmdmgr

import (
	"decen_db/internal/utilities"
)

var helpFileName = map[string]string{
	"localdb": "config/help/localdb_help.cnf",
	"main":    "config/help/main_help.cnf",
}

// HelpCommand reads from related commands file which user wants to know about it and
// returns files contents as string
func HelpCommand(cmd []string) string {

	helpList, err := utilities.ReturnFileLines(returnCommandFileName(cmd))
	if err != nil {
		return err.Error()
	}

	answer := "here is list of available commands:\n"
	for _, i := range helpList {
		answer += i + "\n"
	}

	return answer

}

func returnCommandFileName(cmd []string) string {
	// command Index which user wants to get help about it
	wantedCmdIndex := 1

	if len(cmd) < 2 {
		return helpFileName["main"]
	}

	if helpFileName[cmd[wantedCmdIndex]] == "" {
		return helpFileName["main"]
	}

	return helpFileName[cmd[wantedCmdIndex]]
}
