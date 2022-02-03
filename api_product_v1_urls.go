package gog_atu

import (
	"net/url"
	"strings"
)

func ApiProductV1Url(id string, mt Media) *url.URL {
	apv1url := &url.URL{
		Scheme: HttpsScheme,
		Host:   apiHost,
		Path:   strings.Replace(apiV1PathTemplate, "{id}", id, 1),
	}

	q := apv1url.Query()
	q.Set("expand", strings.Join(expandValues, ","))
	apv1url.RawQuery = q.Encode()

	return apv1url
}