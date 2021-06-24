package gen

import (
	"b-nova-openhub/stapagen/pkg/header"
	"b-nova-openhub/stapagen/pkg/md"
	"b-nova-openhub/stapagen/pkg/url"
	"b-nova-openhub/stapagen/pkg/util"
)

type StaticPage struct {
	Title        string `json:"title"`
	Permalink    string `json:"permalink"`
	Author       string `json:"author"`
	Tags         string `json:"tags"`
	Categories   string `json:"categories"`
	PublishDate  string `json:"publishDate"`
	Description  string `json:"description"`
	ShowComments string `json:"showComments"`
	IsPublished  string `json:"isPublished"`
	Body         string `json:"body"`
}

var GeneratedPages []StaticPage

func Generate(files []string) []StaticPage {
	errors := make([]error, 0)
	pages := make([]StaticPage, 0)

	for _, f := range files {
		var p StaticPage
		frontMatter, err := header.GetFrontMatter(f)
		errors = append(errors, err)
		p.Title = frontMatter["title"]
		p.Permalink = url.GetPermalink(frontMatter["title"])
		p.Author = frontMatter["author"]
		p.Tags = frontMatter["tags"]
		p.Categories = frontMatter["categories"]
		p.PublishDate = frontMatter["publishDate"]
		p.Description = frontMatter["description"]
		p.ShowComments = frontMatter["showComments"]
		p.IsPublished = frontMatter["isPublished"]
		p.Body = md.ConvertBodyToMarkdown(f)
		pages = append(pages, p)
	}
	GeneratedPages = pages

	for _, e := range errors {
		util.HandleError(e)
	}

	return pages
}
