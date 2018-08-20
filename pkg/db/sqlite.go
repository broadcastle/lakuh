package db

import "github.com/jinzhu/gorm"

import _ "github.com/jinzhu/gorm/dialects/sqlite" // Use sqlite3

func sqlite(file string) (*gorm.DB, error) {
	return gorm.Open("sqlite3", file)
}
