package gog_integration

type StoreSortOrder string

const (
	StoreSortByBestselling        StoreSortOrder = "popularity"
	StoreSortByAlphabetically                    = "title"
	StoreSortByUserRating                        = "rating"
	StoreSortByDateAdded                         = "date"
	StoreSortByBestsellingAllTime                = "bestselling"
	StoreSortByOldestFirst                       = "release_asc"
	StoreSortByNewestFirst                       = "release_desc"
)
