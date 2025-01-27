//go:build integration
// +build integration

package prices

import (
	"context"
	_ "github.com/rqlite/gorqlite/stdlib"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"hsfl.de/group6/hsfl-master-ai-cloud-engineering/product-service/prices/model"
	"os"
	"reflect"
	"testing"
	"time"
)

func TestIntegrationRQLiteRepository(t *testing.T) {
	container, err := prepareIntegrationTestRQLiteContainer()
	if err != nil {
		t.Error(err)
	}
	mappedPort, err := container.MappedPort(context.Background(), "4001/tcp")
	if err != nil {
		t.Error(err)
	}

	rqliteRepository := NewRQLiteRepository("http://localhost:" + mappedPort.Port() + "/?disableClusterDiscovery=true")

	t.Run("TestIntegrationRQLiteRepository_Create", func(t *testing.T) {
		price := model.Price{
			UserId:    1,
			ProductId: 1,
			Price:     2.99,
		}

		t.Run("Create price with success", func(t *testing.T) {
			_, err = rqliteRepository.Create(&price)
			if err != nil {
				t.Error(err)
			}
		})

		t.Run("Can't create prices with duplicate ean", func(t *testing.T) {
			_, err = rqliteRepository.Create(&price)
			if err.Error() != ErrorPriceAlreadyExists {
				t.Error(err)
			}
		})

		err := rqliteRepository.cleanTable()
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("TestIntegrationRQLiteRepository_FindAll", func(t *testing.T) {
		prices := []*model.Price{
			{
				UserId:    1,
				ProductId: 1,
				Price:     2.99,
			},
			{
				UserId:    2,
				ProductId: 3,
				Price:     0.55,
			},
		}

		for _, price := range prices {
			rqliteRepository.Create(price)
		}

		t.Run("Successfully fetch all prices", func(t *testing.T) {
			fetchedPrices, err := rqliteRepository.FindAll()

			if err != nil {
				t.Error("Can't fetch prices")
			}

			if len(fetchedPrices) != len(prices) {
				t.Errorf("Unexpected price count. Expected %d, got %d", len(prices), len(fetchedPrices))
			}

			if !reflect.DeepEqual(prices, fetchedPrices) {
				t.Error("Fetched prices do not match expected prices")
			}
		})

		err := rqliteRepository.cleanTable()
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("TestIntegrationRQLiteRepository_FindAllByUser", func(t *testing.T) {
		userIdMerchantA := uint64(1)

		pricesMerchantA := []*model.Price{
			{
				UserId:    userIdMerchantA,
				ProductId: 1,
				Price:     2.99,
			},
			{
				UserId:    userIdMerchantA,
				ProductId: 3,
				Price:     0.55,
			},
		}

		for _, price := range pricesMerchantA {
			rqliteRepository.Create(price)
		}

		t.Run("Successfully fetch all prices from user", func(t *testing.T) {
			fetchedPrices, err := rqliteRepository.FindAllByUser(userIdMerchantA)

			if err != nil {
				t.Error("Can't fetch prices")
			}

			if len(fetchedPrices) != len(pricesMerchantA) {
				t.Errorf("Unexpected price count. Expected %d, got %d", len(pricesMerchantA), len(fetchedPrices))
			}

			if !reflect.DeepEqual(pricesMerchantA, fetchedPrices) {
				t.Error("Fetched prices do not match expected prices")
			}
		})

		err := rqliteRepository.cleanTable()
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("TestIntegrationRQLiteRepository_FindAllByProduct", func(t *testing.T) {
		productIdA := uint64(1)

		pricesMerchantA := []*model.Price{
			{
				UserId:    1,
				ProductId: productIdA,
				Price:     2.99,
			},
			{
				UserId:    2,
				ProductId: productIdA,
				Price:     0.55,
			},
		}

		for _, price := range pricesMerchantA {
			rqliteRepository.Create(price)
		}

		t.Run("Successfully fetch all prices from product", func(t *testing.T) {
			fetchedPrices, err := rqliteRepository.FindAllByProduct(productIdA)

			if err != nil {
				t.Error("Can't fetch prices")
			}

			if len(fetchedPrices) != len(pricesMerchantA) {
				t.Errorf("Unexpected price count. Expected %d, got %d", len(pricesMerchantA), len(fetchedPrices))
			}

			if !reflect.DeepEqual(pricesMerchantA, fetchedPrices) {
				t.Error("Fetched prices do not match expected prices")
			}
		})

		err := rqliteRepository.cleanTable()
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("TestIntegrationRQLiteRepository_FindByIds", func(t *testing.T) {
		price := model.Price{
			UserId:    1,
			ProductId: 1,
			Price:     2.99,
		}

		rqliteRepository.Create(&price)

		t.Run("Successfully fetch price", func(t *testing.T) {
			fetchedPrice, err := rqliteRepository.FindByIds(price.UserId, price.ProductId)
			if err != nil {
				t.Errorf("Can't find expected price with userId %d and productId %d: %v", price.UserId, price.ProductId, err)
			}

			if !reflect.DeepEqual(price, *fetchedPrice) {
				t.Error("Fetched price does not match original price")
			}
		})

		t.Run("Fail to fetch price (price not found)", func(t *testing.T) {
			_, err := rqliteRepository.FindByIds(2, 2)
			if err == nil {
				t.Errorf("there should be an error")
			}
		})

		err = rqliteRepository.cleanTable()
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("TestIntegrationRQLiteRepository_Update", func(t *testing.T) {
		price := model.Price{
			UserId:    1,
			ProductId: 1,
			Price:     3.99,
		}

		changedPrice := model.Price{
			UserId:    1,
			ProductId: 1,
			Price:     2.99,
		}

		rqliteRepository.Create(&price)

		t.Run("Update price with success", func(t *testing.T) {
			updatedProduct, err := rqliteRepository.Update(&changedPrice)
			if reflect.DeepEqual(changedPrice, updatedProduct) || err != nil {
				t.Error("Failed to update price")
			}
		})

		t.Run("Update price with fail (price not found)", func(t *testing.T) {
			unknownPrice := model.Price{
				UserId:    3,
				ProductId: 4,
				Price:     3.99,
			}

			_, err := rqliteRepository.Update(&unknownPrice)
			if err == nil {
				t.Error("there should be an error")
			}
		})

		err = rqliteRepository.cleanTable()
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("TestIntegrationRQLiteRepository_Delete", func(t *testing.T) {
		priceToDelete := model.Price{
			UserId:    1,
			ProductId: 1,
			Price:     2.99,
		}

		rqliteRepository.Create(&priceToDelete)

		t.Run("Delete price with success", func(t *testing.T) {
			err = rqliteRepository.Delete(&priceToDelete)
			if err != nil {
				t.Errorf("Failed to delete price with userId %d and productId %d", priceToDelete.UserId, priceToDelete.ProductId)
			}
		})

		t.Run("Delete price with fail (price not found)", func(t *testing.T) {
			err := rqliteRepository.Delete(&priceToDelete)
			if err == nil {
				t.Error("there should be an error")
			}
		})

		err = rqliteRepository.cleanTable()
		if err != nil {
			t.Error(err)
		}
	})

	t.Cleanup(func() {
		err = container.Terminate(context.Background())
		if err != nil {
			return
		}
		os.Setenv("TESTCONTAINERS_RYUK_DISABLED", "false")
	})
}

func prepareIntegrationTestRQLiteContainer() (testcontainers.Container, error) {
	os.Setenv("TESTCONTAINERS_RYUK_DISABLED", "true")
	request := testcontainers.ContainerRequest{
		Image:        "rqlite/rqlite:8.15.0",
		ExposedPorts: []string{"4001/tcp"},
		WaitingFor: wait.ForAll(
			wait.ForListeningPort("4001/tcp"),
			wait.ForLog(`.*HTTP API available at.*`).AsRegexp(),
			wait.ForLog(".*entering leader state.*").AsRegexp()).
			WithStartupTimeoutDefault(120 * time.Second).
			WithDeadline(360 * time.Second),
	}

	return testcontainers.GenericContainer(
		context.Background(),
		testcontainers.GenericContainerRequest{
			ContainerRequest: request,
			Started:          true,
		})
}
