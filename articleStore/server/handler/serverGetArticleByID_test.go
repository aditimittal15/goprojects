package handler

import (
	"net/http"
	//"fmt"
	//log "github.com/sirupsen/logrus"
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	model "goprojects/articleStore/models"
	"net/http/httptest"
	"testing"
)

func TestGetArticleByIDAPIServiceLogic(t *testing.T) {
	initTest()

	article := model.Article{
		Body:  "This is the body",
		Date:  "2020-02-28",
		ID:    "1",
		Tags:  []string{"health", "potato"},
		Title: "Potato",
	}

	jsonData1, _ := json.Marshal(article)
	createreq1, _ := http.
		NewRequest("POST", "/articles", bytes.NewBuffer(jsonData1))
	rr := httptest.NewRecorder()
	CreateArticleAPIServiceLogic(rr, createreq1)
	//fmt.Println("************\n",rr.Body)
	buf := new(bytes.Buffer)
	buf.ReadFrom(rr.Body)
	json.Unmarshal(buf.Bytes(), &article)
	/*	if err = json.Unmarshal(resp.Body, &article); err != nil {
		        return nil, err
		}
	*/

	req, _ := http.
		NewRequest("GET", "articles/{id}", nil)
	req1, _ := http.
		NewRequest("GET", "articles/{id}", nil)
	req1 = mux.SetURLVars(req1, map[string]string{"id": "hello"})
	req2, _ := http.
		NewRequest("GET", "articles/{id}", nil)
	req2 = mux.SetURLVars(req1, map[string]string{"id": "3782436"})
	req3, _ := http.
		NewRequest("GET", "articles/{id}", nil)
	req3 = mux.SetURLVars(req1, map[string]string{"id": string(article.ID)})

	type args struct {
		req    *http.Request
		status int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "Unsuccesful GetArticleByID, no Id", args: args{req, 400}},
		{name: "Unsuccesful GetArticleByID, invalid Id", args: args{req1, 400}},
		{name: "Unsuccesful GetArticleByID, record doesn't exist", args: args{req2, 404}},
		{name: "Succesful GetArticleByID, record doesn't exist", args: args{req3, 200}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			GetArticleByIDAPIServiceLogic(rr, tt.args.req)
			if rr.Code != tt.args.status {
				t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, tt.args.status)
			}
		})
	}
}
