package cmd

import "strings"

func BuildUse(uses ...string) string {
	return strings.ReplaceAll(strings.Join(uses, " "), " ", " ")
}
