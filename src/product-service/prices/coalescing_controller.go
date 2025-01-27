package prices

import (
	"encoding/json"
	"fmt"
	"golang.org/x/sync/singleflight"
	"hsfl.de/group6/hsfl-master-ai-cloud-engineering/lib/router/middleware/auth"
	"hsfl.de/group6/hsfl-master-ai-cloud-engineering/product-service/prices/model"
	"net/http"
	"strconv"
)

type CoalescingController struct {
	priceRepository Repository
	group           *singleflight.Group
}

func NewCoalescingController(priceRepository Repository) *CoalescingController {
	return &CoalescingController{
		priceRepository,
		&singleflight.Group{}}
}

func (controller *CoalescingController) GetPrices(writer http.ResponseWriter, request *http.Request) {
	msg, err, _ := controller.group.Do("get-all", func() (interface{}, error) {
		return controller.priceRepository.FindAll()
	})

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}

	writer.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(writer).Encode(msg)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (controller *CoalescingController) GetPricesByUser(writer http.ResponseWriter, request *http.Request) {
	userIdAttribute := request.Context().Value("userId").(string)

	msg, err, _ := controller.group.Do("get_id_"+userIdAttribute, func() (interface{}, error) {
		userId, err := strconv.ParseUint(userIdAttribute, 10, 64)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return nil, err
		}

		value, err := controller.priceRepository.FindAllByUser(userId)
		if err != nil {
			if err.Error() == ErrorPriceNotFound {
				http.Error(writer, err.Error(), http.StatusNotFound)
				return nil, err
			}
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}

		return value, nil
	})

	writer.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(writer).Encode(msg)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (controller *CoalescingController) GetPricesByProduct(writer http.ResponseWriter, request *http.Request) {
	productIdAttribute := request.Context().Value("productId").(string)

	msg, err, _ := controller.group.Do("get_id_"+productIdAttribute, func() (interface{}, error) {
		productId, err := strconv.ParseUint(productIdAttribute, 10, 64)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return nil, err
		}

		value, err := controller.priceRepository.FindAllByProduct(productId)
		if err != nil {
			if err.Error() == ErrorPriceNotFound {
				http.Error(writer, err.Error(), http.StatusNotFound)
				return nil, err
			}
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}

		return value, nil
	})

	writer.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(writer).Encode(msg)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (controller *CoalescingController) GetPrice(writer http.ResponseWriter, request *http.Request) {
	userIdAttribute := request.Context().Value("userId").(string)
	productIdAttribute := request.Context().Value("productId").(string)

	msg, err, _ := controller.group.Do(fmt.Sprintf("get_id_%s_%s", userIdAttribute, productIdAttribute), func() (interface{}, error) {
		userId, err := strconv.ParseUint(userIdAttribute, 10, 64)
		productId, err := strconv.ParseUint(productIdAttribute, 10, 64)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return nil, err
		}

		value, err := controller.priceRepository.FindByIds(productId, userId)
		if err != nil {
			if err.Error() == ErrorPriceNotFound {
				http.Error(writer, err.Error(), http.StatusNotFound)
				return nil, err
			}
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}

		return value, nil
	})

	writer.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(writer).Encode(msg)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}

func (controller *CoalescingController) PostPrice(writer http.ResponseWriter, request *http.Request) {
	productId, productIdErr := strconv.ParseUint(request.Context().Value("productId").(string), 10, 64)
	userId, userIdErr := strconv.ParseUint(request.Context().Value("userId").(string), 10, 64)

	if productIdErr != nil || userIdErr != nil {
		http.Error(writer, "Invalid listId or productId", http.StatusBadRequest)
		return
	}

	var requestData JsonFormatCreatePriceRequest
	if err := json.NewDecoder(request.Body).Decode(&requestData); err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	authUserId, _ := request.Context().Value("auth_userId").(uint64)
	authUserRole, _ := request.Context().Value("auth_userRole").(int64)

	if (authUserRole == auth.Administrator) ||
		(authUserRole == auth.Merchant && authUserId == userId) {
		if _, err := controller.priceRepository.Create(&model.Price{
			ProductId: productId,
			UserId:    userId,
			Price:     requestData.Price,
		}); err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
		}
	} else {
		writer.WriteHeader(http.StatusUnauthorized)
	}
}

func (controller *CoalescingController) PutPrice(writer http.ResponseWriter, request *http.Request) {
	userId, err := strconv.ParseUint(request.Context().Value("userId").(string), 10, 64)
	productId, err := strconv.ParseUint(request.Context().Value("productId").(string), 10, 64)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	var requestData JsonFormatUpdatePriceRequest
	if err := json.NewDecoder(request.Body).Decode(&requestData); err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	authUserId, _ := request.Context().Value("auth_userId").(uint64)
	authUserRole, _ := request.Context().Value("auth_userRole").(int64)

	if (authUserRole == auth.Administrator) ||
		(authUserRole == auth.Merchant && authUserId == userId) {
		if _, err := controller.priceRepository.Update(&model.Price{
			UserId:    userId,
			ProductId: productId,
			Price:     requestData.Price,
		}); err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else {
		writer.WriteHeader(http.StatusUnauthorized)
	}
}

func (controller *CoalescingController) DeletePrice(writer http.ResponseWriter, request *http.Request) {
	userId, err := strconv.ParseUint(request.Context().Value("userId").(string), 10, 64)
	productId, err := strconv.ParseUint(request.Context().Value("productId").(string), 10, 64)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	authUserId, _ := request.Context().Value("auth_userId").(uint64)
	authUserRole, _ := request.Context().Value("auth_userRole").(int64)

	if (authUserRole == auth.Administrator) ||
		(authUserRole == auth.Merchant && authUserId == userId) {
		if err := controller.priceRepository.Delete(&model.Price{ProductId: productId, UserId: userId}); err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else {
		writer.WriteHeader(http.StatusUnauthorized)
	}
}
