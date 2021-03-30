package rest

import "time"

// RegisterUserModel ...
type RegisterUserModel struct {
	ID           string     `json:"id" swaggerignore:"true"`
	ClientTypeID string     `json:"client_type_id"`
	RoleID       string     `json:"role_id"`
	Password     string     `json:"password"`
	Active       int8       `json:"active" swaggerignore:"true"`
	ExpiresAt    *time.Time `json:"expires_at"`
	Phones       []string   `json:"phones"`
	Email        string     `json:"email"`
	Login        string     `json:"login"`
}

// UpdateUserModel ...
type UpdateUserModel struct {
	ID           string     `json:"id"`
	ClientTypeID string     `json:"client_type_id"`
	RoleID       string     `json:"role_id"`
	Active       int8       `json:"active" swaggerignore:"true"`
	ExpiresAt    *time.Time `json:"expires_at"`
	CreatedAt    *time.Time `json:"created_at" swaggerignore:"true"`
	UpdatedAt    *time.Time `json:"updated_at" swaggerignore:"true"`
	Phones       []string   `json:"phones"`
	Email        string     `json:"email"`
	Login        string     `json:"login"`
}

// UpdateUserPasswordModel ...
type UpdateUserPasswordModel struct {
	UserID   string `json:"user_id"`
	Password string `json:"password"`
}

// UpdateUserPhoneModel ...
type UpdateUserPhoneModel struct {
	UserID string   `json:"user_id"`
	Phones []string `json:"phones"`
}

// UpdateUserEmailModel ...
type UpdateUserEmailModel struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
}

// UpdateUserLoginModel ...
type UpdateUserLoginModel struct {
	UserID string `json:"user_id"`
	Login  string `json:"login"`
}

// UserModel ...
type UserModel struct {
	ID           string     `json:"id"`
	ClientTypeID string     `json:"client_type_id"`
	RoleID       string     `json:"role_id"`
	Password     string     `json:"password"`
	Active       int8       `json:"active"`
	ExpiresAt    *time.Time `json:"expires_at"`
	CreatedAt    *time.Time `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
}

// UserPhoneModel ...
type UserPhoneModel struct {
	Phone     string     `json:"phone"`
	UserID    string     `json:"user_id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

// UserEmailModel ...
type UserEmailModel struct {
	Email     string     `json:"email"`
	UserID    string     `json:"user_id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

// UserLoginModel ...
type UserLoginModel struct {
	Email     string     `json:"email"`
	UserID    string     `json:"user_id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

// UserKeyModel ...
type UserKeyModel struct {
	Key       string     `json:"key"`
	UserID    string     `json:"user_id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
