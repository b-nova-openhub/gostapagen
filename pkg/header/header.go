package header

import (
	"b-nova-openhub/stapagen/pkg/util"
	"strings"
)

func GetFrontMatter(file string) (map[string]string, error) {
	header, err := getHeader(file)
	return getKeyValues(header, ":"), err
}

func getHeader(file string) (string, error) {
	header, err := util.SubstringBetween(file, "<b-nova-content-header>", "</b-nova-content-header>")
	return string(header), err
}

func getKeyValues(s, sep string) map[string]string {
	lines := strings.Split(s, "\n")
	frontMatter := make(map[string]string)
	for _, pair := range lines {
		if util.StringNotEmpty(pair) {
			kv := strings.Split(pair, sep)
			frontMatter[kv[0]] = strings.TrimSpace(kv[1])
		}
	}
	return frontMatter
}
