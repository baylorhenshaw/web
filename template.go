package web

import (
	"html/template"
	"net/http"
	"path/filepath"
	"time"
)

// Register Page Route
func (s *Server) RegisterPage(page Page, handler func(w http.ResponseWriter, r *http.Request)) {

	// Store Route
	s.Pages[page.PageFile] = page

	// Parse & Cache Page Template
	s.parseAndCachePageTemplate(page)

	// Register HTTP Handler
	http.HandleFunc(page.Route, handler)
}

func (s *Server) parseAndCachePageTemplate(page Page) {
	start := time.Now()

	// Build Template
	tmpl := template.New("")
	tmpl = parseDeepPath("templates/components/", tmpl, 5)
	tmpl = parseFile("templates/layouts/"+page.LayoutFile+".tmpl", tmpl)
	tmpl = parseFile("templates/pages/"+page.PageFile+".tmpl", tmpl)

	// Cache Template
	s.PageTemplates[page.PageFile] = tmpl

	Logger.Info("Parsed and cached template in "+time.Since(start).String(), "template", page)
}

// Parse File
func parseFile(path string, template *template.Template) *template.Template {
	template, err := template.ParseFiles(path)
	if err != nil {
		Logger.Fatal(err)
	}
	return template
}

// Parse Many Paths
func parseDeepPath(path string, template *template.Template, depth int) *template.Template {
	for i := 0; i < depth; i++ {
		path = path + "*/"
		template = parsePath(path, template)
	}
	return template
}

// Parse Path
func parsePath(path string, template *template.Template) *template.Template {
	pattern := filepath.Join(path, "*.tmpl")
	new, err := template.ParseGlob(pattern)
	if err != nil {
		Logger.Debug(err)
		return template
	}
	return new
}
