package main

import (
	"fmt"
	"github.com/b-nova-openhub/gostapagen/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil && err.Error() != "" {
		fmt.Println(err)
	}
}
