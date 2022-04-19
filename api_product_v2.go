package gog_integration

import (
	"path"
	"strings"
	"time"
)

type href struct {
	Href string `json:"href"`
}

type name struct {
	Name string `json:"name"`
}

type slug struct {
	Slug string `json:"slug"`
}

type links struct {
	Self                  href   `json:"self"`
	Store                 href   `json:"store"`
	Support               href   `json:"support"`
	Forum                 href   `json:"forum"`
	Icon                  href   `json:"icon"`
	IconSquare            href   `json:"iconSquare"`
	Logo                  href   `json:"logo"`
	BoxArtImage           href   `json:"boxArtImage"`
	BackgroundImage       href   `json:"backgroundImage"`
	GalaxyBackgroundImage href   `json:"galaxyBackgroundImage"`
	IncludesGames         []href `json:"includesGames"`
	IsIncludedInGames     []href `json:"isIncludedInGames"`
	RequiresGames         []href `json:"requiresGames"`
	IsRequiredByGames     []href `json:"isRequiredByGames"`
}

type supportedOperatingSystem struct {
	OperatingSystem struct {
		Name     string `json:"name"`
		Versions string `json:"versions"`
	} `json:"operatingSystem"`
	SystemRequirements []struct {
		Type         string `json:"type"`
		Description  string `json:"description"`
		Requirements []struct {
			Id          string `json:"id"`
			Name        string `json:"name"`
			Description string `json:"description"`
		} `json:"requirements"`
	} `json:"systemRequirements"`
}

type edition struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Bonuses []struct {
		Id int `json:"id"`
		name
		Type struct {
			name
			slug
		} `json:"type"`
	} `json:"bonuses"`
	HasProductCard bool `json:"hasProductCard"`
	Links          struct {
		Self   href `json:"self"`
		Prices struct {
			href
			Templated bool `json:"templated"`
		} `json:"prices"`
	} `json:"_links"`
}

type product struct {
	IsAvailableForSale bool      `json:"isAvailableForSale"`
	IsVisibleInCatalog bool      `json:"isVisibleInCatalog"`
	IsPreorder         bool      `json:"isPreorder"`
	IsVisibleInAccount bool      `json:"isVisibleInAccount"`
	Id                 int       `json:"id"`
	Title              string    `json:"title"`
	IsInstallable      bool      `json:"isInstallable"`
	GlobalReleaseDate  time.Time `json:"globalReleaseDate"`
	HasProductCard     bool      `json:"hasProductCard"`
	GogReleaseDate     time.Time `json:"gogReleaseDate"`
	IsSecret           bool      `json:"isSecret"`
	Links              struct {
		Image struct {
			href
			Templated  bool     `json:"templated"`
			Formatters []string `json:"formatters"`
		} `json:"image"`
		Checkout href `json:"checkout"`
		Prices   struct {
			href
			Templated bool `json:"templated"`
		} `json:"prices"`
	} `json:"_links"`
}

type localization struct {
	Embedded struct {
		Language struct {
			Code string `json:"code"`
			name
		} `json:"language"`
		LocalizationScope struct {
			Type string `json:"type"`
		} `json:"localizationScope"`
	} `json:"_embedded"`
}

type tag struct {
	Id int `json:"id"`
	name
	Level int `json:"level"`
	slug
}

type property struct {
	name
	slug
}

type screenshot struct {
	Links struct {
		Self struct {
			href
			Templated  bool     `json:"templated"`
			Formatters []string `json:"formatters"`
		} `json:"self"`
	} `json:"_links"`
}

type esrbRating struct {
	Category struct {
		Name string `json:"name"`
		Id   int    `json:"id"`
	} `json:"category"`
	ContentDescriptors []struct {
		Descriptor string `json:"descriptor"`
		Id         int    `json:"id"`
	} `json:"contentDescriptors"`
}

type pegiRating struct {
	AgeRating          int `json:"ageRating"`
	ContentDescriptors []struct {
		Descriptor string `json:"descriptor"`
		Id         int    `json:"id"`
	} `json:"contentDescriptors"`
}

type uskRating struct {
	AgeRating int `json:"ageRating"`
	Category  struct {
		Name string `json:"name"`
	} `json:"category"`
}

type brRating struct {
	AgeRating          int `json:"ageRating"`
	ContentDescriptors []struct {
		Descriptor string `json:"descriptor"`
	} `json:"contentDescriptors"`
}

type bonus struct {
	Id int `json:"id"`
	name
	Type struct {
		name
		slug
	} `json:"type"`
}

type video struct {
	Provider    string `json:"provider"`
	VideoId     string `json:"videoId"`
	ThumbnailId string `json:"thumbnailId"`
	Links       struct {
		Self      href `json:"self"`
		Thumbnail href `json:"thumbnail"`
	} `json:"_links"`
}

type ApiProductV2 struct {
	InDevelopment struct {
		Active bool `json:"active"`
	} `json:"inDevelopment"`
	Copyrights          string `json:"copyrights"`
	IsUsingDosBox       bool   `json:"isUsingDosBox"`
	Description         string `json:"description"`
	Size                int    `json:"size"`
	Overview            string `json:"overview"`
	FeaturesDescription string `json:"featuresDescription"`
	Links               links  `json:"_links"`
	Embedded            struct {
		Product                   product                    `json:"product"`
		ProductType               string                     `json:"productType"`
		Localizations             []localization             `json:"localizations"`
		Videos                    []video                    `json:"videos"`
		Bonuses                   []bonus                    `json:"bonuses"`
		Screenshots               []screenshot               `json:"screenshots"`
		Publisher                 name                       `json:"publisher"`
		Developers                []name                     `json:"developers"`
		SupportedOperatingSystems []supportedOperatingSystem `json:"supportedOperatingSystems"`
		Tags                      []tag                      `json:"tags"`
		Properties                []property                 `json:"properties"`
		ESRBRating                esrbRating                 `json:"esrbRating"`
		PEGIRating                pegiRating                 `json:"pegiRating"`
		USKRating                 uskRating                  `json:"uskRating"`
		BRRating                  brRating                   `json:"brRating"`
		Features                  []struct {
			Id string `json:"id"`
			name
		} `json:"features"`
		Editions []edition `json:"editions"`
		Series   struct {
			Id int `json:"id"`
			name
		} `json:"series"`
	} `json:"_embedded"`
}

func (apv2 *ApiProductV2) GetId() int {
	return apv2.Embedded.Product.Id
}

func (apv2 *ApiProductV2) GetTitle() string {
	return apv2.Embedded.Product.Title
}

func (apv2 *ApiProductV2) GetDevelopers() []string {
	devs := make([]string, 0, len(apv2.Embedded.Developers))
	for _, dev := range apv2.Embedded.Developers {
		devs = append(devs, dev.Name)
	}
	return devs
}

func (apv2 *ApiProductV2) GetPublisher() string {
	return apv2.Embedded.Publisher.Name
}

func (apv2 *ApiProductV2) GetImage() string {
	return apv2.Embedded.Product.Links.Image.Href
}

func (apv2 *ApiProductV2) GetScreenshots() []string {
	screenshots := make([]string, 0)
	for _, screenshot := range apv2.Embedded.Screenshots {
		screenshots = append(screenshots, screenshot.Links.Self.Href)
	}
	return screenshots
}

func idFromHref(href string) string {
	_, urlPath := path.Split(href)
	parts := strings.Split(urlPath, "?")
	if len(parts) > 0 {
		return parts[0]
	}
	return ""
}

func (apv2 *ApiProductV2) getHrefs(hrefProperty []href) []string {
	hrefs := make([]string, 0, len(hrefProperty))
	for _, hp := range hrefProperty {
		id := idFromHref(hp.Href)
		if id != "" {
			hrefs = append(hrefs, id)
		}
	}
	return hrefs
}

func (apv2 *ApiProductV2) GetGenres() []string {
	genres := make([]string, 0, len(apv2.Embedded.Tags))
	for _, tag := range apv2.Embedded.Tags {
		genres = append(genres, tag.Name)
	}
	return genres
}

func (apv2 *ApiProductV2) GetFeatures() []string {
	features := make([]string, 0, len(apv2.Embedded.Features))
	for _, feature := range apv2.Embedded.Features {
		features = append(features, feature.Name)
	}
	return features
}

func (apv2 *ApiProductV2) GetSeries() string {
	return apv2.Embedded.Series.Name
}

func (apv2 *ApiProductV2) GetVideoIds() []string {
	videoIds := make([]string, 0, len(apv2.Embedded.Videos))
	for _, vid := range apv2.Embedded.Videos {
		if vid.Provider != supportedVideoProvider {
			continue
		}
		//1733288799 had video_id specified as "v=...."
		vId := strings.TrimPrefix(vid.VideoId, "v=")
		videoIds = append(videoIds, vId)
	}
	return videoIds
}

func (apv2 *ApiProductV2) GetOperatingSystems() []string {
	os := make([]string, 0, 3)
	for _, sos := range apv2.Embedded.SupportedOperatingSystems {
		switch sos.OperatingSystem.Name {
		case windows:
			os = append(os, windows)
		case osx:
			// GOG.com continues to use "osx" name for macOS
			os = append(os, macos)
		case linux:
			os = append(os, linux)
		}
	}
	return os
}

func (apv2 *ApiProductV2) GetIncludesGames() []string {
	return apv2.getHrefs(apv2.Links.IncludesGames)
}

func (apv2 *ApiProductV2) GetIsIncludedInGames() []string {
	return apv2.getHrefs(apv2.Links.IsIncludedInGames)
}

func (apv2 *ApiProductV2) GetIsRequiredByGames() []string {
	return apv2.getHrefs(apv2.Links.IsRequiredByGames)
}

func (apv2 *ApiProductV2) GetRequiresGames() []string {
	return apv2.getHrefs(apv2.Links.RequiresGames)
}

func (apv2 *ApiProductV2) GetLanguages() map[string]string {
	langs := make(map[string]string, len(apv2.Embedded.Localizations))
	for _, loc := range apv2.Embedded.Localizations {
		langs[loc.Embedded.Language.Code] = loc.Embedded.Language.Name
	}
	return langs
}

func (apv2 *ApiProductV2) GetLanguageCodes() []string {
	lcs := make(map[string]bool, 0)
	for _, loc := range apv2.Embedded.Localizations {
		lcs[loc.Embedded.Language.Code] = true
	}

	codes := make([]string, 0, len(lcs))
	for code, _ := range lcs {
		codes = append(codes, code)
	}
	return codes
}

func (apv2 *ApiProductV2) GetGlobalRelease() int64 {
	return apv2.Embedded.Product.GlobalReleaseDate.Unix()
}

func (apv2 *ApiProductV2) GetGOGRelease() int64 {
	gogReleaseDate := apv2.Embedded.Product.GogReleaseDate.Unix()
	//GOG.com uses "1991-12-25T00:00:00+02:00" as "not released" or "unknown release date".
	//GOG.com uses "2000-12-31T23:59:59+02:00" for games that were removed from the store.
	//Considering GOG was founded in 2008 - neither is a real date for practical purposes.
	if gogReleaseDate == 693612000 ||
		gogReleaseDate == 978299999 {
		return 0
	}
	return gogReleaseDate
}

func (apv2 *ApiProductV2) GetStoreUrl() string {
	return urlPathFromLink(apv2.Links.Store.Href)
}

func (apv2 *ApiProductV2) GetForumUrl() string {
	return urlPathFromLink(apv2.Links.Forum.Href)
}

func (apv2 *ApiProductV2) GetSupportUrl() string {
	return urlPathFromLink(apv2.Links.Support.Href)
}

const emDashCode = "\u2013"
const doubleNewLineChar = "\n\n"
const newLineChar = "\n"

//implicitToExplicitList looks for embedded \u2013 characters
//that GOG.com is using for <ul> lists creation, e.g.
//https://www.gog.com/en/game/deaths_gambit
//and replaces that segment with explicit unordered lists
func implicitToExplicitList(text string) string {
	var items []string
	if strings.Contains(text, emDashCode) {
		items = strings.Split(text, emDashCode)
	} else if strings.Contains(text, doubleNewLineChar) {
		items = strings.Split(text, doubleNewLineChar)
	} else if strings.Contains(text, newLineChar) {
		items = strings.Split(text, newLineChar)
	}

	if len(items) > 0 {
		builder := strings.Builder{}
		builder.WriteString("<ul>")
		for _, item := range items {
			builder.WriteString("<li>" + item + "</li>")
		}
		builder.WriteString("</ul>")
		text = builder.String()
	}

	return text
}

func (apv2 *ApiProductV2) GetDescription() string {
	if apv2.Overview != "" && apv2.FeaturesDescription != "" {
		return apv2.Overview + implicitToExplicitList(apv2.FeaturesDescription)
	} else {
		return apv2.Description
	}
}

func (apv2 *ApiProductV2) GetProductType() string {
	return apv2.Embedded.ProductType
}

func (apv2 *ApiProductV2) GetCopyrights() string {
	return apv2.Copyrights
}

//GetProperties returns values displayed in Store as "Tags".
//Using Properties name here to be consistent with the source data
//and given that Tags is already used for "Genres" (see GetGenres method)
func (apv2 *ApiProductV2) GetProperties() []string {
	properties := make([]string, 0, len(apv2.Embedded.Properties))
	for _, p := range apv2.Embedded.Properties {
		properties = append(properties, p.Name)
	}
	return properties
}
