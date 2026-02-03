package belajar_golang_web

import (
	"fmt"            // Untuk menulis output ke response HTTP
	"net/http"       // Package HTTP server, handler, dan redirect
	"testing"        // Package testing Go
)

// Handler tujuan redirect (endpoint akhir)
func RedirectTo(writer http.ResponseWriter, request *http.Request) {
	// Menulis response sederhana ke client
	fmt.Fprint(writer, "Hello Redirect")
}

// Handler yang melakukan redirect ke endpoint internal
func RedirectFrom(writer http.ResponseWriter, request *http.Request) {
	// Melakukan redirect ke path /redirect-to
	// StatusTemporaryRedirect = HTTP 307
	http.Redirect(writer, request, "/redirect-to", http.StatusTemporaryRedirect)
}

// Handler yang melakukan redirect ke website eksternal
func RedirectOut(writer http.ResponseWriter, request *http.Request) {
	// Redirect ke URL luar (Google)
	http.Redirect(writer, request, "https://google.com", http.StatusTemporaryRedirect)
}

// Menjalankan server HTTP untuk menguji redirect
func TestRedirect(t *testing.T) {
	// Membuat HTTP request multiplexer
	mux := http.NewServeMux()

	// Mendaftarkan handler untuk masing-masing endpoint
	mux.HandleFunc("/redirect-to", RedirectTo)
	mux.HandleFunc("/redirect-from", RedirectFrom)
	mux.HandleFunc("/redirect-out", RedirectOut)

	// Konfigurasi server HTTP
	server := http.Server{
		Addr:    "localhost:8080", // Alamat dan port server
		Handler: mux,              // Handler utama server
	}

	// Menjalankan server
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// Kesimpulan:
// Kode ini mendemonstrasikan mekanisme redirect pada HTTP server Go menggunakan http.Redirect, baik untuk redirect internal antar endpoint maupun redirect ke website eksternal. Penggunaan http.ServeMux memungkinkan pengelolaan beberapa route dalam satu server, sementara status HTTP 307 (Temporary Redirect) memastikan metode request tetap dipertahankan saat proses redirect.
