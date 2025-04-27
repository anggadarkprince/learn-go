package web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateLayout(writter http.ResponseWriter, request *http.Request) {
	t := template.Must(template.New("web").ParseFiles(
		"templates/layout.gohtml",
		"templates/header.gohtml",
		"templates/footer.gohtml",
	))
	t.ExecuteTemplate(writter, "layout.gohtml", map[string]any{
		"Title": "Template Layout",
		"Name":  "Angga",
	})
}

func TestTemplateLayout(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateLayout(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateLayoutName(writter http.ResponseWriter, request *http.Request) {
	t := template.Must(template.New("web").ParseFiles(
		"templates/layout-name.gohtml",
		"templates/header-name.gohtml",
		"templates/footer-name.gohtml",
	))
	t.ExecuteTemplate(writter, "layout", map[string]any{
		"Title": "Template Layout With Name",
		"Name":  "Angga",
	})
}

func TestTemplateLayoutName(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateLayoutName(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}