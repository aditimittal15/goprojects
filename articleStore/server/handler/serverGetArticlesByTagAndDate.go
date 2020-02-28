package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	model "goprojects/articleStore/models"
	"net/http"
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
	articles, err := getArticlesByDate(dateStr, MAX_ARTICLES_FOR_DAY)
	if err != nil {
		errObj.Code = http.StatusInternalServerError
		errObj.Message = fmt.Sprintf("Error in fetching articles for given date:%+v", err)
		writeErrorResp(resp, errObj)
		return

	}
	result.Articles = articles
	count, _ := getTagCountForDate(dateStr, tagName)
	result.Count = count
	relatedTags, _ := getRelatedTagsforDate(dateStr, tagName)
	result.RelatedTags = relatedTags

	resp.Header().Set("Content-Type", "application/json; charset=UTF-8")
	resp.WriteHeader(http.StatusOK)
	js, _ := json.Marshal(result)
	_, err = resp.Write(js)
	if err != nil {
		log.Error("failed to write success response as json ", funcName)
		http.Error(resp, "failed to write response as json", http.StatusInternalServerError)
		return
	}
}
