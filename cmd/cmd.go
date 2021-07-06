package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var (
	cfgFile string

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

	gostapagenCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.jamctl.yaml)")
	viper.BindPFlag("config", gostapagenCmd.PersistentFlags().Lookup("config"))

	gostapagenCmd.AddCommand(serveCmd)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		viper.SetConfigFile(".gostapagen")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Config file used for gostapagen: ", viper.ConfigFileUsed())
	}
}
