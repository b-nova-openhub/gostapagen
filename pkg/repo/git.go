package repo

import (
	"github.com/b-nova-openhub/gostapagen/pkg/file"
	"github.com/b-nova-openhub/gostapagen/pkg/url"
	"github.com/spf13/viper"
	"log"
)

func GetGitRepository(path string) {
	projectPath := viper.GetString("repoClonePath") + "/" + url.GetSlug(viper.GetString("repoUrl"))
	pathExists, pathErr := file.PathExists(projectPath)

	if pathErr != nil {
		log.Fatalf("Error before git clone as path already seems to exist: %+x\n", projectPath)
	}

	if !pathExists {
		cloneErr := Clone(false, projectPath)
		if cloneErr != nil {
			log.Fatalf("Error during git clone. Path: %+x\n", projectPath)
		}
	}
}
