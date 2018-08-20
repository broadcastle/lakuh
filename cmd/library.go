package cmd

import (
	"broadcastle.co/code/lakuh/pkg/library"
	"github.com/spf13/cobra"
)

// libraryCmd represents the library command
var libraryCmd = &cobra.Command{
	Use:   "library",
	Short: "Work on the lakuh library.",
}

var libraryAdd = &cobra.Command{
	Use:   "add",
	Short: "Add a audio file to lakuh.",
	Long: `Add a single audio file to lakuh's system. 

Use the flags to overwrite the tags 
of the audio file.`,
	Args: cobra.ExactArgs(1),
	Run:  library.CobraImport,
}

var libraryRemove = &cobra.Command{
	Use:   "remove",
	Short: "Remove a audio entry from lakuh.",
	Long: `Remove all audio entries that match the 
flags from lakuh.`,
	Run: library.CobraRemove,
}

var libraryEdit = &cobra.Command{
	Use:   "edit",
	Short: "Edit a audio entry.",
	Long: `Edit a single audio entry. Use flags
to change the corresponding tag.`,
	Args: cobra.ExactArgs(1),
	Run:  library.CobraEdit,
}

var libraryView = &cobra.Command{
	Use:   "view",
	Short: "View song(s) from the lakuh library.",
	Run:   library.CobraView,
}

func init() {
	RootCmd.AddCommand(libraryCmd)

	libraryCmd.AddCommand(libraryAdd)
	libraryCmd.AddCommand(libraryRemove)
	libraryCmd.AddCommand(libraryEdit)
	libraryCmd.AddCommand(libraryView)

	libraryFlags(libraryAdd)
	libraryFlags(libraryEdit)

}

func libraryFlags(libraryCmd *cobra.Command) {

	libraryCmd.PersistentFlags().StringP("artist", "a", "", "name of the artist")
	libraryCmd.PersistentFlags().StringP("album", "b", "", "title of the album")
	libraryCmd.PersistentFlags().StringP("title", "t", "", "title of the song")
	libraryCmd.PersistentFlags().StringP("genre", "g", "", "genre of the song")
	libraryCmd.PersistentFlags().IntP("year", "y", 0, "year the song was released")

}
