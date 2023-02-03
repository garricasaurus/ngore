package details

type Details struct {
	Title       string   `json:"title"`
	ReleaseYear string   `json:"releaseYear"`
	Director    string   `json:"director"`
	Actors      string   `json:"actors"`
	Country     string   `json:"country"`
	Labels      string   `json:"labels"`
	Length      string   `json:"length"`
	ImdbRating  string   `json:"imdbRating"`
	ImdbLink    string   `json:"imdbLink"`
	OtherLink   string   `json:"otherLink"`
	CoverImage  string   `json:"coverImage"`
	OtherImages []string `json:"otherImages"`
}
