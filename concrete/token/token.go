package token

import (
	"encoding/json"
	"errors"
	"github.com/mohammadhz98/bazaar/iface"
	"github.com/mohammadhz98/bazaar/response/token"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

var (
	refreshURL string = "https://pardakht.cafebazaar.ir/devapi/v2/auth/token/"
	// mutex to handle access token
	mu sync.RWMutex
)

type Token struct {
	bazaar       iface.Bazaar
	access       string
	expire       time.Time
	refresh      string
	clientID     string
	clientSecret string
}

// NewToken creates and return a token with access and refresh input token
func NewToken(bazaar iface.Bazaar) (*Token, error) {
	// create a token with input
	token := &Token{
		bazaar:       bazaar,
		refresh:      bazaar.Conf().Token.Refresh,
		clientID:     bazaar.Conf().Token.ClientID,
		clientSecret: bazaar.Conf().Token.ClientSecret,
	}

	// initialize access token
	if err := token.newAccess(); err != nil {
		return &Token{}, err
	}

	return token, nil
}

// Access return access token
func (t *Token) Access() (string, error) {
	mu.RLock()
	defer mu.RUnlock()

	return t.access, nil
}

// NewAccess uses refresh token to get new access token
func (t *Token) newAccess() error {
	// prevent reading access token while writing on it
	mu.Lock()
	defer mu.Unlock()

	// requests to get new access token
	client := http.Client{}

	// prepare form values
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

	var ref token.NewAccessResponse
	if err = json.NewDecoder(resp.Body).Decode(&ref); err != nil {
		return err
	}

	// return
	if ref.Error != "" {
		return errors.New(ref.Error)
	}

	// set new access token
	t.access = ref.AccessToken
	// ExpiresIn:3,600,000
	t.expire = time.Now().Add(time.Second * time.Duration(ref.ExpiresIn))

	return nil
}
