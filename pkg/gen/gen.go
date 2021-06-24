package gen

import (
	"b-nova-openhub/stapagen/pkg/header"
	"b-nova-openhub/stapagen/pkg/md"
	"b-nova-openhub/stapagen/pkg/url"
	"b-nova-openhub/stapagen/pkg/util"
	"sort"
	"time"
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

type Status struct {
	TotalPages        int    `json:"totalPages"`
	PublishedPages    int    `json:"publishedPages"`
	UnpublishedPages  int    `json:"unpublishedPages"`
	LastPublishedPage string `json:"lastPublishedPage"`
	LastGeneratedAt   string `json:"lastGeneratedAt"`
}

var GeneratedPages []StaticPage
var CurrentStatus *Status

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
		p.PublishDate = frontMatter["date"]
		p.Description = frontMatter["description"]
		p.ShowComments = frontMatter["showComments"]
		p.IsPublished = frontMatter["publish"]
		p.Body = md.ConvertBodyToMarkdown(f)
		pages = append(pages, p)
	}

	for _, e := range errors {
		util.HandleError(e)
	}

	GeneratedPages = pages
	CurrentStatus = new(Status)
	CurrentStatus.TotalPages = len(GeneratedPages)
	CurrentStatus.PublishedPages = getPublished(GeneratedPages)
	CurrentStatus.UnpublishedPages = getUnpublished(GeneratedPages)
	CurrentStatus.LastPublishedPage = getLastPublished(GeneratedPages)
	CurrentStatus.LastGeneratedAt = time.Now().String()
	return pages
}

func getPublished(pages []StaticPage) (count int) {
	count = 0
	for _, page := range pages {
		if page.IsPublished == "true" {
			count++
		}
	}
	return count
}

func getUnpublished(pages []StaticPage) (count int) {
	count = 0
	for _, page := range pages {
		if page.IsPublished == "false" {
			count++
		}
	}
	return count
}

func getLastPublished(pages []StaticPage) string {
	dates := make([]string, 0)
	for _, page := range pages {
		dates = append(dates, page.PublishDate)
	}
	sort.Strings(dates)
	return dates[len(dates)-1]
}
