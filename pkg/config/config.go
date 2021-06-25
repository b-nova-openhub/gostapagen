package config

import (
	"b-nova-openhub/stapagen/pkg/url"
	"b-nova-openhub/stapagen/pkg/util"
	"flag"
	"fmt"
)

type Config struct {
	AppPort                   string
	ContentDelimiterTag       string
	SourceGitRepositoryUrl    string
	SourceGitRepositoryBranch string
	TargetAbsoluteClonePath   string
	TargetAbsoluteProjectPath string
	TargetAbsoluteContentPath string
}

var AppConfig *Config

func PersistConfig() {
	port := flag.String("port", "8080", "The port being used for the API. Default port is 8080.")
	repo := flag.String("repo", "", "The git repository url to clone from. Fully qualified without .git extension.")
	branch := flag.String("branch", "main", "The git repository branch to clone from. Default branch is 'main'.")
	absolutePath := flag.String("clonePath", "/tmp", "The absolute path to clone the git repository to. Default path is '/tmp'.")
	relativePath := flag.String("contentDir", "/content", "The directory to the content files within the git repository project. Default directory is '/content'.")
	delimiter := flag.String("delimiter", "content-header", "The directory to the content files within the git repository project. Default directory is '/content'.")
	flag.Parse()

	AppConfig = new(Config)
	AppConfig.AppPort = util.DerefString(port)
	AppConfig.ContentDelimiterTag = util.DerefString(delimiter)
	AppConfig.SourceGitRepositoryUrl = util.DerefString(repo)
	AppConfig.SourceGitRepositoryBranch = util.DerefString(branch)
	AppConfig.TargetAbsoluteClonePath = util.DerefString(absolutePath)
	AppConfig.TargetAbsoluteProjectPath = util.DerefString(absolutePath) + "/" + url.GetSlug(util.DerefString(repo))
	AppConfig.TargetAbsoluteContentPath = util.DerefString(absolutePath) + "/" + url.GetSlug(util.DerefString(repo)) + util.DerefString(relativePath)

	fmt.Printf("AppConfig: %+v\n", AppConfig)
}
