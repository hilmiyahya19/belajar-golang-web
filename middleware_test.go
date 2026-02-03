package belajar_golang_web

import (
	"fmt"      // Untuk mencetak log dan menulis response
	"net/http" // Package utama untuk HTTP server dan middleware
	"testing"  // Package testing untuk menjalankan server via Test
)

// Struct middleware untuk logging sebelum dan sesudah handler dieksekusi
type LogMiddleware struct {
	Handler http.Handler // Handler utama yang akan dibungkus middleware
}

// Implementasi interface http.Handler
func (middleware *LogMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Before Execute Handler")              // Log sebelum handler dijalankan
	middleware.Handler.ServeHTTP(writer, request)      // Meneruskan request ke handler berikutnya
	fmt.Println("After Execute Handler")               // Log setelah handler selesai
}

// Struct middleware untuk menangani error (panic)
type ErrorHandler struct {
	Handler http.Handler // Handler yang akan dibungkus dengan error handler
}

// Implementasi http.Handler untuk error handling
func (errorHandler *ErrorHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	// Menangkap panic agar server tidak crash
	defer func() {
		err := recover() // Mengambil panic jika terjadi
		if err != nil {
			fmt.Println("Terjadi Error")                     // Log error ke console
			writer.WriteHeader(http.StatusInternalServerError) // Set status code 500
			fmt.Fprintf(writer, "Error : %s", err)             // Kirim pesan error ke client
		}
	}()

	// Menjalankan handler berikutnya
	errorHandler.Handler.ServeHTTP(writer, request)
}

// Test untuk menjalankan server dengan middleware
func TestMiddleware(t *testing.T) {
	mux := http.NewServeMux() // Router HTTP

	// Handler root
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Handler Executed")          // Log eksekusi handler
		fmt.Fprint(writer, "Hello Middleware")   // Response ke client
	})

	// Handler /foo
	mux.HandleFunc("/foo", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Foo Executed")              // Log eksekusi handler
		fmt.Fprint(writer, "Hello Foo")          // Response ke client
	})

	// Handler /panic untuk simulasi error
	mux.HandleFunc("/panic", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Panic Executed")            // Log sebelum panic
		panic("Panic")                           // Panic disengaja
	})

	// Middleware logging membungkus mux
	logMiddleware := &LogMiddleware{
		Handler: mux,
	}

	// Middleware error handler membungkus middleware logging
	errorHandler := &ErrorHandler{
		Handler: logMiddleware,
	}

	// Konfigurasi HTTP server
	server := http.Server{
		Addr:    "localhost:8080", // Alamat server
		Handler: errorHandler,     // Entry point handler (middleware terluar)
	}

	// Menjalankan server
	err := server.ListenAndServe()
	if err != nil {
		panic(err) // Panic jika server gagal dijalankan
	}
}

// Kesimpulan:
// Kode ini menunjukkan penerapan middleware di Golang dengan membungkus http.Handler secara berlapis, di mana LogMiddleware digunakan untuk logging sebelum dan sesudah handler dijalankan, sedangkan ErrorHandler berfungsi menangkap panic agar server tidak berhenti secara tiba-tiba. Seluruh request masuk melewati ErrorHandler terlebih dahulu, lalu LogMiddleware, dan akhirnya handler utama, sehingga menghasilkan alur eksekusi middleware yang rapi, aman, dan terstruktur.
