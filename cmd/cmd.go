package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile          string
	appPort          string
	repoUrl          string
	repoBranch       string
	repoClonePath    string
	repoContentDir   string
	contentDelimiter string

	gostapagenCmd = &cobra.Command{
		Use:           "gostapagen",
		Short:         "gostapagen – static page generator service",
		Long:          ``,
		Version:       "0.0.1",
		SilenceErrors: true,
		SilenceUsage:  true,
	}
)

func Execute() error {
	return gostapagenCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	gostapagenCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "Path to config file. Default is '$HOME/.gostapagen'.")
	gostapagenCmd.PersistentFlags().StringVarP(&appPort, "appPort", "p", "8080", "The port being used for the API. Default port is 8080.")
	gostapagenCmd.PersistentFlags().StringVarP(&repoUrl, "repoUrl", "r", "", "The git repository url to clone from. Fully qualified without .git extension.l")
	gostapagenCmd.PersistentFlags().StringVarP(&repoBranch, "repoBranch", "b", "main", "The git repository branch to clone from. Default branch is 'main'.")
	gostapagenCmd.PersistentFlags().StringVarP(&repoClonePath, "repoClonePath", "", "/tmp/gostapagen", "The absolute path to clone the git repository to. Default path is '/tmp'.")
	gostapagenCmd.PersistentFlags().StringVarP(&repoContentDir, "repoContentDir", "", "/content", "The directory to the content files within the git repository project. Default directory is '/content'.")
	gostapagenCmd.PersistentFlags().StringVarP(&contentDelimiter, "contentDelimiter", "", "content-header", "Content front matter tag delimiter (default is content-header)")
	viper.BindPFlag("config", gostapagenCmd.PersistentFlags().Lookup("config"))
	viper.BindPFlag("appPort", gostapagenCmd.PersistentFlags().Lookup("appPort"))
	viper.BindPFlag("repoUrl", gostapagenCmd.PersistentFlags().Lookup("repoUrl"))
	viper.BindPFlag("repoBranch", gostapagenCmd.PersistentFlags().Lookup("repoBranch"))
	viper.BindPFlag("repoClonePath", gostapagenCmd.PersistentFlags().Lookup("repoClonePath"))
	viper.BindPFlag("repoContentDir", gostapagenCmd.PersistentFlags().Lookup("repoContentDir"))
	viper.BindPFlag("contentDelimiter", gostapagenCmd.PersistentFlags().Lookup("contentDelimiter"))

	gostapagenCmd.AddCommand(serveCmd)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("/etc/gostapagen/")
		viper.AddConfigPath("$HOME/.gostapagen")
		viper.AddConfigPath(".")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Printf("Config file used for gostapagen: %s", viper.ConfigFileUsed())
	}
}
