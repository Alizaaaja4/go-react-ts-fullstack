package structs

type ProductResponse struct {
	Id        uint    `json:"id"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Stock     int     `json:"stock"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}

type ProductCreateRequest struct {
	Name      string  `json:"name" binding:"required" gorm:"not null"`
	Price     float64 `json:"price" binding:"required" gorm:"not null;min=0"`
	Stock     int     `json:"stock" binding:"required" gorm:"not null;min=0"`
	CreatedBy uint    `json:"created_by" binding:"required"` // Tambahkan ini
}

type ProductUpdateRequest struct {
	Name      string  `json:"name" binding:"required" gorm:"not null"`
	Price     float64 `json:"price" binding:"required" gorm:"not null;min=0"`
	Stock     int     `json:"stock" binding:"required" gorm:"not null;min=0"`
	CreatedBy uint    `json:"created_by" binding:"required"` // Tambahkan ini
}
