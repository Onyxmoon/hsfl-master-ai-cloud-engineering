package router

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"hsfl.de/group6/hsfl-master-ai-cloud-engineering/product-service/prices"
	priceModel "hsfl.de/group6/hsfl-master-ai-cloud-engineering/product-service/prices/model"
	"hsfl.de/group6/hsfl-master-ai-cloud-engineering/product-service/products"
	productModel "hsfl.de/group6/hsfl-master-ai-cloud-engineering/product-service/products/model"
	"hsfl.de/group6/hsfl-master-ai-cloud-engineering/user-service/api/http/middleware"
	"hsfl.de/group6/hsfl-master-ai-cloud-engineering/user-service/auth"
	"hsfl.de/group6/hsfl-master-ai-cloud-engineering/user-service/auth/utils"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRouter(t *testing.T) {
	var pricesRepo = setupMockPriceRepository()
	var productsRepo = setupMockProductRepository()

	var productsController products.Controller = products.NewDefaultController(productsRepo)
	var pricesController prices.Controller = prices.NewDefaultController(pricesRepo)

	tokenGenerator, _ := auth.NewJwtTokenGenerator(auth.JwtConfig{PrivateKey: utils.GenerateRandomECDSAPrivateKeyAsPEM()})
	authMiddleware := middleware.CreateLocalAuthMiddleware(&userRepo, tokenGenerator)

	router := New(&productsController, &pricesController, authMiddleware)

	t.Run("should return 404 NOT FOUND if path is unknown", func(t *testing.T) {
		// given
		w := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "/unknown/route", nil)

		// when
		router.ServeHTTP(w, r)

		// then
		assert.Equal(t, http.StatusNotFound, w.Code)
	})

	t.Run("/api/v1/products", func(t *testing.T) {
		t.Run("should return 404 NOT FOUND if method is not GET", func(t *testing.T) {
			tests := []string{"DELETE", "PUT", "HEAD", "CONNECT", "OPTIONS", "TRACE", "PATCH"}

			for _, test := range tests {
				// given
				w := httptest.NewRecorder()
				r := httptest.NewRequest(test, "/api/v1/products", nil)

				// when
				router.ServeHTTP(w, r)

				// then
				assert.Equal(t, http.StatusNotFound, w.Code)
			}
		})

		t.Run("should call GET handler", func(t *testing.T) {
			// given
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/api/v1/product/", nil)

			// when
			router.ServeHTTP(w, r)

			// then
			assert.Equal(t, http.StatusOK, w.Code)
		})

		t.Run("should call POST handler", func(t *testing.T) {
			// given
			w := httptest.NewRecorder()
			jsonRequest := `{"id": 3, "description": "Test Product", "ean": 12345}`
			r := httptest.NewRequest("POST", "/api/v1/product/", strings.NewReader(jsonRequest))

			// when
			router.ServeHTTP(w, r)

			// then
			assert.Equal(t, http.StatusOK, w.Code)
		})
	})

	t.Run("/api/v1/product/:productid", func(t *testing.T) {
		t.Run("should return 404 NOT FOUND if method is not GET, DELETE or PUT", func(t *testing.T) {
			tests := []string{"POST", "HEAD", "CONNECT", "OPTIONS", "TRACE", "PATCH"}

			for _, test := range tests {
				// given
				w := httptest.NewRecorder()
				r := httptest.NewRequest(test, "/api/v1/product/1", nil)

				// when
				router.ServeHTTP(w, r)

				// then
				assert.Equal(t, http.StatusNotFound, w.Code)
			}
		})

		t.Run("should call GET handler", func(t *testing.T) {
			// given
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/api/v1/product/1", nil)

			// when
			router.ServeHTTP(w, r)

			// then
			assert.Equal(t, http.StatusOK, w.Code)
		})

		t.Run("should call PUT handler", func(t *testing.T) {
			// given
			w := httptest.NewRecorder()
			jsonRequest := `{"description":"New Description","ean":4014819040771}`
			r := httptest.NewRequest("PUT", "/api/v1/product/1", strings.NewReader(jsonRequest))

			// when
			router.ServeHTTP(w, r)

			// then
			assert.Equal(t, http.StatusOK, w.Code)
		})

		t.Run("should call DELETE handler", func(t *testing.T) {
			// given
			w := httptest.NewRecorder()
			r := httptest.NewRequest("DELETE", "/api/v1/product/1", nil)

			// when
			router.ServeHTTP(w, r)

			// then
			assert.Equal(t, http.StatusOK, w.Code)
		})
	})

	t.Run("/api/v1/price", func(t *testing.T) {
		t.Run("should return 404 NOT FOUND if method is not POST", func(t *testing.T) {
			tests := []string{"DELETE", "PUT", "HEAD", "CONNECT", "OPTIONS", "TRACE", "PATCH", "GET"}

			for _, test := range tests {
				// given
				w := httptest.NewRecorder()
				r := httptest.NewRequest(test, "/api/v1/price", nil)

				// when
				router.ServeHTTP(w, r)

				// then
				assert.Equal(t, http.StatusNotFound, w.Code)
			}
		})

		t.Run("should call GET handler", func(t *testing.T) {
			// given
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/api/v1/price/", nil)

			// when
			router.ServeHTTP(w, r)

			// then
			assert.Equal(t, http.StatusOK, w.Code)
		})

		t.Run("should call POST handler", func(t *testing.T) {
			// given
			w := httptest.NewRecorder()
			jsonRequest := `{"price": 0.99}`
			r := httptest.NewRequest("POST", "/api/v1/price/3/3", strings.NewReader(jsonRequest))

			// when
			router.ServeHTTP(w, r)

			// then
			assert.Equal(t, http.StatusOK, w.Code)
		})
	})

	t.Run("/api/v1/price/user/:userId", func(t *testing.T) {
		t.Run("should return 404 NOT FOUND if method is not GET", func(t *testing.T) {
			tests := []string{"HEAD", "CONNECT", "OPTIONS", "TRACE", "PATCH", "POST", "DELETE", "PUT"}

			for _, test := range tests {
				// given
				w := httptest.NewRecorder()
				r := httptest.NewRequest(test, "/api/v1/price/user/1", nil)

				// when
				router.ServeHTTP(w, r)

				// then
				if test == "POST" || test == "PUT" || test == "DELETE" {
					assert.Equal(t, http.StatusBadRequest, w.Code)
				} else {
					assert.Equal(t, http.StatusNotFound, w.Code)
				}
			}
		})

		t.Run("should call GET handler", func(t *testing.T) {
			// given
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/api/v1/price/user/1", nil)

			// when
			router.ServeHTTP(w, r)

			// then
			assert.Equal(t, http.StatusOK, w.Code)
		})
	})

	t.Run("/api/v1/price/:productId/:userId", func(t *testing.T) {
		t.Run("should return 404 NOT FOUND if method is not GET, DELETE, POST or PUT", func(t *testing.T) {
			tests := []string{"HEAD", "CONNECT", "OPTIONS", "TRACE", "PATCH"}

			for _, test := range tests {
				// given
				w := httptest.NewRecorder()
				r := httptest.NewRequest(test, "/api/v1/price/1/1", nil)

				// when
				router.ServeHTTP(w, r)

				fmt.Println(w.Code)

				// then
				assert.Equal(t, http.StatusNotFound, w.Code)
			}
		})

		t.Run("should call GET handler", func(t *testing.T) {
			// given
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/api/v1/price/1/1", nil)

			// when
			router.ServeHTTP(w, r)

			// then
			assert.Equal(t, http.StatusOK, w.Code)
		})

		t.Run("should call PUT handler", func(t *testing.T) {
			// given
			w := httptest.NewRecorder()
			jsonRequest := `{"userId": 1, "productId": 1, "price": 10.99}`
			r := httptest.NewRequest("PUT", "/api/v1/price/1/1", strings.NewReader(jsonRequest))

			// when
			router.ServeHTTP(w, r)

			// then
			assert.Equal(t, http.StatusOK, w.Code)
		})

		t.Run("should call DELETE handler", func(t *testing.T) {
			// given
			w := httptest.NewRecorder()
			r := httptest.NewRequest("DELETE", "/api/v1/price/1/1", nil)

			// when
			router.ServeHTTP(w, r)

			// then
			assert.Equal(t, http.StatusOK, w.Code)
		})
	})
}

func setupMockProductRepository() products.Repository {
	repository := products.NewDemoRepository()
	productSlice := setupDemoProductSlice()
	for _, product := range productSlice {
		_, err := repository.Create(product)
		if err != nil {
			return nil
		}
	}

	return repository
}

func setupDemoProductSlice() []*productModel.Product {
	return []*productModel.Product{
		{
			Id:          1,
			Description: "Strauchtomaten",
			Ean:         4014819040771,
		},
		{
			Id:          2,
			Description: "Lauchzwiebeln",
			Ean:         5001819040871,
		},
	}
}

func setupMockPriceRepository() prices.Repository {
	repository := prices.NewDemoRepository()
	pricesSlice := []*priceModel.Price{
		{
			UserId:    1,
			ProductId: 1,
			Price:     2.99,
		},
		{
			UserId:    2,
			ProductId: 2,
			Price:     5.99,
		},
	}
	for _, price := range pricesSlice {
		_, err := repository.Create(price)
		if err != nil {
			return nil
		}
	}

	return repository
}
