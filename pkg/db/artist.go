package db

import (
	"errors"
	"fmt"

	"github.com/gosimple/slug"
	"github.com/jinzhu/gorm"
)

// Artist for the music
type Artist struct {
	gorm.Model

	Name string
	Slug string

	Songs  []Song
	Albums []Album
}

// Create a song.
func (a *Artist) Create() error {

	a.Slug = slug.Make(a.Name)

	check := Artist{Slug: a.Slug}

	if err := check.Find(); err == nil {
		a = &check
		return nil
	}

	return DB.Create(&a).Error

}

// Update a song.
func (a *Artist) Update() error {

	a.Slug = slug.Make(a.Name)

	return DB.Model(&a).Update(&a).Error

}

// Delete a song.
func (a *Artist) Delete() error {

	if a.ID == 0 {
		return errors.New("id is required to delete artist")
	}

	//// DELETE SONGS

	songs := []Song{}

	// Find all the songs of this artist.
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

	//// Delete Albums

	albums := []Album{}

	// Find all the albums from this artist.
	if err := DB.Model(&a).Related(&albums).Error; err != nil {
		return err
	}

	// Delete all songs.
	for _, album := range albums {
		if err := album.Delete(); err != nil {
			fmt.Println("unable to delete %x", album.ID)
			break
		}
	}

	return DB.Delete(&a).Error
}

// Find a single artist that matches a.
func (a *Artist) Find() error {
	return DB.Where(&a).First(&a).Error
}

// FindAll artists that match a.
func (a *Artist) FindAll(v interface{}) error {
	return DB.Where(&a).Find(&v).Error
}
