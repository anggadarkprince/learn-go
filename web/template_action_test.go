package web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateActionIf(writter http.ResponseWriter, request *http.Request) {
	t := template.Must(template.New("web").ParseFiles("templates/if.gohtml"))
	t.ExecuteTemplate(writter, "if.gohtml", map[string]any{
		"Title": "Template If Action",
		"Name": "Angga",
	})
}

func TestTemplateActionIf(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateActionIf(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateActionComparison(writter http.ResponseWriter, request *http.Request) {
	t := template.Must(template.New("web").ParseFiles("templates/if-comparison.gohtml"))
	t.ExecuteTemplate(writter, "if-comparison.gohtml", map[string]any{
		"Title": "Template Comparison Action",
		"FinalValue": 75,
	})
}

func TestTemplateActionComparison(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateActionComparison(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateActionRange(writter http.ResponseWriter, request *http.Request) {
	t := template.Must(template.New("web").ParseFiles("templates/range.gohtml"))
	t.ExecuteTemplate(writter, "range.gohtml", map[string]any{
		"Title": "Template Range Action",
		"Names": []string{"Angga", "Budi", "Caca"},
	})
}

func TestTemplateActionRange(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateActionRange(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}


func TemplateActionWith(writter http.ResponseWriter, request *http.Request) {
	t := template.Must(template.New("web").ParseFiles("templates/with.gohtml"))
	t.ExecuteTemplate(writter, "with.gohtml", map[string]any{
		"Title": "Template With Action",
		"Name": "Angga",
		"Address": map[string]any{
			"City": "Jakarta",
			"Country": "Indonesia",
		}, 
	})
}

func TestTemplateActionWith(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateActionWith(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}