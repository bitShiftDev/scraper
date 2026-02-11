package scraper

import (
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

type Config struct {
	Headless   bool
	SlowMotion int  // миллисекунды между действиями (0 = без задержки)
	Devtools   bool
	ProxyURL   string
}


func NewBrowser(cfg Config) (*rod.Browser, error) {
	l := launcher.New().
		Headless(cfg.Headless).
		Devtools(cfg.Devtools)

	if cfg.ProxyURL != "" {
		l = l.Proxy(cfg.ProxyURL)
	}

	controlURL, err := l.Launch()
	if err != nil {
		return nil, err
	}

	b := rod.New().ControlURL(controlURL)

	if cfg.SlowMotion > 0 {
		b = b.SlowMotion(time.Duration(cfg.SlowMotion) * time.Millisecond)
	}

	if err := b.Connect(); err != nil {
		return nil, err
	}

	return b, nil
}