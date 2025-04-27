package web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func login(writter http.ResponseWriter, request *http.Request) {
	//cookie := new(http.Cookie)
	//cookieName = "session_id"
	username := request.URL.Query().Get("username")
	if username == "" {
		username = "guest"
	}
	cookie := http.Cookie{
		Name: "session_id",
		Value: username,
		Path: "/",
		Domain: "localhost",
		HttpOnly: true,
		Secure: false,
		SameSite: http.SameSiteLaxMode,
		Expires: time.Now().Add(60 * 60 * 24),
		MaxAge: 60 * 60 * 24,
	}
	http.SetCookie(writter, &cookie)
	writter.WriteHeader(http.StatusOK)
	writter.Write([]byte("Login Success"))
}

func checkLogin(writter http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("session_id")
	if err != nil {
		writter.WriteHeader(http.StatusUnauthorized)
		writter.Write([]byte("Unauthorized"))
		return
	}
	writter.WriteHeader(http.StatusOK)
	writter.Write([]byte("Welcome " + cookie.Value))
}

func TestCookie(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/check", checkLogin)
	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestAccessCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/login?username=angga", nil)
	recorder := httptest.NewRecorder()

	login(recorder, request)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))

	for _, cookie := range response.Cookies() {
		fmt.Println(cookie.Name)
		fmt.Println(cookie.Value)
		fmt.Println(cookie.Path)
		fmt.Println(cookie.Domain)
		fmt.Println(cookie.Expires)
	}
	fmt.Println("===================================")
	//cookie := response.Cookies()[0]
}

func TestGetCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/check", nil)
	cookie := &http.Cookie{
		Name:  "session_id",
		Value: "angga",
	}
	request.AddCookie(cookie)
	recorder := httptest.NewRecorder()

	checkLogin(recorder, request)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}