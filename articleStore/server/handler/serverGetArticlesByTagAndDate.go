package handler

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
	//"fmt"
	"github.com/gorilla/mux"
	model "goprojects/articleStore/models"
	//	db "goprojects/articleStore/server/handler"
)

const MAX_ARTICLES_FOR_DAY = 10

// GetArticlesByTagAndDateAPIServiceLogic ...
// This functions is used for main API handling for GetArticlesByTagAndDateAPIServiceLogic
//
// Input:
//  - resp: HTTP Response to be sent to the network
//  - req: HTTP Request coming from the network
//
func GetArticlesByTagAndDateAPIServiceLogic(resp http.ResponseWriter, req *http.Request) {
	var errObj model.Error
	var result model.ArticlesByTagAndDate

	funcName := GetFuncName()
	log.Debug("Function entered ", funcName)
	defer func() {
		log.Debug("Function Exited ", funcName)
	}()

	//extract path variables
	vars := mux.Vars(req)
	tagName := vars["tagName"]
	dateStr := vars["date"]

	//Validate path variable
	if tagName == "" {
		errObj.Code = http.StatusBadRequest
		errObj.Message = "tagName path variable doesn't exist in request"
		writeErrorResp(resp, errObj)
		return
	}
	if dateStr == "" {
		errObj.Code = http.StatusBadRequest
		errObj.Message = "date path variable doesn't exist in request"
		writeErrorResp(resp, errObj)
		return
	}

	result.Tag = tagName
	result.Articles = getArticlesByDate(dateStr, MAX_ARTICLES_FOR_DAY)
	result.Count = getTagCountForDate(dateStr, tagName)
	result.RelatedTags = getRelatedTagsforDate(dateStr, tagName)

	resp.Header().Set("Content-Type", "application/json; charset=UTF-8")
	resp.WriteHeader(http.StatusOK)
	js, _ := json.Marshal(result)
	_, err := resp.Write(js)
	if err != nil {
		log.Error("failed to write success response as json ", funcName)
		http.Error(resp, "failed to write response as json", http.StatusInternalServerError)
		return
	}
}
