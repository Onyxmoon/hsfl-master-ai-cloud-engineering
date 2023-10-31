package router

import (
	"hsfl.de/group6/hsfl-master-ai-cloud-engineering/shoppinglist-service/userShoppingList"
	"hsfl.de/group6/hsfl-master-ai-cloud-engineering/shoppinglist-service/userShoppingListEntry"
	"net/http"

	"hsfl.de/group6/hsfl-master-ai-cloud-engineering/lib/router"
)

type Router struct {
	router http.Handler
}

func New(shoppingListController userShoppingList.Controller, shoppingListEntryController userShoppingListEntry.Controller) *Router {
	r := router.New()

	r.GET("/api/v1/shoppinglist/:userId", shoppingListController.GetLists)
	r.GET("/api/v1/shoppinglist/:listId/:userId", shoppingListController.GetList)
	r.PUT("/api/v1/shoppinglist/:listId/:userId", shoppingListController.PutList)
	r.POST("/api/v1/shoppinglist/:userId", shoppingListController.PostList)
	r.DELETE("/api/v1/shoppinglist/:listId", shoppingListController.DeleteList)

	r.GET("/api/v1/shoppinglistEntry/:listId", shoppingListEntryController.GetEntries)
	r.GET("/api/v1/shoppinglistEntry/:listId/:productId", shoppingListEntryController.GetEntry)
	r.PUT("/api/v1/shoppinglistEntry/:listId/:productId", shoppingListEntryController.PutEntry)
	r.POST("/api/v1/shoppinglistEntry/:listId/:productId", shoppingListEntryController.PostEntry)
	r.DELETE("/api/v1/shoppinglistEntry/:listId/:productId", shoppingListEntryController.DeleteEntry)

	return &Router{r}
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router.router.ServeHTTP(w, r)
}
