// Copyright (c) 2020. All rights reserved.
// Use of this source code is governed by an MIT-style license that can be
// found in the LICENSE file.

package gog_integration

import (
	"net/url"
	"strings"
)

func DetailsUrl(id string, mt Media) *url.URL {
	path := strings.Replace(detailsPathTemplate, "{mediaType}", mt.String(), 1)
	return &url.URL{
		Scheme: HttpsScheme,
		Host:   GogHost,
		Path:   strings.Replace(path, "{id}", id, 1),
	}
}
