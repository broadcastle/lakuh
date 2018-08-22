package cmd

import (
	"html/template"
	"os"
	"path"
	"time"

	"broadcastle.co/code/lakuh/pkg/utils"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Generate a default config file.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if err := generateDefaults(); err != nil {
			logrus.Fatal(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(initCmd)
}

func generateDefaults() error {

	home, err := homedir.Dir()
	if err != nil {
		return err
	}

	f := path.Join(home, ".lakuh.toml")

	file, err := os.Create(f)
	if err != nil {
		return err
	}

	defer file.Close()

	tmp, err := template.New("lakuhDefaults").Parse(lakuhDefaults)
	if err != nil {
		return err
	}

	type info struct {
		Token string
		DB    string
		Files string
	}

	data := &info{
		Token: utils.SHA256string(time.Now().String()),
		DB:    path.Join(home, ".lakuh", "db", "main.db"),
		Files: path.Join(home, ".lakuh", "audio"),
	}

	if err := tmp.Execute(file, data); err != nil {
		return err
	}

	return viper.ReadInConfig()

}

var lakuhDefaults = `[lakuh]
database = "{{.DB}}"
token = "{{.Token}}"
port = 8080

[audio]
storage = "{{.Files}}"
`
