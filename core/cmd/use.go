package cmd

import "strings"

func BuildUse(args ...string) string {
	var result string

	for _, arg := range args {
		result = result + strings.TrimSpace(arg) + " "
	}

	return strings.TrimSpace(result)
}
