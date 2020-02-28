package handler

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	model "goprojects/articleStore/models"
	"net/http"
	"runtime"
	"strings"
)

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

// writeErrorResp
// writes Error object in HTTP response
// Return:
//  -
func writeErrorResp(resp http.ResponseWriter, errObj model.Error) {
	log.Error(errObj.Message)
	resp.WriteHeader(int(errObj.Code))
	js, _ := json.Marshal(errObj)
	_, err := resp.Write(js)
	if err != nil {
		log.Error("failed to write error response as json")
		http.Error(resp, "failed to write response as json", http.StatusInternalServerError)
		return
	}
	return
}
