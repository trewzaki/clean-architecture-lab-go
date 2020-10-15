package domain

import "time"

// User ..
type User struct {
	ID        int        `json:"user_id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	CreatedAt *time.Time `json:"created_at" gorm:"default:now();"`
	DeletedAt *time.Time `json:"deleted_at"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"default:now();"`
}

// UserWithShop ..
type UserWithShop struct {
	ID        int        `json:"user_id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Shops     []Shop     `json:"shops"`
	CreatedAt *time.Time `json:"created_at" gorm:"default:now();"`
	DeletedAt *time.Time `json:"deleted_at"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"default:now();"`
}

// UserUsecase ..
type UserUsecase interface {
	ListUser() ([]User, error)
	GetUser(userID int) (User, error)
	GetUserWithShop(userID int) (UserWithShop, error)
	CreateUser(User) (User, error)
	UpdateUser(*User) error
	DeleteUser(userID int) error
	UndeleteUser(userID int) error
}

// UserRepository ..
type UserRepository interface {
	List() ([]User, error)
	Get(userID int) (User, error)
	GetWithShop(userID int) (UserWithShop, error)
	Create(User) (User, error)
	Update(*User) error
	Delete(userID int) error
	Undelete(userID int) error
}
