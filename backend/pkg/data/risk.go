package data

import (
	"fmt"
	"github.com/Raajheer1/hackutd-x/m/v2/pkg/config"
	"github.com/Raajheer1/hackutd-x/m/v2/pkg/gpt"
	"math/rand"
)

func CalculateRisk() float64 {
	return RiskScoreFormula(CalculateBaselineRisk(), CalculateRisks(), CalculateTasks())
}

var BaselineRisk float64

func CalculateBaselineRisk() float64 {
	for BaselineRisk == 0 {
		BaselineRisk = gpt.GetBaselineRisk(config.Active.YearsInBusiness, config.Active.Industry).BaselineRisk
	}
	return BaselineRisk
}

type Risk struct {
	Name   string
	Weight float64
	Score  float64
}

type Task struct {
	Name   string
	Weight float64
	Score  float64
}

func RiskScoreFormula(baseline float64, risks []Risk, tasks []Task) float64 {
	fmt.Println("Calculating risk score...")
	var riskScore float64
	for _, risk := range risks {
		fmt.Printf("Risk: %s, Weight: %f, Score: %f\n", risk.Name, risk.Weight, risk.Score)
		if risk.Name == "Finance" {
			risk.Score += risk.Weight * InvertZ(CalculateZ()) / 2.0
		} else {
			riskScore += risk.Weight * risk.Score * 50.0
		}
	}
	var taskScore float64
	for _, task := range tasks {
		taskScore += task.Weight * task.Score * 50.0 / float64(len(tasks))
	}

	limitedZScore := InvertZ(CalculateZ()) * 25
	fmt.Printf("Baseline: %f, RiskScore: %f, TaskScore: %f, ZScore: %f\n", baseline*50, riskScore, taskScore, limitedZScore)
	r := 0.25*(baseline*50) + 0.5*(riskScore-taskScore) - 0.25*(limitedZScore)

	if r <= 0 {
		return rand.Float64() * 10
	}
	if r >= 100 {
		return rand.Float64()*10 + 90
	}
	return r
}

var Weights []gpt.RiskResponse

func CalculateRisks() []Risk {
	var risks []Risk
	for len(Weights) == 0 {
		Weights = gpt.GetRiskWeights(config.Active.YearsInBusiness, config.Active.Industry)
	}
	fmt.Printf("Weights Length: %d", len(Weights))

	tasks := CalculateTasks()
	for _, weight := range Weights {
		r := Risk{
			Name:   weight.RiskFactor,
			Weight: weight.RiskWeight,
			Score:  config.RiskFactorScores[weight.RiskFactor],
		}

		for _, t := range tasks {
			if t.Score == 1 {
				if t.Name == r.Name {
					if r.Score == 2 {
						r.Score = 1
					} else if r.Score == 1 {
						r.Score = 0.5
					} else {
						r.Score = 0
					}
				}
			} else {
				r.Score = 2
			}
		}

		risks = append(risks, r)
	}

	return risks
}

func CalculateTasks() []Task {
	var tasks []Task
	for _, rec := range config.Recs {
		score := 0.0
		if rec.Completed {
			score = 1
		}
		tasks = append(tasks, Task{
			Name:   rec.Type,
			Weight: rec.Weight,
			Score:  score,
		})
	}

	return tasks
}
