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

func main() {
	rootPath := "public"
	blogPath := "public/blog"
	aboutPath := "public/about"
	staticPath := "static"

	// wipe public folder
	os.RemoveAll(rootPath)

	if err := os.Mkdir(rootPath, 0755); err != nil {
		log.Fatalf("failed to create output directory: %v", err)
	}

	// TODO get rootpath to reference files in staticpath instead of copying them over
	err := os.CopyFS(rootPath, os.DirFS(staticPath))
	if err != nil {
		log.Fatal(err)
	}

	// make page for each post
	posts := parseMarkdownPosts()
	// sort posts by date oldest to newest
	sort.Sort(Posts(posts))
	sort.Strings(tags)

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

	// TODO this could be automated somewhat
	name := path.Join(rootPath, "index.html")
	f, err := os.Create(name)
	err = boilerplate(homeContent(), "", "").Render(context.Background(), f)
	if err != nil {
		log.Fatalf("failed to create output file: %v", err)
	}

	if err := os.Mkdir(aboutPath, 0755); err != nil {
		log.Fatalf("failed to create output directory: %v", err)
	}

	name = path.Join(rootPath, "about", "index.html")
	f, err = os.Create(name)
	err = boilerplate(aboutContent(), "", "../").Render(context.Background(), f)
	if err != nil {
		log.Fatalf("failed to create output file: %v", err)
	}

	if err := os.Mkdir("public/resume", 0755); err != nil {
		log.Fatalf("failed to create output directory: %v", err)
	}

	name = path.Join(rootPath, "resume", "index.html")
	f, err = os.Create(name)
	err = boilerplate(resumeContent(), "", "../").Render(context.Background(), f)
	if err != nil {
		log.Fatalf("failed to create output file: %v", err)
	}

	if err := os.Mkdir("public/mail", 0755); err != nil {
		log.Fatalf("failed to create output directory: %v", err)
	}

	name = path.Join(rootPath, "mail", "index.html")
	f, err = os.Create(name)
	err = boilerplate(mailContent(), "", "../").Render(context.Background(), f)
	if err != nil {
		log.Fatalf("failed to create output file: %v", err)
	}

	if err := os.Mkdir("public/projects", 0755); err != nil {
		log.Fatalf("failed to create output directory: %v", err)
	}

	name = path.Join(rootPath, "projects", "index.html")
	f, err = os.Create(name)
	err = boilerplate(projectContent(), "", "../").Render(context.Background(), f)
	if err != nil {
		log.Fatalf("failed to create output file: %v", err)
	}

	if err := os.Mkdir("public/contact", 0755); err != nil {
		log.Fatalf("failed to create output directory: %v", err)
	}

	name = path.Join(rootPath, "contact", "index.html")
	f, err = os.Create(name)
	err = boilerplate(contactContent(), "", "../").Render(context.Background(), f)
	if err != nil {
		log.Fatalf("failed to create output file: %v", err)
	}

	// BLOGPOSTS
	name = path.Join(blogPath, "index.html")
	f, err = os.Create(name)
	err = boilerplate(indexPage(posts, tags), "", "../").Render(context.Background(), f)
	if err != nil {
		log.Fatalf("failed to create output file: %v", err)
	}

}
