package cmd

import (
	"broadcastle.co/code/lakuh/pkg/library"
	"github.com/spf13/cobra"
)

var playerCmd = &cobra.Command{
	Use:   "player",
	Short: "Play music",
}

var playerStart = &cobra.Command{
	Use:   "play",
	Short: "Start playing music.",
	Run:   library.CobraPlay,
}

func init() {
	RootCmd.AddCommand(playerCmd)

	playerCmd.AddCommand(playerStart)
}
