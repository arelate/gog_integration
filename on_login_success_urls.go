package gog_atu

import (
	"net/url"
)

func OnLoginSuccessUrl() *url.URL {
	return &url.URL{
		Scheme: HttpsScheme,
		Host:   WwwGogHost,
		Path:   onLoginSuccessPath,
	}
}
