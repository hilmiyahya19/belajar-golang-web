package belajar_golang_web

import (
	"fmt"                    // Untuk menampilkan output ke console
	"html/template"          // Package template HTML (auto-escape aktif)
	"io"                     // Untuk membaca body response
	"net/http"               // Package HTTP server & handler
	"net/http/httptest"      // Package untuk testing HTTP handler
	"testing"                // Package testing Go
)

// Handler untuk mendemonstrasikan fitur auto-escape pada template
func TemplateAutoEscape(writer http.ResponseWriter, request *http.Request) {
	// Mengeksekusi template dengan data yang mengandung HTML + JavaScript
	myTemplates.ExecuteTemplate(writer, "post.gohtml", map[string]interface{}{
		"Title": "Template Auto Escape", // Judul halaman
		// Body berisi script yang berpotensi XSS
		"Body": "<p>Ini Adalah Body<script>alert('Anda di Hack')</script></p>",
	})
}

// Unit test untuk TemplateAutoEscape menggunakan httptest
func TestTemplateAutoEscape(t *testing.T) {
	// Membuat request palsu
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	// Recorder untuk menangkap response
	recorder := httptest.NewRecorder()

	// Menjalankan handler
	TemplateAutoEscape(recorder, request)

	// Membaca body hasil render
	body, _ := io.ReadAll(recorder.Result().Body)
	// Menampilkan hasil ke console
	fmt.Println(string(body))
}

// Menjalankan server sungguhan untuk melihat hasil auto-escape di browser
func TestTemplateAutoEscapeServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",                    // Alamat server
		Handler: http.HandlerFunc(TemplateAutoEscape), // Handler HTTP
	}

	// Menjalankan server
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// Handler untuk menonaktifkan auto-escape menggunakan template.HTML
func TemplateAutoEscapeDisabled(writer http.ResponseWriter, request *http.Request) {
	myTemplates.ExecuteTemplate(writer, "post.gohtml", map[string]interface{}{
		"Title": "Template Auto Escape Disabled", // Judul halaman
		// template.HTML menandakan bahwa konten dianggap aman (tidak di-escape)
		"Body": template.HTML("<h1>Ini Adalah Body</h1>"),
	})
}

// Unit test untuk TemplateAutoEscapeDisabled
func TestTemplateAutoEscapeDisabled(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateAutoEscapeDisabled(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

// Menjalankan server sungguhan untuk melihat efek auto-escape yang dimatikan
func TestTemplateAutoEscapeDisabledServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",                             // Alamat server
		Handler: http.HandlerFunc(TemplateAutoEscapeDisabled), // Handler HTTP
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// Handler yang mendemonstrasikan potensi celah XSS
func TemplateXSS(writer http.ResponseWriter, request *http.Request) {
	myTemplates.ExecuteTemplate(writer, "post.gohtml", map[string]interface{}{
		"Title": "Template XSS", // Judul halaman
		// Mengambil input user dari query parameter dan mematikannya auto-escape
		"Body": template.HTML(request.URL.Query().Get("body")),
	})
}

// Unit test untuk TemplateXSS
func TestTemplateXSS(t *testing.T) {
	// Request dengan input HTML dari user
	request := httptest.NewRequest(
		http.MethodGet,
		"http://localhost:8080?body=<p>alert</p>",
		nil,
	)
	recorder := httptest.NewRecorder()

	TemplateXSS(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

// Menjalankan server sungguhan untuk menguji XSS lewat browser
func TestTemplateXSSServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",           // Alamat server
		Handler: http.HandlerFunc(TemplateXSS), // Handler HTTP
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// Kesimpulan:
// Kode ini menjelaskan mekanisme auto-escape pada html/template di Go untuk mencegah serangan XSS dengan cara meng-escape konten HTML secara otomatis. Selain itu, ditunjukkan pula bagaimana auto-escape dapat dimatikan menggunakan template.HTML serta risiko keamanan yang muncul jika input user dirender tanpa validasi dan escaping. Contoh-contoh ini diuji baik menggunakan httptest maupun server HTTP sungguhan untuk memperlihatkan dampaknya secara nyata.
