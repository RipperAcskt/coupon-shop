package router

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"shop-smart-api/internal/controller/http/middleware"
	"shop-smart-api/internal/entity"
	"shop-smart-api/pkg"
)

type paymentRouteManager struct {
	group        *echo.Group
	svc          PaymentService
	serverConfig pkg.Server
}

type PaymentService interface {
	CreatePayment(paymentRequest *entity.CreatePaymentRequest, userId string) (interface{}, error)
}

func CreatePaymentRouterManager(g *echo.Group, svc PaymentService, cfg pkg.Server) RouteManager {
	return &paymentRouteManager{
		group:        g,
		svc:          svc,
		serverConfig: cfg,
	}
}

func (r *paymentRouteManager) PopulateRoutes() {
	r.group.Add("POST", "/payment", r.createPayment, middleware.OTPAuthMiddleware(r.serverConfig.Secret))
}

func (r *paymentRouteManager) createPayment(c echo.Context) error {
	paymentRequest := &entity.CreatePaymentRequest{}
	if err := c.Bind(paymentRequest); err != nil {
		return err
	}

	id := c.Get(middleware.CurrentUserKey)

	resp, err := r.svc.CreatePayment(paymentRequest, fmt.Sprint(id.(string)))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, resp)
}
