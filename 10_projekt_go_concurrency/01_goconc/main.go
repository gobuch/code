package main

func main() {

}

type cmdArgs struct {
	name string
	args []string
}

const (
	cmdDelimiter = "::"
)

func parseArgs(args []string) []cmdArgs {
	if len(args) == 0 {
		return nil
	}
	var cmds []cmdArgs
	var cmd cmdArgs
	for _, arg := range args {
		switch {
		case arg == cmdDelimiter:
			cmds = append(cmds, cmd)
			cmd = cmdArgs{}
		case cmd.name == "":
			cmd.name = arg
		default:
			cmd.args = append(cmd.args, arg)
		}
	}
	cmds = append(cmds, cmd)
	return cmds
}
