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
	insertArticleString       = "INSERT INTO articles (Title, Body,Date,Created) VALUES (?,?,?,?)"
	selectArticleString       = "SELECT id,Title,Body,Date FROM articles "
	selectTagsString          = "SELECT Tag FROM tags "
	insertTagString           = "INSERT INTO tags (ArticleId,Tag) VALUES (?,?)"
	dateFormat                = "2006-01-02"
	createdDateFormat         = "2014-09-12T11:45:26.371Z"
	createArticlesTableString = "CREATE TABLE IF NOT EXISTS articles (id INTEGER PRIMARY KEY, Title TEXT, Body TEXT,Date TIMESTAMP, Created TIMESTAMP)"
	createTagsTableString     = "CREATE TABLE IF NOT EXISTS tags (id INTEGER PRIMARY KEY, ArticleId INTEGER, Tag Text,FOREIGN KEY (ArticleId) REFERENCES articles(id))"
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
func InsertTag(id int64, tag string) error {
	statement, err := Conn.Prepare(insertTagString)
	if err != nil {
		return err
	}

	_, err = statement.Exec(id, tag)
	return err
}

func InsertArticle(article model.Article) (int64, error) {
	var createdTime = time.Now().UTC()
	var id int64
	date, err := time.Parse(dateFormat, article.Date)
	if err != nil {
		return id, err
	}
	statement, err := Conn.Prepare(insertArticleString)
	if err != nil {
		return id, err
	}

	result, err := statement.Exec(article.Title, article.Body, date.Format(dateFormat), createdTime.Format(createdDateFormat))
	id, err = result.LastInsertId()
	if err != nil {
		return id, err
	}
	return id, err
}

func getTags(article *model.Article) error {
	stmtStr := selectTagsString + "where ArticleId = " + article.ID
	log.Debug("DB Query executed: ", stmtStr)
	rows, err := Conn.Query(stmtStr)
	if err != nil {
		return err
	}
	for rows.Next() {
		var tag string
		rows.Scan(&tag)
		article.Tags = append(article.Tags, tag)
	}
	return err
}

func GetArticle(id string) (model.Article, error) {
	var (
		article model.Article
	)
	stmtStr := selectArticleString + "where id = " + id
	log.Debug("DB Query executed: ", stmtStr)
	rows, err := Conn.Query(stmtStr)
	if err != nil {
		return article, err
	}
	for rows.Next() {
		var id int
		rows.Scan(&id, &article.Title, &article.Body, &article.Date)
		article.ID = strconv.Itoa(id)
	}
	if article.ID == "" {
		return article, err
	}
	getTags(&article)

	return article, err
}
func getArticlesByDate(date string, limit int) []string {
	var result []string
	stmtStr := fmt.Sprintf("select id from articles where Date = %s ORDER BY id DESC LIMIT %d", date, limit)
	fmt.Println(stmtStr)
	return result

}
func getTagCountForDate(date string, tag string) int32 {
	var result int32
	return result
}
func getRelatedTagsforDate(date string, tag string) []string {

	var result []string
	return result
}
