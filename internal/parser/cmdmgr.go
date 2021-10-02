package parser

import (
	"decen_db/internal/dbasemgr"
	"decen_db/internal/parser/lexer"
	"fmt"
	"strings"
)

// CommandManager Main command manager
func CommandManager(inp string) string {

	mainCmdIndex := 0
	val, tokList := lexer.Lex(inp)
	fmt.Println("valid", val, tokList)

	cmd := strings.Split(inp, " ")

	switch cmd[mainCmdIndex] {
	case "help":
		response := HelpCommand(cmd)
		return response
	case "localdb":
		response := dbasemgr.LocalDbManage(cmd)
		return response
	default:
		return "error: " + cmd[mainCmdIndex] + " is an invalid command please enter help command for more information"
	}

}
