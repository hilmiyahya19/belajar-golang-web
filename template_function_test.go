package belajar_golang_web

import (
	"fmt"                    // Untuk menampilkan output ke console
	"html/template"          // Package template HTML bawaan Go
	"io"                     // Untuk membaca body response
	"net/http"               // Package HTTP server & client
	"net/http/httptest"      // Untuk testing HTTP handler
	"strings"                // Untuk manipulasi string (Uppercase)
	"testing"                // Package testing Go
)

// Struct MyPage sebagai data yang dikirim ke template
type MyPage struct {
	Name string // Field Name untuk menyimpan nama
}

// Method SayHello milik struct MyPage
// Bisa dipanggil langsung dari template
func (myPage MyPage) SayHello(name string) string {
	return "Hello " + name + ", my name is " + myPage.Name
}

// Handler HTTP untuk contoh penggunaan function method di template
func TemplateFunction(writer http.ResponseWriter, request *http.Request) {
	// Membuat template baru dan langsung parse template string
	t := template.Must(
		template.New("FUNCTION").
			Parse(`{{ .SayHello "Nanook" }}`), // Memanggil method SayHello dari struct
	)

	// Menjalankan template dengan data MyPage
	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "Hilmi",
	})
}

// Unit test untuk TemplateFunction
func TestTemplateFunction(t *testing.T) {
	// Membuat request palsu untuk testing
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)

	// Recorder untuk menangkap response dari handler
	recorder := httptest.NewRecorder()

	// Menjalankan handler
	TemplateFunction(recorder, request)

	// Membaca body response
	body, _ := io.ReadAll(recorder.Result().Body)

	// Menampilkan hasil template ke console
	fmt.Println(string(body))
}

// Handler HTTP untuk contoh penggunaan function global bawaan template
func TemplateFunctionGlobal(writer http.ResponseWriter, request *http.Request) {
	// Menggunakan function global bawaan template: len
	t := template.Must(
		template.New("FUNCTION").
			Parse(`{{ len .Name }}`), // Menghitung panjang string Name
	)

	// Menjalankan template dengan data MyPage
	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "Hilmi",
	})
}

// Unit test untuk TemplateFunctionGlobal
func TestTemplateFunctionGlobal(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionGlobal(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

// Handler HTTP untuk membuat function global custom di template
func TemplateFunctionCreateGlobal(writer http.ResponseWriter, request *http.Request) {
	// Membuat template kosong
	t := template.New("FUNCTION")

	// Menambahkan custom function ke template
	t = t.Funcs(map[string]interface{}{
		"upper": func(value string) string { // Function untuk uppercase string
			return strings.ToUpper(value)
		},
	})

	// Parse template yang menggunakan function custom
	t = template.Must(t.Parse(`{{ upper .Name }}`))

	// Menjalankan template
	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "Hilmi Yahya",
	})
}

// Unit test untuk TemplateFunctionCreateGlobal
func TestTemplateFunctonCreateGlobal(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionCreateGlobal(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

// Handler HTTP untuk contoh penggunaan pipeline function di template
func TemplateFunctionCreateGlobalPipeline(writer http.ResponseWriter, request *http.Request) {
	// Membuat template baru
	t := template.New("FUNCTION")

	// Menambahkan beberapa function global
	t = t.Funcs(map[string]interface{}{
		// function sayHello adalah function global template (berdiri sendiri, bukan method struct)
		"sayHello": func(name string) string { // Function untuk greeting
			return "Hello " + name
		},
		"upper": func(value string) string { // Function untuk uppercase
			return strings.ToUpper(value)
		},
	})

	// Menggunakan pipeline: output sayHello diteruskan ke upper
	t = template.Must(t.Parse(`{{ sayHello .Name | upper }}`))

	// Menjalankan template
	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "Hilmi Yahya",
	})
}

// Unit test untuk TemplateFunctionCreateGlobalPipeline
func TestTemplateFunctionCreateGlobalPipeline(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionCreateGlobalPipeline(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

// Kesimpulan:
// Kode ini mendemonstrasikan penggunaan function pada Go HTML template, mulai dari pemanggilan method struct, penggunaan function global bawaan, pembuatan custom function global, hingga penggunaan pipeline function. Selain itu, setiap handler diuji menggunakan httptest untuk memastikan template dirender dengan benar tanpa harus menjalankan server secara nyata.
