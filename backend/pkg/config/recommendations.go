package config

type Recommendation struct {
	Title     string `json:"title"`
	Summary   string `json:"summary"`
	Details   string `json:"details"`
	Type      string `json:"type"`
	Completed bool   `json:"completed"`
}

var Recs []Recommendation = []Recommendation{
	{
		Title:     "Tax Recommendation",
		Summary:   "Lie to the IRS",
		Details:   "Lying to the IRS is an extremely effective way to save money on taxes. So long as you don't get caught...",
		Type:      "Tax",
		Completed: false,
	},
}

type RecommendationTypes string

var RecTypes []RecommendationTypes = []RecommendationTypes{
	"Safety",
	"Tax",
	"Finance",
	"Marketing",
	"Legal",
}
