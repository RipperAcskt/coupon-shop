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
	GetRole(orgId, email string) (string, error)
	GetCouponsPagination(info entity.PaginationInfo) ([]entity.CouponEntity, error)
	GetCategories() ([]entity.CategorySubcategory, error)
	GetRegions() ([]entity.Region, error)
	GetLinks(region string) (entity.Link, error)
	Get(id string) (*entity.User, error)
	UpdateCoupon(coupon entity.CouponEntity) (string, error)
	GetCouponsSearchGRPC(s string) ([]entity.CouponEntity, error)
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
	r.group.Add("GET", "/links/:region", r.getLinks)
	r.group.Add("PUT", "/coupons", r.updateCoupon)
	r.group.Add("GET", "/coupons/:s", r.getCouponsSearch)

}

func (r *subscriptionCouponsRouteManager) getSubscriptions(c echo.Context) error {
	id := c.Get(middleware.CurrentUserKey)
	resp, err := r.svc.GetSubscriptions(fmt.Sprint(id.(string)))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, resp)
}

func (r *subscriptionCouponsRouteManager) getLinks(c echo.Context) error {
	region := c.Param("region")
	resp, err := r.svc.GetLinks(region)
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

	var regionSlice, categorySlice, respSlice, regionCategorySlice, couponsSlice []entity.CouponEntity
	var err error

	if region != "" {
		regionSlice, err = r.svc.GetCouponsByRegion(id.(string), region)
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
		categorySlice, err = r.svc.GetCouponsByCategory(id.(string), cat)
		if err != nil {
			return err
		}
		if len(categorySlice) == 0 {
			return c.JSON(http.StatusOK, categorySlice)
		}
	}

	coupons, err := r.svc.GetCoupons(id.(string))
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
	} else if len(categorySlice) != 0 && len(regionSlice) != 0 {

	} else {
		regionCategorySlice = coupons
	}
	respSlice = regionCategorySlice
	fmt.Println("regionCat", regionCategorySlice)

	if len(regionCategorySlice) != 0 {
		couponsSlice = intersect(regionCategorySlice, coupons)
	} else {
		return c.JSON(http.StatusOK, []entity.CouponEntity{})
	}
	respSlice = couponsSlice

	var limitNum int64
	if limit != "" {
		limitNum, err = strconv.ParseInt(limit, 10, 64)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		if int(limitNum) > len(respSlice) {
			respSlice = respSlice[:len(respSlice)]
		} else if limitNum <= 0 {
			return c.JSON(http.StatusOK, []entity.CouponEntity{})
		}
	}
	if offset != "" {
		offsetNum, err := strconv.ParseInt(offset, 10, 64)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		if int(offsetNum) > len(respSlice) || offsetNum < 0 || int(offsetNum) > len(respSlice) {
			return c.JSON(http.StatusOK, []entity.CouponEntity{})
		} else if int(offsetNum+limitNum) > len(respSlice) {
			respSlice = respSlice[offsetNum:]
		} else if limitNum > 0 {
			respSlice = respSlice[offsetNum : offsetNum+limitNum]
		} else {
			respSlice = respSlice[offsetNum:]
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

	if len(categorySlice) != 0 && len(regionSlice) != 0 {
		regionCategorySlice = intersect(categorySlice, regionSlice)
	} else if len(categorySlice) != 0 && len(regionSlice) == 0 {
		regionCategorySlice = categorySlice
	} else if len(categorySlice) == 0 && len(regionSlice) != 0 {
		regionCategorySlice = regionSlice
	} else if len(categorySlice) != 0 && len(regionSlice) != 0 {

	} else {
		regionCategorySlice = coupons
	}
	respSlice = regionCategorySlice
	fmt.Println("regionCat", regionCategorySlice)

	if len(regionCategorySlice) != 0 {
		couponsSlice = intersect(regionCategorySlice, coupons)
	} else {
		return c.JSON(http.StatusOK, []entity.CouponEntity{})
	}
	respSlice = couponsSlice

	var limitNum int64
	if limit != "" {
		limitNum, err = strconv.ParseInt(limit, 10, 64)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		if int(limitNum) > len(respSlice) {
			respSlice = respSlice[:len(respSlice)]
		} else if limitNum <= 0 {
			return c.JSON(http.StatusOK, []entity.CouponEntity{})
		}
	}
	if offset != "" {
		offsetNum, err := strconv.ParseInt(offset, 10, 64)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		if int(offsetNum) > len(respSlice) || offsetNum < 0 || int(offsetNum) > len(respSlice) {
			return c.JSON(http.StatusOK, []entity.CouponEntity{})
		} else if int(offsetNum+limitNum) > len(respSlice) {
			respSlice = respSlice[offsetNum:]
		} else if limitNum > 0 {
			respSlice = respSlice[offsetNum : offsetNum+limitNum]
		} else {
			respSlice = respSlice[offsetNum:]
		}
	}

	return c.JSON(http.StatusOK, respSlice)
}

func (r *subscriptionCouponsRouteManager) getCouponsSearch(c echo.Context) error {
	coupons, err := r.svc.GetCouponsSearchGRPC(c.Param("s"))
	if err != nil {
		return c.JSON(http.StatusOK, err.Error())
	}

	return c.JSON(http.StatusOK, coupons)
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
	user, err := r.svc.Get(id.(string))
	if err != nil {
		resp.Message = err.Error()
		return c.JSON(http.StatusBadRequest, resp)
	}

	role, err := r.svc.GetRole(organization.ID, user.Email)
	if err != nil {
		resp.Message = err.Error()
		return c.JSON(http.StatusBadRequest, resp)
	}
	fmt.Println(role)
	if role == "" {
		resp.Message = "role is not specified"
		return c.JSON(http.StatusBadRequest, resp)
	}
	message, err := r.svc.UpdateOrganizationInfo(organization, "", fmt.Sprint(id.(string)))
	if err != nil {
		resp.Message = message
		return c.JSON(http.StatusBadRequest, resp)
	}
	resp.Message = message
	return c.JSON(http.StatusOK, resp)
}

func (r *subscriptionCouponsRouteManager) updateCoupon(c echo.Context) error {
	coupon := entity.CouponEntity{}
	var resp UpdateResponse
	if err := c.Bind(&coupon); err != nil {
		return err
	}
	id := c.Get(middleware.CurrentUserKey)
	user, err := r.svc.Get(id.(string))
	if err != nil {
		resp.Message = err.Error()
		return c.JSON(http.StatusBadRequest, resp)
	}

	role, err := r.svc.GetRole(coupon.OrgId, user.Email)
	if err != nil {
		resp.Message = err.Error()
		return c.JSON(http.StatusBadRequest, resp)
	}
	fmt.Println(role)
	if role == "" {
		resp.Message = "role is not specified"
		return c.JSON(http.StatusBadRequest, resp)
	}
	message, err := r.svc.UpdateCoupon(coupon)
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
	user, err := r.svc.Get(id.(string))
	if err != nil {
		resp.Message = err.Error()
		return c.JSON(http.StatusBadRequest, resp)
	}

	for _, v := range members {
		role, err := r.svc.GetRole(v.OrganizationID, user.Email)
		if err != nil {
			resp.Message = err.Error()
			return c.JSON(http.StatusBadRequest, resp)
		}
		fmt.Println(role)
		if role == "" {
			resp.Message = "role is not specified"
			return c.JSON(http.StatusBadRequest, resp)
		}
	}
	message, err := r.svc.UpdateMembersInfo(members, "", fmt.Sprint(id.(string)))
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
