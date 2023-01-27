package activity

type Info struct {
	Rank        Rank              `json:"rank"`
	Stats       Stats             `json:"stats"`
	CanDownload string            `json:"canDownload"`
	History     []TorrentActivity `json:"history"`
}

type Rank struct {
	Daily     string `json:"daily"`
	Weekly    string `json:"weekly"`
	Monthly   string `json:"monthly"`
	PrevMonth string `json:"prevMonth"`
}

type Stats struct {
	Current     string `json:"current"`
	Allowed     string `json:"allowed"`
	PenMonths   string `json:"penMonths"`
	PenTorrents string `json:"penTorrents"`
}

type TorrentActivity struct {
	Name      string `json:"name"`
	Start     string `json:"start"`
	Updated   string `json:"updated"`
	Status    string `json:"status"`
	Up        string `json:"up"`
	Down      string `json:"down"`
	Remaining string `json:"remaining"`
	Ratio     string `json:"ratio"`
}
