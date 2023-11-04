package entity

import (
	"errors"
	"fmt"
)

var (
	ErrNoAnyCoupons       = errors.New("there is not a single coupon")
	ErrCouponDoesNotExist = errors.New("coupon does not exist")
	ErrNoMedia            = fmt.Errorf("no media")
)

type Media struct {
	ID   string `json:"ID,omitempty"`
	Path string `json:"path,omitempty"`
}

type CouponEntity struct {
	ID          string  `json:"ID"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Level       int32   `json:"level"`
	Percent     int32   `json:"percent"`
	ContentUrl  string  `json:"content_url"`
	Media       *Media  `json:"media"`
	Region      string  `json:"region"`
	Category    string  `json:"category"`
	Subcategory string  `json:"subcategory"`
}

type PaginationInfo struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type Category struct {
	Name        string
	Subcategory bool
}
