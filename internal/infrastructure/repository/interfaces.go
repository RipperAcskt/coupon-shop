package repository

import (
	"shop-smart-api/internal/entity"
	"time"
)

type (
	UserRepository interface {
		Get(id string) (*entity.User, error)
		GetByPhone(phone string) (*entity.User, error)
		GetByEmail(email string) (*entity.User, error)
		GetByCode(code string) (*entity.User, error)
		GetByOrganization(id int64) ([]*entity.User, error)
		GetAll() ([]*entity.User, error)
		Store(phone, email, code string, roles []entity.Role) (*entity.User, error)
		UpdateUser(id string, email string) (*entity.User, error)
		AddOrganization(id string, organization int64, role *entity.Role) (*entity.User, error)
	}
	OTPRepository interface {
		GetByOwnerAndCode(owner string, code string) (*entity.OTP, error)
		Store(owner string, code string) (*entity.OTP, error)
		Use(otp *entity.OTP) error
	}
	OrganizationRepository interface {
		Get(id int64) (*entity.Organization, error)
		Store(name, kpp, orgn, inn string, owner int64) (*entity.Organization, error)
	}
	TransactionRepository interface {
		GetByOwner(id int64) ([]*entity.Transaction, error)
		Store(
			owner int64,
			trxNumber string,
			value float64,
			actionedAt *time.Time,
			status bool,
		) (*entity.Transaction, error)
	}
)
