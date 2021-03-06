package auth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/google/uuid"
	"github.com/gorilla/schema"
	"github.com/pkg/browser"
)

const port = ":3000"

func openTwitchAuth(client_id, redirect_uri string, scopes []string, state string) error {
	scope := url.QueryEscape(strings.Join(scopes, " "))
	return browser.OpenURL(fmt.Sprintf(strings.Replace(
		`https://id.twitch.tv/oauth2/authorize
			?response_type=token
			&client_id=%s
			&redirect_uri=%s
			&scope=%s
			&state=%s`,
		"\n", "", -1),
		client_id, redirect_uri, scope, state))
}

type TokenValidationResponse struct {
	ClientID  string   `json:"client_id"`
	Login     string   `json:"login"`
	Scopes    []string `json:"scopes"`
	UserID    string   `json:"user_id"`
	ExpiresIn int64    `json:"expires_in"`
}

func Validate(token_type, token string) (TokenValidationResponse, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://id.twitch.tv/oauth2/validate", nil)

	req.Header.Add("Authorization", fmt.Sprintf(`%s %s`, token_type, token))

	resp, err := client.Do(req)
	data := TokenValidationResponse{}
	if err != nil {
		return data, err
	}
	err = json.NewDecoder(resp.Body).Decode(&data)
	resp.Body.Close()
	if err != nil {
		return data, err
	}

	return data, nil
}

type AuthResult struct {
	AccessToken string   `schema:"access_token"`
	Scope       []string `schema:"scope"`
	State       string   `schema:"state"`
	TokenType   string   `schema:"token_type"`
}

func GetToken(client_id string, scopes []string) (AuthResult, error) {
	redirect_path := "/redirect"
	redirect_uri := "http://localhost" + port + redirect_path
	auth_path := "/auth"

	state := uuid.New().String()
	go func() {
		openTwitchAuth(client_id, redirect_uri, scopes, state)
	}()
	server := &http.Server{
		Addr: port,
	}

	result := AuthResult{}
	var returnError error = nil
	http.HandleFunc(auth_path, func(w http.ResponseWriter, r *http.Request) {
		go func() {
			if err := server.Shutdown(context.Background()); err != nil {
				returnError = err
			}
		}()
		decoder := schema.NewDecoder()
		err := decoder.Decode(&result, r.URL.Query())
		if err != nil {
			returnError = err
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte(`<script>
		window.close()
		</script>`))
	})
	http.HandleFunc(redirect_path, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte(fmt.Sprintf(`<script>
		// console.log(document.location.hash.slice(1))
		document.location = "%s?"+document.location.hash.slice(1)
		</script>`, auth_path)))
	})
	server.ListenAndServe()
	if result.TokenType == "bearer" {
		result.TokenType = "Bearer"
	}
	if result.TokenType == "oauth" {
		result.TokenType = "OAuth"
	}
	return result, returnError
}

func RevokeToken(client_id, token string) error {
	resp, err := http.Post("https://id.twitch.tv/oauth2/revoke",
		"application/x-www-form-urlencoded",
		strings.NewReader(
			fmt.Sprintf(
				"client_id=%s&token=%s",
				client_id, token)))

	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New("wrong token")
	}
	return nil
}
