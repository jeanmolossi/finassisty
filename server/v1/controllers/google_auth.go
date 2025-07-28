// Package controllers holds HTTP handlers for API v1.
package controllers

import (
	"net/http"

	"finassisty/server/config"

	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	oauth2api "google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
)

var googleOAuthConfig = &oauth2.Config{
	ClientID:     config.Env().GoogleOAuth.ClientID,
	ClientSecret: config.Env().GoogleOAuth.ClientSecret,
	RedirectURL:  config.Env().GoogleOAuth.RedirectURL,
	Scopes: []string{
		oauth2api.UserinfoEmailScope,
		oauth2api.UserinfoProfileScope,
	},
	Endpoint: google.Endpoint,
}

// GoogleLogin redirects the user to Google's OAuth consent screen.
func GoogleLogin(c echo.Context) error {
	url := googleOAuthConfig.AuthCodeURL(
		"state-token",
		oauth2.AccessTypeOffline,
	)
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

// GoogleCallback exchanges the OAuth code and returns the user info.
func GoogleCallback(c echo.Context) error {
	ctx := c.Request().Context()
	code := c.QueryParam("code")
	if code == "" {
		return c.JSON(http.StatusBadRequest,
			map[string]string{"error": "code missing"},
		)
	}

	token, err := googleOAuthConfig.Exchange(ctx, code)
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			map[string]string{"error": "token exchange failed"},
		)
	}

	srv, err := oauth2api.NewService(
		ctx,
		option.WithTokenSource(googleOAuthConfig.TokenSource(ctx, token)),
	)
	if err != nil {
		return err
	}

	user, err := srv.Userinfo.Get().Do()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}
