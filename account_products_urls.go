// Copyright (c) 2020. All rights reserved.
// Use of this source code is governed by an MIT-style license that can be
// found in the LICENSE file.

package gog_atu

import (
	"net/url"
	"strconv"
)

func DefaultAccountPageUrl(
	page string,
	mt Media) *url.URL {
	return AccountPageUrl(
		page,
		mt,
		AccountSortByPurchaseDate,
		false,
		false)
}

func AccountPageUrl(
	page string,
	mt Media,
	sortOrder AccountSortOrder,
	updated bool, /* get only updated products */
	hidden bool /* get only hidden products */) *url.URL {

	accountPage := &url.URL{
		Scheme: HttpsScheme,
		Host:   GogHost,
		Path:   accountPagePath,
	}

	q := accountPage.Query()
	q.Add("mediaType", strconv.Itoa(int(mt)))
	if sortOrder != "" {
		q.Add("sortBy", string(sortOrder))
	}
	q.Add("page", page)
	if hidden {
		q.Add("hiddenFlag", "1")
	}
	if updated {
		q.Add("isUpdated", "1")
	}
	accountPage.RawQuery = q.Encode()

	return accountPage
}
