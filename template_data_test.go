package belajar_golang_web // Nama package sesuai modul atau folder project

import (
	"fmt"               // Digunakan untuk menampilkan output ke console
	"html/template"     // Package untuk HTML template Go yang aman (auto escaping)
	"io"                // Digunakan untuk membaca response body
	"net/http"          // Package standar HTTP untuk handler
	"net/http/httptest" // Package untuk testing HTTP handler tanpa server nyata
	"testing"           // Package bawaan Go untuk unit test
)

func TemplateDataMap(writer http.ResponseWriter, request *http.Request) {
	// Mem-parsing file template dari filesystem
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))

	// Menjalankan template dengan data berbentuk map
	t.ExecuteTemplate(writer, "name.gohtml", map[string]interface{}{
		"Title": "Template Data Map", // Data judul halaman
		"Name":  "Hilmi",             // Data nama yang akan ditampilkan di template
		"Address": map[string]interface{}{ // Nested map untuk data alamat
			"Street": "Jalan Belum Ada", // Data alamat jalan
		},
	})
}

func TestTemplateDataMap(t *testing.T) {
	// Membuat HTTP request palsu untuk testing
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)

	// Recorder untuk menangkap response dari handler
	recorder := httptest.NewRecorder()

	// Memanggil handler TemplateDataMap
	TemplateDataMap(recorder, request)

	// Membaca dan menampilkan response body hasil render template
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

// Struct untuk menyimpan data alamat
type Address struct {
	Street string // Field Street untuk alamat jalan
}

// Struct utama untuk data halaman
type Page struct {
	Title   string  // Judul halaman
	Name    string  // Nama pengguna
	Address Address // Data alamat dalam bentuk struct
}

func TemplateDataStruct(writer http.ResponseWriter, request *http.Request) {
	// Mem-parsing file template dari filesystem
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))

	// Menjalankan template dengan data berbentuk struct
	t.ExecuteTemplate(writer, "name.gohtml", Page{
		Title: "Template Data Struct", // Judul halaman
		Name:  "Hilmi",                // Nama pengguna
		Address: Address{              // Data alamat dalam struct
			Street: "Jalan Belum Ada", // Alamat jalan
		},
	})
}

func TestTemplateDataStruct(t *testing.T) {
	// Membuat HTTP request palsu untuk testing
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)

	// Recorder untuk menangkap response dari handler
	recorder := httptest.NewRecorder()

	// Memanggil handler TemplateDataStruct
	TemplateDataStruct(recorder, request)

	// Membaca dan menampilkan response body hasil render template
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

// Kesimpulan:
// Kode ini memperlihatkan dua cara mengirim data ke HTML template di Go, yaitu menggunakan map dan menggunakan struct.
// Pendekatan map bersifat fleksibel dan cepat untuk data dinamis, sedangkan struct lebih terstruktur, aman secara tipe, dan lebih direkomendasikan untuk aplikasi skala besar. Seluruh contoh diuji menggunakan httptest sehingga proses rendering template dapat divalidasi tanpa harus menjalankan server HTTP secara langsung.
