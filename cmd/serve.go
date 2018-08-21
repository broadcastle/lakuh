package cmd

import (
	"broadcastle.co/code/lakuh/pkg/web"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the web interface.",
	Run: func(cmd *cobra.Command, args []string) {

		web.Web()
	},
}

func init() {
	RootCmd.AddCommand(serveCmd)

	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
