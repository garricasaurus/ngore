package activity

type Info struct {
	Rank        Rank
	Stats       Stats
	CanDownload string
	History     []TorrentActivity
}

type Rank struct {
	Daily     string
	Weekly    string
	Monthly   string
	PrevMonth string
}

type Stats struct {
	Current     string
	Allowed     string
	PenMonths   string
	PenTorrents string
}

type TorrentActivity struct {
	Name      string
	Start     string
	Updated   string
	Status    string
	Up        string
	Down      string
	Remaining string
	Ratio     string
}
