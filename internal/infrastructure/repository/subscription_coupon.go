package repository

import (
	"database/sql"
	"errors"
)

type subscriptionCouponRepo struct {
	database *sql.DB
}

func CreateSubscriptionCouponRepo(db *sql.DB) *subscriptionCouponRepo {
	return &subscriptionCouponRepo{db}
}

func (r *subscriptionCouponRepo) GetUserSubscriptionLevel(userId string) (int, error) {
	var subLevel *int
	err := r.database.QueryRow("SELECT subscription FROM users WHERE id = $1", userId).Scan(&subLevel)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return 0, err
		}
	}
	if subLevel == nil {
		return 0, nil
	}

	return *subLevel, nil
}

func (r *subscriptionCouponRepo) GetEmailUser(userID string) (string, error) {
	email := ""
	err := r.database.QueryRow("SELECT email FROM users WHERE id = $1", userID).Scan(&email)
	if err != nil {
		return "", err
	}
	return email, nil
}

func (r *subscriptionCouponRepo) GetOrgSubscriptionLevel(orgID string) (int, error) {
	orgLevel := 0
	err := r.database.QueryRow("SELECT level_subscription FROM organization WHERE id = $1", orgID).Scan(&orgLevel)
	if err != nil {
		return 0, err
	}
	return orgLevel, nil
}

func (r *subscriptionCouponRepo) GetOrgId(email string) (string, error) {
	orgID := ""
	err := r.database.QueryRow("SELECT organization_id FROM members WHERE email = $1", email).Scan(&orgID)
	if err != nil {
		return "", err
	}
	return orgID, nil
}
