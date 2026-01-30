package belajar_golang_web // Nama package sesuai modul atau folder project

import (
	"fmt"               // Digunakan untuk mencetak output ke console
	"html/template"     // Package template HTML Go (aman, auto-escape)
	"io"                // Digunakan untuk membaca response body
	"net/http"          // Package HTTP untuk handler
	"net/http/httptest" // Package untuk testing HTTP handler tanpa server nyata
	"testing"           // Package testing bawaan Go
)

func TemplateActionIf(writer http.ResponseWriter, request *http.Request) {
	// Mem-parsing file template yang berisi action {{if}}
	t := template.Must(template.ParseFiles("./templates/if.gohtml"))

	// Menjalankan template dengan data berbentuk struct Page
	t.ExecuteTemplate(writer, "if.gohtml", Page{
		Title: "Template Action If", // Judul halaman
		Name:  "Hilmi",              // Nama untuk logika if di template
	})
}

func TestTemplateActionIf(t *testing.T) {
	// Membuat HTTP request palsu
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)

	// Recorder untuk menangkap response
	recorder := httptest.NewRecorder()

	// Memanggil handler TemplateActionIf
	TemplateActionIf(recorder, request)

	// Membaca dan menampilkan hasil render template
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateActionOperator(writer http.ResponseWriter, request *http.Request) {
	// Mem-parsing template yang berisi operator perbandingan (eq, gt, lt, dll)
	t := template.Must(template.ParseFiles("./templates/comparator.gohtml"))

	// Menjalankan template dengan data map
	t.ExecuteTemplate(writer, "comparator.gohtml", map[string]interface{}{
		"Title":      "Template Action Operator", // Judul halaman
		"FinalValue": 80,                          // Nilai untuk dibandingkan di template
	})
}

func TestTemplateActionOperator(t *testing.T) {
	// Membuat HTTP request palsu
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)

	// Recorder untuk response
	recorder := httptest.NewRecorder()

	// Memanggil handler TemplateActionOperator
	TemplateActionOperator(recorder, request)

	// Membaca dan menampilkan hasil render template
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateActionRange(writer http.ResponseWriter, request *http.Request) {
	// Mem-parsing template yang menggunakan action {{range}}
	t := template.Must(template.ParseFiles("./templates/range.gohtml"))

	// Menjalankan template dengan data slice
	t.ExecuteTemplate(writer, "range.gohtml", map[string]interface{}{
		"Title": "Template Action Range", // Judul halaman
		"Hobbies": []string{              // Data list untuk dirender dengan range
			"Game", "Watch", "Code",
		},
	})
}

func TestTemplateActionRange(t *testing.T) {
	// Membuat HTTP request palsu
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)

	// Recorder untuk response
	recorder := httptest.NewRecorder()

	// Memanggil handler TemplateActionRange
	TemplateActionRange(recorder, request)

	// Membaca dan menampilkan hasil render template
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateActionWith(writer http.ResponseWriter, request *http.Request) {
	// Mem-parsing template yang menggunakan action {{with}}
	t := template.Must(template.ParseFiles("./templates/address.gohtml"))

	// Menjalankan template dengan data map bersarang
	t.ExecuteTemplate(writer, "address.gohtml", map[string]interface{}{
		"Title": "Template Action With", // Judul halaman
		"Name":  "Hilmi",                // Nama pengguna
		"Address": map[string]interface{}{ // Data nested untuk konteks with
			"Street": "Jalan Belum Ada",
			"City":   "Jakarta",
		},
	})
}

func TestTemplateActionWith(t *testing.T) {
	// Membuat HTTP request palsu
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)

	// Recorder untuk response
	recorder := httptest.NewRecorder()

	// Memanggil handler TemplateActionWith
	TemplateActionWith(recorder, request)

	// Membaca dan menampilkan hasil render template
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

// Kesimpulan:
// Kode ini mendemonstrasikan penggunaan berbagai template action pada Go HTML template, yaitu if untuk percabangan logika, operator untuk perbandingan nilai, range untuk iterasi data slice, dan with untuk mempersempit konteks data bersarang. Seluruh contoh dijalankan dan diuji menggunakan httptest tanpa menjalankan server HTTP sungguhan, sehingga memudahkan pemahaman dan pengujian perilaku template secara terisolasi.
