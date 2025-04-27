package web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateAutoEscape(writter http.ResponseWriter, request *http.Request) {
	t := template.Must(template.New("web").ParseFiles(
		"templates/post.gohtml",
	))
	t.ExecuteTemplate(writter, "post.gohtml", map[string]any{
		"Title": "Template Auto Escape",
		"Body": "<script>alert('XSS')</script>",
	})
}

func TestTemplateAutoEscape(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateAutoEscape(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TestTemplateAutoEscapeServer(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", TemplateAutoEscape)
	server := http.Server{
		Addr: "localhost:8080",
	}
	server.Handler = mux

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TemplateAutoEscapeDisabled(writter http.ResponseWriter, request *http.Request) {
	t := template.Must(template.New("web").ParseFiles(
		"templates/post.gohtml",
	))
	t.ExecuteTemplate(writter, "post.gohtml", map[string]any{
		"Title": "Template Auto Escape",
		"Body": template.HTML("<b>This is a content</b>"),
	})
}

func TestTemplateAutoEscapeDisableServer(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", TemplateAutoEscapeDisabled)
	server := http.Server{
		Addr: "localhost:8080",
	}
	server.Handler = mux

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TemplateXSS(writter http.ResponseWriter, request *http.Request) {
	t := template.Must(template.New("web").ParseFiles(
		"templates/post.gohtml",
	))
	t.ExecuteTemplate(writter, "post.gohtml", map[string]any{
		"Title": "Template XSS",
		"Body": template.HTML(request.URL.Query().Get("body")),
	})
}

func TestTemplateXSS(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/?body=<script>alert('XSS')</script>", nil)
	recorder := httptest.NewRecorder()

	TemplateXSS(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}


func TestTemplateXSSServer(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", TemplateXSS)
	server := http.Server{
		Addr: "localhost:8080",
	}
	server.Handler = mux

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}