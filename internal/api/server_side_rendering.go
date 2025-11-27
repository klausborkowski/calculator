package api

import (
	"net/http"
	"text/template"

	"github.com/mikolabarkouski/calculator/templates"
)

// @Summary Render the application UI
// @Description Serves the HTML UI with available package sizes
// @Tags UI
// @Accept html
// @Produce html
// @Success 200 {string} string "HTML page rendered successfully"
// @Failure 500 {string} string "Failed to render template"
// @Router /app [get]
func (h *Handler) RenderApp(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("index").Parse(templates.IndexHTML))
	packagesDefault := h.app.GetPackagesMap()
	tmpl.Execute(w, packagesDefault)
}
