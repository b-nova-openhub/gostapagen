package header

import (
	"github.com/b-nova-openhub/gostapagen/pkg/util"
	"github.com/spf13/viper"
	"strings"
)

func GetFrontMatter(file string) map[string]string {
	header := getHeader(file)
	return getKeyValues(header, ":")
}

func getHeader(file string) string {
	contentDelimiter := viper.GetString("contentDelimiter")
	header := util.SubstringBetween(file, "<"+contentDelimiter+">", "</"+contentDelimiter+">")
	return string(header)
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
