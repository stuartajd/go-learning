package main

import (
	"encoding/json"
	"fmt"
	"flag"
	"os"
	"io/ioutil"
	"net/http"
)

// Create a custom data structure based off the Reddit API format
type RedditBody struct {
	Data struct {
		Children []struct {
			Data struct {
				Title 	string
				Ups 	int64
				Downs 	int64
			}
		}
	}
}

func main() {
	// Set up a flag for the subreddit name, using floof as default
	var subredditName string
	flag.StringVar(&subredditName, "subreddit", "floof", "Define a subreddit")

	// Parse the flags
	flag.Parse()

	// Create a new instance of RedditBody to parse the JSON into
	var redditBody RedditBody

	// Get the reddit content
	var jsonBody []byte = loadReddit(subredditName)

	// Parse the JSON into the redditBody variable
	err := json.Unmarshal(jsonBody, &redditBody)

	// Check for any unmarshal errors
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}

	// Loop over all the children (posts) within the the returned body, then show a few details.
	for _, child := range redditBody.Data.Children {
		fmt.Println("✉️ ", child.Data.Title, " 👍 ", child.Data.Ups, " 👎 ", child.Data.Downs)
	}
}

// Load the subreddit, returning an array of bytes.
func loadReddit(subreddit string) []byte{
	fmt.Println("Loading Subreddit /r/",subreddit)

	// Create a new HTTP client to send a request from
	client := &http.Client{}

	// Create the request using the subreddit. No body is needed for the request.
	request, _ := http.NewRequest("GET", "https://www.reddit.com/r/"+subreddit+".json", nil)

	// Reddit for some reason didn't like just using a normal GET request, so I've set a custom user agent.
	request.Header.Set("User-Agent", "RedditLookup")

	// Run the request
	response, _ := client.Do(request)

	// Read the entire body of the response
	body, err := ioutil.ReadAll(response.Body)

	// Check for any errors
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}

	// Return the body
	return body
}