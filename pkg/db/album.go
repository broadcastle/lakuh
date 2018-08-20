package db

import (
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
)

// Album has songs.
type Album struct {
	gorm.Model

	Title string

	Songs []Song

	Year int

	// Artist   Artist
	ArtistID uint
}

// Create a album.
func (a *Album) Create() error {

	if a.ArtistID == 0 {
		return errors.New("need artist id")
	}

	return DB.Create(&a).Error
}

// Update a album.
func (a *Album) Update() error {
	return DB.Model(&a).Update(&a).Error
}

// Delete a album.
func (a *Album) Delete() error {

	if a.ID == 0 {
		return errors.New("id is required to delete album")
	}

	songs := []Song{}

	// Find all the songs of this album.
	if err := DB.Model(&a).Related(&songs).Error; err != nil {
		return err
	}

	// Delete all songs.
	for _, song := range songs {
		if err := song.Delete(); err != nil {
			fmt.Println("unable to delete %x", song.ID)
			fmt.Println(err)
			break
		}
	}

	return DB.Delete(&a).Error
}

// Find a single album that matches a.
func (a *Album) Find() error {
	return DB.Where(&a).First(&a).Error
}

// FindAll albums that match a.
func (a *Album) FindAll(v interface{}) error {
	return DB.Where(&a).Find(&v).Error
}
