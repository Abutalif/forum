package pages

import (
	"html/template"
	"net/http"
)

type pageloader struct {
	tmpl  *template.Template
	pages map[string]string
}

func NewPages() (*pageloader, error) {
	pl := &pageloader{
		pages: make(map[string]string),
	}
	pl.initPages()
	if err := pl.loadTemplates(); err != nil {
		return nil, err
	}
	return pl, nil
}

func (p *pageloader) loadTemplates() error {
	tmpls, err := template.ParseGlob("./assets/templates/*.html")
	if err != nil {
		return err
	}
	p.tmpl = tmpls
	return nil
}

func (p *pageloader) initPages() {
	p.pages["/"] = "index.html"
	p.pages["/feed"] = "feed.html"
	p.pages["/post"] = "post.html"
	p.pages["/profile"] = "profile.html"
}

func (p *pageloader) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// TODO: if not method get should give error page with proper status
	page, ok := p.pages[r.URL.Path]
	if !ok {
		page = "error.html"
	}
	if r.Method != http.MethodGet {
		page = "error.html"
	}
	err := p.tmpl.ExecuteTemplate(rw, page, nil)
	if err != nil {
		http.Error(rw, "Cannot parse template", http.StatusInternalServerError)
	}
}
