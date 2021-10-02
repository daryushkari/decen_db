package lexer

import (
	"decen_db/internal/utilities"
	"strings"
)

type token struct {
	Value string
	Type  string
}

var keyword = []string{
	"localdb",
	"ledgerdb",
	"set-path",
	"new",
	"drop",
	"show",
	"query",
	"help",
	"col",
	"new_col",
	"show_col",
	"drop_col",
	"insert",
	"find",
	"where",
	"with",
	"update",
	"delete",
}

var separator = []string{
	"{",
	"}",
	":",
	";",
	"[",
	"]",
	"(",
	")",
	",",
}

var operator = []string{
	"+",
	"<",
	">",
	"<=",
	">=",
	"-",
	"=",
	"*",
	"/",
}

var identifierChars = []string{
	"1", "2", "3", "4", "5", "6", "7", "8", "9", "0",
	"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "x", "y", "w", "z", "u", "v",
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "X", "Y", "W", "Z", "U", "V",
	"_",
}

func lexConstTokens(s string) (tk *token) {
	// constant tokens like keyword, separator and operator
	tk = &token{}
	tk.Value = s
	if utilities.CheckStringInSlice(s, keyword) {
		tk.Type = "keyword"
		return tk
	}
	if utilities.CheckStringInSlice(s, separator) {
		tk.Type = "separator"
		return tk
	}
	if utilities.CheckStringInSlice(s, operator) {
		tk.Type = "operator"
		return tk
	}
	return nil
}

func checkIsString(s string) bool {
	if s[0] == '"' && s[len(s)-1] == '"' {
		if strings.Count(s, "\"") != 2 {
			return false
		}
		return true
	}
	return false
}

func checkIsNumber(s string) bool {
	numbers := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	for i := 0; i < len(s); i++ {
		if !utilities.CheckStringInSlice(string(s[i]), numbers) {
			return false
		}
	}
	return true
}

func checkIsBoolean(s string) bool {
	if s == "false" || s == "true" {
		return true
	}
	return false
}

func checkLiteral(s string) (tk *token) {
	tk = &token{}
	tk.Value = s
	if checkIsString(s) {
		tk.Type = "literal_str"
		return tk
	}
	if checkIsNumber(s) {
		tk.Type = "literal_int"
		return tk
	}
	if checkIsBoolean(s) {
		tk.Type = "literal_bool"
		return tk
	}
	return nil
}

func checkIdentifier(s string) (tk *token) {
	for i := 0; i < len(s); i++ {
		if !utilities.CheckStringInSlice(string(s[i]), identifierChars) {
			return nil
		}
	}
	tk = &token{}
	tk.Value = s
	tk.Type = "identifier"
	return tk
}

func getTokenType(s string) (isValid bool, tk *token) {
	tk = lexConstTokens(s)
	if tk != nil {
		return true, tk
	}

	tk = checkLiteral(s)
	if tk != nil {
		return true, tk
	}

	tk = checkIdentifier(s)
	if tk != nil {
		return true, tk
	}

	return false, nil
}

func Lex(cmd string) (isValid bool, tkList []*token) {
	cmdList := strings.Split(cmd, " ")
	for i := 0; i < len(cmdList); i++ {
		isValid, tk := getTokenType(cmdList[i])
		if isValid {
			tkList = append(tkList, tk)
		} else {
			return false, nil
		}
	}
	return true, tkList
}
