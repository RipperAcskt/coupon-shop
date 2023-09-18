// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Organization struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Orgn      string `json:"orgn"`
	Kpp       string `json:"kpp"`
	Inn       string `json:"inn"`
	OwnerID   string `json:"ownerId"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type Transaction struct {
	ID         string  `json:"id"`
	OwnerID    string  `json:"ownerId"`
	Value      float64 `json:"value"`
	TrxNumber  string  `json:"trxNumber"`
	Status     bool    `json:"status"`
	CreatedAt  string  `json:"createdAt"`
	UpdatedAt  string  `json:"updatedAt"`
	ActionedAt *string `json:"actionedAt,omitempty"`
}

type UpdateUser struct {
	Email string `json:"email"`
}

type User struct {
	ID             string   `json:"id"`
	Email          *string  `json:"email,omitempty"`
	Phone          string   `json:"phone"`
	Roles          []string `json:"roles"`
	OrganizationID *string  `json:"organizationId,omitempty"`
	CreatedAt      string   `json:"createdAt"`
	UpdatedAt      string   `json:"updatedAt"`
}
