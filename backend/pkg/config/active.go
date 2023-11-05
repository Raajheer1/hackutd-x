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
	PresetLevel        string  `json:"presetLevel,omitempty"`
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

var Configs []ActiveConfig = []ActiveConfig{
	{
		//company that is doing exceptional
		CompanyName:        "Sue's Bakery",
		PresetLevel:        "Exceptional",
		Industry:           "Bakery",
		City:               "Southlake",
		Revenue:            float64(800000),
		Cogs:               float64(300000),
		Depreciation:       float64(15000),
		LongTermAssets:     float64(150000),
		ShortTermAssets:    float64(250000),
		LongTermLiability:  float64(50000),
		ShortTermLiability: float64(100000),
		OperatingExpense:   float64(200000),
		RetainedEarnings:   float64(100000),
		YearsInBusiness:    uint(7),
	},
	{
		//company that is doing Average
		CompanyName:        "Sue's Bakery",
		PresetLevel:        "Average",
		Industry:           "Bakery",
		City:               "Southlake",
		Revenue:            float64(400000),
		Cogs:               float64(300000),
		Depreciation:       float64(15000),
		LongTermAssets:     float64(150000),
		ShortTermAssets:    float64(250000),
		LongTermLiability:  float64(50000),
		ShortTermLiability: float64(100000),
		OperatingExpense:   float64(200000),
		RetainedEarnings:   float64(100000),
		YearsInBusiness:    uint(7),
	},
	{
		//company that is doing Average
		CompanyName:        "Sue's Bakery",
		PresetLevel:        "Terrible",
		Industry:           "Bakery",
		City:               "Southlake",
		Revenue:            float64(100000),
		Cogs:               float64(300000),
		Depreciation:       float64(15000),
		LongTermAssets:     float64(150000),
		ShortTermAssets:    float64(250000),
		LongTermLiability:  float64(50000),
		ShortTermLiability: float64(100000),
		OperatingExpense:   float64(200000),
		RetainedEarnings:   float64(100000),
		YearsInBusiness:    uint(1),
	},
	{
		//company that is doing really bad
		CompanyName:        "Joe's Plumbing",
		PresetLevel:        "Terrible",
		Industry:           "Plumbing",
		City:               "Richardson",
		Revenue:            float64(400000),
		Cogs:               float64(250000),
		Depreciation:       float64(20000),
		LongTermAssets:     float64(50000),
		ShortTermAssets:    float64(100000),
		LongTermLiability:  float64(80000),
		ShortTermLiability: float64(120000),
		OperatingExpense:   float64(300000),
		RetainedEarnings:   float64(-50000),
		YearsInBusiness:    uint(3),
	},
	{
		//company that is going middle-of-the-road well
		CompanyName:        "Bob's Electrician Shop",
		PresetLevel:        "Average",
		Industry:           "Electrician",
		City:               "Dallas",
		Revenue:            float64(900000),
		Cogs:               float64(400000),
		Depreciation:       float64(30000),
		LongTermAssets:     float64(150000),
		ShortTermAssets:    float64(280000),
		LongTermLiability:  float64(70000),
		ShortTermLiability: float64(110000),
		OperatingExpense:   float64(250000),
		RetainedEarnings:   float64(100000),
		YearsInBusiness:    uint(5),
	},
}
