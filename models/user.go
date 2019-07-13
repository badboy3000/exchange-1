package models

// User @doc http://gorm.io/docs/models.html
type User struct {
	BaseModel
	Name    string
	Email   string
	Role    string
	Address string
}
