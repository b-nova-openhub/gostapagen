package file

import (
	"github.com/b-nova-openhub/gostapagen/pkg/url"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func GetAllMdFilesInPath() ([]string, error) {
	contentPath := viper.GetString("repoClonePath") + "/" + url.GetSlug(viper.GetString("repoUrl")) + viper.GetString("repoContentDir")
	return walkMatch(contentPath, "*.md")
}

func walkMatch(root, pattern string) ([]string, error) {
	var matches []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if matched, err := filepath.Match(pattern, filepath.Base(path)); err != nil {
			return err
		} else if matched {
			matches = append(matches, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return matches, nil
}
