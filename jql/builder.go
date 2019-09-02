package jql

import "strings"

func BuildJQL(args []string) string {
	return strings.Join(args, " AND ")
}
