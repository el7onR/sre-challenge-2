package users

import (
	"backend/internal/config"
	"time"
)

//User struct
type User struct {
	ID        int64     `json:"id,string,omitempty"`
	Username  string    `json:"username" validate:"required" gorm:"uniqueIndex:idx_username,length:20"`
	FullName  string    `json:"full_name" validate:"required" gorm:"not null"`
	Email     string    `json:"email" validate:"required,email" gorm:"not null"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

const (
	errorQueryMessage   = "Error while querying database, details %s"
	userNotFoundMessage = "User with id &d not found"
)

//List query users on database
func List(env *config.Env) (*[]User, error) {
	users := []User{}
	result := env.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return &users, nil
}

//ByID query an User by its ID
func ByID(env *config.Env, id *int64) (*User, error) {
	u := User{}
	result := env.DB.Take(&u, &id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &u, nil
}

//ByUsername query an User by its ID
func ByUsername(env *config.Env, username *string) (*User, error) {
	u := User{}
	result := env.DB.Where("username = ?", &username).Take(&u)
	if result.Error != nil {
		return nil, result.Error
	}
	return &u, nil
}

//New creates an new User
func New(env *config.Env, u *User) (*User, error) {
	result := env.DB.Omit("id").Create(&u)
	if result.Error != nil {
		return nil, result.Error
	}
	return u, nil
}

//Update updates an User by its ID
func Update(env *config.Env, u *User, id *int64) (*User, error) {
	forUpdate := User{}
	result := env.DB.First(&forUpdate, &id).Updates(User{FullName: u.FullName, Email: u.Email, Username: u.Username})
	if result.Error != nil {
		return nil, result.Error
	}
	return &forUpdate, nil
}

//Delete deletes an User by its ID
func Delete(env *config.Env, id *int64) error {
	result := env.DB.Delete(User{}, &id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
