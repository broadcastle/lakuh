package library

import (
	"errors"
	"os"
	"strconv"
	"time"

	"broadcastle.co/code/lakuh/pkg/db"
	"github.com/faiface/beep"
	"github.com/faiface/beep/flac"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/vorbis"
	"github.com/faiface/beep/wav"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// CobraPlay plays music from the command line.
func CobraPlay(cmd *cobra.Command, args []string) {

	db.Init()
	defer db.Close()

	id, err := strconv.Atoi(args[0])
	if err != nil {
		logrus.Fatal(err)
	}

	audio := Audio{ID: id}

	if err := audio.Play(); err != nil {
		logrus.Fatal(err)
	}

}

// Play a audio file.
func (a Audio) Play() error {

	if a.ID == 0 {
		return errors.New("need a id in order to play audio")
	}

	audio, artist, _, err := db.AllSongInfo(a.ID)
	if err != nil {
		return err
	}

	logrus.Infof("playing %s by %s", audio.Title, artist.Name)

	file, err := os.Open(audio.Location)
	if err != nil {
		return err
	}

	var (
		stream beep.StreamSeekCloser
		format beep.Format
	)

	switch audio.MIME {
	case "audio/mpeg":
		stream, format, err = mp3.Decode(file)
	case "audio/x-wav":
		stream, format, err = wav.Decode(file)
	case "audio/x-flac":
		stream, format, err = flac.Decode(file)
	case "audio/ogg":
		stream, format, err = vorbis.Decode(file)
	default:
		return errors.New("MIME is missing")
	}

	if err := speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10)); err != nil {
		return err
	}

	done := make(chan struct{})

	speaker.Play(beep.Seq(stream, beep.Callback(func() {
		close(done)
	})))

	<-done

	// if err := speaker.Play(stream); err != nil {
	// 	return err
	// }

	// s, format, err := wav.Decode(file)
	// if err != nil {
	// 	return err
	// }

	return nil

}
