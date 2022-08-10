package models

// Create package models, import "time",
import "time"

// User struct, UsersProfileResponse  struct,
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name" gorm:"type: varchar(255)"`
	Email     string    `json:"email" gorm:"type: varchar(255)"`
	Password  string    `json:"-" gorm:"type: varchar(255)"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

// and UsersProfileResponse TableName method here ...
type UsersProfileResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (UsersProfileResponse) TableName() string {
	return "users"
}
