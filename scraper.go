package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
)

type Scraper struct {
	BaseURL   string
	Collector *colly.Collector
	DB        *Database
}

func NewScraper(baseURL string, db *Database) *Scraper {
	c := colly.NewCollector(
		colly.AllowedDomains("komiku.id", "data.komiku.id"),
		colly.Async(true),
	)

	c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 2,
		Delay:       2 * time.Second,
	})

	return &Scraper{
		BaseURL:   baseURL,
		Collector: c,
		DB:        db,
	}
}

func (s *Scraper) GetAllComics(filter string) ([]Comic, error) {
	var comics []Comic
	url := s.BaseURL + "daftar-komik/"
	if filter != "" {
		url = s.BaseURL + "daftar-komik/?tipe=" + filter
	}

	s.Collector.OnHTML("div.ls4", func(e *colly.HTMLElement) {
		var comic Comic
		e.ForEach("div.ls4v", func(i int, element *colly.HTMLElement) {
			comic.Endpoint = element.ChildAttr("a", "href")
			comic.Image = strings.TrimSuffix(element.ChildAttr("img", "data-src"), "?resize=240,170")
		})
		e.ForEach("div.ls4j", func(i int, element *colly.HTMLElement) {
			comic.Title = element.ChildText("h4")
		})
		comic.Type = filter
		comics = append(comics, comic)
	})

	err := s.Collector.Visit(url)
	if err != nil {
		return nil, err
	}

	s.Collector.Wait()
	return comics, nil
}

func (s *Scraper) GetComicInfo(endpoint string) (*Comic, error) {
	var comicInfo Comic
	url := s.BaseURL + endpoint

	s.Collector.OnHTML("table.inftable", func(e *colly.HTMLElement) {
		e.ForEach("tbody", func(i int, element *colly.HTMLElement) {
			comicInfo.Title = element.ChildText("tr:nth-child(1) > td:nth-child(2)")
			comicInfo.Type = element.ChildText("tr:nth-child(2) > td:nth-child(2)")
			comicInfo.Author = element.ChildText("tr:nth-child(4) > td:nth-child(2)")
			comicInfo.Status = element.ChildText("tr:nth-child(5) > td:nth-child(2)")
			comicInfo.Rating = element.ChildText("tr:nth-child(6) > td:nth-child(2)")
		})
	})

	s.Collector.OnHTML("ul.genre", func(e *colly.HTMLElement) {
		e.ForEach("li.genre", func(i int, element *colly.HTMLElement) {
			comicInfo.Genre = append(comicInfo.Genre, element.Text)
		})
	})

	s.Collector.OnHTML("div.ims", func(e *colly.HTMLElement) {
		comicInfo.Thumbnail = strings.TrimSuffix(e.ChildAttr("img", "src"), "?w=225&quality=60")
	})

	s.Collector.OnHTML("div.desc", func(e *colly.HTMLElement) {
		comicInfo.Desc = e.Text
	})

	comicInfo.Endpoint = endpoint
	err := s.Collector.Visit(url)
	if err != nil {
		return nil, err
	}

	s.Collector.Wait()
	return &comicInfo, nil
}

func (s *Scraper) GetChapterList(endpoint string) ([]Chapter, error) {
	var chapters []Chapter
	url := s.BaseURL + endpoint

	s.Collector.OnHTML("tbody._3Rsjq", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(i int, element *colly.HTMLElement) {
			var chapter Chapter
			if element.ChildText("td.judulseries") != "" {
				chapter.Endpoint = element.ChildAttr("a", "href")
				chapter.Name = element.ChildText("td.judulseries")
				chapters = append(chapters, chapter)
			}
		})
	})

	err := s.Collector.Visit(url)
	if err != nil {
		return nil, err
	}

	s.Collector.Wait()
	return chapters, nil
}

func (s *Scraper) GetChapterDetail(endpoint string) (*ChapterDetail, error) {
	var detail ChapterDetail
	url := s.BaseURL + endpoint

	s.Collector.OnHTML("section[id=Baca_Komik]", func(e *colly.HTMLElement) {
		imageList := e.ChildAttrs("img", "src")
		detail.Images = imageList
	})

	s.Collector.OnHTML("header[id=Judul]", func(e *colly.HTMLElement) {
		if e.Index == 0 {
			detail.Title = e.ChildText("h1")
		}
	})

	err := s.Collector.Visit(url)
	if err != nil {
		return nil, err
	}

	s.Collector.Wait()
	return &detail, nil
}

func (s *Scraper) GetNewestComics(page int) ([]Comic, error) {
	var comics []Comic
	url := fmt.Sprintf("https://data.komiku.id/pustaka/")
	if page != 1 {
		url = fmt.Sprintf("https://data.komiku.id/pustaka/page/%d/", page)
	}

	s.Collector.OnHTML("div.bge", func(e *colly.HTMLElement) {
		var comic Comic
		comic.Desc = e.ChildText("p")
		e.ForEach("div.bgei", func(i int, e2 *colly.HTMLElement) {
			comic.Image = strings.TrimSuffix(e2.ChildAttr("img", "data-src"), "?resize=450,235&quality=60")
			comic.Endpoint = strings.Replace(e2.ChildAttr("a", "href"), s.BaseURL, "", 1)
			comic.Type = e2.ChildText("b")
			comic.Endpoint = "/" + comic.Endpoint
		})
		e.ForEach("div.kan", func(i int, e2 *colly.HTMLElement) {
			comic.Title = e2.ChildText("h3")
		})
		comics = append(comics, comic)
	})

	err := s.Collector.Visit(url)
	if err != nil {
		return nil, err
	}

	s.Collector.Wait()
	return comics, nil
}
