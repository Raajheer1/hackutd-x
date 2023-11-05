package data

import (
	"github.com/Raajheer1/hackutd-x/m/v2/pkg/config"
	"github.com/Raajheer1/hackutd-x/m/v2/pkg/gpt"
)

func CalculateRisk() float64 {
	return RiskScoreFormula(CalculateBaselineRisk(), CalculateRisks(), CalculateTasks())
}

func CalculateBaselineRisk() float64 {
	return gpt.GetBaselineRisk(config.Active.YearsInBusiness, config.Active.Industry).BaselineRisk
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
	var riskScore float64
	for _, risk := range risks {
		if risk.Name == "Finance" {
			risk.Score += risk.Weight * InvertZ(CalculateZ()) / 2.0
		} else {
			riskScore += risk.Weight * risk.Score
		}
	}
	var taskScore float64
	for _, task := range tasks {
		taskScore += task.Weight * task.Score
	}
	return baseline + riskScore - taskScore
}

func CalculateRisks() []Risk {
	var risks []Risk
	weights := gpt.GetRiskWeights(config.Active.YearsInBusiness, config.Active.Industry)

	tasks := CalculateTasks()
	for _, weight := range weights {
		r := Risk{
			Name:   weight.RiskFactor,
			Weight: weight.RiskWeight,
			Score:  config.RiskFactorScores[weight.RiskFactor],
		}

		for _, t := range tasks {
			if t.Name == r.Name {
				if r.Score == 2 {
					r.Score = 1
				} else if r.Score == 1 {
					r.Score = 0.5
				} else {
					r.Score = 0
				}
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
