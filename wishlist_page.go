// Copyright (c) 2020. All rights reserved.
// Use of this source code is governed by an MIT-style license that can be
// found in the LICENSE file.

package gog_atu

type WishlistPage struct {
	Page
	SortBy          string `json:"sortBy"`
	TotalProducts   int    `json:"totalProducts"`
	ProductsPerPage int    `json:"productsPerPage"`
	// TODO: find data examples where ContentSystemCompatibility is not empty and create a type from that
	ContentSystemCompatibility interface{}    `json:"contentSystemCompatibility"`
	Products                   []StoreProduct `json:"products"`
}

func (wp *WishlistPage) GetProducts() []IdGetter {
	idGetters := make([]IdGetter, 0)
	for _, sp := range wp.Products {
		sp := sp
		idGetters = append(idGetters, &sp)
	}
	return idGetters
}
