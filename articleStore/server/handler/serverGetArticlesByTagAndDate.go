package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	model "goprojects/articleStore/models"
	"net/http"
	"strconv"
)

const (
	MAX_ARTICLES_FOR_DAY = 10
	DATE_STR_LENGTH      = 8
	MONTH_START_INDEX    = 4
	DAY_START_INDEX      = 6
)

//convertDateString
//This function validates and convert the date string in request URL path to
//required date format YYYY-MM-DD
//Input:
// - dateStr : date string received from URL path
//Return:
// - date: required formatted date
// - error: error in case date str is not valid
func convertDateString(dateStr string) (string, bool) {
	var date string
	if len(dateStr) != DATE_STR_LENGTH {
		return date, false
	}
	_, err := strconv.Atoi(dateStr)
	if err != nil {
		return date, false
	}
	date = dateStr[:MONTH_START_INDEX] + "-" + dateStr[MONTH_START_INDEX:DAY_START_INDEX] + "-" + dateStr[DAY_START_INDEX:]
	return date, true
}

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

	date, valid := convertDateString(dateStr)
	if valid != true {
		errObj.Code = http.StatusBadRequest
		errObj.Message = fmt.Sprintf("dateStr %s not in required format YYYYMMDD", dateStr)
		writeErrorResp(resp, errObj)
		return
	}

	result.Tag = tagName
	articles, err := getArticlesByDate(date, MAX_ARTICLES_FOR_DAY)
	if err != nil {
		errObj.Code = http.StatusInternalServerError
		errObj.Message = fmt.Sprintf("Error in fetching articles for given date:%+v", err)
		writeErrorResp(resp, errObj)
		return

	}
	result.Articles = articles
	count, _ := getTagCountForDate(date, tagName)
	result.Count = count
	relatedTags, _ := getRelatedTagsforDate(date, tagName)
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
