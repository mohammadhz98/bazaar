package token

/*
token used to create a bazaar token api handler.
it implement handler for token related APIs such as creating new access token.
*/

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/mohammadhz98/bazaar/iface"
	"github.com/mohammadhz98/bazaar/response/token"
)

var (
	refreshURL string = "https://pardakht.cafebazaar.ir/devapi/v2/auth/token/"
	// mutex to handle access token
	mu sync.RWMutex
)

// Token include authentication requirement while calling bazaar APIs
type Token struct {
	bazaar       iface.Bazaar
	access       string
	expire       time.Time
	refresh      string
	clientID     string
	clientSecret string
}

func NewToken(bazaar iface.Bazaar) (*Token, error) {
	token := &Token{
		bazaar:       bazaar,
		refresh:      bazaar.Conf().Token.Refresh,
		clientID:     bazaar.Conf().Token.ClientID,
		clientSecret: bazaar.Conf().Token.ClientSecret,
	}

	// initialize access token
	if err := token.refreshToken(); err != nil {
		return &Token{}, err
	}

	return token, nil
}

// Access return access token
//
// it also requests for new access if token is expired
func (t *Token) Access() (token string, err error) {
	mu.RLock()
	defer mu.RUnlock()

	// request new access token if the token is expired
	err = t.refreshTokenIfNeed()

	return t.access, nil
}

// refreshToken uses refresh token to get new access token
func (t *Token) refreshToken() error {
	// prevent reading access token while writing on it
	mu.Lock()
	defer mu.Unlock()

	client := http.Client{}

	form := url.Values{}
	form.Add("grant_type", "refresh_token")
	form.Add("client_id", t.clientID)
	form.Add("client_secret", t.clientSecret)
	form.Add("refresh_token", t.refresh)

	req, _ := http.NewRequest("POST", refreshURL, strings.NewReader(form.Encode()))
	req.PostForm = form
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// unmarshal json response to go struct
	var ref token.NewAccessResponse
	if err = json.NewDecoder(resp.Body).Decode(&ref); err != nil {
		return err
	}

	// return if getting response error
	if resp.StatusCode != 200 {
		return fmt.Errorf("status code %d. %s, %s", resp.StatusCode, ref.Error, ref.ErrorDescription)
	}

	t.access = ref.AccessToken
	// ExpiresIn:3,600,000
	t.expire = time.Now().Add(time.Second * time.Duration(ref.ExpiresIn))

	return nil
}

// refreshTokenIfNeeded request for new token if access token expired
func (t *Token) refreshTokenIfNeed() (err error) {
	// request new access token if the token is expired
	if time.Now().After(t.expire) {
		if err = t.refreshToken(); err != nil {
			return
		}
	}

	return
}
