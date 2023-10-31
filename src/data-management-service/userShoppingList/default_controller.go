package userShoppingList

import (
	"encoding/json"
	"hsfl.de/group6/hsfl-master-ai-cloud-engineering/data-management-service/userShoppingList/model"
	"net/http"
	"strconv"
)

type defaultController struct {
	userShoppingListRepository Repository
}

func NewDefaultController(userShoppingListRepository Repository) *defaultController {
	return &defaultController{userShoppingListRepository}
}

func (controller defaultController) GetList(writer http.ResponseWriter, request *http.Request) {
	userId, err := strconv.ParseUint(request.Context().Value("userId").(string), 10, 64)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	value, err := controller.userShoppingListRepository.FindById(userId)
	if err != nil {
		if err.Error() == ErrorListNotFound {
			http.Error(writer, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}

	writer.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(writer).Encode(value)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}

func (controller defaultController) PutList(writer http.ResponseWriter, request *http.Request) {
	userId, err := strconv.ParseUint(request.Context().Value("userId").(string), 10, 64)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	if _, err := controller.userShoppingListRepository.Create(&model.UserShoppingList{
		UserId: userId,
	}); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusCreated)
}

func (controller defaultController) DeleteList(writer http.ResponseWriter, request *http.Request) {
	userId, err := strconv.ParseUint(request.Context().Value("userId").(string), 10, 64)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	if err := controller.userShoppingListRepository.Delete(&model.UserShoppingList{UserId: userId}); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}