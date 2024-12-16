package Schema

import "time"

type Image struct {
	Thumbnail string `json:"thumbnail"`
	Mobile    string `json:"mobile"`
	Tablet    string `json:"tablet"`
	Desktop   string `json:"desktop"`
}

type FoodItem struct {
	ID       string  `json:"id"`
	Image    Image   `json:"image"`
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Price    float64 `json:"price"`
}

type Product struct {
	Id        uint
	Image     Image     `json:"image"`
	Name      string    `json:"name"`
	Category  string    `json:"category"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `sql:"DEFAULT:current_timestamp"`
	UpdatedAt time.Time
}

func (b *Product) TableName() string {
	return "products"
}
