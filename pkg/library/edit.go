package library

import (
	"strconv"

	"broadcastle.co/code/lakuh/pkg/db"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// CobraEdit removes audio through the CLI.
func CobraEdit(cmd *cobra.Command, args []string) {

	db.Init()
	defer db.Close()

	audio, err := cflag(cmd)
	if err != nil {
		logrus.Fatal(err)
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		logrus.Fatal(err)
	}

	audio.ID = id

	if err := audio.Edit(); err != nil {
		logrus.Fatal(err)
	}

}

// Edit audio
func (a Audio) Edit() error {

	// Get the song and all related info.
	song, artist, album, err := db.AllSongInfo(a.ID)
	if err != nil {
		return err
	}

	if a.Year != 0 {
		song.Year = a.Year
	}

	// Update the title
	if a.Title != "" {
		song.Title = a.Title
	}

	// Update the artist
	if a.Artist != "" && a.Artist != artist.Name {

		new := db.Artist{Name: a.Artist}

		if err := new.Create(); err != nil {
			return err
		}

		song.ArtistID = new.ID

	}

	if a.Album != "" || song.ArtistID != album.ArtistID {

		new := db.Album{Title: a.Album, ArtistID: song.ArtistID, Year: song.Year}

		if a.Album == "" {
			new.Title = album.Title
		}

		if err := new.Create(); err != nil {
			return err
		}

		song.AlbumID = new.ID

	}

	return song.Update()

}
