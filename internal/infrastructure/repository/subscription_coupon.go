package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"shop-smart-api/internal/entity"
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

func (r *subscriptionCouponRepo) GetOrgSubscriptionLevel(email string) (int, error) {
	orgLevel := 0
	err := r.database.QueryRow("WITH subscription(organization_id) AS (SELECT organization_id FROM members WHERE email = '$1') SELECT MAX(level_subscription) FROM organization JOIN subscription ON subscription.organization_id = organization.id", email).Scan(&orgLevel)
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

func (r *subscriptionCouponRepo) GetRole(orgId, email string) (string, error) {
	role := ""
	err := r.database.QueryRow("SELECT role FROM members WHERE email=$1 AND organization_id = $2", email, orgId).Scan(&role)
	if err != nil {
		return "", err
	}
	if role == "" {
		return "", fmt.Errorf("role for this user is not defined")
	}
	return role, nil
}

func (r *subscriptionCouponRepo) GetCouponsPagination(pagination entity.PaginationInfo) ([]entity.CouponEntity, error) {

	rows, err := r.database.Query("SELECT * FROM coupons LIMIT $1 OFFSET $2", pagination.Limit, pagination.Offset)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, entity.ErrNoAnyCoupons
		}
		return nil, fmt.Errorf("query context failed: %w", err)
	}
	defer rows.Close()

	coupons := make([]entity.CouponEntity, 0)

	for rows.Next() {
		coupon := entity.CouponEntity{}
		err := rows.Scan(&coupon.ID, &coupon.Name, &coupon.Description, &coupon.Price, &coupon.Percent, &coupon.Level)
		if err != nil {
			return nil, fmt.Errorf("scan failed: %w", err)
		}

		media, err := r.getMyMedia(coupon.ID)
		if err != nil {
			return nil, fmt.Errorf("get media failed: %w", err)
		}

		coupon.Media = media

		coupons = append(coupons, coupon)
	}

	return coupons, nil
}

func (r *subscriptionCouponRepo) getMyMedia(id string) (entity.Media, error) {
	rows, err := r.database.Query("SELECT m.id, m.path FROM media m INNER JOIN coupons c ON m.coupon_id = c.id WHERE c.id = $1", id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.Media{}, entity.ErrNoMedia
		}

		return entity.Media{}, fmt.Errorf("query context media failed: %w", err)
	}
	rows.Next()
	media := entity.Media{}
	err = rows.Scan(&media.ID, &media.Path)
	if err != nil {
		return entity.Media{}, fmt.Errorf("scan failed: %w", err)
	}

	return media, nil
}

func (r *subscriptionCouponRepo) Get(id string) (*entity.User, error) {
	return r.executeQueryRow("SELECT * FROM users WHERE id = $1", id)
}

func (r *subscriptionCouponRepo) executeQueryRow(query string, args ...any) (*entity.User, error) {
	var user entity.User

	err := r.database.QueryRow(query, args...).Scan(
		&user.ID,
		&user.Email,
		&user.Phone,
		&user.Code,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.Subscription,
		&user.SubscriptionTime,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
