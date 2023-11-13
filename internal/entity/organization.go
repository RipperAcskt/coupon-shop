package entity

import "time"

type Organization struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	ID        int64
	OwnerID   int64
	Name      string
	ORGN      string
	KPP       string
	INN       string
}

type OrganizationEntity struct {
	ID                string   `json:"id"`
	Name              string   `json:"name"`
	EmailAdmin        string   `json:"email_admin"`
	LevelSubscription int      `json:"level_subscription"`
	ORGN              string   `json:"orgn"`
	KPP               string   `json:"kpp"`
	INN               string   `json:"inn"`
	Address           string   `json:"address"`
	Members           []Member `json:"members"`
	ContentUrl        string   `json:"content_url"`
	//image
}

type Member struct {
	ID             string `json:"id"`
	Email          string `json:"email"`
	FirstName      string `json:"name"`
	SecondName     string `json:"second_name"`
	OrganizationID string `json:"organization_ID"`
	Role           string `json:"role"`
}
