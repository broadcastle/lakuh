package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"broadcastle.co/code/lakuh/pkg/utils"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	debug   bool
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "lakuh",
	Short: "Distributed radio automation.",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {

	if debug {
		logrus.SetLevel(logrus.DebugLevel)
		logrus.Info("debug text enabled")
	}

	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.lakuh.toml)")
	RootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "show debug messages")

	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			// fmt.Println(err)
			// os.Exit(1)
			logrus.Fatal(err)
		}

		// Search config in home directory with name ".lakuh" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".lakuh")
	}

	viper.AutomaticEnv() // read in environment variables that match

	if err := viper.ReadInConfig(); err != nil {

		logrus.Debug(err)

		if e := generateDefaults(); e != nil {
			logrus.Fatal(e)
		}

	}

	defaults()
}

func defaults() {

	storage, err := utils.FullPath(viper.GetString("audio.storage"))
	if err != nil {
		logrus.Fatal(err)
	}

	if err := os.MkdirAll(storage, os.ModePerm); err != nil {
		logrus.Fatal(err)
	}

	database, err := utils.FullPath(viper.GetString("lakuh.database"))
	if err != nil {
		logrus.Fatal(err)
	}

	database = filepath.Dir(database)

	if err := os.MkdirAll(database, os.ModePerm); err != nil {
		logrus.Fatal(err)
	}

}
