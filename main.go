package main

import (
	"fmt"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/joho/godotenv"
)

func setup_env() {
	godotenv.Load()
}

func get_env_value(key string) string {
	return os.Getenv(key)
}

func main() {
	setup_env()
	CONSUMER_KEY := get_env_value("CONSUMER_KEY")
	CONSUMER_SECRET := get_env_value("CONSUMER_SECRET")

	ACCESS_TOKEN := get_env_value("ACCESS_TOKEN")
	ACCESS_TOKEN_SECRET := get_env_value("ACCESS_TOKEN_SECRET")

	config := oauth1.NewConfig(CONSUMER_KEY, CONSUMER_SECRET)
	token := oauth1.NewToken(ACCESS_TOKEN, ACCESS_TOKEN_SECRET)

	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)
	username := "geeky_bhavani"
	search, _, _ := client.PremiumSearch.SearchFullArchive(&twitter.PremiumSearchTweetParams{
		Query:      fmt.Sprintf("from:%s", username),
		FromDate:   "201907010000",
		ToDate:     "201907012359",
		MaxResults: 10,
	}, "searchText")

	for _, s := range search.Results {
		if !s.Retweeted && s.FavoriteCount > 5 {
			fmt.Printf("https://twitter.com/%s/status/%s\n", username, s.IDStr)
		}
	}

}
