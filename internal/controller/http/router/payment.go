package router

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"shop-smart-api/internal/entity"
)

type paymentRouteManager struct {
	group *echo.Group
	svc   PaymentService
}

type PaymentService interface {
	CreatePayment(paymentRequest *entity.CreatePaymentRequest) (interface{}, error)
}

func CreatePaymentRouterManager(g *echo.Group, svc PaymentService) RouteManager {
	return &paymentRouteManager{
		group: g,
		svc:   svc,
	}
}

func (r *paymentRouteManager) PopulateRoutes() {
	r.group.Add("POST", "/payment", r.createPayment)
}

func (r *paymentRouteManager) createPayment(c echo.Context) error {
	paymentRequest := &entity.CreatePaymentRequest{}
	if err := c.Bind(paymentRequest); err != nil {
		return err
	}

	resp, err := r.svc.CreatePayment(paymentRequest)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, resp)
}
