package repo

import (
	"b-nova-openhub/stapagen/pkg/config"
	"b-nova-openhub/stapagen/pkg/file"
)

func GetGitRepository(path string) (error, error) {
	projectPath := config.AppConfig.TargetAbsoluteProjectPath
	pathExists, err := file.PathExists(projectPath)
	if !pathExists {
		return err, Clone(false, projectPath)
	}
	return err, nil
}
