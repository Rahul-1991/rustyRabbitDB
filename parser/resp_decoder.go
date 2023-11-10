package parser

import "strings"

func DecodeArrayString(reqStr string) []string {
	var commandArgList []string
	strComponents := strings.Split(reqStr, "\r\n")
	for i := 1; i < len(strComponents); i++ {
		if i%2 == 0 {
			commandArgList = append(commandArgList, strComponents[i])
		}
	}
	return commandArgList
}

func DecodeSimpleString(reqStr string) string {
	modifiedStr := strings.TrimRight(reqStr, "\r\n")
	return strings.TrimLeft(modifiedStr, "+")
}
