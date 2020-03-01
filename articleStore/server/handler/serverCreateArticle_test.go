package handler

import (
	log "github.com/sirupsen/logrus"
	model "goprojects/articleStore/models"
	"net/http"
	"testing"
	"encoding/json"
	"net/http/httptest"
	"bytes"
)

func initLog() {
	formatter := &log.TextFormatter{
		TimestampFormat:        "02-01-2006 15:04:05", // the "time" field configuratiom
		FullTimestamp:          true,
		DisableLevelTruncation: true, // log level field configuration
	}
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(formatter)
}

func initTest() {
	initLog()
	CreateDbConnection()
	CreateDBSchema()
}

/*func TestCreateArticleDbConnError (t *testing.T){
	initLog()
	CreateDbConnection()
	article := model.Article{
                Body:  "This is the body",
                Date:  "2020-02-28",
                ID:    "1",
                Title: "Potato",
        }
        jsonData, _ := json.Marshal(article)
	req, _ := http.
                NewRequest("POST", "/articles", bytes.NewBuffer(jsonData))
	type args struct {
                req    *http.Request
                status int
        }
        tests := []struct {
                name string
                args args
        }{
                {name: "UnSuccessful CreateArticle test case,no db Connection", args: args{req, 500}},
        }
        for _, tt := range tests {
                t.Run(tt.name, func(t *testing.T) {
                        rr := httptest.NewRecorder()
                        CreateArticleAPIServiceLogic(rr, tt.args.req)
                        if rr.Code != tt.args.status {
                                t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, tt.args.status)
                        }
                })
        }
}*/
func TestCreateArticleAPIServiceLogic(t *testing.T) {
	initTest()
	//var article model.Article
	article := model.Article{
		Body:  "This is the body",
		Date:  "2020-02-28",
		ID:    "1",
		Title: "Potato",
	}
	jsonData, _ := json.Marshal(article)

	article = model.Article{
		Body:  "This is the body",
		Date:  "2020-02-28",
		ID:    "1",
		Tags:  []string{"health", "potato"},
		Title: "Potato",
	}

	jsonData1, _ := json.Marshal(article)
	jsonData2, _ := json.Marshal("hello")
	article3 := model.Article{
                Body:  "This is the body",
                Date:  "2020-02-28",
                ID:    "1",
                Tags:  []string{"health", "potato"},
        }
	jsonData3, _ := json.Marshal(article3)
	article4 := model.Article{
                Body:  "This is the body",
                ID:    "1",
		Title: "Potato",
        }
	jsonData4, _ := json.Marshal(article4)
	article5 := model.Article{
                Body:  "This is the body",
                Date:  "02-28-2020",
                ID:    "1",
		Title: "Potato",
        }
	jsonData5, _ := json.Marshal(article5)
	req, _ := http.
		NewRequest("POST", "/articles", bytes.NewBuffer(jsonData))
	req1, _ := http.
		NewRequest("POST", "/articles", bytes.NewBuffer(jsonData1))
	req2, _ := http.
		NewRequest("POST", "/articles", bytes.NewBuffer(jsonData2))
	req3, _ := http.
		NewRequest("POST", "/articles", bytes.NewBuffer(jsonData3))
	req4, _ := http.
		NewRequest("POST", "/articles", bytes.NewBuffer(jsonData4))
	req5, _ := http.
		NewRequest("POST", "/articles", bytes.NewBuffer(jsonData5))

	type args struct {
		req    *http.Request
		status int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "Successful CreateArticle test case,no tags", args: args{req, 200}},
		{name: "Successful CreateArticle test case", args: args{req1, 200}},
		{name: "Unsuccessful CreateArticle test case,request body json not valid", args: args{req2, 400}},
		{name: "Unsuccessful CreateArticle test case,mandatory field title missing", args: args{req3, 400}},
		{name: "Unsuccessful CreateArticle test case,mandatory field date missing", args: args{req4, 400}},
		{name: "Unsuccessful CreateArticle test case,date format incorrect", args: args{req5, 400}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			CreateArticleAPIServiceLogic(rr, tt.args.req)
			if rr.Code != tt.args.status {
				t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, tt.args.status)
			}
		})
	}
}
