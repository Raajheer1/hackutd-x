package gpt

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Raajheer1/hackutd-x/m/v2/pkg/config"
	"github.com/sashabaranov/go-openai"
	"os"
)

type RecommendationResponse struct {
	Title    string  `json:"title"`
	Summary  string  `json:"summary"`
	Details  string  `json:"details"`
	Category string  `json:"category"`
	Weight   float64 `json:"weight"`
}

type ReccomendationResponseTwo struct {
	Recommendations []RecommendationResponse `json:"recommendations"`
}

func GetRecommendation(msg string, category config.RecommendationTypes) []RecommendationResponse {
	client := openai.NewClient(os.Getenv("CHATGPT_API_KEY"))

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: fmt.Sprintf("You are a recommendation assistant API that returns JSON ONLY. I am going to give you a company and you need to provide me a list of 2 recommendations related to %s to decrease overall risk for that company additionally you need to weight on a scale of 0-1 how critical that recommendation is with 1 being very critical, the weights should sum to 1. Return a list of JSON recommendations with the keys: 'title', 'summary', 'details', 'weight'. Return JSON only.", category),
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: msg,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return nil
	}

	response := resp.Choices[0].Message.Content
	var recommendations []RecommendationResponse
	err = json.Unmarshal([]byte(response), &recommendations)
	if err != nil {
		fmt.Println("ChatCompletion error: Failing back to method 2")
		var recs ReccomendationResponseTwo
		err = json.Unmarshal([]byte(response), &recs)
		if err != nil {
			fmt.Printf("ChatCompletion error: %v\n", err)
			fmt.Printf("Given response:\n %s\n\n", response)
			return nil
		}
		return recs.Recommendations
	}

	return recommendations
}

type BaselineRiskResponse struct {
	BaselineRisk float64 `json:"baseline_risk"`
}

func GetBaselineRisk(age uint, industry string) BaselineRiskResponse {
	client := openai.NewClient(os.Getenv("CHATGPT_API_KEY"))

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "You are a risk assessor assistant API that returns JSON ONLY. I am going to give you a company's age and industry and you need to provide me a baseline risk factor on a 0-2 scale with 2 being the riskiest. Return a JSON with the key: 'baseline_risk'. Return JSON only.",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: fmt.Sprintf("Age: %d, Industry: %s", age, industry),
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return BaselineRiskResponse{
			BaselineRisk: 1.0,
		}
	}

	response := resp.Choices[0].Message.Content
	var baselineRisk BaselineRiskResponse
	err = json.Unmarshal([]byte(response), &baselineRisk)
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		fmt.Printf("Given response:\n %s\n\n", response)
		return BaselineRiskResponse{
			BaselineRisk: 1.0,
		}
	}

	return baselineRisk
}

type RiskResponse struct {
	RiskFactor string  `json:"risk_factor"`
	RiskWeight float64 `json:"risk_weight"`
}

type RiskResponseTwo struct {
	RiskWeights []RiskResponse `json:"risk_weights"`
}

type RiskResponseThree struct {
	RiskWeights []RiskResponse `json:"risk_factors"`
}

func GetRiskWeights(age uint, industry string) []RiskResponse {
	client := openai.NewClient(os.Getenv("CHATGPT_API_KEY"))

	riskFactors := ""
	for _, factor := range config.RecTypes {
		riskFactors += string(factor) + ", "
	}
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "You are a risk assessor assistant API that returns JSON ONLY. I am going to give you a company's age, industry and risk factors and you need to provide me the weights for the risk factors that will help me combine it into a single risk metric. The risk weights should sum to 1.0. Return a JSON list of objects with the keys: 'risk_factor' and 'risk_weight'. Finance should always be rated at 0.5. Return JSON only.",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: fmt.Sprintf("Age: %d, Industry: %s, Risk Factors:%s", age, industry, riskFactors),
				},
			},
		},
	)

	var riskWeights []RiskResponse

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return []RiskResponse{
			{
				RiskFactor: "Finance",
				RiskWeight: 0.5,
			},
			{
				RiskFactor: "Safety",
				RiskWeight: 0.1,
			},
			{
				RiskFactor: "Security Measures",
				RiskWeight: 0.1,
			},
			{
				RiskFactor: "Accident History",
				RiskWeight: 0.1,
			},
			{
				RiskFactor: "Employee Training",
				RiskWeight: 0.1,
			},
			{
				RiskFactor: "Legal",
				RiskWeight: 0.1,
			},
		}
	}

	response := resp.Choices[0].Message.Content
	err = json.Unmarshal([]byte(response), &riskWeights)
	if err != nil {
		fmt.Println("Weight ChatCompletion error: Failing back to method 2")
		fmt.Println(response)
		var recs RiskResponseTwo
		err = json.Unmarshal([]byte(response), &recs)
		if err != nil {
			fmt.Println("Weight ChatCompletion error: Failing back to method 3")
			var recs RiskResponseThree
			err = json.Unmarshal([]byte(response), &recs)
			if err != nil {
				fmt.Printf("Weight ChatCompletion error: %v\n", err)
				fmt.Printf("Given response:\n %s\n\n", response)
				return []RiskResponse{
					{
						RiskFactor: "Finance",
						RiskWeight: 0.5,
					},
					{
						RiskFactor: "Safety",
						RiskWeight: 0.1,
					},
					{
						RiskFactor: "Security Measures",
						RiskWeight: 0.1,
					},
					{
						RiskFactor: "Accident History",
						RiskWeight: 0.1,
					},
					{
						RiskFactor: "Employee Training",
						RiskWeight: 0.1,
					},
					{
						RiskFactor: "Legal",
						RiskWeight: 0.1,
					},
				}
			}
		}
		return recs.RiskWeights
	}

	return riskWeights
}

func Conversation(pastMessages []openai.ChatCompletionMessage) []openai.ChatCompletionMessage {
	client := openai.NewClient(os.Getenv("CHATGPT_API_KEY"))

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: append([]openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "You are a Insurance Agent at State Farm here to help answer questions and provide guidance to your clients.",
				},
			}, pastMessages...),
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return nil
	}

	return append(pastMessages, resp.Choices[0].Message)
}
