package cmd

import (
	"github.com/b-nova-openhub/gostapagen/pkg/config"
	"github.com/b-nova-openhub/gostapagen/pkg/rest"
	"github.com/spf13/cobra"
)

var (
	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "Serve rest endpoints",
		Long:  ``,
		Run:   serve,
	}
)

func serve(ccmd *cobra.Command, args []string) {
	config.PersistConfig()
	rest.HandleRequests()
}
