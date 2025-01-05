package cmd

const (
	UseArgPackages = " [<packages>] "
	UseArgPath     = " [<path>] "

	DefaultPath     = "."
	DefaultPackages = DefaultPath + "/..."
)

func ArgPackages(defaultPackages string, args []string) string {
	if len(args) == 0 {
		return defaultPackages
	}

	return args[len(args)-1]
}
