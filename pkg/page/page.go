package page

import (
	"net/http"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

func Main(r *http.Request) ([]byte, int, error) {
	p, err := loadPage("content/main.html")
	if err != nil {
		return []byte("Not found main.html"), http.StatusInternalServerError, err
	}
	return p.Body, http.StatusOK, nil
}

func loadPage(title string) (*Page, error) {
	filename := title
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}
