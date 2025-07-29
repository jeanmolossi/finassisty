// Package controllers holds HTTP handlers for API v1.
package controllers

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"net/http"

	"finassisty/server/config"

	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	oauth2api "google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
)

const (
	stateByteLen = 16
	emptyStr     = ""
)

func newGoogleOAuthConfig() (*oauth2.Config, error) {
	creds := config.Env().GoogleOAuth
	if creds.ClientID == emptyStr {
		return nil, errors.New("google client id missing")
	}
	if creds.ClientSecret == emptyStr {
		return nil, errors.New("google client secret missing")
	}
	if creds.RedirectURL == emptyStr {
		return nil, errors.New("google redirect url missing")
	}

	cfg := &oauth2.Config{
		ClientID:     creds.ClientID,
		ClientSecret: creds.ClientSecret,
		RedirectURL:  creds.RedirectURL,
		Scopes: []string{
			oauth2api.UserinfoEmailScope,
			oauth2api.UserinfoProfileScope,
		},
		Endpoint: google.Endpoint,
	}

	return cfg, nil
}

func randomState() (string, error) {
	b := make([]byte, stateByteLen)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}

// GoogleLogin redirects the user to Google's OAuth consent screen.
func GoogleLogin(c echo.Context) error {
	oauthCfg, err := newGoogleOAuthConfig()
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"error": "oauth not configured"})
	}

	state, err := randomState()
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"error": "could not start oauth"})
	}

	cookie := &http.Cookie{
		Name:     "oauth_state",
		Value:    state,
		Path:     "/",
		HttpOnly: true,
		Secure:   c.IsTLS(),
		SameSite: http.SameSiteLaxMode,
	}
	c.SetCookie(cookie)

	url := oauthCfg.AuthCodeURL(state, oauth2.AccessTypeOffline)
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

// GoogleCallback exchanges the OAuth code and returns the user info.
func GoogleCallback(c echo.Context) error {
	oauthCfg, err := newGoogleOAuthConfig()
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"error": "oauth not configured"})
	}

	state := c.QueryParam("state")
	cookie, cErr := c.Cookie("oauth_state")
	if state == "" || cErr != nil || cookie.Value != state {
		return c.JSON(http.StatusBadRequest,
			map[string]string{"error": "invalid state"})
	}

	// clear state cookie
	c.SetCookie(&http.Cookie{
		Name:   "oauth_state",
		Value:  emptyStr,
		Path:   "/",
		MaxAge: -1,
	})

	code := c.QueryParam("code")
	if code == emptyStr {
		return c.JSON(http.StatusBadRequest,
			map[string]string{"error": "code missing"})
	}

	ctx := c.Request().Context()
	token, err := oauthCfg.Exchange(ctx, code)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest,
			map[string]string{"error": "token exchange failed"})
	}

	srv, err := oauth2api.NewService(
		ctx,
		option.WithTokenSource(oauthCfg.TokenSource(ctx, token)),
	)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"error": "failed to get user info"})
	}

	user, err := srv.Userinfo.Get().Do()
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"error": "failed to get user info"})
	}

	return c.JSON(http.StatusOK, user)
}
