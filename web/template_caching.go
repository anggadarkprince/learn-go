package web

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Take out parsing out of the function
// only parse once and use multiple times

//go:embed templates/*.gohtml
var templateFiles embed.FS
var myTemplates = template.Must(template.ParseFS(templateFiles, "templates/*.gohtml"))

func TemplateCaching(writter http.ResponseWriter, request *http.Request) {
	myTemplates.ExecuteTemplate(writter, "simple.gohtml", "Hello template caching")
}

func TestTemplateCaching(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateCaching(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}