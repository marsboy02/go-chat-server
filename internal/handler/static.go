package handler

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

type StaticHandler struct {
	templateDir string
	staticDir   string
}

func NewStaticHandler(templateDir, staticDir string) *StaticHandler {
	return &StaticHandler{
		templateDir: templateDir,
		staticDir:   staticDir,
	}
}

func (h *StaticHandler) ServeHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	tmplPath := filepath.Join(h.templateDir, "index.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		log.Printf("Template parsing error: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := tmpl.Execute(w, nil); err != nil {
		log.Printf("Template execution error: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func (h *StaticHandler) ServeStatic() http.Handler {
	return http.StripPrefix("/static/", http.FileServer(http.Dir(h.staticDir)))
}