package main

import (
	"fmt"
	"github.com/Raajheer1/hackutd-x/m/v2/pkg/config"
	"github.com/Raajheer1/hackutd-x/m/v2/pkg/gpt"
	"sync"
)

type container struct {
	mu   sync.Mutex
	recs []config.Recommendation
}

func main() {
	fmt.Println("Im online...")

	config.Active = config.ActiveConfig{
		CompanyName:        "Bob's Electrician Shop",
		Industry:           "Electrician",
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

	// TODO - call GPT-3 to get recommendations
	// Recommendation types: Safety, Tax, Finance, Marketing, Legal
	prompt := fmt.Sprintf("Company: %s, Industry: %s, City: %s", config.Active.CompanyName, config.Active.Industry, config.Active.City)

	cont := container{
		recs: []config.Recommendation{},
	}

	var wg = &sync.WaitGroup{}
	for _, rec := range config.RecTypes {

		rec := rec
		wg.Add(1)
		go func(recType config.RecommendationTypes) {
			defer wg.Done()
			resp := gpt.GetRecommendation(prompt, recType)
			for _, r := range resp {
				cont.mu.Lock()
				cont.recs = append(cont.recs, config.Recommendation{
					Title:     r.Title,
					Summary:   r.Summary,
					Details:   r.Details,
					Type:      string(rec),
					Completed: false,
				})
				cont.mu.Unlock()
			}
		}(rec)
	}

	wg.Wait()

	fmt.Println(len(cont.recs))
	for _, rec := range cont.recs {
		fmt.Println(rec.Title)
		fmt.Println(rec.Summary)
		fmt.Println(rec.Details)
		fmt.Println("----")
	}
}
