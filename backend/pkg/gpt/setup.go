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
	Title    string `json:"title"`
	Summary  string `json:"summary"`
	Details  string `json:"details"`
	Category string `json:"category"`
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
					Content: fmt.Sprintf("You are a recommendation assistant API that returns JSON ONLY. I am going to give you a company and you need to provide me a list of 2 recommendations related to %s to decrease overall risk for that company. Return a list of JSON recommendations with the keys: 'title', 'summary', 'details'. Return JSON only.", category),
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
