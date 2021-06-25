package repo

import (
	"b-nova-openhub/stapagen/pkg/config"
	"b-nova-openhub/stapagen/pkg/file"
	"io/ioutil"
	"log"
)

func RepoContents() []string {
	contentFiles := make([]string, 0)
	GetGitRepository(config.AppConfig.TargetAbsoluteClonePath)

	files, mdErr := file.GetAllMdFilesInPath()

	if mdErr != nil {
		log.Fatalln("Error during markdown files parsing.")
	}

	for _, f := range files {
		readFile, _ := ioutil.ReadFile(f)
		contentFiles = append(contentFiles, string(readFile))
	}

	return contentFiles
}
