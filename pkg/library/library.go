package library

import "github.com/spf13/cobra"

// Audio has the input information.
type Audio struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Album  string `json:"album"`
	Genre  string `json:"genre"`
	Year   int    `json:"year"`
}

// cflag returns a Audio struct or error from cobra flags.
func cflag(cmd *cobra.Command) (result Audio, err error) {

	result.Title, err = cmd.Flags().GetString("title")
	if err != nil {
		return
	}

	result.Artist, err = cmd.Flags().GetString("artist")
	if err != nil {
		return
	}

	result.Album, err = cmd.Flags().GetString("album")
	if err != nil {
		return
	}

	result.Genre, err = cmd.Flags().GetString("genre")
	if err != nil {
		return
	}

	result.Year, err = cmd.Flags().GetInt("year")

	return

}
