package data

// Tanpa GORM
type User struct {
	ID       string `json:"id" validate:"required,max=100"`
	NIM      string `json:"nim" validate:"required,max=100"`
	Email    string `json:"email" validate:"required,max=100"`
	Password string `json:"password" validate:"required,max=100"`
}
