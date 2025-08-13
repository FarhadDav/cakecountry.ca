package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

// TODO extract common logic for rendering templates
func instructions(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	templates := []string{
		"./ui/html/base.tmpl",
		"./ui/html/pages/instructions.tmpl",
	}

	ts, err := template.ParseFiles(templates...)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Parse template
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func menu(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	templates := []string{
		"./ui/html/base.tmpl",
		"./ui/html/pages/menu.tmpl",
	}

	ts, err := template.ParseFiles(templates...)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Parse template
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func pricing(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	templates := []string{
		"./ui/html/base.tmpl",
		"./ui/html/pages/pricing.tmpl",
	}

	ts, err := template.ParseFiles(templates...)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func weddings(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	templates := []string{
		"./ui/html/base.tmpl",
		"./ui/html/pages/weddings.tmpl",
	}

	ts, err := template.ParseFiles(templates...)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Parse template
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	templates := []string{
		"./ui/html/base.tmpl",
		"./ui/html/pages/home.tmpl",
	}

	ts, err := template.ParseFiles(templates...)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Scan carousel images folder
	imgDir := "./ui/static/images/carousel"
	files, err := os.ReadDir(imgDir)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	var imgURLs []string
	for _, file := range files {
		if !file.IsDir() {
			ext := strings.ToLower(filepath.Ext(file.Name()))
			if ext == ".jpg" || ext == ".jpeg" || ext == ".png" || ext == ".gif" {
				imgURLs = append(imgURLs, "/static/images/carousel/"+file.Name())
			}
		}
	}

	// Pass image URLs to template
	data := struct {
		CarouselImgs []string
	}{
		CarouselImgs: imgURLs,
	}

	err = ts.ExecuteTemplate(w, "base", data)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
