package handler

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	model "goprojects/articleStore/models"
	"strconv"
	"time"
	//db "goprojects/articleStore/server/handler"
	"net/http"
)

// validateArticle ...
// This functions is used to validate Article object received
//
// Input:
//  - article: Obj to be validated
// Return:
//  - err: returns err with respect to validation failure
func validateArticle(article model.Article) error {
	var err error
	if article.Title == "" {
		err = fmt.Errorf("Title is empty, provide title for article being stored")
		return err
	}
	if article.Date == "" {
		err = fmt.Errorf("Date is empty, provide Date for article being stored")
		return err
	}
	_, err = time.Parse(dateFormat, article.Date)
	if err != nil {
		err = fmt.Errorf("Date %s is not as per allowed format %s", article.Date, dateFormat)
		return err
	}

	return err
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

	id, errMsg := InsertArticle(article)
	if errMsg != nil {
		err := fmt.Errorf("Db Insert operation error %+v", errMsg)
		errObj.Code = http.StatusInternalServerError
		errObj.Message = err.Error()
		writeErrorResp(resp, errObj)
		return
	}
	article.ID = strconv.Itoa(int(id))
	for _, tag := range article.Tags {
		errMsg := InsertTag(id, tag)
		if errMsg != nil {
			err := fmt.Errorf("Db Insert operation error %+v", errMsg)
			errObj.Code = http.StatusInternalServerError
			errObj.Message = err.Error()
			writeErrorResp(resp, errObj)
			return
		}
	}
	log.Info("Article posted: ", article)
	resp.Header().Set("Content-Type", "application/json; charset=UTF-8")
	resp.WriteHeader(http.StatusOK)
	js, _ := json.Marshal(article)
	resp.Write(js)
}
