package belajar_golang_web

import (
	"fmt"        // Untuk menulis response teks ke client
	"net/http"   // Package utama untuk HTTP server dan request handling
	"testing"    // Package testing untuk menjalankan fungsi Test
)

// Handler untuk mendownload file dari server
func DownloadFile(writer http.ResponseWriter, request *http.Request) {
	// Mengambil parameter query ?file= dari URL
	file := request.URL.Query().Get("file")

	// Validasi jika parameter file kosong
	if file == "" {
		writer.WriteHeader(http.StatusBadRequest) // Set status code 400
		fmt.Fprint(writer, "Bad Request")          // Tulis pesan error
		return                                    // Hentikan eksekusi handler
	}

	// Menambahkan header agar browser menganggap response sebagai file download
	writer.Header().Add(
		"Content-Disposition",
		"attachment; filename=\""+file+"\"",
	)

	// Mengirimkan file dari folder resources ke client
	http.ServeFile(writer, request, "./resources/"+file)
}

// Test untuk menjalankan HTTP server secara manual
func TestDownloadFile(t *testing.T) {
	// Konfigurasi server HTTP
	server := http.Server{
		Addr: "localhost:8080",                 // Alamat dan port server
		Handler: http.HandlerFunc(DownloadFile), // Handler langsung ke fungsi DownloadFile
	}

	// Menjalankan server
	err := server.ListenAndServe()
	if err != nil {
		panic(err) // Panic jika server gagal dijalankan
	}
}

// Kesimpulan:
// Kode ini mengimplementasikan fitur download file di Golang dengan memanfaatkan query parameter URL untuk menentukan nama file yang akan diunduh, memvalidasi input agar tidak kosong, lalu menggunakan header Content-Disposition supaya browser memicu proses download. File dikirimkan langsung dari server menggunakan http.ServeFile, dan server dijalankan secara manual melalui fungsi test untuk mempermudah pengujian.
