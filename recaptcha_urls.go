package gog_atu

import (
	"net/url"
)

func ReCaptchaUrl() *url.URL {
	return &url.URL{
		Scheme: HttpsScheme,
		Host:   reCaptchaHost,
		Path:   reCaptchaPath}
}
