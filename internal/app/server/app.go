package server

import (
	"database/sql"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"shop-smart-api/internal/app"
	"shop-smart-api/internal/controller"
	di "shop-smart-api/internal/infrastructure/container"
	"shop-smart-api/internal/infrastructure/repository"
	"shop-smart-api/internal/service/payment"
	"shop-smart-api/internal/service/subscription_coupon"
	"shop-smart-api/pkg"
	adminpb "shop-smart-api/proto"
)

type application struct {
	database *sql.DB
	appCfg   *pkg.AppConfig
}

func CreateApplication(db *sql.DB, a *pkg.AppConfig) app.Application {
	return &application{db, a}
}

func (a *application) Run() error {
	container := di.CreateContainer(a.database, a.appCfg.Server, a.appCfg.Mailer)

	userService := container.ProvideUserService()
	otpService := container.ProvideOTPService()
	organizationService := container.ProvideOrganizationService()
	transactionService := container.ProvideTransactionService()

	transactionRepository := repository.CreatePaymentRepository(a.database)
	subscriptionCouponRepository := repository.CreateSubscriptionCouponRepo(a.database)
	paymentSvc := payment.New(transactionRepository, a.appCfg)

	conn, err := grpc.Dial("localhost:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	defer conn.Close()

	client := adminpb.NewAdminServiceClient(conn)
	subscriptionCouponService := subscription_coupon.New(subscriptionCouponRepository, a.appCfg, client)

	httpServer := controller.CreateServer(
		a.appCfg.Server,
		otpService,
		userService,
		organizationService,
		paymentSvc,
		subscriptionCouponService,
		transactionService,
		a.appCfg,
	)

	return httpServer.RunServer()
}
