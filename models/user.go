package models

import (
	"github.com/jinzhu/gorm"
)

// User @doc http://gorm.io/docs/models.html
type User struct {
	gorm.Model
	Name    string
	Email   string
	Role    string
	Address string
}
