package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
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

type CuratedResult struct {
	Page     int32   `json:"page"`
	PerPage  int32   `json:"per_page"`
	Photos   []Photo `json:"photos"`
	NextPage int32   `json:"next_page"`
}

type PhotoSource struct {
	Original  string `json:"original"`
	Large     string `json:"large"`
	Large2X   string `json:"large2x"`
	Medium    string `json:"medium"`
	Small     string `json:"small"`
	Portrait  string `json:"portrait"`
	Square    string `json:"square"`
	Landscape string `json:"landscape"`
	Tiny      string `json:"tiny"`
}

type VideoSearchResult struct {
	Page         int32   `json:"page"`
	PerPage      int32   `json:"per_page"`
	TotalResults int32   `json:"total_results"`
	NextPage     int32   `json:"next_page"`
	Videos       []Video `json:"videos"`
}

type Video struct {
	Id            int32           `json:"id"`
	Width         int32           `json:"width"`
	Height        int32           `json:"height"`
	Url           string          `json:"url"`
	Image         string          `json:"image"`
	FullRes       interface{}     `json:"full_res"`
	Duration      float64         `json:"duration"`
	VideoFiles    []VideoFiles    `json:"video_files"`
	VideoPictures []VideoPictures `json:"video_pictures"`
}

type PopularVideos struct {
}

type VideoFiles struct {
}

type VideoPictures struct {
}

func (c *Client) SearchPhotos(query string, perPage, page int32) (*SearchResult, error) {
	url := fmt.Sprintf(PhotoAPI+"/search?query=%s&per_page=%d&page=%d", query, perPage, page)
	resp, err := c.requestDoWithAuth("GET", url)
	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}
	var result SearchResult
	err = json.Unmarshal(data, &result)
	return &result, err
}

func (c *Client) CuratePhotos(perPage, page int32) (*CuratedResult, error) {
	url := fmt.Sprintf(PhotoAPI+"curated?per_page=%d&page=%d", perPage, page)
	resp, err := c.requestDoWithAuth("GET", url)
	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var result CuratedResult
	err = json.Unmarshal(data, &result)
	return &result, err
}

func (c *Client) requestDoWithAuth(method, url string) (*http.Response, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", c.Token)
	resp, err := c.hc.Do(req)

	if err != nil {
		return nil, err
	}

	times, err := strconv.Atoi(resp.Header.Get("X-RateLimit-Remaining"))

	if err != nil {
		return resp, nil
	} else {
		c.RemainingTimes = int32(times)
	}
	return resp, nil
}

func (c *Client) GetPhoto(id int32) (*Photo, error) {
	url := fmt.Sprintf(PhotoAPI+"photos/%d", id)
	resp, err := c.requestDoWithAuth("GET", url)
	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result Photo
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) GetRandomPhoto() (*Photo, error) {
	// rand.Seed(time.Now().Unix())
	num, _ := rand.Int(rand.Reader, big.NewInt(1001))
	randNum := int32(num.Int64())
	result, err := c.CuratePhotos(1, randNum)

	if err == nil && len(result.Photos) == 1 {
		return &result.Photos[0], nil
	}
	return nil, err
}

func (c *Client) SearchVideo(query, perPage, page int) (*VideoSearchResult, error) {

}

func (c *Client) PopularVideo(perPage, page int) (*PopulatVideos, error)

func (c *Client) GetRandomVideo() (*Video, error) {

}

func main() {
	os.Setenv("PexelsToken", "TIjXAUMb1XrFzEvhq7kWxHPhlZcgOsSBWT2iFUNiXzwvQszRkefqRRyv")
	TOKEN := os.Getenv("PexelsToken")
	var c = NewClient(TOKEN)

	result, err := c.SearchPhotos("waves", 15, 1)

	if err != nil {
		fmt.Errorf("Error searching photos: %v", err)
	}
	if result.Page == 0 {
		fmt.Errorf("No results found")
	}
	fmt.Println(result)
}
