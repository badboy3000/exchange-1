package models

import (
	"database/sql"
	"time"

	"github.com/jinzhu/gorm"
)

// User @doc http://gorm.io/docs/models.html
type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(100);unique_index"`
	Age      sql.NullInt64
	Birthday *time.Time
	Email    string `gorm:"type:varchar(100);unique_index"`
	Role     string `gorm:"size:255"`   // set field size to 255
	Address  string `gorm:"index:addr"` // create index with name `addr` for address
}
