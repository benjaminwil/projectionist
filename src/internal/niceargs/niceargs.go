package niceargs

import (
	"strings"
)

func List(args []string) []string {
	flags := make([]string, 0, len(args))
	arguments := make([]string, 0, len(args))

	for index, arg := range args {
		if isFlagPair(arg) {
			flags = append(flags, arg)
			continue
		}

		if isFlag(arg) && !hasEquals(arg) {
			flags = append(flags, arg)
			continue
		}

		if !isFlag(arg) && index > 0 {
			previousArg := args[index-1]

			if isFlag(previousArg) && !hasEquals(previousArg) {
				flags = append(flags, arg)
			} else {
				arguments = append(arguments, arg)
			}
			continue
		}

		if !isFlag(arg) {
			arguments = append(arguments, arg)
			continue
		}
	}

	return append(flags, arguments...)
}

func isFlag(s string) bool {
	if string(s[0]) == "-" {
		return true
	}

	return false
}

func isFlagPair(s string) bool {
	if isFlag(s) && hasEquals(s) {
		return true
	}

	return false
}

func hasEquals(s string) bool {
	if strings.Contains(s, "=") {
		return true
	}

	return false
}
