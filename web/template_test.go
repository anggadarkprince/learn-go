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

func SimpleHtml(writter http.ResponseWriter, request *http.Request) {
	templateText := `<html><body>{{.}}</body></html>`
	//t, err := template.New("web").Parse(templateText)
	//if err != nil {
	//	http.Error(writter, err.Error(), http.StatusInternalServerError)
	//}
	t := template.Must(template.New("web").Parse(templateText))
	t.ExecuteTemplate(writter, "web", "Hello World")
}

func TestSimpleHtml(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	SimpleHtml(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func SimpleHtmlFile(writter http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("templates/simple.gohtml"))
	t.ExecuteTemplate(writter, "simple.gohtml", "Hello HTML World")
}

func TestSimpleHtmlTemplate(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	SimpleHtmlFile(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

// Test for template directory
func TemplateDirectory(writter http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseGlob("templates/*.gohtml"))
	t.ExecuteTemplate(writter, "simple.gohtml", "Hello HTML World")
}

func TestSimpleHtmlTemplateDirectory(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateDirectory(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

//go:embed templates/*.gohtml
var templates embed.FS

func TemplateEmbed(writter http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFS(templates, "templates/*.gohtml"))
	t.ExecuteTemplate(writter, "simple.gohtml", "Hello HTML World")
}
func TestSimpleHtmlTemplateEmbed(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateEmbed(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateDataMap(writter http.ResponseWriter, request *http.Request) {
	data := map[string]any{
		"Title": "Template Data Map",
		"Name": "Angga",
		"Address": map[string]any{
			"City": "Jakarta",
		},
	}
	t := template.Must(template.ParseFS(templates, "templates/*.gohtml"))
	t.ExecuteTemplate(writter, "data.gohtml", data)
}

func TestSimpleHtmlTemplateDataMap(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateDataMap(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateDataStruct(writter http.ResponseWriter, request *http.Request) {
	type Address struct {
		City string
	}
	type Data struct {
		Title string
		Name string
		Address Address
	}
	data := Data{
		Title: "Template Data Struct",
		Name: "Angga",
		Address: Address{
			City: "Jakarta",
		},
	}
	t := template.Must(template.ParseFS(templates, "templates/*.gohtml"))
	t.ExecuteTemplate(writter, "data.gohtml", data)
}

func TestSimpleHtmlTemplateDataStruct(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateDataStruct(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
