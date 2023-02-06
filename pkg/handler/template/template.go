package template

import (
	tmpHtml "github.com/taaaaakahiro/golang-rest-example/template"
	"go.uber.org/zap"
	"html/template"
	"log"
	"net/http"
	"time"
)

const (
	index = "index.html"
)

type Handler struct {
	logger   *zap.Logger
	res      *response
	template *tmpHtml.Template
}

type response struct {
}

func NewHandler(template *tmpHtml.Template) *Handler {
	return &Handler{
		template: template,
	}
}

func (h *Handler) IndexTemplateHandler(w http.ResponseWriter, r *http.Request) {
	filePath := h.template.Path + "/" + index
	t, err := template.ParseFiles(filePath)
	if err != nil {
		log.Fatal(err)
	}
	if err := t.Execute(w, struct {
		Title   string
		Message string
		Time    time.Time
	}{
		Title:   "テストページ",
		Message: "こんにちは！",
		Time:    time.Now(),
	}); err != nil {
		log.Printf("failed to execute template: %v", err)
	}
}
