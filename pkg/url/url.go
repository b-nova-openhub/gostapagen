package url

import "github.com/gosimple/slug"

func GetPermalink(s string) string {
	return "b-nova.com/home/content/" + slug.Make(s)
}
