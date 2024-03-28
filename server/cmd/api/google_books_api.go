package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"os"
)

const googleBooksAPIURL = "https://www.googleapis.com/books/v1/volumes"

type googleBooksResponse struct {
	Items []struct {
		VolumeInfo struct {
			Title               string   `json:"title"`
			Authors             []string `json:"authors"`
			Description         string   `json:"description"`
			PublishedDate       string   `json:"publishedDate"`
			PageCount           int      `json:"pageCount"`
			Categories          []string `json:"categories"`
			IndustryIdentifiers []struct {
				Type       string `json:"type"`
				Identifier string `json:"identifier"`
			} `json:"industryIdentifiers"`
		} `json:"volumeInfo"`
	} `json:"items"`
}

func (app *application) searchBooksHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if query == "" {
		app.badRequestResponse(w, r, errors.New("missing search query parameter"))
		return
	}

	booksResponse, err := app.searchBooks(query)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"success": true, "data": booksResponse.Items}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

// I think this should move somewhere else but I'm not sure where
// ... it's not a database operation, but it's not really a handler either

func (app *application) searchBooks(query string) (*googleBooksResponse, error) {
	escapedQuery := url.QueryEscape(query)
	url := googleBooksAPIURL + "?q=" + escapedQuery + "&key=" + os.Getenv("PAGEWISE_GOOGLE_BOOKS_API_KEY")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := app.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("unexpected status code")
	}

	var booksResponse googleBooksResponse
	err = json.NewDecoder(resp.Body).Decode(&booksResponse)
	if err != nil {
		return nil, err
	}

	return &booksResponse, nil
}
