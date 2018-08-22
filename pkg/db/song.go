package db

import (
	"errors"
	"os"
	"time"

	"broadcastle.co/code/lakuh/pkg/utils"
	"github.com/jinzhu/gorm"
)

// Song holds the basic
type Song struct {
	gorm.Model

	Title    string
	ArtistID uint
	AlbumID  uint
	SHA256   string
	Year     int
	Length   string
	Location string
	MIME     string
	Genre    string

	Tags []Tag `gorm:"many2many:song_tags"`
}

// Create a song.
func (s *Song) Create() error {

	if s.ArtistID == 0 {
		return errors.New("need artist id for song")
	}

	if s.AlbumID == 0 {
		album := Album{
			ArtistID: s.ArtistID,
			Title:    "Unknown Album",
			Year:     time.Now().Year(),
		}

		if err := album.Create(); err != nil {
			return err
		}

		s.AlbumID = album.ID
	}

	if s.SHA256 == "" {

		hash, err := utils.SHA256(s.Location)
		if err != nil {
			return err
		}

		s.SHA256 = hash

	}

	if s.MIME == "" {

		mime, err := utils.AudioCheck(s.Location)
		if err != nil {
			return err
		}

		s.MIME = mime

	}

	return DB.Create(&s).Error
}

// Update a song.
func (s *Song) Update() error {
	return DB.Model(&s).Update(&s).Error
}

// Delete a song.
func (s *Song) Delete() error {

	if err := s.Find(); err != nil {
		return err
	}

	os.Remove(s.Location)

	return DB.Delete(&s).Error
}

// Find a single song that matches s.
func (s *Song) Find() error {
	return DB.Where(&s).First(&s).Error
}

// FindAll songs that match s.
func (s *Song) FindAll(v interface{}) error {
	return DB.Where(&s).Find(&v).Error
}

// AllSongs returns all songs in the database.
func AllSongs() ([]Song, error) {

	song := []Song{}

	err := DB.Find(&song).Error
	return song, err

}
