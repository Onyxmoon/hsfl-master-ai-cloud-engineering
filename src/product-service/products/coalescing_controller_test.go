package products

import (
	"context"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"golang.org/x/sync/singleflight"
	"hsfl.de/group6/hsfl-master-ai-cloud-engineering/product-service/products/model"
	"net/http"
	"net/http/httptest"
	"sort"
	"strings"
	"testing"
)

func TestNewCoalescingController(t *testing.T) {
	demoRepo := NewDemoRepository()

	controller := NewCoalescingController(demoRepo)

	assert.NotNil(t, controller)
	assert.Equal(t, demoRepo, controller.productRepository)
	assert.IsType(t, &singleflight.Group{}, controller.group)
}

func TestCoalescingController_GetProducts(t *testing.T) {
	t.Run("should return all products", func(t *testing.T) {
		controller := NewCoalescingController(GenerateExampleDemoRepository())

		writer := httptest.NewRecorder()
		request := httptest.NewRequest("GET", "/api/v1/product", nil)

		controller.GetProducts(writer, request)

		res := writer.Result()
		var response []model.Product
		err := json.NewDecoder(res.Body).Decode(&response)

		if err != nil {
			t.Error(err)
		}

		if writer.Code != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, writer.Code)
		}

		if writer.Header().Get("Content-Type") != "application/json" {
			t.Errorf("Expected content type %s, got %s",
				"application/json", writer.Header().Get("Content-Type"))
		}

		products := GenerateExampleProductSlice()

		sort.Slice(response, func(i, j int) bool {
			return response[i].Id < response[j].Id
		})

		if len(response) != len(products) {
			t.Errorf("Expected count of product is %d, got %d",
				2, len(response))
		}

		for i, product := range products {
			if product.Id != response[i].Id {
				t.Errorf("Expected id of product %d, got %d", product.Id, response[i].Id)
			}

			if product.Description != response[i].Description {
				t.Errorf("Expected description of product %s, got %s", product.Description, response[i].Description)
			}

			if product.Ean != response[i].Ean {
				t.Errorf("Expected ean of product %s, got %s", product.Ean, response[i].Ean)
			}
		}

	})
}

func TestCoalescingController_GetProductById(t *testing.T) {
	type fields struct {
		productRepository Repository
	}
	type args struct {
		writer  *httptest.ResponseRecorder
		request *http.Request
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantStatus int
	}{
		{
			name: "Bad non-numeric request (expect 400)",
			fields: fields{
				productRepository: GenerateExampleDemoRepository(),
			},
			args: args{
				writer: httptest.NewRecorder(),
				request: func() *http.Request {
					var request = httptest.NewRequest("GET", "/api/v1/product/abc", nil)
					request = request.WithContext(context.WithValue(request.Context(), "productId", "abc"))
					return request
				}(),
			},
			wantStatus: http.StatusBadRequest,
		},
		{
			name: "Unknown product (expect 404)",
			fields: fields{
				productRepository: GenerateExampleDemoRepository(),
			},
			args: args{
				writer: httptest.NewRecorder(),
				request: func() *http.Request {
					var request = httptest.NewRequest("GET", "/api/v1/product/4", nil)
					request = request.WithContext(context.WithValue(request.Context(), "productId", "4"))
					return request
				}(),
			},
			wantStatus: http.StatusNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			controller := NewCoalescingController(tt.fields.productRepository)
			controller.GetProductById(tt.args.writer, tt.args.request)
			if tt.args.writer.Code != tt.wantStatus {
				t.Errorf("Expected status code %d, got %d", tt.wantStatus, tt.args.writer.Code)
			}
		})
	}

	t.Run("Successfully get existing product (expect 200 and product)", func(t *testing.T) {
		writer := httptest.NewRecorder()
		request := httptest.NewRequest("GET", "/api/v1/product/1", nil)
		request = request.WithContext(context.WithValue(request.Context(), "productId", "1"))

		controller := NewCoalescingController(GenerateExampleDemoRepository())
		controller.GetProductById(writer, request)

		if writer.Code != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, writer.Code)
		}

		if writer.Header().Get("Content-Type") != "application/json" {
			t.Errorf("Expected content type %s, got %s",
				"application/json", writer.Header().Get("Content-Type"))
		}

		result := writer.Result()
		var response model.Product
		err := json.NewDecoder(result.Body).Decode(&response)
		if err != nil {
			t.Fatal(err.Error())
		}

		if response.Id != 1 {
			t.Errorf("Expected id of product %d, got %d", 1, response.Id)
		}

		if response.Description != "Strauchtomaten" {
			t.Errorf("Expected description of product %s, got %s", "Strauchtomaten", response.Description)
		}

		if response.Ean != "4014819040771" {
			t.Errorf("Expected ean of product %s, got %s", "4014819040771", response.Ean)
		}

	})
}

func TestCoalescingController_GetProductByEan(t *testing.T) {
	t.Run("Invalid product ean (expect 400)", func(t *testing.T) {
		controller := NewCoalescingController(GenerateExampleDemoRepository())

		writer := httptest.NewRecorder()
		request := httptest.NewRequest("GET", "/api/v1/products/ean?ean=123", nil)
		request = request.WithContext(context.WithValue(request.Context(), "productEan", "123"))

		controller.GetProductByEan(writer, request)

		if writer.Code != http.StatusBadRequest {
			t.Errorf("Expected status code %d, got %d", http.StatusNotFound, writer.Code)
		}
	})

	t.Run("Unknown product (expect 404)", func(t *testing.T) {
		controller := NewCoalescingController(GenerateExampleDemoRepository())

		writer := httptest.NewRecorder()
		request := httptest.NewRequest("GET", "/api/v1/products/ean?ean=12345670", nil)
		request = request.WithContext(context.WithValue(request.Context(), "productEan", "12345670"))

		controller.GetProductByEan(writer, request)

		if writer.Code != http.StatusNotFound {
			t.Errorf("Expected status code %d, got %d", http.StatusNotFound, writer.Code)
		}
	})

	t.Run("Should return product by EAN (expect 200 and product)", func(t *testing.T) {
		controller := NewCoalescingController(GenerateExampleDemoRepository())

		writer := httptest.NewRecorder()
		request := httptest.NewRequest("GET", "/api/v1/products/ean?ean=4014819040771", nil)
		request = request.WithContext(context.WithValue(request.Context(), "productEan", "4014819040771"))

		controller.GetProductByEan(writer, request)

		if writer.Code != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, writer.Code)
		}

		res := writer.Result()
		var response model.Product
		err := json.NewDecoder(res.Body).Decode(&response)
		if err != nil {
			t.Error(err)
		}

		if response.Id != 1 {
			t.Errorf("Expected id of product %d, got %d", 1, response.Id)
		}

		if response.Description != "Strauchtomaten" {
			t.Errorf("Expected description of product %s, got %s", "Strauchtomaten", response.Description)
		}

		if response.Ean != "4014819040771" {
			t.Errorf("Expected ean of product %s, got %s", "4014819040771", response.Ean)
		}
	})
}

func TestCoalescingController_PostProduct(t *testing.T) {
	type fields struct {
		productRepository Repository
	}
	type args struct {
		writer  *httptest.ResponseRecorder
		request *http.Request
	}
	tests := []struct {
		name             string
		fields           fields
		args             args
		expectedStatus   int
		expectedResponse string
	}{
		{
			name: "Unauthorized (expect 401)",
			fields: fields{
				productRepository: GenerateExampleDemoRepository(),
			},
			args: args{
				writer: httptest.NewRecorder(),
				request: func() *http.Request {
					var request = httptest.NewRequest(
						"POST",
						"/api/v1/product",
						strings.NewReader(`{"id": 3, "description": "Test Product", "ean": "12345"}`))
					return request
				}(),
			},
			expectedStatus:   http.StatusUnauthorized,
			expectedResponse: "",
		},
		{
			name: "Invalid create (expect 400)",
			fields: fields{
				productRepository: GenerateExampleDemoRepository(),
			},
			args: args{
				writer: httptest.NewRecorder(),
				request: func() *http.Request {
					var request = httptest.NewRequest(
						"POST",
						"/api/v1/product",
						strings.NewReader(`{"id": 3, "description": "Test Product", "ean": "12345"}`))
					ctx := context.WithValue(request.Context(), "auth_userRole", int64(1))
					return request.WithContext(ctx)
				}(),
			},
			expectedStatus:   http.StatusBadRequest,
			expectedResponse: "",
		},
		{
			name: "Valid create (expect 200)",
			fields: fields{
				productRepository: GenerateExampleDemoRepository(),
			},
			args: args{
				writer: httptest.NewRecorder(),
				request: func() *http.Request {
					var request = httptest.NewRequest(
						"POST",
						"/api/v1/product",
						strings.NewReader(`{"id": 3, "description": "Test Product", "ean": "12345670"}`))
					ctx := context.WithValue(request.Context(), "auth_userRole", int64(1))
					return request.WithContext(ctx)
				}(),
			},
			expectedStatus:   http.StatusOK,
			expectedResponse: "",
		},
		{
			name: "Malformed JSON (expect 400)",
			fields: fields{
				productRepository: GenerateExampleDemoRepository(),
			},
			args: args{
				writer: httptest.NewRecorder(),
				request: func() *http.Request {
					var request = httptest.NewRequest(
						"POST",
						"/api/v1/product",
						strings.NewReader(`{"description": "Test Product"`))
					ctx := context.WithValue(request.Context(), "auth_userRole", int64(1))
					return request.WithContext(ctx)
				}(),
			},
			expectedStatus:   http.StatusBadRequest,
			expectedResponse: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			controller := coalescingController{
				productRepository: tt.fields.productRepository,
			}
			controller.PostProduct(tt.args.writer, tt.args.request)

			if tt.args.writer.Code != tt.expectedStatus {
				t.Errorf("Expected status code %d, but got %d", tt.expectedStatus, tt.args.writer.Code)
			}

			if tt.expectedResponse != "" {
				actualResponse := tt.args.writer.Body.String()
				if actualResponse != tt.expectedResponse {
					t.Errorf("Expected response: %s, but got: %s", tt.expectedResponse, actualResponse)
				}
			}
		})
	}
}

func TestCoalescingController_PutProduct(t *testing.T) {
	type fields struct {
		productRepository Repository
	}
	type args struct {
		writer  *httptest.ResponseRecorder
		request *http.Request
	}

	tests := []struct {
		name             string
		fields           fields
		args             args
		expectedStatus   int
		expectedResponse string
	}{
		{
			name: "Unauthorized (expect 401)",
			fields: fields{
				productRepository: GenerateExampleDemoRepository(),
			},
			args: args{
				writer: httptest.NewRecorder(),
				request: func() *http.Request {
					var request = httptest.NewRequest(
						"PUT",
						"/api/v1/product/1",
						strings.NewReader(`{"id": 1, "description": "Updated Product", "ean": 54321}`))
					ctx := context.WithValue(request.Context(), "productId", "1")
					return request.WithContext(ctx)
				}(),
			},
			expectedStatus:   http.StatusUnauthorized,
			expectedResponse: "",
		},
		{
			name: "Valid update (expect 200)",
			fields: fields{
				productRepository: GenerateExampleDemoRepository(),
			},
			args: args{
				writer: httptest.NewRecorder(),
				request: func() *http.Request {
					var request = httptest.NewRequest(
						"PUT",
						"/api/v1/product/1",
						strings.NewReader(`{"id": 1, "description": "Updated Product", "ean": "12345670"}`))
					ctx := context.WithValue(request.Context(), "auth_userRole", int64(1))
					ctx = context.WithValue(ctx, "productId", "1")
					return request.WithContext(ctx)
				}(),
			},
			expectedStatus:   http.StatusOK,
			expectedResponse: "",
		},
		{
			name: "Invalid update (expect 400)",
			fields: fields{
				productRepository: GenerateExampleDemoRepository(),
			},
			args: args{
				writer: httptest.NewRecorder(),
				request: func() *http.Request {
					var request = httptest.NewRequest(
						"PUT",
						"/api/v1/product/2",
						strings.NewReader(`{"description": "Suppe", "ean": "12345"}`))
					ctx := context.WithValue(request.Context(), "auth_userRole", int64(1))
					ctx = context.WithValue(ctx, "productId", "2")
					return request.WithContext(ctx)
				}(),
			},
			expectedStatus:   http.StatusBadRequest,
			expectedResponse: "",
		},
		{
			name: "Malformed JSON",
			fields: fields{
				productRepository: GenerateExampleDemoRepository(),
			},
			args: args{
				writer: httptest.NewRecorder(),
				request: func() *http.Request {
					var request = httptest.NewRequest(
						"PUT",
						"/api/v1/product/2",
						strings.NewReader(`{"description": "Incomplete Update"`))
					ctx := context.WithValue(request.Context(), "auth_userRole", int64(1))
					ctx = context.WithValue(ctx, "productId", "2")
					return request.WithContext(ctx)
				}(),
			},
			expectedStatus:   http.StatusBadRequest,
			expectedResponse: "",
		},
		{
			name: "Incorrect type for EAN (Non-numeric)",
			fields: fields{
				productRepository: GenerateExampleDemoRepository(),
			},
			args: args{
				writer: httptest.NewRecorder(),
				request: func() *http.Request {
					var request = httptest.NewRequest(
						"PUT",
						"/api/v1/product/2",
						strings.NewReader(`{"ean": "Wrong Type"`))
					ctx := context.WithValue(request.Context(), "auth_userRole", int64(1))
					ctx = context.WithValue(ctx, "productId", "2")
					return request.WithContext(ctx)
				}(),
			},
			expectedStatus:   http.StatusBadRequest,
			expectedResponse: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			controller := coalescingController{
				productRepository: tt.fields.productRepository,
			}
			controller.PutProduct(tt.args.writer, tt.args.request)

			if tt.args.writer.Code != tt.expectedStatus {
				t.Errorf("Expected status code %d, but got %d", tt.expectedStatus, tt.args.writer.Code)
			}

			if tt.expectedResponse != "" {
				actualResponse := tt.args.writer.Body.String()
				if actualResponse != tt.expectedResponse {
					t.Errorf("Expected response: %s, but got: %s", tt.expectedResponse, actualResponse)
				}
			}
		})
	}
}

func TestCoalescingController_DeleteProduct(t *testing.T) {
	type fields struct {
		productRepository Repository
	}
	type args struct {
		writer  *httptest.ResponseRecorder
		request *http.Request
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		expectedStatus int
	}{
		{
			name: "Unauthorized (expect 401)",
			fields: fields{
				productRepository: GenerateExampleDemoRepository(),
			},
			args: args{
				writer: httptest.NewRecorder(),
				request: func() *http.Request {
					var request = httptest.NewRequest(
						"DELETE",
						"/api/v1/product/1",
						nil)
					ctx := context.WithValue(request.Context(), "productId", "1")
					return request.WithContext(ctx)
				}(),
			},
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name: "Valid delete (expect 200)",
			fields: fields{
				productRepository: GenerateExampleDemoRepository(),
			},
			args: args{
				writer: httptest.NewRecorder(),
				request: func() *http.Request {
					var request = httptest.NewRequest(
						"DELETE",
						"/api/v1/product/1",
						nil)
					ctx := context.WithValue(request.Context(), "auth_userRole", int64(2))
					ctx = context.WithValue(ctx, "productId", "1")
					return request.WithContext(ctx)
				}(),
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "Invalid delete, non-numeric request (expect 400)",
			fields: fields{
				productRepository: GenerateExampleDemoRepository(),
			},
			args: args{
				writer: httptest.NewRecorder(),
				request: func() *http.Request {
					var request = httptest.NewRequest(
						"DELETE",
						"/api/v1/product/abc",
						nil)
					ctx := context.WithValue(request.Context(), "auth_userRole", int64(2))
					ctx = context.WithValue(ctx, "productId", "abc")
					return request.WithContext(ctx)
				}(),
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "Invalid delete, unknown product (expect 500)",
			fields: fields{
				productRepository: GenerateExampleDemoRepository(),
			},
			args: args{
				writer: httptest.NewRecorder(),
				request: func() *http.Request {
					var request = httptest.NewRequest(
						"DELETE",
						"/api/v1/product/10",
						nil)
					ctx := context.WithValue(request.Context(), "auth_userRole", int64(2))
					ctx = context.WithValue(ctx, "productId", "10")
					return request.WithContext(ctx)
				}(),
			},
			expectedStatus: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			controller := coalescingController{
				productRepository: tt.fields.productRepository,
			}
			controller.DeleteProduct(tt.args.writer, tt.args.request)
			if tt.args.writer.Code != tt.expectedStatus {
				t.Errorf("Expected status code %d, got %d", tt.expectedStatus, tt.args.writer.Code)
			}
		})
	}
}
