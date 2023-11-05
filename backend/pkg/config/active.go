package config

type ActiveConfig struct {
	CompanyName        string  `json:"companyName"`
	Industry           string  `json:"industry"`
	City               string  `json:"city"`
	Revenue            float64 `json:"revenue"`
	Cogs               float64 `json:"cogs"`
	Depreciation       float64 `json:"depreciation"`
	LongTermAssets     float64 `json:"longTermAssets"`
	ShortTermAssets    float64 `json:"shortTermAssets"`
	LongTermLiability  float64 `json:"longTermLiability"`
	ShortTermLiability float64 `json:"shortTermLiability"`
	OperatingExpense   float64 `json:"operatingExpense"`
	RetainedEarnings   float64 `json:"retainedEarnings"`
	YearsInBusiness    uint    `json:"yearsInBusiness"`
}

var Active ActiveConfig = ActiveConfig{
	CompanyName:        "State Farm",
	Industry:           "Insurance",
	City:               "Dallas",
	Revenue:            float64(1000000),
	Cogs:               float64(500000),
	Depreciation:       float64(100000),
	LongTermAssets:     float64(100000),
	ShortTermAssets:    float64(100000),
	LongTermLiability:  float64(100000),
	ShortTermLiability: float64(100000),
	OperatingExpense:   float64(100000),
	RetainedEarnings:   float64(100000),
	YearsInBusiness:    uint(5),
}
