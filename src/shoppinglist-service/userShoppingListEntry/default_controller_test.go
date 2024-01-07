package userShoppingListEntry

import (
	"context"
	"hsfl.de/group6/hsfl-master-ai-cloud-engineering/shoppinglist-service/userShoppingList"
	listModel "hsfl.de/group6/hsfl-master-ai-cloud-engineering/shoppinglist-service/userShoppingList/model"
	entryModel "hsfl.de/group6/hsfl-master-ai-cloud-engineering/shoppinglist-service/userShoppingListEntry/model"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func TestNewDefaultController(t *testing.T) {
	type args struct {
		userShoppingListEntryRepository Repository
		userShoppingListRepository      userShoppingList.Repository
	}
	tests := []struct {
		name string
		args args
		want *DefaultController
	}{
		{
			name: "Test construction with DemoRepository",
			args: args{userShoppingListEntryRepository: NewDemoRepository(), userShoppingListRepository: userShoppingList.NewDemoRepository()},
			want: &DefaultController{userShoppingListEntryRepository: NewDemoRepository(), userShoppingListRepository: userShoppingList.NewDemoRepository()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDefaultController(tt.args.userShoppingListEntryRepository, tt.args.userShoppingListRepository); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDefaultController() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDefaultController_GetEntries(t *testing.T) {
	t.Run("Get entries for a shopping list (expect 200)", func(t *testing.T) {
		controller := DefaultController{
			userShoppingListEntryRepository: setupMockEntryRepository(),
		}
		writer := httptest.NewRecorder()
		request := httptest.NewRequest("GET", "/api/v1/shoppinglistentries/1", nil)
		ctx := request.Context()
		ctx = context.WithValue(ctx, "listId", "1")
		request = request.WithContext(ctx)

		controller.GetEntries(writer, request)

		if writer.Code != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, writer.Code)
		}

		// Add assertions for the response data if needed.
	})

	t.Run("Malformed request (expect 400)", func(t *testing.T) {
		controller := DefaultController{
			userShoppingListEntryRepository: setupMockEntryRepository(),
		}
		writer := httptest.NewRecorder()
		request := httptest.NewRequest("GET", "/api/v1/shoppinglistentries/abc", nil)
		ctx := request.Context()
		ctx = context.WithValue(ctx, "listId", "abc")
		request = request.WithContext(ctx)

		controller.GetEntries(writer, request)

		if writer.Code != http.StatusBadRequest {
			t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, writer.Code)
		}
	})
}

func TestDefaultController_GetEntry(t *testing.T) {
	t.Run("Get an existing shopping list entry (expect 200)", func(t *testing.T) {
		controller := DefaultController{
			userShoppingListEntryRepository: setupMockEntryRepository(),
		}
		writer := httptest.NewRecorder()
		request := httptest.NewRequest("GET", "/api/v1/shoppinglistentries/1/2", nil)
		ctx := request.Context()
		ctx = context.WithValue(ctx, "listId", "1")
		ctx = context.WithValue(ctx, "productId", "2")
		request = request.WithContext(ctx)

		controller.GetEntry(writer, request)

		if writer.Code != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, writer.Code)
		}
	})
}

func TestDefaultController_PostEntry(t *testing.T) {
	t.Run("Create a shopping list entry (expect 201)", func(t *testing.T) {
		controller := DefaultController{
			userShoppingListEntryRepository: setupMockEntryRepository(),
		}
		writer := httptest.NewRecorder()
		requestBody := `{"count": 2, "note": "Test entry", "checked": false}`
		request := httptest.NewRequest("POST", "/api/v1/shoppinglistentries/1/2", strings.NewReader(requestBody))
		ctx := request.Context()
		ctx = context.WithValue(ctx, "listId", "2")
		ctx = context.WithValue(ctx, "productId", "2")
		request = request.WithContext(ctx)

		controller.PostEntry(writer, request)

		if writer.Code != http.StatusCreated {
			t.Errorf("Expected status code %d, got %d", http.StatusCreated, writer.Code)
		}
	})

	t.Run("Malformed JSON request (expect 400)", func(t *testing.T) {
		controller := DefaultController{
			userShoppingListEntryRepository: setupMockEntryRepository(),
		}
		writer := httptest.NewRecorder()
		request := httptest.NewRequest("POST", "/api/v1/shoppinglistentries/1/2", strings.NewReader(`{"count": 2`))
		ctx := request.Context()
		ctx = context.WithValue(ctx, "listId", "1")
		ctx = context.WithValue(ctx, "productId", "2")
		request = request.WithContext(ctx)

		controller.PostEntry(writer, request)

		if writer.Code != http.StatusBadRequest {
			t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, writer.Code)
		}
	})
}

func TestDefaultController_PutEntry(t *testing.T) {
	t.Run("Update an existing shopping list entry (expect 200)", func(t *testing.T) {
		controller := DefaultController{
			userShoppingListEntryRepository: setupMockEntryRepository(),
		}
		writer := httptest.NewRecorder()
		requestBody := `{"count": 5, "note": "Updated entry", "checked": true}`
		request := httptest.NewRequest("PUT", "/api/v1/shoppinglistentries/1/2", strings.NewReader(requestBody))
		ctx := request.Context()
		ctx = context.WithValue(ctx, "listId", "1")
		ctx = context.WithValue(ctx, "productId", "2")
		request = request.WithContext(ctx)
		controller.PutEntry(writer, request)

		if writer.Code != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, writer.Code)
		}
	})

	t.Run("Malformed JSON request (expect 400)", func(t *testing.T) {
		controller := DefaultController{
			userShoppingListEntryRepository: setupMockEntryRepository(),
		}
		writer := httptest.NewRecorder()
		request := httptest.NewRequest("PUT", "/api/v1/shoppinglistentries/1/2", strings.NewReader(`{"count": 5`))
		ctx := request.Context()
		ctx = context.WithValue(ctx, "listId", "1")
		ctx = context.WithValue(ctx, "productId", "2")
		request = request.WithContext(ctx)

		controller.PutEntry(writer, request)

		if writer.Code != http.StatusBadRequest {
			t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, writer.Code)
		}
	})
}

func TestDefaultController_DeleteEntry(t *testing.T) {
	t.Run("Delete an existing shopping list entry (expect 200)", func(t *testing.T) {
		controller := DefaultController{
			userShoppingListEntryRepository: setupMockEntryRepository(),
		}
		writer := httptest.NewRecorder()
		request := httptest.NewRequest("DELETE", "/api/v1/shoppinglistentries/1/2", nil)
		ctx := request.Context()
		ctx = context.WithValue(ctx, "listId", "1")
		ctx = context.WithValue(ctx, "productId", "2")
		request = request.WithContext(ctx)

		controller.DeleteEntry(writer, request)

		if writer.Code != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, writer.Code)
		}
	})
}
func setupMockEntryRepository() Repository {
	repository := NewDemoRepository()
	entries := setupMockEntrySlice()
	for _, entry := range entries {
		repository.Create(entry)
	}
	return repository
}

func setupMockEntrySlice() []*entryModel.UserShoppingListEntry {
	return []*entryModel.UserShoppingListEntry{
		{ShoppingListId: 1, ProductId: 1, Count: 3, Note: "Sample entry 1", Checked: false},
		{ShoppingListId: 1, ProductId: 2, Count: 2, Note: "Sample entry 2", Checked: true},
		{ShoppingListId: 2, ProductId: 1, Count: 1, Note: "Sample entry 3", Checked: false},
	}
}

func setupMockListRepository() Repository {
	repository := NewDemoRepository()
	lists := setupMockListSlice()
	for _, list := range lists {
		repository.Create(list)
	}
	return repository
}

func setupMockListSlice() []*listModel.UserShoppingList {
	return []*listModel.UserShoppingList{
		{Id: 1, UserId: 2, Description: "Frühstück mit Anne", Completed: false},
		{Id: 2, UserId: 2, Description: "Suppe", Completed: false},
		{Id: 3, UserId: 4, Description: "Geburtstagskuchen", Completed: false},
	}
}
