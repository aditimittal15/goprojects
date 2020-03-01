package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	//log "github.com/sirupsen/logrus"
	model "goprojects/articleStore/models"
	//"strconv"
	//"time"
	"github.com/gorilla/mux"
	"net/http/httptest"
	"testing"
)

func TestGetArticlesByTagAndDateAPIServiceLogic(t *testing.T) {
	initTest()
	article := model.Article{
		Body:  "GetArticlesByTagAndDate",
		Date:  "2020-02-01",
		ID:    "1",
		Tags:  []string{"test", "get"},
		Title: "test",
	}

	jsonData1, _ := json.Marshal(article)
	createreq1, _ := http.
		NewRequest("POST", "/articles", bytes.NewBuffer(jsonData1))
	rr := httptest.NewRecorder()
	CreateArticleAPIServiceLogic(rr, createreq1)
	req, _ := http.
		NewRequest("GET", "/tags/{tagName}/{date}", nil)
	req1, _ := http.
		NewRequest("GET", "/tags/{tagName}/{date}", nil)
	req1 = mux.SetURLVars(req1, map[string]string{"tagName": "hello"})
	req2, _ := http.
		NewRequest("GET", "/tags/{tagName}/{date}", nil)
	req2 = mux.SetURLVars(req2, map[string]string{"tagName": "test",
		"date": "20200201"})
	req3, _ := http.
		NewRequest("GET", "/tags/{tagName}/{date}", nil)
	req3 = mux.SetURLVars(req3, map[string]string{"tagName": "test",
		"date": "2020020"})
	req4, _ := http.
		NewRequest("GET", "/tags/{tagName}/{date}", nil)
	req4 = mux.SetURLVars(req4, map[string]string{"tagName": "test",
		"date": "202002df"})

	type args struct {
		req    *http.Request
		status int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "Unsuccesful GetArticlesByTagAndDate, tagname not given", args: args{req, 400}},
		{name: "Unsuccesful GetArticlesByTagAndDate, datestr not given", args: args{req1, 400}},
		{name: "Succesful GetArticlesByTagAndDate", args: args{req2, 200}},
		{name: "Unsuccesful GetArticlesByTagAndDate,dateStr invalid length", args: args{req3, 400}},
		{name: "Unsuccesful GetArticlesByTagAndDate,dateStr invalid value", args: args{req4, 400}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			GetArticlesByTagAndDateAPIServiceLogic(rr, tt.args.req)
			if rr.Code != tt.args.status {
				t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, tt.args.status)
				return
			}
			buf := new(bytes.Buffer)
			buf.ReadFrom(rr.Body)
			var result model.ArticlesByTagAndDate
			json.Unmarshal(buf.Bytes(), &result)
			fmt.Println(result)
		})
	}
}
