package library

import (
	"fmt"
	"strconv"

	"broadcastle.co/code/lakuh/pkg/db"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// CobraView presents songs to the CLI.
func CobraView(cmd *cobra.Command, args []string) {

	db.Init()
	defer db.Close()

	for x := range args {

		id, err := strconv.Atoi(args[x])
		if err != nil {
			logrus.Warn(err)
			break
		}

		audio := Audio{ID: id}
		if err := audio.View(); err != nil {
			logrus.Warn(err)
			break
		}

		fmt.Printf("Title: %s\nArtist: %s\nAlbum: %s\nGenre: %s\nYear: %v\n", audio.Title, audio.Artist, audio.Album, audio.Genre, audio.Year)

	}

}

// View audio files.
func (a *Audio) View() error {

	song, artist, album, err := db.AllSongInfo(a.ID)
	if err != nil {
		return err
	}

	a.Title = song.Title
	a.Album = album.Title
	a.Artist = artist.Name
	a.Genre = song.Genre
	a.Year = album.Year

	return nil
}

// AllAudio return multiple audio files.
func AllAudio() ([]Audio, error) {

	all := []Audio{}

	songs, err := db.AllSongs()
	if err != nil {
		return all, err
	}

	for _, x := range songs {
		song, artist, album, err := db.AllSongInfo(int(x.ID))
		if err != nil {
			return all, err
		}

		single := Audio{
			ID:     int(song.ID),
			Title:  song.Title,
			Album:  album.Title,
			Artist: artist.Name,
			Year:   album.Year,
			Genre:  song.Genre,
		}

		all = append(all, single)
	}

	return all, nil

}
