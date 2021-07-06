package md

import (
	"github.com/b-nova-openhub/gostapagen/pkg/util"
	"github.com/gomarkdown/markdown"
	"github.com/spf13/viper"
)

func ConvertBodyToMarkdown(s string) string {
	return string(markdown.ToHTML([]byte(getBody(s)), nil, nil))
}

func getBody(s string) string {
	return util.SubstringAfter(s, "</"+viper.GetString("contentDelimiter")+">")
}
