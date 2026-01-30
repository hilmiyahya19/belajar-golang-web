package belajar_golang_web // Nama package sesuai dengan modul atau folder project

import (
	"embed"                 // Package untuk fitur embed file/template ke dalam binary
	"fmt"                   // Digunakan untuk output ke console
	"html/template"         // Package untuk HTML templating yang aman (auto-escape)
	"io"                    // Digunakan untuk membaca body response
	"net/http"              // Package standar untuk HTTP server dan handler
	"net/http/httptest"     // Package untuk testing HTTP handler tanpa server sungguhan
	"testing"               // Package testing bawaan Go
)

func SimpleHTML(writer http.ResponseWriter, request *http.Request) {
	// Template HTML sederhana dalam bentuk string
	templateText := `<html><body>{{.}}</body></html>`

	//t, err := template.New("SIMPLE").Parse(templateText)
	//if err != nil {
	//	panic(err)
	//}

	// Membuat dan mem-parse template, panic otomatis jika terjadi error
	t := template.Must(template.New("SIMPLE").Parse(templateText))

	// Menjalankan template bernama "SIMPLE" dengan data string
	t.ExecuteTemplate(writer, "SIMPLE", "Hello HTML Template")
}

func TestSimpleHTML(t *testing.T) {
	// Membuat HTTP request palsu untuk keperluan testing
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)

	// Recorder untuk menangkap response dari handler
	recorder := httptest.NewRecorder()

	// Memanggil handler secara langsung
	SimpleHTML(recorder, request)

	// Membaca hasil response body
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body)) // Menampilkan hasil render template ke console
}

func SimpleHTMLFile(writer http.ResponseWriter, request *http.Request) {
	// Mem-parsing satu file template dari filesystem
	t := template.Must(template.ParseFiles("./templates/simple.gohtml"))

	// Menjalankan template berdasarkan nama file
	t.ExecuteTemplate(writer, "simple.gohtml", "Hello HTML Template")
}

func TestSimpleHTMLFile(t *testing.T) {
	// Membuat HTTP request palsu
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)

	// Recorder untuk menangkap response
	recorder := httptest.NewRecorder()

	// Memanggil handler
	SimpleHTMLFile(recorder, request)

	// Membaca dan menampilkan response body
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateDirectory(writer http.ResponseWriter, request *http.Request) {
	// Mem-parsing semua file template dengan ekstensi .gohtml dalam satu folder
	t := template.Must(template.ParseGlob("./templates/*.gohtml"))

	// Menjalankan salah satu template dari kumpulan template
	t.ExecuteTemplate(writer, "simple.gohtml", "Hello HTML Template")
}

func TestTemplateDirectory(t *testing.T) {
	// Membuat HTTP request palsu
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)

	// Recorder untuk response
	recorder := httptest.NewRecorder()

	// Memanggil handler
	TemplateDirectory(recorder, request)

	// Membaca dan menampilkan response body
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

//go:embed templates/*.gohtml
// Menyematkan semua file template .gohtml ke dalam binary aplikasi
var templates embed.FS

func TemplateEmbed(writer http.ResponseWriter, request *http.Request) {
	// Mem-parsing template langsung dari embedded filesystem
	t := template.Must(template.ParseFS(templates, "templates/*.gohtml"))

	// Menjalankan template dari hasil embed
	t.ExecuteTemplate(writer, "simple.gohtml", "Hello HTML Template")
}

func TestTemplateEmbed(t *testing.T) {
	// Membuat HTTP request palsu
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)

	// Recorder untuk response
	recorder := httptest.NewRecorder()

	// Memanggil handler embed template
	TemplateEmbed(recorder, request)

	// Membaca dan menampilkan response body
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

// Kesimpulan:
// Kode ini mendemonstrasikan beberapa cara penggunaan HTML template di Go, mulai dari template berbasis string, template dari satu file, template dari satu direktori, hingga template yang di-embed langsung ke dalam binary menggunakan fitur embed. Seluruh contoh diuji menggunakan httptest tanpa menjalankan server sungguhan, sehingga memudahkan pengujian dan debugging. Pendekatan embed sangat cocok untuk deployment karena tidak bergantung pada file eksternal dan membuat aplikasi lebih portable.
