package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

const (
	PhotoAPI = "https://api.pexels.com/v1/"
	VideoAPI = "https://api.pexels.com/videos/"
)

type Client struct {
	Token          string
	hc             http.Client
	RemainingTimes int32
}

func NewClient(token string) *Client {
	c := http.Client{}
	return &Client{Token: token, hc: c}
}

type SearchResult struct {
	Page         int32   `json:"page"`
	PerPage      int32   `json:"per_page"`
	TotalResults int32   `json:"total_Results"`
	NextPage     string  `json:"next_page"`
	Photos       []Photo `json:"photos"`
}

type Photo struct {
	Id              int32       `json:"id"`
	Width           int32       `json:"width"`
	Height          int32       `json:"height"`
	Url             string      `json:"url"`
	Photographer    string      `json:"photographer"`
	PhotographerUrl string      `json:"photographer_url"`
	Src             PhotoSource `json:"src"`
}

func main() {
	os.Setenv("PexelsToken", "TIjXAUMb1XrFzEvhq7kWxHPhlZcgOsSBWT2iFUNiXzwvQszRkefqRRyv")
	TOKEN := os.Getenv("PexelsToken")
	var c = NewClient(TOKEN)

	result, err := c.SearchPhotos("waves")

	if err != nil {
		fmt.Errorf("Error searching photos: %v", err)
	}
	if result.Page == 0 {
		fmt.Errorf("No results found")
	}
	fmt.Println(result)
}
