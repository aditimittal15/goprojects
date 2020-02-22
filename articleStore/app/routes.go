package main

import (
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	articleStoreService "goprojects/articleStore/server/handler"
	"net/http"
	"strings"
)

// Route defines the service router names, methods, patterns and handlers
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	{
		"CreateArticleAPI",
		strings.ToUpper("Post"),
		"/articles",
		articleStoreService.CreateArticleAPIServiceLogic,
	},
	{
		"GetArticleByIDAPI",
		strings.ToUpper("Get"),
		"/articles/{id}",
		articleStoreService.GetArticleByIDAPIServiceLogic,
	},

	{
		"GetArticlesByTagAndDateAPI",
		strings.ToUpper("Get"),
		"/tags/{tagName}/{date}",
		articleStoreService.GetArticlesByTagAndDateAPIServiceLogic,
	},
}

// NewRouter Creates a New Router
//
// Returns:
//  -  router: New ruter that has been created of type *mux.Router
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler = route.HandlerFunc

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}
