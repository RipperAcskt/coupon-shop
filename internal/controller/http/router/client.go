package router

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"shop-smart-api/internal/controller/http/middleware"
	"shop-smart-api/internal/entity"
	"shop-smart-api/pkg"
)

type subscriptionCouponsRouteManager struct {
	group        *echo.Group
	svc          SubscriptionCouponService
	serverConfig pkg.Server
}

type SubscriptionCouponService interface {
	GetSubscriptions(userId string) ([]entity.SubscriptionEntity, error)
	GetCoupons(userId string) ([]entity.CouponEntity, error)
	GetOrganizationInfo(userId string) (entity.OrganizationEntity, error)
}

func CreateSubscriptionCouponService(g *echo.Group, svc SubscriptionCouponService, cfg pkg.Server) RouteManager {
	return &subscriptionCouponsRouteManager{
		group:        g,
		svc:          svc,
		serverConfig: cfg,
	}
}

func (r *subscriptionCouponsRouteManager) PopulateRoutes() {
	r.group.Add("GET", "/subscriptions", r.getSubscriptions, middleware.OTPAuthMiddleware(r.serverConfig.Secret))
	r.group.Add("GET", "/coupons", r.getCoupons, middleware.OTPAuthMiddleware(r.serverConfig.Secret))
	r.group.Add("GET", "/organizationInfo", r.getOrganizationInfo, middleware.OTPAuthMiddleware(r.serverConfig.Secret))
}

func (r *subscriptionCouponsRouteManager) getSubscriptions(c echo.Context) error {
	id := c.Get(middleware.CurrentUserKey)
	resp, err := r.svc.GetSubscriptions(fmt.Sprint(id.(string)))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, resp)
}

func (r *subscriptionCouponsRouteManager) getCoupons(c echo.Context) error {
	id := c.Get(middleware.CurrentUserKey)
	resp, err := r.svc.GetCoupons(fmt.Sprint(id.(string)))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, resp)
}

func (r *subscriptionCouponsRouteManager) getOrganizationInfo(c echo.Context) error {
	id := c.Get(middleware.CurrentUserKey)
	resp, err := r.svc.GetOrganizationInfo(fmt.Sprint(id.(string)))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, resp)
}
