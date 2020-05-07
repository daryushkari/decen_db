package collectionmgr


func CollectionManager(cmd []string)(msg string){
	colSubCmdIndex := 3

	switch cmd[colSubCmdIndex] {
	case "new_col":
		return MakeNewCollection(cmd)

	default:
		return ""
	}
}