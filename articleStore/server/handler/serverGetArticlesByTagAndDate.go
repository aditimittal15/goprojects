package handler

import (
	log "github.com/sirupsen/logrus"
	"net/http"
)

// GetArticlesByTagAndDateAPIServiceLogic ...
// This functions is used for main API handling for GetArticlesByTagAndDateAPIServiceLogic
//
// Input:
//  - resp: HTTP Response to be sent to the network
//  - req: HTTP Request coming from the network
//
func GetArticlesByTagAndDateAPIServiceLogic(resp http.ResponseWriter, req *http.Request) {
	var ptr interface{}
	log.Debug("ptr", ptr)

	resp.Header().Set("Content-Type", "application/json; charset=UTF-8")
	resp.WriteHeader(http.StatusOK)
}
