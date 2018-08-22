package library

import (
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"

	"broadcastle.co/code/lakuh/pkg/db"
	"broadcastle.co/code/lakuh/pkg/utils"
	"github.com/dhowden/tag"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// CobraImport takes commands from the CLI.
func CobraImport(cmd *cobra.Command, args []string) {

	db.Init()
	defer db.Close()

	location := args[0]

	if location == "" {
		logrus.Fatal("need a valid path")
	}

	audio, err := cflag(cmd)
	if err != nil {
		logrus.Fatal(err)
	}

	if err := audio.Import(location); err != nil {
		logrus.Fatal(err)
	}

}

// Import Audio
func (a Audio) Import(src string) error {

	// Import check
	if src == "" {
		return errors.New("library.Audio.Import: need a file")
	}

	// Get MIME info
	mime, err := utils.AudioCheck(src)
	if err != nil {
		return err
	}

	data, err := os.Open(src)
	if err != nil {
		return err
	}

	tag, err := tag.ReadFrom(data)
	if err != nil {
		return err
	}

	if a.Title == "" {
		a.Title = tag.Title()
	}

	if a.Artist == "" {
		a.Artist = tag.Artist()
	}

	if a.Album == "" {
		a.Album = tag.Album()
	}

	if a.Year == 0 {
		a.Year = tag.Year()
	}

	if a.Genre == "" {
		a.Genre = tag.Genre()
	}

	song := db.Song{
		Title: a.Title,
		MIME:  mime,
		Genre: a.Genre,
		Year:  a.Year,
	}

	// Get storage location.
	storage, err := utils.FullPath(viper.GetString("audio.storage"))
	if err != nil {
		return err
	}

	// Get the extension of the source file.
	ext := filepath.Ext(filepath.Base(src))

	// Create SHA256 hash.
	song.SHA256, err = utils.SHA256(src)
	if err != nil {
		return err
	}

	song.Location = path.Join(storage, song.SHA256+ext)

	old := db.Song{SHA256: song.SHA256}

	if err := old.Find(); err == nil {
		fmt.Printf("this file already exists under id: %v\n", old.ID)
		return nil
	}

	// Copy file to destination.
	if err := utils.CopyFile(src, song.Location); err != nil {
		return err
	}

	logrus.Info("importing ", a)

	// CREATE - GET ARTIST
	artist := db.Artist{Name: a.Artist}

	// Find Artist or else create a new one
	if err := artist.Find(); err != nil {
		if err := artist.Create(); err != nil {
			return err
		}
	}

	song.ArtistID = artist.ID

	// CREATE - GET ALBUM

	album := db.Album{Title: a.Album, ArtistID: song.ArtistID}

	if err := album.Find(); err != nil {

		album.Year = a.Year

		if err := album.Create(); err != nil {
			return err
		}
	}

	song.AlbumID = album.ID

	return song.Create()

}
