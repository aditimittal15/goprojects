package handler

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
	model "goprojects/articleStore/models"
	"strconv"
	"time"
)

const (
	databaseName              = "sqlite3"
	databaseFileName          = "./articles.db"
	insertArticleString       = "INSERT INTO articles (Title, Body,Date) VALUES (?,?,?)"
	selectArticleString       = "SELECT * FROM articles "
	dateFormat                = "2006-01-02"
	createArticlesTableString = "CREATE TABLE IF NOT EXISTS articles (id INTEGER PRIMARY KEY, Title TEXT, Body TEXT,Date TIMESTAMP)"
	createTagsTableString     = "CREATE TABLE IF NOT EXISTS tags (id INTEGER PRIMARY KEY, Article_Id INTEGER, Tag Text)"
)

var Conn *sql.DB

func CreateDbConnection() {
	var err error
	Conn, err = sql.Open(databaseName, databaseFileName)
	if err != nil {
		err = fmt.Errorf("DB connection error %+v", err)
		panic(err.Error())
	}
}

func checkPanicError(err error) {
	if err != nil {
		panic(err.Error())
	}

}
func CreateDBSchema() {
	statement, err := Conn.Prepare(createArticlesTableString)
	checkPanicError(err)
	_, err = statement.Exec()
	checkPanicError(err)
	statement, err = Conn.Prepare(createTagsTableString)
	checkPanicError(err)
	_, err = statement.Exec()
	checkPanicError(err)
	//var datetime = time.Now().UTC()
	/*statement, _ = Conn.Prepare("INSERT INTO articles (Title, Body,Date) VALUES (?,?,?)")
	  statement.Exec("hello", "this is me",datetime.Format("2006-01-02"))
	  rows, _ := Conn.Query("SELECT * FROM articles")*/
}

func Insert(article model.Article) error {
	var datetime = time.Now().UTC()
	statement, err := Conn.Prepare(insertArticleString)
	if err != nil {
		return err
	}
	_, err = statement.Exec(article.Title, article.Body, datetime.Format(dateFormat))
	return err
}

func GetArticle(whereClause string) ([]model.Article, error) {
	var (
		articles = make([]model.Article, 0)
	)
	stmtStr := selectArticleString
	if whereClause != "" {
		stmtStr = stmtStr + "where " + whereClause
	}
	log.Debug("DB Query executed: ", stmtStr)
	rows, err := Conn.Query(stmtStr)
	if err != nil {
		return articles, err
	}
	for rows.Next() {
		var article model.Article
		var id int
		rows.Scan(&id, &article.Title, &article.Body, &article.Date)
		article.ID = strconv.Itoa(id)
		articles = append(articles, article)

	}
	log.Debug("rows count returned ", len(articles))
	return articles, err
}
