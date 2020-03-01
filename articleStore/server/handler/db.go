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
	databaseName     = "sqlite3"
	databaseFileName = "./articles.db"

	createArticlesTableString = "CREATE TABLE IF NOT EXISTS articles (id INTEGER PRIMARY KEY, Title TEXT, Body TEXT,Date TIMESTAMP)"
	createTagsTableString     = "CREATE TABLE IF NOT EXISTS tags (id INTEGER PRIMARY KEY, ArticleId INTEGER, Tag Text,FOREIGN KEY (ArticleId) REFERENCES articles(id))"
	insertArticleString       = "INSERT INTO articles (Title, Body,Date) VALUES (?,?,?)"
	insertTagString           = "INSERT INTO tags (ArticleId,Tag) VALUES (?,?)"
	selectArticleString       = "SELECT id,Title,Body,Date FROM articles "
	selectTagsString          = "SELECT Tag FROM tags "
	articlesByDateQuery       = "select id from articles where Date = \"%s\" ORDER BY id DESC LIMIT %d"
	tagCountForDateQuery      = "select count(*) from articles inner join tags on articles.id = tags.ArticleId where Date = \"%s\" and Tag = \"%s\";"
	relatedTagsForDateQuery   = "select DISTINCT Tag from tags where ArticleId IN (select articles.id from articles inner join tags on articles.id = tags.ArticleId where Date = \"%s\" and Tag = \"%s\") and Tag != \"%s\""

	dateFormat = "2006-01-02"
)

//sqlite DB connection
//Can be used concurrently
var Conn *sql.DB

//CreateDbConnection
// Function to create and store DB connection handle
func CreateDbConnection() {
	var err error
	funcName := GetFuncName()
	log.Debug("Function entered ", funcName)
	defer func() {
		log.Debug("Function exited ", funcName)
	}()
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

//CreateDBSchema..
//Function to create required DBschema/tables
//server panic when error occurs in db schema creation
func CreateDBSchema() {
	funcName := GetFuncName()
	log.Debug("Function entered ", funcName)
	defer func() {
		log.Debug("Function exited ", funcName)
	}()
	statement, err := Conn.Prepare(createArticlesTableString)
	checkPanicError(err)
	_, err = statement.Exec()
	checkPanicError(err)
	statement, err = Conn.Prepare(createTagsTableString)
	checkPanicError(err)
	_, err = statement.Exec()
	checkPanicError(err)
}

//InsertTag
// Function to insert Tag record in DB
// Input:
// - id : Article Id to which tag belongs
// - tag: tag to store
// Return:
// - err: failure occured in DB query execution
func InsertTag(id int64, tag string) error {
	funcName := GetFuncName()
	log.Debug("Function entered ", funcName)
	defer func() {
		log.Debug("Function exited ", funcName)
	}()
	statement, err := Conn.Prepare(insertTagString)
	if err != nil {
		return err
	}
	log.Debug("DB Query executed: ", statement)

	_, err = statement.Exec(id, tag)
	return err
}

//InsertArticle
//Function inserts Article in DB
//Input:
// - article : Article to be inserted
//Returns:
// - id : Article Id generated
// - err : error in case failure in db query execution
func InsertArticle(article model.Article) (int64, error) {
	var id int64
	funcName := GetFuncName()
	log.Debug("Function entered ", funcName)
	defer func() {
		log.Debug("Function exited ", funcName)
	}()
	date, _ := time.Parse(dateFormat, article.Date)
	statement, err := Conn.Prepare(insertArticleString)
	if err != nil {
		return id, err
	}
	log.Debug("DB Query executed: ", statement)

	result, err := statement.Exec(article.Title, article.Body, date.Format(dateFormat))
	id, err = result.LastInsertId()
	if err != nil {
		return id, err
	}
	return id, err
}

//getTags
//Function to get/store tags for/into given article object
//Input:
// - article: Article object in which tags to store
//Return
// - err : incase failure in DB query execution
func getTags(article *model.Article) error {
	funcName := GetFuncName()
	log.Debug("Function entered ", funcName)
	defer func() {
		log.Debug("Function exited ", funcName)
	}()
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

//GetArticle
//Function to get article from DB for given article id
//Input:
// - id : id of the article to fetch from DB
//Return
// - err : incase failure in DB query execution
func GetArticle(id string) (model.Article, error) {
	var (
		article model.Article
	)
	funcName := GetFuncName()
	log.Debug("Function entered ", funcName)
	defer func() {
		log.Debug("Function exited ", funcName)
	}()
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

//getArticlesByDate
//Function to get latest article ids by date from DB with given limit/count
//Input:
// - date : date for which articles to get
// - limit : limit or count of latest articles to get
//Return
// - result : array of article id strings
// - err : incase failure in DB query execution
func getArticlesByDate(date string, limit int) ([]string, error) {
	var result []string
	funcName := GetFuncName()
	log.Debug("Function entered ", funcName)
	defer func() {
		log.Debug("Function exited ", funcName)
	}()
	stmtStr := fmt.Sprintf(articlesByDateQuery, date, limit)
	log.Debug("DB Query executed: ", stmtStr)
	rows, err := Conn.Query(stmtStr)
	if err != nil {
		return result, err
	}
	for rows.Next() {
		var id int
		rows.Scan(&id)
		result = append(result, strconv.Itoa(id))
	}
	return result, err
}

//getTagCountForDate
//Function to get count of given tag on given date
//Input:
// - date : date for which tag count to get
// - tag : tag for which count to calculate
//Return
// - result : desired count of tags
// - err : incase failure in DB query execution
func getTagCountForDate(date string, tag string) (int32, error) {
	var result int32
	funcName := GetFuncName()
	log.Debug("Function entered ", funcName)
	defer func() {
		log.Debug("Function exited ", funcName)
	}()
	stmtStr := fmt.Sprintf(tagCountForDateQuery, date, tag)
	log.Debug("DB Query executed: ", stmtStr)
	rows, err := Conn.Query(stmtStr)
	if err != nil {
		return result, err
	}
	for rows.Next() {
		rows.Scan(&result)
	}
	return result, err
}

//getRelatedTagsforDate
//Function to get unique related tags for given tag and date
//Input:
// - date : date for which tags to be fetched
// - tag : tag for which related tags to be fetched
//Return
// - result : desired list of unique tags
// - err : incase failure in DB query execution
func getRelatedTagsforDate(date string, tag string) ([]string, error) {
	var result []string
	funcName := GetFuncName()
	log.Debug("Function entered ", funcName)
	defer func() {
		log.Debug("Function exited ", funcName)
	}()
	stmtStr := fmt.Sprintf(relatedTagsForDateQuery, date, tag, tag)
	log.Debug("DB Query executed: ", stmtStr)
	rows, err := Conn.Query(stmtStr)
	if err != nil {
		return result, err
	}
	for rows.Next() {
		var tag string
		rows.Scan(&tag)
		result = append(result, tag)
	}
	return result, err
}
