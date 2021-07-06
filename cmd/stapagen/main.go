package main

import (
	"github.com/b-nova-openhub/stapagen/pkg/config"
	"github.com/b-nova-openhub/stapagen/pkg/rest"
)

func main() {
	config.PersistConfig()
	rest.HandleRequests()
}
