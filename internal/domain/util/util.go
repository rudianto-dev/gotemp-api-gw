package util

type Service string

var (
	USER_SERVICE        Service = "user"
	BUSINESS_SERVICE    Service = "business"
	TRANSACTION_SERVICE Service = "transaction"
)

type GetHealthRequest struct{}

type GetHealthResponse struct {
	Status  int
	Message string
}

type GetServiceHealthRequest struct {
	Service string `json:"service" validate:"required"`
}

type GetServiceHealthResponse struct {
	Service string `json:"service"`
	Status  int    `json:"status"`
	Message string `json:"message"`
}
