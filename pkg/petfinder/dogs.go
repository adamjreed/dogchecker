package petfinder

import (
	"encoding/json"
	"fmt"
	"github.com/friendsofgo/errors"
	"net/http"
	"net/url"
	"strconv"
)

const DogsPrefix = "dogs:"

type DogsResponse struct {
	Dogs       []*Dog `json:"animals"`
	Pagination `json:"pagination"`
}

type Dog struct {
	Id     int64      `json:"id"`
	Name   string     `json:"name"`
	URL    string     `json:"url"`
	Photos []DogPhoto `json:"photos"`
}

type DogPhoto struct {
	Medium string `json:"medium"`
}

var filters = url.Values{
	"type":     []string{"Dog"},
	"breed":    []string{"Cairn Terrier,Cardigan Welsh Corgi,Corgi,Dachshund,Miniature Dachshund,Pembroke Welsh Corgi,Shih Tzu,Terrier,Wirehaired Dachshund"},
	"size":     []string{"small"},
	"age":      []string{"young,adult"},
	"location": []string{"80223"},
	"distance": []string{"100"},
	"sort":     []string{"recent"},
	"limit":    []string{"100"},
}

func (c *Client) GetDogs() ([]*Dog, error) {
	var dogs []*Dog

	pager := Pager{
		Limit: 100,
		Page:  1,
	}

	moreRecords := true
	for moreRecords {
		dogsResponse, err := c.getDogsByPage(pager)
		if err != nil {
			return nil, err
		}

		dogs = append(dogs, dogsResponse.Dogs...)

		if dogsResponse.Pagination.CurrentPage >= dogsResponse.Pagination.TotalPages {
			moreRecords = false
		}

		pager.Page = pager.Page + 1
	}

	return dogs, nil
}

func (c *Client) getDogsByPage(pager Pager) (*DogsResponse, error) {
	parsedUrl, err := url.Parse(fmt.Sprintf("%s/animals", c.baseUrl))
	if err != nil {
		return nil, err
	}

	filters.Set("limit", strconv.Itoa(pager.Limit))
	filters.Set("page", strconv.Itoa(pager.Page))
	parsedUrl.RawQuery = filters.Encode()

	req, err := http.NewRequest("GET", parsedUrl.String(), nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.authToken))

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("received a non-OK status from auth endpoint: %s", res.Status))
	}

	var dogsResponse DogsResponse
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&dogsResponse)
	if err != nil {
		return nil, errors.Wrap(err, "decoding json response")
	}

	return &dogsResponse, nil
}
