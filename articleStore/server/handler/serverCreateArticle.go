package handler

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	db "goprojects/articleStore/dbwrapper"
	model "goprojects/articleStore/models"
	"net/http"
)

func validateArticle(article model.Article) error {

	return nil
}

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

// CreateArticleAPIServiceLogic ...
// This functions is used for main API handling for CreateArticleAPIServiceLogic
//
// Input:
//  - resp: HTTP Response to be sent to the network
//  - req: HTTP Request coming from the network
//
func CreateArticleAPIServiceLogic(resp http.ResponseWriter, req *http.Request) {
	var (
		article model.Article
		errObj  model.Error
	)
	funcName := GetFuncName()
	log.Debug("Function entered ", funcName)
	defer func() {
		log.Debug("Function Exited", funcName)
	}()

	//Decode request body
	decoder := json.NewDecoder(req.Body)
	errMsg := decoder.Decode(&article)
	if errMsg != nil {
		err := fmt.Errorf("request body decoding error %+v", errMsg)
		errObj.Code = http.StatusBadRequest
		errObj.Message = err.Error()
		writeErrorResp(resp, errObj)
		return
	}

	errMsg = validateArticle(article)
	if errMsg != nil {
		err := fmt.Errorf("request body validation error %+v", errMsg)
		errObj.Code = http.StatusBadRequest
		errObj.Message = err.Error()
		writeErrorResp(resp, errObj)
		return
	}

	errMsg = db.Insert(article)
	if errMsg != nil {
		err := fmt.Errorf("Db Insert operation error %+v", errMsg)
		errObj.Code = http.StatusInternalServerError
		errObj.Message = err.Error()
		writeErrorResp(resp, errObj)
		return
	}
	log.Info("Article posted: ", article)

	resp.Header().Set("Content-Type", "application/json; charset=UTF-8")
	resp.WriteHeader(http.StatusCreated)
}
