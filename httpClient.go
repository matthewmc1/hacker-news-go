package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"

	client "github.com/bozd4g/go-http-client"
)

type Story struct {
	Id int
}

type Article struct {
	Id    int
	Score int
	Title string
	Text  string
	Type  string
}

func main() {

	httpClient := client.New("https://hacker-news.firebaseio.com/")
	request, err := httpClient.Get("v0/topstories.json")

	if err != nil {
		panic(err)
	}

	resp, err := httpClient.Do(request)

	if err != nil {
		panic(err)
	}

	start := time.Now()
	if resp.Get().StatusCode == 200 {
		var wg sync.WaitGroup
		fmt.Printf("Size of the returned list %d", len(resp.Get().Body))
		for _, story := range resp.Get().Body {
			wg.Add(1) // add one to wait group every time we enter the loop and wait will wait for all these to edn

			go func(x int) { //doing this for multi-threading to do most of the requests in parrallel - MIT 6.824 recommendation
				HackerNews(httpClient, x)
				wg.Done()
			}(int(story)) //assignment to inner function from story to x value if we didn't do this we would pass in an incorrect value
		}
		wg.Wait()
	}

	elapsed := time.Since(start)
	log.Printf("Start time %s, Elapsed time %s", start, elapsed)
}

//function to handle query specific article
func HackerNews(client client.IHttpClient, value int) {
	request, err := client.Get(fmt.Sprintf("/v0/item/%d.json?print=pretty", value))

	fmt.Printf("Request being made for id %d", value)

	if err != nil {
		panic(err)
	}

	resp, err := client.Do(request)

	if err != nil {
		panic(err)
	}

	var article Article
	resp.To(&article)
	fmt.Printf("Hacker news article contains title, %s and text, %s - current score is %d", article.Title, article.Text, article.Score)
}

func defaultHttpRequest() {

	resp, err := http.Get("https://hacker-news.firebaseio.com/v0/topstories.json?print=pretty")

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	fmt.Println(body)
}
