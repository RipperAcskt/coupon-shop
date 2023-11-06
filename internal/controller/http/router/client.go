package router

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"shop-smart-api/internal/controller/http/middleware"
	"shop-smart-api/internal/entity"
	"shop-smart-api/pkg"
	"strconv"
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
	GetCategories() ([]entity.CategorySubcategory, error)
	GetRegions() ([]entity.Region, error)
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
	r.group.Add("GET", "/categories", r.getCategories)
	r.group.Add("GET", "/regions", r.getRegions)
	//r.group.Add("POST", "/coupons/pagination", r.getCouponsPagination, middleware.AuthMiddleware(r.serverConfig.Secret))
	r.group.Add("GET", "/coupons/standard", r.getCouponsStandard)
	//r.group.Add("GET", "/coupons/standard/filter/region/:region", r.getCouponsStandardByRegion)
	//r.group.Add("GET", "/coupons/standard/filter/category/:category", r.getCouponsStandardByCategory)
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

func (r *subscriptionCouponsRouteManager) getRegions(c echo.Context) error {
	resp, err := r.svc.GetRegions()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, resp)
}

func (r *subscriptionCouponsRouteManager) getCoupons(c echo.Context) error {
	id := c.Get(middleware.CurrentUserKey)
	region := c.QueryParam("region")
	category := c.QueryParam("category")
	subcategory := c.Request().Header["Subcategory"][0]
	limit := c.QueryParam("limit")
	offset := c.QueryParam("offset")

	var regionSlice, categorySlice, respSlice, couponsSlice, regionCategorySlice []entity.CouponEntity
	var err error

	if region != "" {
		regionSlice, err = r.svc.GetCouponsByRegion(fmt.Sprint(id.(string)), region)
		if err != nil {
			return err
		}
	}
	if category != "" {
		var cat entity.Category
		cat.Name = category
		if subcategory == "true" {
			cat.Subcategory = true
		} else {
			cat.Subcategory = false
		}
		categorySlice, err = r.svc.GetCouponsByCategory(fmt.Sprint(id.(string)), cat)
		if err != nil {
			return err
		}
	}
	coupons, err := r.svc.GetCoupons(fmt.Sprint(id.(string)))
	if err != nil {
		return err
	}
	respSlice = coupons

	if len(categorySlice) != 0 && len(regionSlice) != 0 {
		regionCategorySlice = intersect(categorySlice, regionSlice)
	} else if len(categorySlice) != 0 && len(regionSlice) == 0 {
		regionCategorySlice = categorySlice
	} else if len(categorySlice) == 0 && len(regionSlice) != 0 {
		regionCategorySlice = regionSlice
	} else {
		regionCategorySlice = coupons
	}
	respSlice = regionCategorySlice

	if len(regionCategorySlice) != 0 {
		couponsSlice = intersect(regionCategorySlice, coupons)
	} else {
		couponsSlice = coupons
	}
	respSlice = couponsSlice

	var limitNum int64
	if limit != "" {
		limitNum, err = strconv.ParseInt(limit, 10, 64)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		if int(limitNum) >= len(couponsSlice) {
			respSlice = couponsSlice[:len(couponsSlice)-1]
		} else if limitNum < 0 {
			respSlice = couponsSlice
		} else {
			respSlice = couponsSlice[:limitNum]
		}
	}
	if offset != "" {
		offsetNum, err := strconv.ParseInt(offset, 10, 64)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		if int(offsetNum) >= len(couponsSlice) || int(offsetNum+limitNum) > len(couponsSlice) {
			respSlice = couponsSlice[:len(couponsSlice)]
		} else if offsetNum < 0 {
			respSlice = couponsSlice
		} else {
			respSlice = couponsSlice[offsetNum : offsetNum+limitNum]
		}
	}

	return c.JSON(http.StatusOK, respSlice)
}

func (r *subscriptionCouponsRouteManager) getCouponsStandard(c echo.Context) error {
	region := c.QueryParam("region")
	category := c.QueryParam("category")
	subcategory := c.Request().Header["Subcategory"][0]
	limit := c.QueryParam("limit")
	offset := c.QueryParam("offset")

	var regionSlice, categorySlice, respSlice, couponsSlice, regionCategorySlice []entity.CouponEntity
	var err error

	if region != "" {
		regionSlice, err = r.svc.GetCouponsStandardByRegion(region)
		if err != nil {
			return err
		}
		if len(regionSlice) == 0 {
			return c.JSON(http.StatusOK, regionSlice)
		}
	}
	if category != "" {
		var cat entity.Category
		cat.Name = category
		if subcategory == "true" {
			cat.Subcategory = true
		} else {
			cat.Subcategory = false
		}
		categorySlice, err = r.svc.GetCouponsStandardByCategory(cat)
		if err != nil {
			return err
		}
		if len(categorySlice) == 0 {
			return c.JSON(http.StatusOK, categorySlice)
		}
	}
	coupons, err := r.svc.GetCouponsStandard()
	if err != nil {
		return err
	}
	respSlice = coupons
	fmt.Println("cat", categorySlice)
	fmt.Println("region", regionSlice)

	if len(categorySlice) != 0 && len(regionSlice) != 0 {
		regionCategorySlice = intersect(categorySlice, regionSlice)
	} else if len(categorySlice) != 0 && len(regionSlice) == 0 {
		regionCategorySlice = categorySlice
	} else if len(categorySlice) == 0 && len(regionSlice) != 0 {
		regionCategorySlice = regionSlice
	} else {
		regionCategorySlice = coupons
	}
	respSlice = regionCategorySlice
	fmt.Println("regionCat", regionCategorySlice)

	if len(regionCategorySlice) != 0 {
		couponsSlice = intersect(regionCategorySlice, coupons)
	} else {
		couponsSlice = coupons
	}
	respSlice = couponsSlice

	var limitNum int64
	if limit != "" {
		limitNum, err = strconv.ParseInt(limit, 10, 64)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		if int(limitNum) > len(couponsSlice) {
			respSlice = couponsSlice[:len(couponsSlice)]
		} else if limitNum < 0 {
			respSlice = couponsSlice
		} else {
			respSlice = couponsSlice[:limitNum]
		}
	}
	if offset != "" {
		offsetNum, err := strconv.ParseInt(offset, 10, 64)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		if int(offsetNum) > len(couponsSlice) {
			respSlice = []entity.CouponEntity{}
		} else if offsetNum < 0 {
			respSlice = couponsSlice
		} else if limitNum == 0 {
			respSlice = couponsSlice[offsetNum:]
		} else if int(offsetNum+limitNum) > len(couponsSlice) {
			respSlice = []entity.CouponEntity{}
		} else {
			respSlice = couponsSlice[offsetNum : offsetNum+limitNum]
		}
	}
	fmt.Println("resp", respSlice)

	return c.JSON(http.StatusOK, respSlice)
}

func (r *subscriptionCouponsRouteManager) getCategories(c echo.Context) error {
	categories, err := r.svc.GetCategories()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprint(err))
	}

	return c.JSON(http.StatusOK, categories)
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

func intersect(slice1 []entity.CouponEntity, slice2 []entity.CouponEntity) []entity.CouponEntity {
	var slice []entity.CouponEntity
	for i1 := 0; i1 < len(slice1); i1++ {
		for i2 := 0; i2 < len(slice2); i2++ {
			if slice1[i1] == slice2[i2] {
				slice = append(slice, slice1[i1])
			}
		}
	}
	return slice
}
