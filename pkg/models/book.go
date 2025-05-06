package models

import (
	"Book-Management/pkg/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"type:varchar(255);not null" json:"name"`
	Author      string `gorm:"type:varchar(255);not null" json:"author"`
	Publication string `gorm:"type:varchar(255);not null" json:"publication"`
}

// type City struct {
// 	ID          uint `gorm:"primaryKey"`
// 	Name        string
// 	CountryCode string
// }

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).First(&getBook)
	return &getBook, db
}

// func DeleteBook(ID int64) Book {
// 	var book Book
// 	db.Where("ID=?", ID).Delete(&Book{})
// 	return book
// }

func DeleteBook(ID int64) error {
	result := db.Unscoped().Where("id = ?", ID).Delete(&Book{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// func GetCities() []string {
// 	var cities []string
// 	db := config.GetDB() // ✅ Ensure we're using the initialized DB

// 		log.Fatal("Database connection is nil. Ensure Connect() has been called before using DB.")
// 	}
// 	db.Table("city").Where("CountryCode = ?", "NLD").Pluck("Name", &cities)
// 	return cities
// }
