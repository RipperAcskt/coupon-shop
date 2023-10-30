package controller

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"shop-smart-api/internal/controller/graphql/directives"
	"shop-smart-api/internal/controller/graphql/graph"
	"shop-smart-api/internal/controller/graphql/transformers"
	http_context "shop-smart-api/internal/controller/http/context"
	"shop-smart-api/internal/controller/http/router"
	http_validator "shop-smart-api/internal/controller/http/validator"
	"shop-smart-api/internal/service"
	"shop-smart-api/internal/service/transaction"
	"shop-smart-api/pkg"
)

type Server interface {
	RunServer() error
}

type http struct {
	otpService              service.OTPService
	userService             service.UserService
	userTransformer         transformers.UserTransformer
	organizationService     service.OrganizationService
	organizationTransformer transformers.OrganizationTransformer
	transactionService      transaction.PaymentService
	paymentSvc              router.PaymentService
	subscriptionCouponSvc   router.SubscriptionCouponService
	serverConfig            pkg.Server
	validator               *http_validator.Validator
	echo                    *echo.Echo
	cfg                     *pkg.AppConfig
}

func CreateServer(sc pkg.Server, ots service.OTPService, us service.UserService, ogs service.OrganizationService, paymnetSvc router.PaymentService, subscriptionCouponSvc router.SubscriptionCouponService, ts transaction.PaymentService, cfg *pkg.AppConfig) Server {
	v := http_validator.CreateValidator(validator.New())
	e := echo.New()
	e.Validator = v
	e.Use(http_context.EchoContextToContextMiddleware())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	ut := transformers.CreateUserTransformer()
	ot := transformers.CreateOrganizationTransformer()

	return &http{
		otpService:              ots,
		userService:             us,
		userTransformer:         ut,
		organizationService:     ogs,
		organizationTransformer: ot,
		transactionService:      ts,
		subscriptionCouponSvc:   subscriptionCouponSvc,
		serverConfig:            sc,
		paymentSvc:              paymnetSvc,
		validator:               v,
		echo:                    e,
		cfg:                     cfg,
	}
}

func (h *http) RunServer() error {
	h.appendRestRoutes(h.echo)
	h.appendGraphqlRoutes(h.echo)

	host := fmt.Sprintf(":%s", h.serverConfig.Port)
	return h.echo.Start(host)
}

func (h *http) appendRestRoutes(e *echo.Echo) {
	apiGroup := e.Group("/api")
	authRouter := router.CreateAuthRouterManager(h.subscriptionCouponSvc, apiGroup, h.validator, h.userService, h.otpService)
	authRouter.PopulateRoutes()
	paymentRouter := router.CreatePaymentRouterManager(apiGroup, h.paymentSvc, h.cfg.Server)
	paymentRouter.PopulateRoutes()

	subscriptionCouponRouter := router.CreateSubscriptionCouponService(apiGroup, h.subscriptionCouponSvc, h.cfg.Server)
	subscriptionCouponRouter.PopulateRoutes()

	otpGroup := apiGroup.Group("/otp")
	otpRouter := router.CreateOTPRouterManager(otpGroup, h.validator, h.userService, h.otpService, h.serverConfig)
	otpRouter.PopulateRoutes()
}

func (h *http) appendGraphqlRoutes(e *echo.Echo) {
	resolver := graph.CreateResolver(
		h.userService,
		h.userTransformer,
		h.organizationService,
		h.organizationTransformer,
	)
	c := graph.Config{Resolvers: resolver}
	c.Directives.Auth = directives.Auth
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(c))
	pg := playground.Handler("GraphQL playground", "/query")

	graphqlRouter := router.CreateGraphqlRouterManager(e.Group(""), srv, pg, h.serverConfig)

	graphqlRouter.PopulateRoutes()
}
