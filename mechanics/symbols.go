package mechanics

import "strings"

func IsValidSystemSymbol(sysSymbol string) bool {
	cmp := strings.SplitN(sysSymbol, "-", 2)
	if len(cmp) < 2 {
		return false
	}
	if len(cmp[0]) != 2 {
		return false
	}
	if len(cmp[1]) < 2 || len(cmp[1]) > 4 {
		return false
	}
	return true
}
