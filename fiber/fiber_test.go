package main

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/client"
	"github.com/gofiber/fiber/v3/middleware/static"
	"github.com/gofiber/template/mustache/v2"
	"github.com/stretchr/testify/assert"
)

var engine = mustache.New("./template", ".mustache")
var app = fiber.New(fiber.Config{
	Views: engine,
})

func TestRouting(t *testing.T) {

	// Define a route for the GET method on the root path '/'
	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	request := httptest.NewRequest("GET", "/", nil)

	// Test the GET request to the root path
	resp, err := app.Test(request)
	if err != nil {
		t.Fatalf("Failed to make GET request: %v", err)
	}

	assert.Equal(t, 200, resp.StatusCode, "Expected status code 200")
	bytes, err := io.ReadAll(resp.Body)
	assert.Nil(t, err, "Failed to read response body")
	assert.Equal(t, "Hello, World 👋!", string(bytes), "Expected response body to be 'Hello, World 👋!'")
}

func TestCtx(t *testing.T) {

	// Define a route for the GET method on the root path '/'
	app.Get("/", func(c fiber.Ctx) error {
		name := c.Query("name", "Guest")
		return c.SendString("Hello, " + name)
	})

	request := httptest.NewRequest("GET", "/?name=Angga", nil)

	resp, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	bytes, err := io.ReadAll(resp.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello, Angga", string(bytes))


	request = httptest.NewRequest("GET", "/", nil)

	resp, err = app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	bytes, err = io.ReadAll(resp.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello, Guest", string(bytes))
}

func TestHttpRequest(t *testing.T) {
	app.Get("/request", func(c fiber.Ctx) error {
		first := c.Get("firstname")
		last := c.Cookies("lastname")
		return c.SendString("Hello, " + first + " " + last)
	})

	request := httptest.NewRequest("GET", "/request", nil)
	request.Header.Set("firstname", "Angga")
	request.AddCookie(&http.Cookie{Name: "lastname", Value: "Ari"})

	resp, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	bytes, err := io.ReadAll(resp.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello, Angga Ari", string(bytes))
}

func TestRouteParameter(t *testing.T) {
	app.Get("/:company/user/:id", func(c fiber.Ctx) error {
		company := c.Params("company")
		id := c.Params("id")
		return c.SendString("Company: " + company + ", User ID: " + id)
	})

	request := httptest.NewRequest("GET", "/sketch-project/user/123", nil)

	resp, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	bytes, err := io.ReadAll(resp.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Company: sketch-project, User ID: 123", string(bytes))
}

func TestRequestForm(t *testing.T) {
	app.Post("/form", func(c fiber.Ctx) error {
		name := c.FormValue("name")
		age := c.FormValue("age")
		return c.SendString("Name: " + name + ", Age: " + age)
	})

	formData := strings.NewReader("name=Angga&age=30")
	request := httptest.NewRequest("POST", "/form", formData)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	bytes, err := io.ReadAll(resp.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Name: Angga, Age: 30", string(bytes))
}

//go:embed assets/file.txt
var fileData []byte
func TestMultipartForm(t *testing.T) {
	app.Post("/upload", func(c fiber.Ctx) error {
		file, err := c.FormFile("file")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("File upload failed")
		}

		err = c.SaveFile(file, "./uploads/" + file.Filename)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to save file")
		}

		return c.SendString("File uploaded: " + file.Filename)
	})

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	file, err := writer.CreateFormFile("file", "file.txt")
	assert.Nil(t, err, "Failed to create form file")
	file.Write(fileData) // Writing content to the file
	writer.Close()
	
	request := httptest.NewRequest("POST", "/upload", body)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	resp, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	bytes, err := io.ReadAll(resp.Body)
	assert.Nil(t, err)
	assert.Equal(t, "File uploaded: file.txt", string(bytes))
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func TestRequestBody(t *testing.T) {
	app.Post("/login", func(c fiber.Ctx) error {
		body := c.Body()
		request := new(LoginRequest)
		err := json.Unmarshal(body, request)
		if err != nil {
			return err
		}
		return c.SendString("Hello: " + request.Username)
	})

	body := strings.NewReader(`{"username": "angga", "password": "secret"}`)
	request := httptest.NewRequest("POST", "/login", body)
	request.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	bytes, err := io.ReadAll(resp.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello: angga", string(bytes))
}

type RegisterRequest struct {
	Username string `json:"username" xml:"username" form:"username"`
	Password string `json:"password" xml:"password" form:"password"`
	Name string `json:"name" xml:"name" form:"name"`
}

func TestBodyParserJson(t *testing.T) {
	app.Post("/register", func(c fiber.Ctx) error {
		request := new(RegisterRequest)
		if err := c.Bind().Body(request);
		 err != nil {
			return err
		}
		return c.SendString("Hello: " + request.Name)
	})

	body := strings.NewReader(`{"username": "angga", "password": "secret", "name": "Angga Ari"}`)
	request := httptest.NewRequest("POST", "/register", body)
	request.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	bytes, err := io.ReadAll(resp.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello: Angga Ari", string(bytes))
}

func TestBodyParserForm(t *testing.T) {
	app.Post("/register", func(c fiber.Ctx) error {
		request := new(RegisterRequest)
		if err := c.Bind().Body(request);
		 err != nil {
			return err
		}
		return c.SendString("Hello: " + request.Name)
	})

	body := strings.NewReader(`username=angga&password=secret&name=Angga+Ari`)
	request := httptest.NewRequest("POST", "/register", body)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	bytes, err := io.ReadAll(resp.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello: Angga Ari", string(bytes))
}

func TestBodyParserXML(t *testing.T) {
	app.Post("/register", func(c fiber.Ctx) error {
		request := new(RegisterRequest)
		if err := c.Bind().Body(request);
		 err != nil {
			return err
		}
		return c.SendString("Hello: " + request.Name)
	})

	body := strings.NewReader(`
		<RegisterRequest>
			<username>angga</username>
			<password>secret</password>
			<name>Angga Ari</name>
		</RegisterRequest>
	`)
	request := httptest.NewRequest("POST", "/register", body)
	request.Header.Set("Content-Type", "application/xml")

	resp, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	bytes, err := io.ReadAll(resp.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello: Angga Ari", string(bytes))
}

func TestHttpResponseJson(t *testing.T) {
	app.Get("/greeting", func(c fiber.Ctx) error {
		c.Set("Content-Type", "application/json")
		c.Status(fiber.StatusOK)
		return c.JSON(fiber.Map{
			"message": "Hello, World 👋!",
			"status":  "success",
		})
	})

	request := httptest.NewRequest("GET", "/greeting", nil)

	resp, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	bytes, err := io.ReadAll(resp.Body)
	assert.Nil(t, err)

	expected := `{"message":"Hello, World 👋!","status":"success"}`
	assert.JSONEq(t, expected, string(bytes), "Expected JSON response to match")
}

func TestHttpResponseRedirect(t *testing.T) {
	app.Get("/redirect", func(c fiber.Ctx) error {
		c.Status(fiber.StatusFound)
		c.Location("https://example.com")
		return nil
	})

	request := httptest.NewRequest("GET", "/redirect", nil)

	resp, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 302, resp.StatusCode)

	location := resp.Header.Get("Location")
	assert.Equal(t, "https://example.com", location, "Expected redirect location to be 'https://example.com'")
}

func TestDownloadFile(t *testing.T) {
	app.Get("/download", func(c fiber.Ctx) error {
		c.Set("Content-Disposition", "attachment; filename=\"file.txt\"")
		//return c.SendString("this is an example file")
		return c.Download("./assets/file.txt", "file.txt")
	})

	request := httptest.NewRequest("GET", "/download", nil)

	resp, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	contentDisposition := resp.Header.Get("Content-Disposition")
	assert.Contains(t, contentDisposition, "attachment; filename=\"file.txt\"", "Expected Content-Disposition header to indicate file download")

	bytes, err := io.ReadAll(resp.Body)
	assert.Nil(t, err)
	assert.Equal(t, "this is an example file", string(bytes), "Expected response body to match file content")
}

func TestRoutingGroup(t *testing.T) {
	// Create a new group for API routes
	api := app.Group("/api")
	api.Get("/hello", func(c fiber.Ctx) error {
		return c.SendString("Hello from API 👋!")
	})

	// Create a new group for Web routes
	web := app.Group("/web")
	web.Get("/hello", func(c fiber.Ctx) error {
		return c.SendString("Hello from Web 👋!")
	})

	request := httptest.NewRequest("GET", "/api/hello", nil)

	// Test the GET request to the group path
	resp, err := app.Test(request)
	if err != nil {
		t.Fatalf("Failed to make GET request: %v", err)
	}

	assert.Equal(t, 200, resp.StatusCode, "Expected status code 200")
	bytes, err := io.ReadAll(resp.Body)
	assert.Nil(t, err, "Failed to read response body")
	assert.Equal(t, "Hello from API 👋!", string(bytes), "Expected response body to be 'Hello from API 👋!'")
}

func TestStatic(t *testing.T) {
	// Serve static files from the "assets" directory
	app.Use("/public", static.New("./assets"))

	request := httptest.NewRequest("GET", "/public/file.txt", nil)

	resp, err := app.Test(request)
	if err != nil {
		t.Fatalf("Failed to make GET request: %v", err)
	}

	assert.Equal(t, 200, resp.StatusCode, "Expected status code 200")
	bytes, err := io.ReadAll(resp.Body)
	assert.Nil(t, err, "Failed to read response body")
	assert.Equal(t, "this is an example file", string(bytes), "Expected response body to be 'this is an example file'")
}

func TestErrorHandler(t *testing.T) {
	// Custom error handler
	app.Use(func(c fiber.Ctx) error {
		return fiber.NewError(fiber.StatusInternalServerError, "Custom error occurred")
	})

	request := httptest.NewRequest("GET", "/", nil)

	resp, err := app.Test(request)
	if err != nil {
		t.Fatalf("Failed to make GET request: %v", err)
	}

	assert.Equal(t, 500, resp.StatusCode, "Expected status code 500")
	bytes, err := io.ReadAll(resp.Body)
	assert.Nil(t, err, "Failed to read response body")
	assert.Equal(t, "Custom error occurred", string(bytes), "Expected response body to be 'Custom error occurred'")
}

func TestTemplateEngine(t *testing.T) {
	// Define a route that renders a template
	app.Get("/view", func(c fiber.Ctx) error {
		data := fiber.Map{
			"title": "Fiber Template Example",
			"header": "Hello from Fiber Template!",
			"content": "This is a simple example of using templates in Fiber.",
		}
		return c.Render("index", data)
	})

	request := httptest.NewRequest("GET", "/view", nil)

	resp, err := app.Test(request)
	if err != nil {
		t.Fatalf("Failed to make GET request: %v", err)
	}

	assert.Equal(t, 200, resp.StatusCode, "Expected status code 200")
	bytes, err := io.ReadAll(resp.Body)
	assert.Nil(t, err, "Failed to read response body")
	assert.Contains(t, string(bytes), "Hello from Fiber Template!")
	assert.Contains(t, string(bytes), "This is a simple example of using templates in Fiber.")
}

func TestMiddleware(t *testing.T) {
	// Middleware to log request method and path
	app.Use(func(c fiber.Ctx) error {
		method := c.Method()
		path := c.Path()
		c.Locals("method", method)
		c.Locals("path", path)
		
		fmt.Println("Before")
		err := c.Next()
		fmt.Println("After")

		return err
	})

	// Middleware to handle API routes (only affecting routes under /api)
	app.Use("/api", func(c fiber.Ctx) error {
		fmt.Println("API Middleware")
		return c.Next()
	})

	app.Get("/middleware", func(c fiber.Ctx) error {
		method := c.Locals("method").(string)
		path := c.Locals("path").(string)
		fmt.Println("Inside handler")
		return c.SendString("Method: " + method + ", Path: " + path)
	})

	request := httptest.NewRequest("GET", "/middleware", nil)

	resp, err := app.Test(request)
	if err != nil {
		t.Fatalf("Failed to make GET request: %v", err)
	}

	assert.Equal(t, 200, resp.StatusCode, "Expected status code 200")
	bytes, err := io.ReadAll(resp.Body)
	assert.Nil(t, err, "Failed to read response body")
	assert.Equal(t, "Method: GET, Path: /middleware", string(bytes), "Expected response body to match middleware output")
}

func TestFiberHttpClient(t *testing.T) {
	cc := client.AcquireRequest()
    cc.SetTimeout(10 * time.Second)
	defer client.ReleaseRequest(cc)

    // Send a GET request
    resp, err := cc.Get("https://httpbin.org/get")
    if err != nil {
        panic(err)
    }

    fmt.Printf("Status: %d\n", resp.StatusCode())
    fmt.Printf("Body: %s\n", string(resp.Body()))
}