// Define a book model
package models

import (
	"github.com/jinzhu/gorm"
)

// Book model
type Book struct {
	gorm.Model
	Title  string
	Author string
	Rating int
}
