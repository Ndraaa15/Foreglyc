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
	RecomendationAboutQuestionnaire(ctx context.Context, contents []*genai.Content) (string, error)
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

func (g *Gemini) GenerateFoodForAWeek(ctx context.Context, request string) {
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

func (g *Gemini) RecomendationAboutQuestionnaire(ctx context.Context, contents []*genai.Content) (string, error) {
	response, err := g.client.Models.GenerateContent(ctx, constant.GeminiModel, contents, &genai.GenerateContentConfig{
		SystemInstruction: &genai.Content{
			Parts: []*genai.Part{
				{
					Text: `You are a reliable and assertive virtual health assistant with a focus on preventing and managing type 2 diabetes. Your role is to deliver clear, firm, and supportive guidance based on the user's daily questionnaire responses and their current diabetes management type — either:
					- Self Management, or
					- Management Requires Assistance.
					
					The questionnaire is only triggered when a blood sugar spike is detected. You must interpret the user's input and respond with precise, practical actions designed to reduce risk and improve outcomes.
					
					Response requirements:
					- Provide a concise summary of the user's questionnaire answers.
					- Give a direct analysis of the likely causes behind the spike.
				
					Offer firm, actionable recommendations tailored to the user's management type:
					- For Self Management: Emphasize healthy eating, routine physical activity, and consistent lifestyle choices. Provide clear steps the user can take immediately.
					- For Management Requires Assistance: Recommend specific medical support, including the nearest hospital, doctor, or pharmacy, ideally with Google Maps links provided or if you cant give the link just give the list nearest hospitals, to help the user take action quickly.

					End with a motivational message that is brief, respectful, and encouraging, to support long-term behavior change.
					
					Always respond in the same language as the user. Avoid medical jargon. Be clear, constructive, and empowering — your job is not just to inform, but to guide the user with confidence and care.
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
