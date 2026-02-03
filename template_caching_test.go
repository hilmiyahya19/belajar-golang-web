package belajar_golang_web

import (
	"embed"                  // Package untuk embed file ke dalam binary
	"fmt"                    // Untuk menampilkan output ke console
	"html/template"          // Package template HTML bawaan Go
	"io"                     // Untuk membaca body response
	"net/http"               // Package HTTP server & client
	"net/http/httptest"      // Package untuk testing HTTP handler
	"testing"                // Package testing Go
)

// Directive untuk meng-embed semua file template .gohtml di folder templates
//go:embed templates/*.gohtml
var templates embed.FS // File system virtual berisi template yang di-embed

// Parsing seluruh template sekali di awal (template caching)
var myTemplates = template.Must(
	template.ParseFS(templates, "templates/*.gohtml"), // Membaca template dari embed FS
)

// Handler HTTP untuk menggunakan template yang sudah di-cache
func TemplateCaching(writer http.ResponseWriter, request *http.Request) {
	// Menjalankan template tanpa parsing ulang (lebih efisien)
	myTemplates.ExecuteTemplate(
		writer,
		"simple.gohtml",            // Nama file template yang dieksekusi
		"Hello Template Caching",   // Data yang dikirim ke template
	)
}

// Unit test untuk TemplateCaching
func TestTemplateCaching(t *testing.T) {
	// Membuat request palsu untuk keperluan testing
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)

	// Recorder untuk menangkap response dari handler
	recorder := httptest.NewRecorder()

	// Menjalankan handler TemplateCaching
	TemplateCaching(recorder, request)

	// Membaca body hasil render template
	body, _ := io.ReadAll(recorder.Result().Body)

	// Menampilkan hasil ke console
	fmt.Println(string(body))
}

// Kesimpulan:
// Kode ini menunjukkan penerapan template caching pada Go dengan memanfaatkan embed.FS untuk menyimpan file template di dalam binary aplikasi. Seluruh template diparse satu kali di awal aplikasi sehingga handler HTTP dapat mengeksekusi template tanpa parsing ulang, yang meningkatkan performa dan efisiensi aplikasi. Proses ini diuji menggunakan httptest untuk memastikan template berhasil dirender dengan benar tanpa menjalankan server secara nyata.
