package scraper

import (
	"strings"
	"fmt"
	"sync"
	"time"
	"encoding/json"

	"github.com/go-rod/rod"
)

const (
	catalogPageURL = "https://lenta.com/catalog/"
	categoriesAPI  = "/api-gateway/v1/catalog/categories"
)

type apiCategory struct {
	ID          int             `json:"id"`
	Name        string          `json:"name"`
	Slug        string          `json:"slug"`
	ParentID    int             `json:"parentId"`
	ParentName  string          `json:"parentName"`
	Level       int             `json:"level"`
	HasChildren bool            `json:"hasChildren"`
	IsAdult     bool            `json:"isAdult"`
	ImageURL    string          `json:"imageUrl"`
	IconURL     string          `json:"iconUrl"`
	Badges      json.RawMessage `json:"badges"`
}

// apiResponse — структура ответа API Ленты.
type apiResponse struct {
	Categories []apiCategory `json:"categories"`
}

// FetchCategories открывает страницу каталога, перехватывает API-ответ
// с категориями и возвращает распарсенный список
func FetchCategories(page *rod.Page) ([]Category, error) {
	var (
		mu       sync.Mutex
		rawJSON  []byte
		captured bool
	)

	router := page.HijackRequests()
	router.MustAdd("*", func(ctx *rod.Hijack) {
		ctx.MustLoadResponse()
		if strings.Contains(ctx.Request.URL().Path, categoriesAPI) {
			mu.Lock()
			if !captured {
				rawJSON = []byte(ctx.Response.Body())
				captured = true
			}
			mu.Unlock()
		}
	})
	go router.Run()

	
	if err := page.Navigate(catalogPageURL); err != nil {
		router.MustStop()
		return nil, fmt.Errorf("navigate: %w", err)
	}
	
	deadline := time.Now().Add(30 * time.Second)
	for time.Now().Before(deadline) {
	    mu.Lock()
	    done := captured
	    mu.Unlock()
	    if done {
	        break
	    }
		time.Sleep(500 * time.Millisecond)
	}
	router.MustStop()

	if rawJSON == nil {
		return nil, fmt.Errorf("categories API response not captured (endpoint: %s)", categoriesAPI)
	}

	return parseCategories(rawJSON)
}

func parseCategories(data []byte) ([]Category, error) {
	var resp apiResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, fmt.Errorf("unmarshal categories: %w", err)
	}

	out := make([]Category, 0, len(resp.Categories))
	for _, c := range resp.Categories {
		if c.Slug == "" || c.Name == "" {
			continue
		}
		out = append(out, Category{
			Name:     c.Name,
			Slug:     c.Slug,
			URL:      catalogPageURL + c.Slug + "/",
		})
	}

	return out, nil
}