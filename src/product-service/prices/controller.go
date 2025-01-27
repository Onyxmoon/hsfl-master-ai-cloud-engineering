package prices

import "net/http"

type JsonFormatCreatePriceRequest struct {
	UserId    uint64  `json:"userId,omitempty"`
	ProductId uint64  `json:"productId,omitempty"`
	Price     float32 `json:"price,omitempty"`
}

type JsonFormatUpdatePriceRequest struct {
	Price float32 `json:"price,omitempty"`
}

type Controller interface {
	GetPrices(http.ResponseWriter, *http.Request)
	GetPricesByUser(http.ResponseWriter, *http.Request)
	GetPricesByProduct(http.ResponseWriter, *http.Request)
	GetPrice(http.ResponseWriter, *http.Request)
	PostPrice(http.ResponseWriter, *http.Request)
	PutPrice(http.ResponseWriter, *http.Request)
	DeletePrice(http.ResponseWriter, *http.Request)
}
