package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	model "goprojects/articleStore/models"
	//	db "goprojects/articleStore/server/handler"
	"net/http"
)

// GetArticleByIDAPIServiceLogic ...
// This functions is used for main API handling for GetArticleByIDAPIServiceLogic
//
// Input:
//  - resp: HTTP Response to be sent to the network
//  - req: HTTP Request coming from the network
//
func GetArticleByIDAPIServiceLogic(resp http.ResponseWriter, req *http.Request) {
	var errObj model.Error
	funcName := GetFuncName()
	log.Debug("Function entered ", funcName)
	defer func() {
		log.Debug("Function Exited ", funcName)
	}()

	//extract path variables
	vars := mux.Vars(req)
	ID := vars["id"]

	//Get article from DB
	article, errMsg := GetArticle(ID)
	if errMsg != nil {

		err := fmt.Errorf("Db Get operation error %+v ", errMsg)
		errObj.Code = http.StatusInternalServerError
		errObj.Message = err.Error()
		writeErrorResp(resp, errObj)
		return
	}
	if article.ID == "" {
		errObj.Code = http.StatusNotFound
		errObj.Message = "Record doesn't exist"
		writeErrorResp(resp, errObj)
		return
	}

	resp.WriteHeader(http.StatusOK)
	js, _ := json.Marshal(article)
	_, err := resp.Write(js)
	if err != nil {
		log.Error("failed to write success response as json ", funcName)
		http.Error(resp, "failed to write response as json", http.StatusInternalServerError)
		return
	}

	resp.Header().Set("Content-Type", "application/json; charset=UTF-8")
}
