package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"sort"
	"time"

	"github.com/a-h/templ"
	"github.com/gosimple/slug"
	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/parser"
)

type TemplTemplate func() templ.Component

type Post struct {
	id        string
	title     string
	permaLink string
	tags      []string
	date      time.Time
	dateStr   string
	rawMd     bytes.Buffer
	html      templ.Component
}

type Posts []*Post

var tagMap map[string]bool = make(map[string]bool)
var tags []string

func (e Posts) Len() int {
	return len(e)
}

func (e Posts) Less(i, j int) bool {
	return e[i].date.After(e[j].date)
}

func (e Posts) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func mdToHtml(html string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		_, err = io.WriteString(w, html)
		return
	})
}

func parseMarkdownPosts() (posts []*Post) {

	// turn markdown file into struct
	markdown := goldmark.New(
		goldmark.WithExtensions(
			meta.Meta,
		),
	)
	files, err := os.ReadDir("blog_posts")
	if err != nil {
		log.Fatalf("failed to read directory %v", err)
	}
	for _, f := range files {
		context := parser.NewContext()
		byteArray, err := os.ReadFile(fmt.Sprintf("blog_posts/%s", f.Name()))
		var buf bytes.Buffer

		if err != nil {
			log.Fatalf("failed to read file %v", err)
		}

		if err := markdown.Convert([]byte(byteArray[:]), &buf, parser.WithContext(context)); err != nil {
			panic(err)
		}

		metaData := meta.Get(context)

		var post Post
		post.id = f.Name()
		post.title = metaData["title"].(string)

		for _, tag := range metaData["tags"].([]any) {
			s := tag.(string)
			if !tagMap[s] {
				tags = append(tags, s)
				tagMap[s] = true
			}

			post.tags = append(post.tags, s)
		}

		post.dateStr = metaData["date"].(string)
		post.date, err = time.Parse("2006-01-02", metaData["date"].(string))
		if err != nil {
			log.Fatal(err)
		}
		post.rawMd = buf
		post.html = mdToHtml(post.rawMd.String())
		post.permaLink = path.Join(post.date.Format("2006/01/02"), slug.Make(post.title), "/")

		posts = append(posts, &post)
	}
	return posts
}
func GenerateStaticPage(pathStr, relativePathToRoot string, template TemplTemplate, makeFolder bool) {
	if makeFolder {
		if err := os.Mkdir(pathStr, 0755); err != nil {
			log.Fatalf("failed to create output directory: %v", err)
		}
	}

	name := path.Join(pathStr, "index.html")
	f, err := os.Create(name)
	err = boilerplate(template(), "", relativePathToRoot).Render(context.Background(), f)
	if err != nil {
		log.Fatalf("failed to create output file: %v", err)
	}
}

func main() {
	rootPath := "public"
	blogPath := "public/blog"
	staticPath := "static"

	// wipe public folder (should a makefile handle this?)
	os.RemoveAll(rootPath)

	if err := os.Mkdir(rootPath, 0755); err != nil {
		log.Fatalf("failed to create output directory: %v", err)
	}

	err := os.CopyFS(rootPath, os.DirFS(staticPath))
	if err != nil {
		log.Fatal(err)
	}

	//non blog pages
	GenerateStaticPage(rootPath, "", homeContent, false)
	GenerateStaticPage(path.Join(rootPath, "about"), "../", aboutContent, true)
	GenerateStaticPage(path.Join(rootPath, "resume"), "../", resumeContent, true)
	GenerateStaticPage(path.Join(rootPath, "mail"), "../", mailContent, true)
	GenerateStaticPage(path.Join(rootPath, "projects"), "../", projectContent, true)
	GenerateStaticPage(path.Join(rootPath, "contact"), "../", contactContent, true)

	// make page for each post
	posts := parseMarkdownPosts()
	// sort posts by date oldest to newest
	sort.Sort(Posts(posts))
	sort.Strings(tags)
	// every blog post
	for _, post := range posts {
		dir := path.Join(blogPath, post.date.Format("2006/01/02"), slug.Make(post.title))
		if err := os.MkdirAll(dir, 0755); err != nil && err != os.ErrExist {
			log.Fatalf("failed to create dir %q: %v", dir, err)
		}
		name := path.Join(dir, "index.html")
		f, err := os.Create(name)
		if err != nil {
			log.Fatalf("failed to create output file: %v", err)
		}

		err = boilerplate(contentPage(post), "blog", "../../../../../").Render(context.Background(), f)

		if err != nil {
			log.Fatalf("failed to write output file: %v", err)
		}
	}

	// // BLOG index page (all)
	name := path.Join(blogPath, "index.html")
	f, err := os.Create(name)
	err = boilerplate(blogPostsContent(posts, tags, "", true), "", "../").Render(context.Background(), f)
	if err != nil {
		log.Fatalf("failed to create output file: %v", err)
	}

	// TAG collections pages
	collections := path.Join("public", "blog", "collection")

	if err := os.Mkdir(collections, 0755); err != nil {
		log.Fatalf("failed to create output directory: %v", err)
	}
	for _, tag := range tags {
		collectionPath := path.Join("public", "blog", "collection", tag)

		if err := os.Mkdir(collectionPath, 0755); err != nil {
			log.Fatalf("failed to create output directory: %v", err)
		}

		name = path.Join(collectionPath, "index.html")
		f, err = os.Create(name)
		err = boilerplate(blogPostsContent(posts, tags, tag, false), "", "../../../").Render(context.Background(), f)
		if err != nil {
			log.Fatalf("failed to create output file: %v", err)
		}

	}

}
