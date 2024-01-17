package rpc

import (
	"context"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/errgo.v2/errors"
	proto "hsfl.de/group6/hsfl-master-ai-cloud-engineering/lib/rpc/product"
	"hsfl.de/group6/hsfl-master-ai-cloud-engineering/product-service/prices"
	pricesMock "hsfl.de/group6/hsfl-master-ai-cloud-engineering/product-service/prices/_mock"
	priceModel "hsfl.de/group6/hsfl-master-ai-cloud-engineering/product-service/prices/model"
	"hsfl.de/group6/hsfl-master-ai-cloud-engineering/product-service/products"
	productsMock "hsfl.de/group6/hsfl-master-ai-cloud-engineering/product-service/products/_mock"
	productModel "hsfl.de/group6/hsfl-master-ai-cloud-engineering/product-service/products/model"
	"testing"
)

func TestNewProductServiceServer(t *testing.T) {
	mockProductRepo := productsMock.NewMockRepository(t)
	mockPriceRepo := pricesMock.NewMockRepository(t)
	var productRepo products.Repository = mockProductRepo
	var priceRepo prices.Repository = mockPriceRepo

	server := NewProductServiceServer(&productRepo, &priceRepo)

	assert.NotNil(t, server)
	assert.NotNil(t, server.productRepository)
	assert.NotNil(t, server.priceRepository, priceRepo)
}

func TestProductServiceServer_CreateProduct(t *testing.T) {
	mockProductRepo := productsMock.NewMockRepository(t)
	mockPriceRepo := pricesMock.NewMockRepository(t)
	var productRepo products.Repository = mockProductRepo
	var priceRepo prices.Repository = mockPriceRepo
	server := NewProductServiceServer(&productRepo, &priceRepo)

	ctx := context.Background()
	req := &proto.CreateProductRequest{
		Product: &proto.Product{
			Id:          123,
			Description: "Test",
			Ean:         "12345670",
		},
	}

	t.Run("Successful create", func(t *testing.T) {
		mockProductRepo.EXPECT().Create(
			&productModel.Product{Id: 123, Description: "Test", Ean: "12345670"}).
			Return(
				&productModel.Product{Id: 123, Description: "Test", Ean: "12345670"}, nil).Once()

		resp, err := server.CreateProduct(ctx, req)
		if err != nil {
			t.Errorf("Error creating product: %v", err)
		}

		if resp.Product.Id != 123 {
			t.Errorf("Expected product ID '123', got '%d'", resp.Product.Id)
		}
	})

	t.Run("Unsuccessful create", func(t *testing.T) {
		mockProductRepo.EXPECT().Create(
			&productModel.Product{Id: 123, Description: "Test", Ean: "12345670"}).
			Return(nil, errors.New(products.ErrorProductAlreadyExists)).Once()

		_, err := server.CreateProduct(ctx, req)
		if e, ok := status.FromError(err); ok {
			if e.Code() != codes.Internal {
				t.Errorf("expected error: %v", err)
			}
		}
	})
}

func TestProductServiceServer_GetProduct(t *testing.T) {
	mockProductRepo := productsMock.NewMockRepository(t)
	mockPriceRepo := pricesMock.NewMockRepository(t)
	var productRepo products.Repository = mockProductRepo
	var priceRepo prices.Repository = mockPriceRepo
	server := NewProductServiceServer(&productRepo, &priceRepo)

	t.Run("Successful get product", func(t *testing.T) {
		mockProductRepo.EXPECT().FindById(uint64(2)).Return(
			&productModel.Product{Id: 2, Description: "Test", Ean: "12345670"}, nil).Once()

		ctx := context.Background()
		getReq := &proto.GetProductRequest{Id: 2}
		resp, err := server.GetProduct(ctx, getReq)
		if err != nil {
			t.Errorf("Error getting product: %v", err)
		}

		if resp.Product.Id != 2 {
			t.Errorf("Expected product ID '123', got '%d'", resp.Product.Id)
		}
	})

	t.Run("Unsuccessful get product", func(t *testing.T) {
		mockProductRepo.EXPECT().FindById(uint64(2)).Return(nil, errors.New(products.ErrorProductNotFound)).Once()

		ctx := context.Background()
		getReq := &proto.GetProductRequest{Id: 2}
		_, err := server.GetProduct(ctx, getReq)
		if e, ok := status.FromError(err); ok {
			if e.Code() != codes.NotFound {
				t.Errorf("expected error: %v", err)
			}
		}
	})
}

func TestProductServiceServer_GetAllProducts(t *testing.T) {
	mockProductRepo := productsMock.NewMockRepository(t)
	mockPriceRepo := pricesMock.NewMockRepository(t)
	var productRepo products.Repository = mockProductRepo
	var priceRepo prices.Repository = mockPriceRepo
	server := NewProductServiceServer(&productRepo, &priceRepo)

	products := []*productModel.Product{
		{Id: 1, Description: "Banane", Ean: "12345670"},
		{Id: 2, Description: "Apfel", Ean: "13828523"},
	}

	mockProductRepo.EXPECT().FindAll().Return(products, nil)

	ctx := context.Background()
	getReq := &proto.GetAllProductsRequest{}
	resp, err := server.GetAllProducts(ctx, getReq)
	if err != nil {
		t.Errorf("Error getting all products: %v", err)
	}

	for i, product := range products {
		if product.Id != resp.Products[i].Id {
			t.Errorf("Expected product ID '%d', got '%d'", product.Id, resp.Products[i].Id)
		}
		if product.Description != resp.Products[i].Description {
			t.Errorf("Expected product Description '%s', got '%s'", product.Description, resp.Products[i].Description)
		}
		if product.Ean != resp.Products[i].Ean {
			t.Errorf("Expected product Ean '%s', got '%s'", product.Ean, resp.Products[i].Ean)
		}
	}
}

func TestProductServiceServer_UpdateProduct(t *testing.T) {
	mockProductRepo := productsMock.NewMockRepository(t)
	mockPriceRepo := pricesMock.NewMockRepository(t)
	var productRepo products.Repository = mockProductRepo
	var priceRepo prices.Repository = mockPriceRepo
	server := NewProductServiceServer(&productRepo, &priceRepo)

	t.Run("Successful Update", func(t *testing.T) {
		changedProduct := &productModel.Product{Id: 1, Description: "Erdbeere", Ean: "12345670"}

		mockProductRepo.EXPECT().Update(changedProduct).Return(changedProduct, nil).Once()

		ctx := context.Background()
		updReq := &proto.UpdateProductRequest{Product: &proto.Product{
			Id:          1,
			Description: "Erdbeere",
			Ean:         "12345670",
		}}
		resp, err := server.UpdateProduct(ctx, updReq)
		if err != nil {
			t.Errorf("Error getting all products: %v", err)
		}

		if changedProduct.Id != resp.Product.Id {
			t.Errorf("Expected product ID '%d', got '%d'", changedProduct.Id, resp.Product.Id)
		}
		if changedProduct.Description != resp.Product.Description {
			t.Errorf("Expected product Description '%s', got '%s'", changedProduct.Description, resp.Product.Description)
		}
		if changedProduct.Ean != resp.Product.Ean {
			t.Errorf("Expected product Ean '%s', got '%s'", changedProduct.Ean, resp.Product.Ean)
		}
	})

	t.Run("Unsuccessful Update", func(t *testing.T) {
		changedProduct := &productModel.Product{Id: 1, Description: "Erdbeere", Ean: "12345670"}
		mockProductRepo.EXPECT().Update(changedProduct).Return(nil, errors.New(products.ErrorProductUpdate)).Once()

		ctx := context.Background()
		updReq := &proto.UpdateProductRequest{Product: &proto.Product{
			Id:          1,
			Description: "Erdbeere",
			Ean:         "12345670",
		}}
		_, err := server.UpdateProduct(ctx, updReq)
		if e, ok := status.FromError(err); ok {
			if e.Code() != codes.Internal {
				t.Errorf("expected error: %v", err)
			}
		}
	})
}

func TestProductServiceServer_DeleteProduct(t *testing.T) {
	mockProductRepo := productsMock.NewMockRepository(t)
	mockPriceRepo := pricesMock.NewMockRepository(t)
	var productRepo products.Repository = mockProductRepo
	var priceRepo prices.Repository = mockPriceRepo
	server := NewProductServiceServer(&productRepo, &priceRepo)

	t.Run("Successful Delete", func(t *testing.T) {
		productToDelete := &productModel.Product{Id: 1}

		mockProductRepo.EXPECT().Delete(productToDelete).Return(nil).Once()

		ctx := context.Background()
		delReq := &proto.DeleteProductRequest{Id: 1}
		_, err := server.DeleteProduct(ctx, delReq)
		if err != nil {
			t.Errorf("Error getting deleting product: %v", err)
		}
	})

	t.Run("Unsuccessful Delete", func(t *testing.T) {
		productToDelete := &productModel.Product{Id: 1}

		mockProductRepo.EXPECT().Delete(productToDelete).Return(errors.New(products.ErrorProductDeletion)).Once()

		ctx := context.Background()
		delReq := &proto.DeleteProductRequest{Id: 1}
		_, err := server.DeleteProduct(ctx, delReq)
		if e, ok := status.FromError(err); ok {
			if e.Code() != codes.Internal {
				t.Errorf("expected error: %v", err)
			}
		}
	})
}

func TestProductServiceServer_CreatePrice(t *testing.T) {
	mockProductRepo := productsMock.NewMockRepository(t)
	mockPriceRepo := pricesMock.NewMockRepository(t)
	var productRepo products.Repository = mockProductRepo
	var priceRepo prices.Repository = mockPriceRepo
	server := NewProductServiceServer(&productRepo, &priceRepo)

	ctx := context.Background()
	req := &proto.CreatePriceRequest{
		Price: &proto.Price{
			UserId:    1,
			ProductId: 2,
			Price:     2.99,
		},
	}

	t.Run("Successful create", func(t *testing.T) {
		mockPriceRepo.EXPECT().Create(
			&priceModel.Price{UserId: 1, ProductId: 2, Price: 2.99}).
			Return(
				&priceModel.Price{UserId: 1, ProductId: 2, Price: 2.99}, nil).Once()

		resp, err := server.CreatePrice(ctx, req)
		if err != nil {
			t.Errorf("Error creating product: %v", err)
		}

		if resp.Price.UserId != 1 {
			t.Errorf("Expected user ID '1', got '%d'", resp.Price.UserId)
		}

		if resp.Price.ProductId != 2 {
			t.Errorf("Expected product ID '2', got '%d'", resp.Price.ProductId)
		}

		if resp.Price.Price != 2.99 {
			t.Errorf("Expected price '2.99', got '%f'", resp.Price.Price)
		}
	})

	t.Run("Unsuccessful create", func(t *testing.T) {
		mockPriceRepo.EXPECT().Create(
			&priceModel.Price{UserId: 1, ProductId: 2, Price: 2.99}).
			Return(nil, errors.New(prices.ErrorPriceAlreadyExists)).Once()

		_, err := server.CreatePrice(ctx, req)
		if e, ok := status.FromError(err); ok {
			if e.Code() != codes.Internal {
				t.Errorf("expected error: %v", err)
			}
		}
	})
}

func TestProductServiceServer_FindAllPrices(t *testing.T) {
	mockProductRepo := productsMock.NewMockRepository(t)
	mockPriceRepo := pricesMock.NewMockRepository(t)
	var productRepo products.Repository = mockProductRepo
	var priceRepo prices.Repository = mockPriceRepo
	server := NewProductServiceServer(&productRepo, &priceRepo)

	availablePrices := []*priceModel.Price{
		{UserId: 1, ProductId: 2, Price: 2.99},
		{UserId: 1, ProductId: 1, Price: 3.99},
	}

	mockPriceRepo.EXPECT().FindAll().Return(availablePrices, nil)

	ctx := context.Background()
	getReq := &proto.FindAllPricesRequest{}
	resp, err := server.FindAllPrices(ctx, getReq)
	if err != nil {
		t.Errorf("Error getting all products: %v", err)
	}

	for i, price := range availablePrices {
		if price.UserId != resp.Price[i].UserId {
			t.Errorf("Expected price UserID '%d', got '%d'", price.UserId, resp.Price[i].UserId)
		}
		if price.ProductId != resp.Price[i].ProductId {
			t.Errorf("Expected price ProductID '%d', got '%d'", price.ProductId, resp.Price[i].ProductId)
		}
		if price.Price != resp.Price[i].Price {
			t.Errorf("Expected price'%f', got '%f'", price.Price, resp.Price[i].Price)
		}
	}
}

func TestProductServiceServer_FindAllPricesFromUser(t *testing.T) {
	mockProductRepo := productsMock.NewMockRepository(t)
	mockPriceRepo := pricesMock.NewMockRepository(t)
	var productRepo products.Repository = mockProductRepo
	var priceRepo prices.Repository = mockPriceRepo
	server := NewProductServiceServer(&productRepo, &priceRepo)

	availablePrices := []*priceModel.Price{
		{UserId: 1, ProductId: 2, Price: 2.99},
		{UserId: 1, ProductId: 1, Price: 3.99},
	}

	mockPriceRepo.EXPECT().FindAllByUser(uint64(1)).Return(availablePrices, nil)

	ctx := context.Background()
	getReq := &proto.FindAllPricesFromUserRequest{UserId: 1}
	resp, err := server.FindAllPricesFromUser(ctx, getReq)
	if err != nil {
		t.Errorf("Error getting all products: %v", err)
	}

	for i, price := range availablePrices {
		if price.UserId != resp.Price[i].UserId {
			t.Errorf("Expected price UserID '%d', got '%d'", price.UserId, resp.Price[i].UserId)
		}
		if price.ProductId != resp.Price[i].ProductId {
			t.Errorf("Expected price ProductID '%d', got '%d'", price.ProductId, resp.Price[i].ProductId)
		}
		if price.Price != resp.Price[i].Price {
			t.Errorf("Expected price'%f', got '%f'", price.Price, resp.Price[i].Price)
		}
	}
}

func TestProductServiceServer_FindPrice(t *testing.T) {
	mockProductRepo := productsMock.NewMockRepository(t)
	mockPriceRepo := pricesMock.NewMockRepository(t)
	var productRepo products.Repository = mockProductRepo
	var priceRepo prices.Repository = mockPriceRepo
	server := NewProductServiceServer(&productRepo, &priceRepo)

	t.Run("Successful find Price", func(t *testing.T) {
		mockPriceRepo.EXPECT().FindByIds(uint64(1), uint64(1)).Return(
			&priceModel.Price{UserId: 1, ProductId: 1, Price: 3.99}, nil).Once()

		ctx := context.Background()
		getReq := &proto.FindPriceRequest{UserId: 1, ProductId: 1}
		resp, err := server.FindPrice(ctx, getReq)
		if err != nil {
			t.Errorf("Error getting product: %v", err)
		}

		if resp.Price.UserId != 1 {
			t.Errorf("Expected User ID '1', got '%d'", resp.Price.UserId)
		}
		if resp.Price.ProductId != 1 {
			t.Errorf("Expected Product ID '1', got '%d'", resp.Price.ProductId)
		}
	})

	t.Run("Unsuccessful find Price", func(t *testing.T) {
		mockPriceRepo.EXPECT().FindByIds(uint64(1), uint64(5)).Return(nil, errors.New(prices.ErrorPriceNotFound)).Once()

		ctx := context.Background()
		getReq := &proto.FindPriceRequest{UserId: 5, ProductId: 1}
		_, err := server.FindPrice(ctx, getReq)
		if e, ok := status.FromError(err); ok {
			if e.Code() != codes.NotFound {
				t.Errorf("expected error: %v", err)
			}
		}
	})
}

func TestProductServiceServer_UpdatePrice(t *testing.T) {
	mockProductRepo := productsMock.NewMockRepository(t)
	mockPriceRepo := pricesMock.NewMockRepository(t)
	var productRepo products.Repository = mockProductRepo
	var priceRepo prices.Repository = mockPriceRepo
	server := NewProductServiceServer(&productRepo, &priceRepo)

	t.Run("Successful Update", func(t *testing.T) {
		changedPrice := &priceModel.Price{UserId: 1, ProductId: 1, Price: 5.55}

		mockPriceRepo.EXPECT().Update(changedPrice).Return(changedPrice, nil).Once()

		ctx := context.Background()
		updReq := &proto.UpdatePriceRequest{Price: &proto.Price{
			UserId:    1,
			ProductId: 1,
			Price:     5.55,
		}}
		resp, err := server.UpdatePrice(ctx, updReq)
		if err != nil {
			t.Errorf("Error getting all products: %v", err)
		}

		if changedPrice.UserId != resp.Price.UserId {
			t.Errorf("Expected price UserID '%d', got '%d'", changedPrice.UserId, resp.Price.UserId)
		}
		if changedPrice.ProductId != resp.Price.ProductId {
			t.Errorf("Expected price ProductID '%d', got '%d'", changedPrice.ProductId, resp.Price.ProductId)
		}
		if changedPrice.Price != resp.Price.Price {
			t.Errorf("Expected price '%f', got '%f'", changedPrice.Price, resp.Price.Price)
		}
	})

	t.Run("Unsuccessful Update", func(t *testing.T) {
		changedPrice := &priceModel.Price{UserId: 4, ProductId: 1, Price: 5.55}
		mockPriceRepo.EXPECT().Update(changedPrice).Return(nil, errors.New(prices.ErrorPriceUpdate)).Once()

		ctx := context.Background()
		updReq := &proto.UpdatePriceRequest{Price: &proto.Price{
			UserId:    4,
			ProductId: 1,
			Price:     5.55,
		}}
		_, err := server.UpdatePrice(ctx, updReq)
		if e, ok := status.FromError(err); ok {
			if e.Code() != codes.Internal {
				t.Errorf("expected error: %v", err)
			}
		}
	})

}

func TestProductServiceServer_DeletePrice(t *testing.T) {
	mockProductRepo := productsMock.NewMockRepository(t)
	mockPriceRepo := pricesMock.NewMockRepository(t)
	var productRepo products.Repository = mockProductRepo
	var priceRepo prices.Repository = mockPriceRepo
	server := NewProductServiceServer(&productRepo, &priceRepo)

	t.Run("Successful Delete", func(t *testing.T) {
		priceToDelete := &priceModel.Price{UserId: 1, ProductId: 1}

		mockPriceRepo.EXPECT().Delete(priceToDelete).Return(nil).Once()

		ctx := context.Background()
		delReq := &proto.DeletePriceRequest{UserId: 1, ProductId: 1}
		_, err := server.DeletePrice(ctx, delReq)
		if err != nil {
			t.Errorf("Error getting deleting price: %v", err)
		}
	})

	t.Run("Unsuccessful Delete", func(t *testing.T) {
		priceToDelete := &priceModel.Price{UserId: 1, ProductId: 1}

		mockPriceRepo.EXPECT().Delete(priceToDelete).Return(errors.New(products.ErrorProductDeletion)).Once()

		ctx := context.Background()
		delReq := &proto.DeletePriceRequest{UserId: 1, ProductId: 1}
		_, err := server.DeletePrice(ctx, delReq)
		if e, ok := status.FromError(err); ok {
			if e.Code() != codes.Internal {
				t.Errorf("expected error: %v", err)
			}
		}
	})

}
