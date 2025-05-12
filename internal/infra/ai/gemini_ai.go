package ai

import (
	"context"
	"fmt"

	"github.com/Ndraaa15/foreglyc-server/pkg/constant"
	"github.com/sirupsen/logrus"
	"google.golang.org/genai"
)

type IGemini interface {
	ChatForeglycExpert(ctx context.Context, contents []*genai.Content) (string, error)
}

type Gemini struct {
	client *genai.Client
	log    *logrus.Logger
}

func New(client *genai.Client, log *logrus.Logger) IGemini {
	return &Gemini{
		client: client,
		log:    log,
	}
}

func (g *Gemini) ChatForeglycExpert(ctx context.Context, contents []*genai.Content) (string, error) {
	fmt.Println("contents", contents)

	response, err := g.client.Models.GenerateContent(ctx, constant.GeminiModel, contents, &genai.GenerateContentConfig{
		SystemInstruction: &genai.Content{
			Parts: []*genai.Part{
				{
					Text: `You are a helpful health assistant specialized in diabetes prevention. Provide informative, practical, and friendly advice about healthy eating, physical activity, and lifestyle choices to help users prevent type 2 diabetes. Always respond clearly and avoid medical jargon.
					
					Please answer the question with the same language as the user.
					`,
				},
			},
		},
	})
	if err != nil {
		g.log.WithError(err).Error("failed to generate content")
		return "", err
	}

	return response.Text(), nil
}

func (g *Gemini) TestN8N(ctx context.Context, request string) {
	partA := &genai.Part{
		Text: `
		Your task is to generate a response to the following request. You are not allowed to use any external resources or APIs. You must generate the response based on your own knowledge and understanding of the topic.
		`,
		FunctionResponse: &genai.FunctionResponse{
			ID:   "testN8N",
			Name: "testN8N",
			Response: map[string]any{
				"prompt":   "",
				"response": "",
			},
		},
	}

	instruction := &genai.Content{
		Parts: []*genai.Part{partA},
	}

	resp, err := g.client.Models.GenerateContent(ctx, constant.GeminiModel, genai.Text("generate food recomendatation based on the data that you got in database with table foods"), &genai.GenerateContentConfig{
		SystemInstruction: instruction,
	})

	if err != nil {
		logrus.WithError(err).Error("failed to generate content")
	}

	fmt.Println(resp.Text())
}
