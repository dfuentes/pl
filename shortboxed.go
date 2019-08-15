package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"time"
)

const APIDomain = "api.shortboxed.com"

type APIResponse struct {
	Comics []ComicDetails `json:"comics"`
}

type ComicDetails struct {
	Publisher   string `json:"publisher"`
	Description string `json:"description"`
	Title       string `json:"title"`
	Price       string `json:"price"`
	Creators    string `json:"creators"`
	ReleaseDate string `json:"release_date"`
	DiamondID   string `json:"diamond_id"`
}

type Query struct {
	ReleaseDate string
	Publisher   string
	Title       string
	Creators    string
}

type Client struct {
	httpClient *http.Client
}

func NewClient() *Client {
	httpClient := &http.Client{
		Timeout: 10 * time.Second,
	}

	return &Client{
		httpClient: httpClient,
	}
}

func (c *Client) doGet(endpoint string, v interface{}) error {
	response, err := c.httpClient.Get("https://" + path.Join(APIDomain, endpoint))
	if err != nil {
		return err
	}
	defer response.Body.Close()

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(contents, &v)
}

func (c *Client) GetNewReleases() ([]ComicDetails, error) {
	var result APIResponse
	if err := c.doGet("/comics/v1/new", &result); err != nil {
		return nil, err
	}

	return result.Comics, nil
}

func (c *Client) GetPreviousReleases() ([]ComicDetails, error) {
	var result APIResponse
	if err := c.doGet("/comics/v1/previous", &result); err != nil {
		return nil, err
	}

	return result.Comics, nil
}

func (c *Client) GetFutureReleases() ([]ComicDetails, error) {
	var result APIResponse
	if err := c.doGet("/comics/v1/future", &result); err != nil {
		return nil, err
	}

	return result.Comics, nil
}

func (c *Client) Query(q Query) ([]ComicDetails, error) {
	var result APIResponse
	if err := c.doGet("/comics/v1/query"+q.Encode(), &result); err != nil {
		return nil, err
	}

	return result.Comics, nil
}

func (c *Client) GetComicsForDate(date string) ([]ComicDetails, error) {
	var result APIResponse
	if err := c.doGet("/comics/v1/release_date/"+date, &result); err != nil {
		return nil, err
	}

	return result.Comics, nil
}

func (c *Client) GetReleaseDates() ([]string, error) {
	var result struct {
		Dates []string `json:"dates"`
	}
	if err := c.doGet("/comics/v1/releases/available", &result); err != nil {
		return nil, err
	}

	return result.Dates, nil
}

func (q Query) Encode() string {
	v := url.Values{}
	if q.ReleaseDate != "" {
		v.Set("release_date", q.ReleaseDate)
	}
	if q.Creators != "" {
		v.Set("creators", q.Creators)
	}
	if q.Title != "" {
		v.Set("title", q.Title)
	}
	if q.Publisher != "" {
		v.Set("publisher", q.Publisher)
	}
	return v.Encode()
}
