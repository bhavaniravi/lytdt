package main

import (
	"fmt"
	"os"
	"time"
	"io/ioutil"
	"strings"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/joho/godotenv"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}


func setup_env() {
	godotenv.Load()
}

func get_env_value(key string) string {
	return os.Getenv(key)
}


func listTweets(){
	setup_env()
	CONSUMER_KEY := get_env_value("CONSUMER_KEY")
	CONSUMER_SECRET := get_env_value("CONSUMER_SECRET")

	ACCESS_TOKEN := get_env_value("ACCESS_TOKEN")
	ACCESS_TOKEN_SECRET := get_env_value("ACCESS_TOKEN_SECRET")

	config := oauth1.NewConfig(CONSUMER_KEY, CONSUMER_SECRET)
	token := oauth1.NewToken(ACCESS_TOKEN, ACCESS_TOKEN_SECRET)

	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)
	username := getUsername()
	last_year  := time.Now().AddDate(-1, 0, 0).Format("20060102")
	search, res, _:= client.PremiumSearch.SearchFullArchive(&twitter.PremiumSearchTweetParams{
		Query:      fmt.Sprintf("from:%s", username),
		FromDate:   fmt.Sprintf("%s0000", last_year),
		ToDate:     fmt.Sprintf("%s2359", last_year),
		MaxResults: 10,
	}, "searchText")

	if res.StatusCode == 429 {
		fmt.Printf("Limit Exceeded for today")
		return 
	}

	for _, s := range search.Results {
		if !s.Retweeted && s.FavoriteCount > 5 {
			fmt.Printf("https://twitter.com/%s/status/%s\n", username, s.IDStr)
		}
	}
}

func addUsername(username string) {
	f, err := os.Create("/tmp/lytdt.txt")
	check(err)
	f.WriteString(username)
}

func getUsername() string {
	data, _ := ioutil.ReadFile("/tmp/lytdt.txt")
	return strings.TrimSuffix(string(data), "\n")
}

func main() {
	option := "default"
	if len(os.Args) >= 1 {
		option = os.Args[1]
	}

	switch ; option {
	case "add":
		if len(os.Args) < 2 {
			fmt.Printf("Supported format is `./main.go add <username>`")
		} else {
			username := os.Args[2]
			addUsername(username)
		}
	case "list":
		listTweets()
	default:
		fmt.Printf("Supported options are 1. add 2. list \n `./main.go add <username>` \n `./main.go list`")
	}
}
