package library

import (
	"strconv"

	"broadcastle.co/code/lakuh/pkg/db"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// CobraRemove removes audio through the CLI.
func CobraRemove(cmd *cobra.Command, args []string) {

	db.Init()
	defer db.Close()

	for x := range args {

		i, err := strconv.Atoi(args[x])
		if err != nil {
			logrus.Warn(err)
			break
		}

		single := Audio{ID: i}

		if err := single.Remove(); err != nil {
			logrus.Warn(err)
		}

	}

}

// Remove audio
func (a Audio) Remove() error {

	song := db.Song{}
	song.ID = uint(a.ID)

	return song.Delete()
}
