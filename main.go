package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			fmt.Fprintf(w, "Mendapatkan daftar order")
		case "POST":
			fmt.Fprintf(w, "Membuat order baru")
		default:
			http.Error(w, "Metode tidak diperbolehkan.", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/orders/:orderId", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "PUT":
			fmt.Fprintf(w, "Memperbarui order yang ada")
		case "DELETE":
			fmt.Fprintf(w, "Menghapus order")
		default:
			http.Error(w, "Metode tidak diperbolehkan.", http.StatusMethodNotAllowed)
		}
	})

	http.ListenAndServe(":8080", nil)
}
