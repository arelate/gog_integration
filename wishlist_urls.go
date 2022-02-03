// Copyright (c) 2020. All rights reserved.
// Use of this source code is governed by an MIT-style license that can be
// found in the LICENSE file.

package gog_atu

import (
	"net/url"
	"strconv"
	"strings"
)

func DefaultWishlistPageUrl(
	page string,
	mt Media) *url.URL {
	return WishlistPageUrl(page, mt, WishlistSortByDateAdded, false)
}

func WishlistPageUrl(
	page string,
	mt Media,
	sortOrder WishlistSortOrder,
	hidden bool) *url.URL {

	wishlistPage := &url.URL{
		Scheme: HttpsScheme,
		Host:   GogHost,
		Path:   wishlistSearchPath,
	}

	q := wishlistPage.Query()
	q.Add("mediaType", strconv.Itoa(int(mt)))
	if sortOrder != "" {
		q.Add("sortBy", string(sortOrder))
	}
	if hidden {
		q.Add("hiddenFlag", "1")
	}
	q.Add("page", page)
	wishlistPage.RawQuery = q.Encode()

	return wishlistPage
}

func AddToWishlistUrl(id string) *url.URL {
	return &url.URL{
		Scheme: HttpsScheme,
		Host:   GogHost,
		Path:   strings.Replace(wishlistAddPathTemplate, "{id}", id, 1),
	}
}

func RemoveFromWishlistUrl(id string) *url.URL {
	return &url.URL{
		Scheme: HttpsScheme,
		Host:   GogHost,
		Path:   strings.Replace(wishlistRemovePathTemplate, "{id}", id, 1),
	}
}
