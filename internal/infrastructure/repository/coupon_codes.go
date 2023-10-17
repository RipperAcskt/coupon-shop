package repository

import (
	"database/sql"
	"shop-smart-api/internal/entity"
)

type codesRepository struct {
	database *sql.DB
}

func CreateCodesRepository(db *sql.DB) *codesRepository {
	return &codesRepository{db}
}

func (r *codesRepository) CreateCode(code entity.UsersCodes) (*entity.UsersCodes, error) {
	return r.executeQueryRow(`INSERT INTO coupon_codes (id, user_id, code) VALUES ($1, $2, $3) RETURNING id, user_id, code`, code.Id, code.UserId, code.Coupon)
}

func (r *codesRepository) GetCodeByCoupon(coupon string) (*entity.UsersCodes, error) {
	return r.executeQueryRow("SELECT * FROM coupon_codes WHERE code = $1", coupon)
}

func (r *codesRepository) GetCodeByUserID(userID string) ([]entity.UsersCodes, error) {
	return r.executeQuery("SELECT * FROM coupon_codes WHERE user_id = $1", userID)
}

func (r *codesRepository) DeleteCode(coupon string) (*entity.UsersCodes, error) {
	row := r.database.QueryRow("DELETE FROM coupon_codes WHERE code = $1", coupon)
	if row.Err() != nil {
		return nil, row.Err()
	}
	return nil, nil
}

func (r *codesRepository) executeQuery(query string, args ...any) ([]entity.UsersCodes, error) {
	rows, err := r.database.Query(query, args...)
	if err != nil {
		return nil, err
	}

	codes := make([]entity.UsersCodes, 0)
	for rows.Next() {
		var code entity.UsersCodes

		err := rows.Scan(
			&code.Id,
			&code.UserId,
			&code.Coupon,
		)
		if err != nil {
			return nil, err
		}

		codes = append(codes, code)
	}

	return codes, nil
}

func (r *codesRepository) executeQueryRow(query string, args ...any) (*entity.UsersCodes, error) {
	var code entity.UsersCodes

	err := r.database.QueryRow(query, args...).Scan(
		&code.Id,
		&code.UserId,
		&code.Coupon,
	)
	if err != nil {
		return nil, err
	}

	return &code, nil
}
