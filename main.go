package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/joho/godotenv"
)

var filePath string = "/tmp/lytdt.txt"

// How to load .env variables
func setupEnv() {
	godotenv.Load()
}

// os.environ equivalent in python
func getEnvValue(key string) string {
	return os.Getenv(key)
}

// how to read data from a file
func readFromFile() string {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("File reading error", err)
		return ""
	}
	return string(data)
}

// how to write data to a file
// how to convert string to byte
func wrtieFile(data string) {
	strData := []byte(data)
	err := ioutil.WriteFile(filePath, strData, 0644)
	if err != nil {
		fmt.Println(err)
	}

}

func getTweets(username string) {
	consumerKey := getEnvValue("CONSUMER_KEY")
	consumerSecret := getEnvValue("CONSUMER_SECRET")

	accessToken := getEnvValue("ACCESS_TOKEN")
	accessTokenSecret := getEnvValue("ACCESS_TOKEN_SECRET")

	// How to use go twitter library
	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessTokenSecret)

	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	// Date formatting YeeYaah!
	fromDate := time.Now().AddDate(-1, 0, -1).Format("20060102")
	toDate := time.Now().AddDate(-1, 0, 0).Format("20060102")
	fromDate = fmt.Sprintf("%s0000", fromDate)
	toDate = fmt.Sprintf("%s2359", toDate)

	// Using twitter search SDK
	search, _, _ := client.PremiumSearch.SearchFullArchive(&twitter.PremiumSearchTweetParams{
		Query: fmt.Sprintf("from:%s", username),
		// Starting  12 am to 12 pm
		FromDate:   fromDate,
		ToDate:     toDate,
		MaxResults: 100,
	}, "searchText")

	// Looping through and ignoring favcount
	for _, s := range search.Results {
		if !s.Retweeted && s.FavoriteCount > 10 {
			fmt.Printf("https://twitter.com/%s/status/%s\n", username, s.IDStr)
		}
	}
}

func main() {
	// Hey switch, welcome back after 4 years of Python
	switch command := os.Args[1]; command {
	case "add":
		wrtieFile(os.Args[2])
	case "list":
		setupEnv()
		username := readFromFile()
		getTweets(username)
	}
}
