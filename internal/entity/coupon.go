package entity

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
}
