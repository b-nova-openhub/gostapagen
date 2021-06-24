package repo

import "b-nova-openhub/stapagen/pkg/file"

func GetGitRepository(path string) (error, error) {
	pathExists, err := file.PathExists(path)
	if !pathExists {
		return err, Clone(false, path)
	}
	return err, nil
}
