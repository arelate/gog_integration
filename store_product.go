// Copyright (c) 2020. All rights reserved.
// Use of this source code is governed by an MIT-style license that can be
// found in the LICENSE file.

package gog_integration

import "strconv"

type StoreProduct struct {
	CustomAttributes interface{} `json:"customAttributes"`
	Developer        string      `json:"developer"`
	Publisher        string      `json:"publisher"`
	Gallery          []string    `json:"gallery"`
	Video            struct {
		Id       string `json:"id"`
		Provider string `json:"provider"`
	} `json:"video"`
	SupportedOperatingSystems []string `json:"supportedOperatingSystems"`
	Genres                    []string `json:"genres"`
	GlobalReleaseDate         int      `json:"globalReleaseDate"`
	IsTBA                     bool     `json:"isTBA"`
	Price                     struct {
		Amount                     string  `json:"amount"`
		BaseAmount                 string  `json:"baseAmount"`
		FinalAmount                string  `json:"finalAmount"`
		IsDiscounted               bool    `json:"isDiscounted"`
		DiscountPercentage         int     `json:"discountPercentage"`
		DiscountDifference         string  `json:"discountDifference"`
		Symbol                     string  `json:"symbol"`
		IsFree                     bool    `json:"isFree"`
		Discount                   int     `json:"discount"`
		IsBonusStoreCreditIncluded bool    `json:"isBonusStoreCreditIncluded"`
		BonusStoreCreditAmount     string  `json:"bonusStoreCreditAmount"`
		PromoId                    *string `json:"promoId"`
	} `json:"price"`
	IsDiscounted    bool `json:"isDiscounted"`
	IsInDevelopment bool `json:"isInDevelopment"`
	Id              int  `json:"id"`
	ReleaseDate     int  `json:"releaseDate"`
	Availability    struct {
		IsAvailable          bool `json:"isAvailable"`
		IsAvailableInAccount bool `json:"isAvailableInAccount"`
	} `json:"availability"`
	SalesVisibility struct {
		IsActive   bool `json:"isActive"`
		FromObject struct {
			Date         string `json:"date"`
			TimezoneType int    `json:"timezone_type"`
			Timezone     string `json:"timezone"`
		} `json:"fromObject"`
		From     int `json:"from"`
		ToObject struct {
			Date         string `json:"date"`
			TimezoneType int    `json:"timezone_type"`
			Timezone     string `json:"timezone"`
		} `json:"toObject"`
		To int `json:"to"`
	} `json:"salesVisibility"`
	Buyable    bool   `json:"buyable"`
	Title      string `json:"title"`
	Image      string `json:"image"`
	Url        string `json:"url"`
	SupportUrl string `json:"supportUrl"`
	ForumUrl   string `json:"forumUrl"`
	WorksOn    struct {
		Windows bool `json:"Windows"`
		Mac     bool `json:"Mac"`
		Linux   bool `json:"Linux"`
	} `json:"worksOn"`
	Category         string `json:"category"`
	OriginalCategory string `json:"originalCategory"`
	Rating           int    `json:"rating"`
	Type             int    `json:"type"`
	IsComingSoon     bool   `json:"isComingSoon"`
	IsPriceVisible   bool   `json:"isPriceVisible"`
	IsMovie          bool   `json:"isMovie"`
	IsGame           bool   `json:"isGame"`
	Slug             string `json:"slug"`
	IsWishlistable   bool   `json:"isWishlistable"`
}

func (sp *StoreProduct) GetId() int {
	return sp.Id
}

func (sp *StoreProduct) GetTitle() string {
	return sp.Title
}

func (sp *StoreProduct) GetDevelopers() []string {
	return []string{sp.Developer}
}

func (sp *StoreProduct) GetPublisher() string {
	return sp.Publisher
}

func (sp *StoreProduct) GetImage() string {
	return sp.Image
}

func (sp *StoreProduct) GetScreenshots() []string {
	return sp.Gallery
}

func (sp *StoreProduct) GetRating() string {
	return strconv.Itoa(sp.Rating)
}

func (sp *StoreProduct) GetGenres() []string {
	return sp.Genres
}

func (sp *StoreProduct) GetVideoIds() []string {
	if sp.Video.Provider != supportedVideoProvider {
		return []string{}
	}
	return []string{sp.Video.Id}
}

func (sp *StoreProduct) GetOperatingSystems() []string {
	os := make([]string, 0, 3)
	if sp.WorksOn.Windows {
		os = append(os, windows)
	}
	if sp.WorksOn.Mac {
		os = append(os, macos)
	}
	if sp.WorksOn.Linux {
		os = append(os, linux)
	}
	return os
}

func (sp *StoreProduct) GetSlug() string {
	return sp.Slug
}

func (sp *StoreProduct) GetGlobalRelease() int64 {
	return int64(sp.GlobalReleaseDate)
}

func (sp *StoreProduct) GetGOGRelease() int64 {
	//we're using .ReleaseDate as GOG release date
	if sp.ReleaseDate == sp.GlobalReleaseDate {
		return 0
	}
	return int64(sp.ReleaseDate)
}

func (sp *StoreProduct) GetStoreUrl() string {
	return sp.Url
}

func (sp *StoreProduct) GetForumUrl() string {
	return sp.ForumUrl
}

func (sp *StoreProduct) GetSupportUrl() string {
	return sp.SupportUrl
}

func (sp *StoreProduct) GetTBA() bool {
	return sp.IsTBA
}

func (sp *StoreProduct) GetComingSoon() bool {
	return sp.IsComingSoon
}
