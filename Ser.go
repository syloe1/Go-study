package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

const uploadPath = "./uploads"

func main() {
	http.HandleFunc("/upload", uploadHandler)
	http.HandleFunc("download/", downloadHandler)

	if err := os.MkdirAll(uploadPath, os.ModePerm); err != nil {
		fmt.Println("Failed to create upload directory!", err)
		return
	}
	fmt.Println("Server is running on: 8080")
	http.ListenAndServe(":8080", nil)
}
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Failed to get file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	filePath := uploadPath + "/" + "uploaded_file" // You can generate unique names
	out, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "Failed to save file", http.StatusInternalServerError)
		return
	}
	defer out.Close()

	if _, err := io.Copy(out, file); err != nil {
		http.Error(w, "Failed to save file", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "File uploaded successfully")
}

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	fileName := r.URL.Path[len("/download/"):]
	filePath := uploadPath + "/" + fileName

	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	defer file.Close()

	w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
	w.Header().Set("Content-Type", "application/octet-stream")
	if _, err := io.Copy(w, file); err != nil {
		http.Error(w, "Failed to send file", http.StatusInternalServerError)
	}
}
