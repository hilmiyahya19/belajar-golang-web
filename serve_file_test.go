package belajar_golang_web // Nama package sesuai modul atau folder project

import (
	_ "embed"      // Digunakan untuk mengaktifkan directive //go:embed
	"fmt"          // Digunakan untuk menulis output ke ResponseWriter
	"net/http"     // Package standar untuk membuat HTTP server dan handler
	"testing"      // Package untuk menjalankan fungsi test di Go
)

func ServeFile(writer http.ResponseWriter, request *http.Request) {
	// Mengecek apakah query parameter "name" ada dan tidak kosong
	if request.URL.Query().Get("name") != "" {
		// Jika ada parameter "name", kirim file ok.html ke client
		http.ServeFile(writer, request, "./resources/ok.html")
	} else {
		// Jika tidak ada parameter "name", kirim file notfound.html
		http.ServeFile(writer, request, "./resources/notfound.html")
	}
}

func TestServeFileServer(t *testing.T) {
	// Konfigurasi HTTP server
	server := http.Server{
		Addr: "localhost:8080",             // Alamat dan port server
		Handler: http.HandlerFunc(ServeFile), // Handler HTTP menggunakan fungsi ServeFile
	}

	// Menjalankan server HTTP
	err := server.ListenAndServe()
	if err != nil {
		panic(err) // Menghentikan program jika server gagal dijalankan
	}
}

//go:embed resources/ok.html
// Menyematkan isi file ok.html ke dalam binary sebagai string
var resourceOk string

//go:embed resources/notfound.html
// Menyematkan isi file notfound.html ke dalam binary sebagai string
var resourceNotFound string

func ServeFileEmbed(writer http.ResponseWriter, request *http.Request) {
	// Mengecek apakah query parameter "name" ada dan tidak kosong
	if request.URL.Query().Get("name") != "" {
		// Menulis langsung konten HTML dari hasil embed ke response
		fmt.Fprint(writer, resourceOk)
	} else {
		// Menulis konten HTML notfound dari hasil embed ke response
		fmt.Fprint(writer, resourceNotFound)
	}
}

func TestServeFileEmbedServer(t *testing.T) {
	// Konfigurasi HTTP server
	server := http.Server{
		Addr: "localhost:8080",                   // Alamat dan port server
		Handler: http.HandlerFunc(ServeFileEmbed), // Handler HTTP menggunakan fungsi ServeFileEmbed
	}

	// Menjalankan server HTTP
	err := server.ListenAndServe()
	if err != nil {
		panic(err) // Menghentikan program jika server gagal dijalankan
	}
}

// Kesimpulan:
// Kode ini menunjukkan dua pendekatan dalam melayani file HTML di Go, yaitu menggunakan file fisik di filesystem dengan http.ServeFile dan menggunakan fitur embed untuk menyematkan file HTML langsung ke dalam binary aplikasi.
// Pendekatan embed membuat aplikasi lebih mudah dideploy karena tidak bergantung pada file eksternal,
// sementara penggunaan query parameter digunakan sebagai logika sederhana untuk menentukan response yang dikirim ke client.
