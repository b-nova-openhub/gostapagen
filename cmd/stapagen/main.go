package main

import (
	"b-nova-openhub/stapagen/pkg/config"
	"b-nova-openhub/stapagen/pkg/rest"
)

func main() {
	config.PersistsFlags()
	rest.HandleRequests()
}
