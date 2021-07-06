package main

import (
	"github.com/b-nova-openhub/gostapagen/pkg/config"
	"github.com/b-nova-openhub/gostapagen/pkg/rest"
)

func main() {
	config.PersistConfig()
	rest.HandleRequests()
}
