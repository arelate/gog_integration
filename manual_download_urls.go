package gog_atu

import (
	"net/url"
)

func ManualDownloadUrl(manualDownload string) *url.URL {
	return &url.URL{
		Scheme: HttpsScheme,
		Host:   WwwGogHost,
		Path:   manualDownload,
	}
}
