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
	GetCouponsByRegion(userId, region string) ([]entity.CouponEntity, error)
	GetCouponsByCategory(userId string, category entity.Category) ([]entity.CouponEntity, error)
	GetOrganizationInfo(userId string) (entity.OrganizationEntity, error)
	GetCouponsStandard() ([]entity.CouponEntity, error)
	GetCouponsStandardByRegion(region string) ([]entity.CouponEntity, error)
	GetCouponsStandardByCategory(category entity.Category) ([]entity.CouponEntity, error)
	UpdateOrganizationInfo(organizationEntity entity.OrganizationEntity, role string, id string) (string, error)
	UpdateMembersInfo(members []entity.Member, role string, id string) (string, error)
	GetRole(email string) (string, error)
	GetCouponsPagination(info entity.PaginationInfo) ([]entity.CouponEntity, error)
}
type UpdateResponse struct {
	Message string `json:"message"`
}

func CreateSubscriptionCouponService(g *echo.Group, svc SubscriptionCouponService, cfg pkg.Server) RouteManager {
	return &subscriptionCouponsRouteManager{
		group:        g,
		svc:          svc,
		serverConfig: cfg,
	}
}

func (r *subscriptionCouponsRouteManager) PopulateRoutes() {
	r.group.Add("GET", "/subscriptions", r.getSubscriptions, middleware.AuthMiddleware(r.serverConfig.Secret))
	r.group.Add("GET", "/coupons", r.getCoupons, middleware.AuthMiddleware(r.serverConfig.Secret))
	r.group.Add("GET", "/coupons/filter/region/:region", r.getCouponsByRegion, middleware.AuthMiddleware(r.serverConfig.Secret))
	r.group.Add("GET", "/coupons/filter/category/:category", r.getCouponsByCategory, middleware.AuthMiddleware(r.serverConfig.Secret))
	r.group.Add("POST", "/coupons/pagination", r.getCouponsPagination, middleware.AuthMiddleware(r.serverConfig.Secret))
	r.group.Add("GET", "/coupons/standard", r.getCouponsStandard)
	r.group.Add("GET", "/coupons/standard/filter/region/:region", r.getCouponsStandardByRegion)
	r.group.Add("GET", "/coupons/standard/filter/category/:category", r.getCouponsStandardByCategory)
	r.group.Add("GET", "/organizationInfo", r.getOrganizationInfo, middleware.AuthMiddleware(r.serverConfig.Secret))
	r.group.Add("PUT", "/organizationInfo", r.updateOrganizationInfo, middleware.AuthMiddleware(r.serverConfig.Secret))
	r.group.Add("PUT", "/membersInfo", r.updateMembersInfo, middleware.AuthMiddleware(r.serverConfig.Secret))
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

func (r *subscriptionCouponsRouteManager) getCouponsByRegion(c echo.Context) error {
	id := c.Get(middleware.CurrentUserKey)
	region := c.Param("region")
	resp, err := r.svc.GetCouponsByRegion(fmt.Sprint(id.(string)), region)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, resp)
}

func (r *subscriptionCouponsRouteManager) getCouponsByCategory(c echo.Context) error {
	id := c.Get(middleware.CurrentUserKey)
	category := c.Param("category")
	subcategory := c.Request().Header["Subcategory"][0]
	var cat entity.Category
	cat.Name = category
	if subcategory == "true" {
		cat.Subcategory = true
	} else {
		cat.Subcategory = false
	}
	resp, err := r.svc.GetCouponsByCategory(fmt.Sprint(id.(string)), cat)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, resp)
}

func (r *subscriptionCouponsRouteManager) getCouponsStandardByCategory(c echo.Context) error {
	category := c.Param("category")
	subcategory := c.Request().Header["Subcategory"][0]
	var cat entity.Category
	cat.Name = category
	if subcategory == "true" {
		cat.Subcategory = true
	} else {
		cat.Subcategory = false
	}
	resp, err := r.svc.GetCouponsStandardByCategory(cat)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, resp)
}

func (r *subscriptionCouponsRouteManager) getCouponsPagination(c echo.Context) error {
	info := entity.PaginationInfo{}
	if err := c.Bind(&info); err != nil {
		return err
	}
	resp, err := r.svc.GetCouponsPagination(info)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, resp)
}

func (r *subscriptionCouponsRouteManager) getCouponsStandard(c echo.Context) error {
	resp, err := r.svc.GetCouponsStandard()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, resp)
}

func (r *subscriptionCouponsRouteManager) getCouponsStandardByRegion(c echo.Context) error {
	region := c.Param("region")
	resp, err := r.svc.GetCouponsStandardByRegion(region)
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

func (r *subscriptionCouponsRouteManager) updateOrganizationInfo(c echo.Context) error {
	organization := entity.OrganizationEntity{}
	var resp UpdateResponse
	if err := c.Bind(&organization); err != nil {
		return err
	}
	id := c.Get(middleware.CurrentUserKey)
	role := c.Get(middleware.CurrentUserRole)
	fmt.Println(role.(string))
	if role == nil {
		resp.Message = "role is not specified"
		return c.JSON(http.StatusBadRequest, resp)
	}
	message, err := r.svc.UpdateOrganizationInfo(organization, fmt.Sprint(role.(string)), fmt.Sprint(id.(string)))
	if err != nil {
		resp.Message = message
		return c.JSON(http.StatusBadRequest, resp)
	}
	resp.Message = message
	return c.JSON(http.StatusOK, resp)
}

func (r *subscriptionCouponsRouteManager) updateMembersInfo(c echo.Context) error {
	var members []entity.Member
	var resp UpdateResponse
	if err := c.Bind(&members); err != nil {
		return err
	}
	id := c.Get(middleware.CurrentUserKey)
	role := c.Get(middleware.CurrentUserRole)
	fmt.Println(role.(string))
	if role == nil {
		resp.Message = "role is not specified"
		return c.JSON(http.StatusBadRequest, resp)
	}
	message, err := r.svc.UpdateMembersInfo(members, fmt.Sprint(role.(string)), fmt.Sprint(id.(string)))
	if err != nil {
		resp.Message = message
		return c.JSON(http.StatusBadRequest, resp)
	}
	resp.Message = message
	return c.JSON(http.StatusOK, resp)
}
