package querymgr

func QueryManager(cmd []string) string {
	queryCmdIndex := 3
	switch cmd[queryCmdIndex] {
	case "update":
		return ""
	case "insert":
		return ""
	case "delete":
		return ""
	case "find":
		return ""
	default:
		return ""
	}
}
