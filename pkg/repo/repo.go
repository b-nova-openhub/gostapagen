package repo

import (
	"b-nova-openhub/stapagen/pkg/file"
	"b-nova-openhub/stapagen/pkg/util"
	"io/ioutil"
	"log"
)

func RepoContents() []string {
	contentFiles := make([]string, 0)
	err2, err3 := GetGitRepository("/Users/rschneider/tmp/")
	if err2 == nil && err3 == nil {
		files, err := file.GetAllMdFilesInPath()
		for _, f := range files {
			readFile, _ := ioutil.ReadFile(f)
			contentFiles = append(contentFiles, string(readFile))
		}
		util.HandleError(err)

	} else {
		log.Fatalln(err2)
		log.Fatalln(err3)
	}

	return contentFiles
}
