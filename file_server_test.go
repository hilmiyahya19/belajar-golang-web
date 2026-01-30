package belajar_golang_web // Nama package, biasanya disesuaikan dengan folder atau modul

import (
	"embed"      // Package untuk menyematkan file/folder ke dalam binary Go
	"io/fs"      // Package untuk bekerja dengan filesystem abstraction
	"net/http"   // Package standar untuk membuat web server HTTP
	"testing"    // Package untuk membuat unit test di Go
)

func TestFileServer(t *testing.T) {
	// Menentukan direktori fisik di filesystem lokal
	directory := http.Dir("./resources")

	// Membuat file server untuk melayani file statis dari directory
	fileServer := http.FileServer(directory)

	// Membuat ServeMux sebagai router HTTP
	mux := http.NewServeMux()

	// Mengatur route /static/ agar mengarah ke file server
	// StripPrefix digunakan agar "/static" tidak ikut dicari di folder
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Konfigurasi HTTP server
	sever := http.Server{
		Addr:    "localhost:8080", // Alamat dan port server
		Handler: mux,              // Handler utama server
	}

	// Menjalankan server HTTP
	err := sever.ListenAndServe()
	if err != nil {
		panic(err) // Menghentikan program jika server gagal dijalankan
	}
}

//go:embed resources
// Menyematkan folder "resources" ke dalam binary aplikasi
var resources embed.FS

func TestFileServerGolangEmbed(t *testing.T) {
	// Mengambil sub-folder "resources" dari embedded filesystem
	directory, _ := fs.Sub(resources, "resources")

	// Membuat file server dari filesystem hasil embed
	fileServer := http.FileServer(http.FS(directory))

	// Membuat ServeMux sebagai router HTTP
	mux := http.NewServeMux()

	// Mengatur route /static/ untuk melayani file dari embedded filesystem
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Konfigurasi HTTP server
	server := http.Server{
		Addr:    "localhost:8080", // Alamat dan port server
		Handler: mux,              // Handler utama server
	}

	// Menjalankan server HTTP
	err := server.ListenAndServe()
	if err != nil {
		panic(err) // Menghentikan program jika server gagal dijalankan
	}
}

// Kesimpulan:
// Kode ini menunjukkan dua cara menjalankan static file server di Go: cara pertama menggunakan folder fisik di filesystem lokal,
// sedangkan cara kedua menggunakan fitur embed untuk menyematkan folder resources langsung ke dalam binary aplikasi.
// Pendekatan embed sangat berguna untuk deployment karena tidak membutuhkan file eksternal,
// sementara konfigurasi ServeMux, FileServer, dan StripPrefix digunakan untuk mengatur routing
// agar file statis dapat diakses melalui endpoint /static/.
