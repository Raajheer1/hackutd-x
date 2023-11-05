package ipad

import (
	"fmt"
	"github.com/Raajheer1/hackutd-x/m/v2/pkg/config"
	"github.com/Raajheer1/hackutd-x/m/v2/pkg/data"
	"github.com/Raajheer1/hackutd-x/m/v2/pkg/database/models"
	"github.com/Raajheer1/hackutd-x/m/v2/pkg/echo/dto"
	"github.com/Raajheer1/hackutd-x/m/v2/pkg/gpt"
	"github.com/labstack/echo/v4"
	"github.com/sashabaranov/go-openai"
	"net/http"
	"sync"
)

type Config struct {
	CompanyName string          `json:"companyName"`
	Industry    string          `json:"industry"`
	City        string          `json:"city"`
	Presets     []models.Preset `json:"presets"`
}

// For the iPad show all the configs
// getConfigs godoc
// @Summary Get all configs
// @Tags iPad
// @Accept json
// @Produce json
// @Success 200 {object} []models.Preset
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /configs [get]
func getConfigs(c echo.Context) error {
	return c.JSON(http.StatusOK, config.Configs)
}

// From the iPad sets the current company with values

type container struct {
	mu   sync.Mutex
	recs []config.Recommendation
}

func setActiveConfig(c echo.Context) error {
	var req config.ActiveConfig
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: err.Error(),
		})
	}

	config.Active = req

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
					Weight:    r.Weight,
					Completed: false,
				})
				cont.mu.Unlock()
			}
		}(rec)
	}
	wg.Wait()

	config.Recs = cont.recs
	data.Weights = nil
	data.BaselineRisk = 0

	return c.JSON(http.StatusOK, dto.MessageResponse{
		Message: "Active config updated successfully",
	})
}

// For the user show the current config

type overall struct {
	RiskValue          float64             `json:"riskValue"`
	ZValue             float64             `json:"zValue"`
	BankruptcyRisk     string              `json:"bankruptcyRisk"`
	WeightedBaseline   float64             `json:"weightedBaseline"`
	WeightedRecs       float64             `json:"weightedRecs"`
	WeightedBankruptcy float64             `json:"weightedBankruptcy"`
	Active             config.ActiveConfig `json:"active"`
}

func getActiveConfig(c echo.Context) error {
	var o overall
	o.Active = config.Active
	o.RiskValue = data.CalculateRisk()
	o.ZValue = data.CalculateZ()
	o.BankruptcyRisk = data.Situation(o.ZValue)
	if o.ZValue < 1.23 {
		added := false
		for _, rec := range config.Recs {
			if rec.Title == "Get Liability Insurance" {
				added = true
			}
		}
		if !added {
			config.Recs = append(config.Recs, config.Recommendation{
				Title:     "Get Liability Insurance",
				Summary:   "Protect Against the Unexpected: Liability Insurance as Your Business Safety Net",
				Details:   "For small businesses navigating financial challenges, liability insurance is a critical tool for risk management. It offers protection against the high costs associated with legal claims and lawsuits, which can be the difference between survival and closure. When facing potential bankruptcy, liability insurance acts as a shield, guarding your hard-earned business against the threats that could exacerbate financial strain.",
				Type:      "Legal",
				Weight:    1,
				Completed: false,
			})
		}
	}

	o.WeightedBankruptcy = data.BankruptcyWeight
	o.WeightedBaseline = data.BaselineWeight
	o.WeightedRecs = data.RecWeight

	return c.JSON(http.StatusOK, o)
}

// For the user show all the recommendations
func getRecommendations(c echo.Context) error {
	return c.JSON(http.StatusOK, config.Recs)
}

// For the user mark a recommendation as complete
// Should update Z value and possibly create new recommendations
// Passes in title of recommendation

type CompleteRecommendationRequest struct {
	Title string `json:"title"`
}

func toggleRecommendation(c echo.Context) error {
	var req CompleteRecommendationRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: err.Error(),
		})
	}

	for i, rec := range config.Recs {
		if rec.Title == req.Title {
			config.Recs[i].Completed = !config.Recs[i].Completed
			return c.JSON(http.StatusOK, dto.MessageResponse{
				Message: "Recommendation updated successfully",
			})
		}
	}

	return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
		Error: "Recommendation not found",
	})
}

func timeSeriesBeauty(c echo.Context) error {
	return nil
}

func ChatCompletion(c echo.Context) error {
	var req []openai.ChatCompletionMessage
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: err.Error(),
		})
	}

	resp := gpt.Conversation(req)
	return c.JSON(http.StatusOK, resp)
}
