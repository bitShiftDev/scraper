package scraper

import (
	"github.com/go-rod/rod"
)

type PagePool struct {
	pool    rod.Pool[rod.Page]
	browser *rod.Browser
}


func NewPagePool(browser *rod.Browser, maxPages int) *PagePool {
	return &PagePool{
		pool:    rod.NewPagePool(maxPages),
		browser: browser,
	}
}

func (pp *PagePool) Get() *rod.Page {
	create := func() *rod.Page {
		return pp.browser.MustIncognito().MustPage()
	}
	return pp.pool.MustGet(create)
}

func (pp *PagePool) Put(page *rod.Page) {
	pp.pool.Put(page)
}


func (pp *PagePool) Cleanup() {
	pp.pool.Cleanup(func(p *rod.Page) {
		p.MustClose()
	})
}

