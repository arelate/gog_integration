// Copyright (c) 2020. All rights reserved.
// Use of this source code is governed by an MIT-style license that can be
// found in the LICENSE file.

package gog_atu

type Page struct {
	TotalPages
	Page int `json:"page"`
}

type TotalPages struct {
	TotalPages int `json:"totalPages"`
}
