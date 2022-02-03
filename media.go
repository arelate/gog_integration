// Copyright (c) 2020. All rights reserved.
// Use of this source code is governed by an MIT-style license that can be
// found in the LICENSE file.

package gog_atu

type Media int

const (
	Unknown Media = iota
	Game
	Movie
)

var mediaStrings = map[Media]string{
	Unknown: "unknown",
	Game:    "game",
	Movie:   "movie",
}

func (mt Media) String() string {
	if media, ok := mediaStrings[mt]; ok {
		return media
	}
	return ""
}

func ParseMedia(media string) Media {
	for mt, str := range mediaStrings {
		if str == media {
			return mt
		}
	}
	return Unknown
}

func ValidMedia(mt Media) bool {
	for media, _ := range mediaStrings {
		if media == mt && media != Unknown {
			return true
		}
	}
	return false
}

func AllMedia() []Media {
	return []Media{
		Game,
		Movie,
	}
}
