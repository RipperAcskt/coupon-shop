package server

import (
	"database/sql"
	"shop-smart-api/internal/app"
	"shop-smart-api/internal/controller"
	di "shop-smart-api/internal/infrastructure/container"
	"shop-smart-api/internal/infrastructure/repository"
	"shop-smart-api/internal/service/payment"
	"shop-smart-api/pkg"
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
	paymentSvc := payment.New(transactionRepository, a.appCfg)

	httpServer := controller.CreateServer(
		a.appCfg.Server,
		otpService,
		userService,
		organizationService,
		paymentSvc,
		transactionService,
		a.appCfg,
	)

	return httpServer.RunServer()
}
