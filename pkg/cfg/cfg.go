package cfg

import (
	"fmt"
	"go.uber.org/config"
	"strings"
)

type appConfig struct {
	source_git_repository_url           string
	source_git_repository_branch        string
	source_git_repository_username      string
	source_git_repository_password      string
	target_git_repository_absolute_path string
	target_git_repository_relative_path string
	content_permalink_base              string
	content_frontmatter_open            string
	content_frontmatter_close           string
}

func ProvideAppConfig() {
	readConfig()
}

func getDefaultConfigYaml() config.YAMLOption {
	return config.Source(strings.NewReader(""))
}

func getAppConfigYaml() config.YAMLOption {
	return config.Source(strings.NewReader(""))
}

func readConfig() {
	yaml, err := config.NewYAML(getDefaultConfigYaml(), getAppConfigYaml())
	if err != nil {
		panic(err)
	}

	var cfg appConfig

	if err := yaml.Get("").Populate(&cfg); err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", cfg)
}
