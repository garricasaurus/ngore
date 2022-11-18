package ncgore

import (
	"errors"
	"golang.org/x/net/html"
	"net/http"
	"net/url"
	"strconv"
)

func (a *api) Search(params *SearchParams) ([]*SearchResult, error) {
	res, err := a.client.PostForm(a.baseUrl+urlSearch, searchForm(params))
	if err != nil {
		return nil, err
	}
	if isLoginRequired(res) {
		return nil, errors.New(errUserNotLoggedIn)
	}
	return parseSearchResults(res)
}

func parseSearchResults(res *http.Response) ([]*SearchResult, error) {
	doc, err := html.Parse(res.Body)
	if err != nil {
		return nil, err
	}
	result := make([]*SearchResult, 0)
	elements := GetElementsByClass(doc, "torrent_txt")
	for _, element := range elements {
		a := GetFirstChild(element, "a")
		if a != nil {
			title, ok := findAttr(a, "title")
			if ok {
				result = append(result, &SearchResult{
					Title: title,
				})
			}
		}
	}
	return result, nil
}

type SearchParams struct {
	SearchPhrase string
	Field        SearchField
	Category     SearchCategory
	SortField    SortField
	SortMode     SortMode
	Page         int
}

func searchForm(s *SearchParams) url.Values {
	return url.Values{
		"mire":      {s.SearchPhrase},
		"miben":     {s.Field.String()},
		"tipus":     {s.Category.String()},
		"miszerint": {s.SortField.String()},
		"hogyan":    {s.SortMode.String()},
		"oldal":     {strconv.Itoa(s.Page + 1)},
	}
}

type SearchResult struct {
	Title    string
	Uploaded string
	Size     string
	Health   string
	Seeds    string
	Peers    string
}

type SearchField int

const (
	Name SearchField = iota
	Description
	Imdb
	Label
)

func (s SearchField) String() string {
	switch s {
	case Name:
		return "name"
	case Description:
		return "leiras"
	case Imdb:
		return "imdb"
	case Label:
		return "cimke"
	default:
		return "unknown"
	}
}

type SearchCategory int

const (
	MovieSdEn SearchCategory = iota
	MovieSdHu
	MovieDvdEn
	MovieDvdHu
	MovieDvd9En
	MovieDvd9Hu
	MovieHdEn
	MovieHdHu
	SeriesSdEn
	SeriesSdHu
	SeriesDvdEn
	SeriesDvdHu
	SeriesHdEn
	SeriesHdHu
	Mp3En
	Mp3Hu
	LosslessEn
	LosslessHu
	Clip
	GameIso
	GameRip
	Console
	EbookEn
	EbookHu
	Iso
	Misc
	Mobile
	XImg
	XSd
	XDvd
	XHd
	AllOwn
)

func (s SearchCategory) String() string {
	switch s {
	case MovieSdEn:
		return "xvid"
	case MovieSdHu:
		return "xvid_hun"
	case MovieDvdEn:
		return "dvd"
	case MovieDvdHu:
		return "dvd_hun"
	case MovieDvd9En:
		return "dvd9"
	case MovieDvd9Hu:
		return "dvd9_hun"
	case MovieHdEn:
		return "hd"
	case MovieHdHu:
		return "hd_hun"
	case SeriesSdEn:
		return "xvidser"
	case SeriesSdHu:
		return "xvidser_hun"
	case SeriesDvdEn:
		return "dvdser"
	case SeriesDvdHu:
		return "dvdser_hun"
	case SeriesHdEn:
		return "hdser"
	case SeriesHdHu:
		return "hdser_hun"
	case Mp3En:
		return "mp3"
	case Mp3Hu:
		return "mp3_hun"
	case LosslessEn:
		return "lossless"
	case LosslessHu:
		return "lossless_hun"
	case Clip:
		return "clip"
	case GameIso:
		return "game_iso"
	case GameRip:
		return "game_rip"
	case Console:
		return "console"
	case EbookEn:
		return "ebook"
	case EbookHu:
		return "ebook_hun"
	case Iso:
		return "iso"
	case Misc:
		return "misc"
	case Mobile:
		return "mobil"
	case XImg:
		return "xxx_imageset"
	case XSd:
		return "xxx_xvid"
	case XDvd:
		return "xxx_dvd"
	case XHd:
		return "xxx_hd"
	case AllOwn:
		return "all_own"
	default:
		return "unknown"
	}
}

type SortField int

const (
	ByName SortField = iota
	ByUpload
	BySize
	ByDownloaded
	BySeeders
	ByLeechers
)

func (s SortField) String() string {
	switch s {
	case ByName:
		return "name"
	case ByUpload:
		return "fid"
	case BySize:
		return "size"
	case ByDownloaded:
		return "times_completed"
	case BySeeders:
		return "seeders"
	case ByLeechers:
		return "leechers"
	default:
		return "unknown"
	}
}

type SortMode int

const (
	Ascending SortMode = iota
	Descending
)

func (s SortMode) String() string {
	switch s {
	case Ascending:
		return "ASC"
	case Descending:
		return "DESC"
	default:
		return "unknown"
	}
}
