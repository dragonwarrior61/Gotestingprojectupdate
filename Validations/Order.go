package Validations

type Image struct {
	Thumbnail string `json:"thumbnail"`
	Mobile    string `json:"mobile"`
	Tablet    string `json:"tablet"`
	Desktop   string `json:"desktop"`
}

type ProductCreate struct {
	Image    Image  `json:"image"`
	Name     string `json:"name" binding:"required"`
	Category string `json:"category" binding:"required"`
	Price    string `json:"price" binding:"required"`
}

type ProductUpdate struct {
	Image    Image  `json:"image"`
	Name     string `json:"name" binding:"required"`
	Category string `json:"category" binding:"required"`
	Price    string `json:"price" binding:"required"`
}
