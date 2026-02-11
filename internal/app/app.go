package app

import (
	"context"
	"fmt"
	"log/slog"
	
	"scraper/internal/logger"
	"scraper/internal/scraper"
)

func Run(ctx context.Context) error {
	
	opts := logger.Options{
		Level: slog.LevelDebug,
		AddSource: false,
	}
	
	log := logger.New(opts)
	log.Info("app works!")
	
	browserConfig := scraper.Config{
		Headless: true,
	}
	
	
	b, err := scraper.NewBrowser(browserConfig)
	if err != nil {
		log.Error("browser launch: ", err)
	}
	defer b.MustClose()
	
	
	pagePool := scraper.NewPagePool(b, 5)
	page := pagePool.Get()
	
	categories, err := scraper.FetchCategories(page)
	if err != nil {
		log.Error("Fetch categories err: ", err)
		return err
	}
	
	for _, v := range categories {
		fmt.Println(v)
	}
	
	
	return nil
}