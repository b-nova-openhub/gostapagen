package repo

import (
	"b-nova-openhub/stapagen/pkg/config"
	"b-nova-openhub/stapagen/pkg/file"
	"log"
)

func GetGitRepository(path string) {
	projectPath := config.AppConfig.TargetAbsoluteProjectPath
	pathExists, pathErr := file.PathExists(projectPath)

	if pathErr != nil {
		log.Fatalln("Error before git clone as path already seems to exist: %+x\n", projectPath)
	}

	if !pathExists {
		cloneErr := Clone(false, projectPath)
		if cloneErr != nil {
			log.Fatalln("Error during git clone. Path: %+x\n", projectPath)
		}
	}
}
