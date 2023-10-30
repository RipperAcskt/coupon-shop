package entity

import "errors"

type SubscriptionEntity struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Level       int32   `json:"level"`
}

var ErrUserSubEmpty = errors.New("user subscription is empty")
