package petfinder

import (
	"encoding/json"
	"fmt"
	"github.com/friendsofgo/errors"
	"net/http"
	"net/url"
)

type AuthResponse struct {
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	AccessToken string `json:"access_token"`
}

func (c *Client) Authenticate() error {
	res, err := http.PostForm(fmt.Sprintf("%s/oauth2/token", c.baseUrl), url.Values{
		"grant_type":    []string{"client_credentials"},
		"client_id":     []string{c.clientId},
		"client_secret": []string{c.clientSecret},
	})
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("received a non-OK status from auth endpoint: %s", res.Status))
	}

	var authResponse AuthResponse
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&authResponse)
	if err != nil {
		return errors.Wrap(err, "decoding json response")
	}

	c.authToken = authResponse.AccessToken

	return nil
}
