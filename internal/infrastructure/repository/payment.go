package repository

import (
	"database/sql"
	"fmt"
	"shop-smart-api/internal/entity"
)

type paymentRepository struct {
	database *sql.DB
}

func CreatePaymentRepository(db *sql.DB) *paymentRepository {
	return &paymentRepository{db}
}

func (r *paymentRepository) CreatePayment(payment entity.Payment) (*entity.Payment, error) {
	return r.executeQueryRow(`INSERT INTO transactions (id, value, status, owner_id, payment_type, entity_id) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, value, status, updated_at, owner_id, payment_type, entity_id`, payment.Id, payment.Value, payment.Status, payment.UserID, payment.PaymentType, payment.TypeID)
}

func (r *paymentRepository) GetPayments(userId string) ([]entity.Payment, error) {
	return r.executeQuery("SELECT * FROM transactions WHERE owner_id = $1", userId)
}

func (r *paymentRepository) UpdatePayment(id string) (*entity.Payment, error) {
	return r.executeQueryRow(`
		UPDATE transactions SET status = true WHERE id = $1
		RETURNING id, value, status, updated_at, owner_id, payment_type, entity_id
	`, id)
}

func (r *paymentRepository) UpdateSubscription(id, userId string) error {
	row := r.database.QueryRow("SELECT level FROM subscriptions WHERE id = $1", id)
	if row.Err() != nil {
		return row.Err()
	}

	var level int
	err := row.Scan(&level)
	if err != nil {
		return fmt.Errorf("scan failed: %w", err)
	}

	_, err = r.database.Exec("UPDATE users SET subscription = $1 WHERE id = $2", level, userId)
	if err != nil {
		return fmt.Errorf("exec failed: %w", err)
	}

	return nil
}

func (r *paymentRepository) executeQuery(query string, args ...any) ([]entity.Payment, error) {
	rows, err := r.database.Query(query, args...)
	if err != nil {
		return nil, err
	}

	payments := make([]entity.Payment, 0)
	for rows.Next() {
		var payment entity.Payment

		err := rows.Scan(
			&payment.Id,
			&payment.Value,
			&payment.Status,
			&payment.UpdatedAt,
			&payment.UserID,
			&payment.PaymentType,
			&payment.TypeID,
		)
		if err != nil {
			return nil, err
		}

		payments = append(payments, payment)
	}

	return payments, nil
}

func (r *paymentRepository) executeQueryRow(query string, args ...any) (*entity.Payment, error) {
	var payment entity.Payment

	err := r.database.QueryRow(query, args...).Scan(
		&payment.Id,
		&payment.Value,
		&payment.Status,
		&payment.UpdatedAt,
		&payment.UserID,
		&payment.PaymentType,
		&payment.TypeID,
	)
	if err != nil {
		return nil, err
	}

	return &payment, nil
}
