package db

import (
	"broadcastle.co/code/lakuh/pkg/utils"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// DB is publicly accessable.
var DB *gorm.DB

// Init starts the database.
func Init() {

	var err error

	location, err := utils.FullPath(viper.GetString("lakuh.database"))
	if err != nil {
		logrus.Fatal(err)
	}

	DB, err = sqlite(location)
	if err != nil {
		logrus.Fatal(err)
	}

	DB.AutoMigrate(&Song{})
	DB.AutoMigrate(&Artist{})
	DB.AutoMigrate(&Album{})
	DB.AutoMigrate(&User{})
}

// Close the database.
func Close() {
	DB.Close()
}

// AllSongInfo will return all of the information of a song given an ID.
func AllSongInfo(songID int) (song Song, artist Artist, album Album, err error) {

	song.ID = uint(songID)
	err = song.Find()
	if err != nil {
		return
	}

	artist.ID = song.ArtistID
	err = artist.Find()
	if err != nil {
		return
	}

	album.ID = song.AlbumID
	err = album.Find()
	if err != nil {
		return
	}

	return
}
