package main

import (
	"fmt"
	"os"
)

var ApiKey = os.Getenv("NEWS_APIKEY")

const BASE_URL_V2 = "https://newsapi.org/v2"
const TOP_HEADLINES_URL = BASE_URL_V2 + "/top-headlines"
const EVERYTHING_URL = BASE_URL_V2 + "/everything"

func topHeadlinesURL() string {
	return url(TOP_HEADLINES_URL)
}

func everythingURL() string {
	return url(EVERYTHING_URL)
}

func url(path string) string {
	return fmt.Sprintf("%s?apiKey=%s", path, ApiKey)
}
