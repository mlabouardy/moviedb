package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"golang.org/x/net/html"
)

const (
	URL             = "https://www.themoviedb.org/%s/%s"
	IMG_TAG         = "img"
	CLASS_ATTRIBUTE = "class"
	ITEM_CLASS      = "poster lazyload fade"
)

type Show struct {
	Cover string
	Title string
}

type MovieDB struct {
	client *http.Client
}

func NewMovieDB() *MovieDB {
	return &MovieDB{
		client: &http.Client{},
	}
}

func (m MovieDB) fetchShows(showType string, baseURL string) []Show {
	url := fmt.Sprintf(URL, showType, baseURL)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := m.client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	return crawl(resp.Body)
}

func (m MovieDB) GetPopularMovies() []Show {
	return m.fetchShows("movie", "")
}

func (m MovieDB) GetUpcomingMovies() []Show {
	return m.fetchShows("movie", "upcoming")
}

func (m MovieDB) GetTopRatedMovies() []Show {
	return m.fetchShows("movie", "top-rated")
}

func (m MovieDB) GetNowPlayingMovies() []Show {
	return m.fetchShows("movie", "now-playing")
}

func (m MovieDB) GetPopularTVShows() []Show {
	return m.fetchShows("tv", "")
}

func (m MovieDB) GetTopRatedShows() []Show {
	return m.fetchShows("tv", "top-rated")
}

func (m MovieDB) GetOnTVShows() []Show {
	return m.fetchShows("tv", "on-the-air")
}

func (m MovieDB) GetAiringTodayShows() []Show {
	return m.fetchShows("tv", "airing-today")
}

func crawl(body io.Reader) []Show {
	shows := make([]Show, 0)
	tokenizer := html.NewTokenizer(body)
	for {
		token := tokenizer.Next()
		switch token {
		case html.ErrorToken:
			return shows
		case html.StartTagToken:
			t := tokenizer.Token()
			if t.Data == IMG_TAG {
				for _, attr := range t.Attr {
					if attr.Key == CLASS_ATTRIBUTE && attr.Val == ITEM_CLASS {
						show := Show{
							Cover: t.Attr[2].Val,
							Title: t.Attr[4].Val,
						}
						shows = append(shows, show)
					}
				}
			}
		}
	}
}

func main() {
	db := NewMovieDB()
	fmt.Println("upcoming movies")
	for _, show := range db.GetUpcomingMovies() {
		fmt.Println(show)
	}
	fmt.Println("popular movies")
	for _, show := range db.GetPopularMovies() {
		fmt.Println(show)
	}
	fmt.Println("top rated movies")
	for _, show := range db.GetTopRatedMovies() {
		fmt.Println(show)
	}
	fmt.Println("now playing movies")
	for _, show := range db.GetNowPlayingMovies() {
		fmt.Println(show)
	}
	fmt.Println("top rated shows")
	for _, show := range db.GetTopRatedShows() {
		fmt.Println(show)
	}
	fmt.Println("airing today shows")
	for _, show := range db.GetAiringTodayShows() {
		fmt.Println(show)
	}
	fmt.Println("on tv shows")
	for _, show := range db.GetOnTVShows() {
		fmt.Println(show)
	}
	fmt.Println("popular shows")
	for _, show := range db.GetPopularTVShows() {
		fmt.Println(show)
	}
}
