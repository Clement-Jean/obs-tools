package pkg

func ParseCommands(args []string, separator string) [][]string {
	var cmdParts []string
	var cmdList [][]string

	for _, arg := range args {
		if arg == separator {
			if len(cmdParts) > 0 {
				cmdList = append(cmdList, cmdParts)
				cmdParts = []string{}
			}
		} else {
			cmdParts = append(cmdParts, arg)
		}
	}

	return cmdList
}
