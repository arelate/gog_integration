// Copyright (c) 2020. All rights reserved.
// Use of this source code is governed by an MIT-style license that can be
// found in the LICENSE file.

package gog_atu

import (
	"net/url"
)

func DefaultStorePageUrl(
	page string,
	mt Media) *url.URL {
	return StorePageUrl(page, mt, StoreSortByNewestFirst)
}

func StorePageUrl(
	page string,
	mt Media,
	sortOrder StoreSortOrder) *url.URL {

	storePage := &url.URL{
		Scheme: HttpsScheme,
		Host:   GogHost,
		Path:   storePagePath,
	}

	q := storePage.Query()
	q.Add("mediaType", mt.String())
	if sortOrder != "" {
		q.Add("sort", string(sortOrder))
	}
	q.Add("page", page)
	storePage.RawQuery = q.Encode()

	return storePage
}
