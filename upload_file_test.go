package belajar_golang_web

import (
	"bytes"                 // Untuk membuat buffer data (dipakai saat test upload)
	_ "embed"                // Untuk embed file ke dalam binary Go
	"fmt"                    // Untuk print output ke console
	"io"                     // Untuk operasi copy data stream
	"mime/multipart"         // Untuk membuat form multipart (upload file)
	"net/http"               // Package utama HTTP server & client
	"net/http/httptest"      // Untuk testing HTTP tanpa server sungguhan
	"os"                     // Untuk operasi file system
	"testing"                // Untuk unit testing
)

// Handler untuk menampilkan form upload
func UploadForm(writer http.ResponseWriter, request *http.Request) {
	// Render template form upload
	myTemplates.ExecuteTemplate(writer, "upload.form.gohtml", nil)
}

// Handler untuk menerima dan memproses file upload
func Upload(writer http.ResponseWriter, request *http.Request) {
	// Mengambil file dari form dengan name="file"
	file, fileHeader, err := request.FormFile("file")
	if err != nil {
		panic(err) // Panic jika file tidak ditemukan / error parsing
	}
	defer file.Close() // Menutup file setelah selesai digunakan

	// Membuat file baru di folder resources sesuai nama file upload
	fileDestination, err := os.Create("./resources/" + fileHeader.Filename)
	if err != nil {
		panic(err) // Panic jika gagal membuat file
	}
	defer fileDestination.Close() // Menutup file tujuan setelah selesai

	// Menyalin isi file upload ke file tujuan
	_, err = io.Copy(fileDestination, file)
	if err != nil {
		panic(err) // Panic jika gagal copy file
	}

	// Mengambil data text dari input name
	name := request.PostFormValue("name")

	// Menampilkan halaman sukses upload
	myTemplates.ExecuteTemplate(writer, "upload.success.gohtml", map[string]interface{}{
		"Name": name,                                // Nama yang diinput user
		"File": "/static/" + fileHeader.Filename,   // URL file untuk diakses via browser
	})
}

// Test manual menggunakan server sungguhan
func TestUpload(t *testing.T) {
	mux := http.NewServeMux() // Router HTTP

	// Route halaman form
	mux.HandleFunc("/", UploadForm)

	// Route upload file
	mux.HandleFunc("/upload", Upload)

	// Static file server untuk mengakses file hasil upload
	mux.Handle("/static/", http.StripPrefix(
		"/static",
		http.FileServer(http.Dir("./resources")),
	))
	// Kode routing static file di atas bekerja dengan cara mendefinisikan sendiri prefix URL `/static/` (bukan bawaan Go) sebagai namespace untuk mengakses file statis, lalu menggunakan http.FileServer untuk membaca file dari folder fisik `./resources`. Karena FileServer hanya memahami path relatif ke direktori yang diberikan, sementara request dari browser masih mengandung prefix `/static`, maka http.StripPrefix digunakan untuk menghapus bagian `/static` dari URL sebelum diteruskan ke FileServer, sehingga path URL dan struktur folder menjadi sesuai dan file dapat ditemukan serta ditampilkan dengan benar.

	// Konfigurasi server
	server := http.Server{
		Addr:    "localhost:8080", // Alamat server
		Handler: mux,              // Handler router
	}

	// Menjalankan server
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// Embed file gambar untuk kebutuhan test upload
//go:embed resources/dottore.png
var uploadFileTest []byte // Data file dalam bentuk byte

// Test upload tanpa menjalankan server sungguhan
func TestUploadFile(t *testing.T) {
	body := new(bytes.Buffer) // Buffer sebagai body request

	// Membuat multipart writer
	writer := multipart.NewWriter(body)

	// Menambahkan field text "name"
	writer.WriteField("name", "Hilmi Yahya")

	// Menambahkan file upload ke form
	file, _ := writer.CreateFormFile("file", "contoh-upload.png")
	file.Write(uploadFileTest) // Menulis data file embed ke form

	writer.Close() // Menutup writer agar form valid

	// Membuat request POST palsu
	request := httptest.NewRequest(
		http.MethodPost,
		"http://localhost:8080/upload",
		body,
	)

	// Set header multipart/form-data
	request.Header.Set("Content-Type", writer.FormDataContentType())

	// Recorder untuk menangkap response
	recorder := httptest.NewRecorder()

	// Memanggil handler Upload langsung
	Upload(recorder, request)

	// Membaca response body
	bodyResponse, _ := io.ReadAll(recorder.Result().Body)

	// Menampilkan response ke console
	fmt.Println(string(bodyResponse))
}

// Kesimpulan:
// Kode ini mendemonstrasikan implementasi upload file di Golang menggunakan net/http dan template HTML, mulai dari menampilkan form upload, memproses file multipart, menyimpan file ke server, hingga menampilkan hasil upload. Selain itu, disertakan dua jenis pengujian, yaitu menjalankan server secara manual dan unit testing menggunakan httptest, sehingga memastikan fitur upload berjalan dengan benar tanpa harus menjalankan server sungguhan.
