package web

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func UploadForm(writter http.ResponseWriter, request *http.Request) {
	myTemplates.ExecuteTemplate(writter, "upload-form.gohtml", nil)
}

func Upload(writter http.ResponseWriter, request *http.Request) {
	//request.ParseMultipartForm(10 << 20) // 10 MB limit

	file, fileHeader, err := request.FormFile("file")
	if err != nil {
		http.Error(writter, "Error retrieving the file", http.StatusBadRequest)
		return
	}
	fileDestination, err := os.Create("statics/" + fileHeader.Filename)
	if err != nil {
		http.Error(writter, "Error creating the file", http.StatusInternalServerError)
		return
	}
	_, err = io.Copy(fileDestination, file)
	if err != nil {
		http.Error(writter, "Error saving the file", http.StatusInternalServerError)
		return
	}

	name := request.PostFormValue("name")
	myTemplates.ExecuteTemplate(writter, "upload-response.gohtml", map[string]any{
		"Name": name,
		"File": "/statics/" + fileHeader.Filename,
	})
}

func TestUploadForm(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/form", UploadForm)
	mux.HandleFunc("/upload", Upload)
	mux.Handle("/statics/", http.StripPrefix("/statics", http.FileServer(http.Dir("statics"))))
	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//go:embed resources/gear.png
var uploadFileTest []byte

func TestUploadFile(t *testing.T) {
	body := new(bytes.Buffer)

	writer := multipart.NewWriter(body)
	writer.WriteField("name", "Angga Ari")
	file, _ := writer.CreateFormFile("file", "image.png")
	file.Write(uploadFileTest)
	writer.Close()

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/upload", body)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	recorder := httptest.NewRecorder()

	Upload(recorder, request)

	bodyResponse, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(bodyResponse))
}