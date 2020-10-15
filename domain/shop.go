package domain

import "time"

// Shop ..
type Shop struct {
	ID        int        `json:"shop_id"`
	UserID    int        `json:"user_id"`
	Name      string     `json:"name"`
	CreatedAt *time.Time `json:"created_at" gorm:"default:now();"`
	DeletedAt *time.Time `json:"deleted_at"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"default:now();"`
}

// ShopUsecase ..
type ShopUsecase interface {
}

// ShopRepository ..
type ShopRepository interface {
}
