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
	LevelSubscription int      `json:"levelSubscription"`
	Members           []Member `json:"members"`
}

type Member struct {
	ID             string `json:"id"`
	Email          string `json:"email"`
	FirstName      string `json:"name"`
	SecondName     string `json:"secondName"`
	OrganizationID string `json:"organizationID"`
}
