package repo

import (
	"github.com/b-nova-openhub/gostapagen/pkg/file"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
)

func RepoContents() []string {
	GetGitRepository(viper.GetString("repoClonePath"))
	files, mdErr := file.GetAllMdFilesInPath()
	if mdErr != nil {
		log.Fatalln("Error during markdown files parsing.")
	}

	contentFiles := make([]string, 0)
	for _, f := range files {
		readFile, _ := ioutil.ReadFile(f)
		contentFiles = append(contentFiles, string(readFile))
	}
	return contentFiles
}
