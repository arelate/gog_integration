package gog_integration

type ProductsGetter interface {
	GetProducts() []IdGetter
}

type IdGetter interface {
	GetId() int
}

type TitleGetter interface {
	GetTitle() string
}

type DevelopersGetter interface {
	GetDevelopers() []string
}

type PublisherGetter interface {
	GetPublisher() string
}

type ImageGetter interface {
	GetImage() string
}

type ScreenshotsGetter interface {
	GetScreenshots() []string
}

type RatingGetter interface {
	GetRating() string
}

type GenresGetter interface {
	GetGenres() []string
}

type FeaturesGetter interface {
	GetFeatures() []string
}

type SeriesGetter interface {
	GetSeries() string
}

type accountTagsGetter interface {
	getAccountTags() []accountTag
}

type TagIdsGetter interface {
	GetTagIds() []string
}

type TagNamesGetter interface {
	GetTagNames([]string) map[string]string
}

type VideoIdsGetter interface {
	GetVideoIds() []string
}

type OperatingSystemsGetter interface {
	GetOperatingSystems() []string
}

type IncludesGamesGetter interface {
	GetIncludesGames() []string
}

type IsIncludedInGamesGetter interface {
	GetIsIncludedInGames() []string
}

type IsRequiredByGamesGetter interface {
	GetIsRequiredByGames() []string
}

type RequiresGamesGetter interface {
	GetRequiresGames() []string
}

type LanguagesGetter interface {
	GetLanguages() map[string]string
}

type NativeLanguagesGetter interface {
	GetNativeLanguages() map[string]string
}

type LanguageCodesGetter interface {
	GetLanguageCodes() []string
}

type SlugGetter interface {
	GetSlug() string
}

type GlobalReleaseGetter interface {
	GetGlobalRelease() int64
}

type GOGReleaseGetter interface {
	GetGOGRelease() int64
}
