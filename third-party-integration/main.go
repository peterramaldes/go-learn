package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	_, e := findArticles(topHeadlinesURL())
	fmt.Printf("%+v\n", e)
}

func findArticles(url string) (Response, Error) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode > 299 {
		e := Error{}
		err := json.NewDecoder(res.Body).Decode(&e)
		if err != nil {
			log.Fatal(err)
		}
		return Response{}, e
	}

	r := Response{}
	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		log.Fatal(err)
	}

	err = res.Body.Close()
	return r, Error{}
}
