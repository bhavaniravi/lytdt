package main

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
	"github.com/dghubble/go-twitter/twitter"
    "github.com/dghubble/oauth1"
)

func setup_env(){
	godotenv.Load()
}

func get_env_value(key string) string{
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
	
	search, _, _ := client.Search.Tweets(&twitter.SearchTweetParams{
		Query: "from:@geeky_bhavani",
	})
	
	for  i, s := range search.Statuses {
		fmt.Println(i, s.IDStr, s.Text)
	}


}