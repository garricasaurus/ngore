package recommended

type Recommendations struct {
	Movies *RecommendationGroup `json:"movies"`
	Series *RecommendationGroup `json:"series"`
	Games  *RecommendationGroup `json:"games"`
	Music  *RecommendationGroup `json:"music"`
	Apps   *RecommendationGroup `json:"apps"`
	Books  *RecommendationGroup `json:"books"`
}

type RecommendationGroup struct {
	Staff  []*RecommendationWithImage `json:"staff"`
	Active []*Recommendation          `json:"active"`
}

type Recommendation struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type RecommendationWithImage struct {
	Recommendation
	Image string `json:"image"`
}
