// Copyright (c) 2020. All rights reserved.
// Use of this source code is governed by an MIT-style license that can be
// found in the LICENSE file.

package gog_integration

type StorePage struct {
	Page
	Products []StoreProduct `json:"products"`
	// TODO: find data examples where Ts is not empty and create a type from that
	Ts               interface{} `json:"ts"`
	TotalResults     string      `json:"totalResults"`
	TotalGamesFound  int         `json:"totalGamesFound"`
	TotalMoviesFound int         `json:"totalMoviesFound"`
}

func (spp *StorePage) GetProducts() []IdGetter {
	idGetters := make([]IdGetter, 0)
	for _, sp := range spp.Products {
		sp := sp
		idGetters = append(idGetters, &sp)
	}
	return idGetters
}
