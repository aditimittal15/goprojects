package handler

import (
	"runtime"
	"strings"
)

/*****************************************************************************/
// GetFuncName
// Return:
//  - funcName: Name of an API function

func GetFuncName() string {
	var funcName string
	programCounter, _, _, ok := runtime.Caller(1)
	if !ok {
		return ""
	}
	details := runtime.FuncForPC(programCounter)
	funcDetails := strings.Split(details.Name(), ".")
	funcName = funcDetails[len(funcDetails)-1]
	return funcName
}
