package ai

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/Ndraaa15/foreglyc-server/internal/domain/chatbot/dto"
	fooddto "github.com/Ndraaa15/foreglyc-server/internal/domain/food/dto"
	"github.com/Ndraaa15/foreglyc-server/pkg/constant"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/genai"
)

type IGemini interface {
	ChatForeglycExpert(ctx context.Context, contents []*genai.Content) (string, error)
	RecomendationAboutQuestionnaire(ctx context.Context, contents []*genai.Content) (string, error)
	GenerateFoodInformation(ctx context.Context, contents []*genai.Content) (fooddto.FoodInformationResponse, error)
	GlucosePredictionN8N(ctx context.Context, userId string, glucometerMonitoringIds []int64) ([]dto.ScenarioResponse, error)
	GenerateDietaryInformation(ctx context.Context, contents []*genai.Content) (fooddto.DietaryInformationChatbotResponse, error)
	FoodRecomendationsN8N(ctx context.Context, userId string) ([]fooddto.MenuChatBotResponse, error)
	ChatForeglycExpertWithPrediction(ctx context.Context, contents []*genai.Content) (string, error)
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

					If the user asks about diabetes management, offer firm, actionable recommendations for Self Management or Management Requires Assistance.
					
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

func (g *Gemini) ChatForeglycExpertWithPrediction(ctx context.Context, contents []*genai.Content) (string, error) {
	response, err := g.client.Models.GenerateContent(ctx, constant.GeminiModel, contents, &genai.GenerateContentConfig{
		SystemInstruction: &genai.Content{
			Parts: []*genai.Part{
				{
					Text: `You are a helpful health assistant specialized in diabetes prevention. Provide informative, practical, and friendly advice about healthy eating, physical activity, and lifestyle choices to help users prevent type 2 diabetes. Always respond clearly and avoid medical jargon.

					There are a prediction about the user's glucose level, please answer the question using knowledge from the prediction.
					
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

func (g *Gemini) GlucosePredictionN8N(ctx context.Context, userId string, glucometerMonitoringIDs []int64) ([]dto.ScenarioResponse, error) {
	basicAuthHeader := basicAuth(
		viper.GetString("n8n.username"),
		viper.GetString("n8n.password"),
	)

	payload := map[string]interface{}{
		"userId":                  userId,
		"monitoringGlucometerIds": glucometerMonitoringIDs,
	}

	bodyBytes, err := json.Marshal(payload)
	if err != nil {
		g.log.WithError(err).Error("failed to marshal request payload")
		return nil, err
	}

	url := viper.GetString("n8n.url") + viper.GetString("n8n.prediction_uri")

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(bodyBytes))
	if err != nil {
		g.log.WithError(err).Error("failed to create request")
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", basicAuthHeader)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		g.log.WithError(err).WithField("url", url).Error("request failed")
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		data, _ := io.ReadAll(resp.Body)
		g.log.Errorf("prediction API returned status %d: %s", resp.StatusCode, string(data))
		return nil, fmt.Errorf("prediction API error: status %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		g.log.WithError(err).Error("failed to read response body")
		return nil, err
	}

	if len(data) == 0 {
		g.log.Error("prediction API returned empty response body")
		return nil, fmt.Errorf("empty response from prediction API")
	}

	g.log.Debugf("Prediction response payload: %s", string(data))

	var scenarios []dto.ScenarioResponse
	if err := json.Unmarshal(data, &scenarios); err != nil {
		g.log.WithError(err).
			WithField("body", string(data)).
			Error("failed to unmarshal scenarios")
		return nil, fmt.Errorf("invalid JSON format: %w", err)
	}

	return scenarios, nil
}

func (g *Gemini) FoodRecomendationsN8N(ctx context.Context, userId string) ([]fooddto.MenuChatBotResponse, error) {
	basicAuthHeader := basicAuth(
		viper.GetString("n8n.username"),
		viper.GetString("n8n.password"),
	)

	payload := map[string]interface{}{
		"userId": userId,
	}
	bodyBytes, err := json.Marshal(payload)
	if err != nil {
		g.log.WithError(err).Error("failed to marshal request payload")
		return nil, err
	}

	var menuResponses []fooddto.MenuChatBotResponse

	for i := 0; i < 7; i++ {
		var menuResponsen8n []fooddto.FoodRecommendationChatBotResponse

		url := viper.GetString("n8n.url") + viper.GetString("n8n.food_recomendation_uri")
		req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(bodyBytes))
		if err != nil {
			g.log.WithError(err).Error("failed to create request")
			return nil, err
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", basicAuthHeader)

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			g.log.WithError(err).WithField("url", url).Error("request failed")
			return nil, err
		}
		defer resp.Body.Close()

		if resp.StatusCode < 200 || resp.StatusCode >= 300 {
			data, _ := io.ReadAll(resp.Body)
			g.log.Errorf("prediction API returned status %d: %s", resp.StatusCode, string(data))
			return nil, fmt.Errorf("prediction API error: status %d", resp.StatusCode)
		}

		data, err := io.ReadAll(resp.Body)
		if err != nil {
			g.log.WithError(err).Error("failed to read response body")
			return nil, err
		}

		if len(data) == 0 {
			g.log.Error("prediction API returned empty response body")
			return nil, fmt.Errorf("empty response from prediction API")
		}

		if err := json.Unmarshal(data, &menuResponsen8n); err != nil {
			g.log.WithError(err).
				WithField("body", string(data)).
				Error("failed to unmarshal scenarios")
			return nil, fmt.Errorf("invalid JSON format: %w", err)
		}

		menuResponses = append(menuResponses, fooddto.MenuChatBotResponse{
			Date:               time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day()+i, 0, 0, 0, 0, time.Local).Format("Mon, 02 Jan 2006"),
			FoodRecomendations: menuResponsen8n,
		})
	}

	return menuResponses, nil
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
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

func (g *Gemini) GenerateFoodInformation(ctx context.Context, contents []*genai.Content) (fooddto.FoodInformationResponse, error) {
	response, err := g.client.Models.GenerateContent(ctx, constant.GeminiModel, contents, &genai.GenerateContentConfig{
		SystemInstruction: &genai.Content{
			Parts: []*genai.Part{
				{
					Text: `
					You are a food analysis assistant. Given an image of a dish, identify each macronutrient group and its components, estimate portions and calories, and output _only_ a JSON object with this exact structure:

					{
					"foodName": "<dish name>",
					"nutritions": [
						{
							"type": "<macronutrient name, e.g. carbohydrate>",
							"components": [
							{
								"name": "<food item name>",
								"portion": <estimated portion size in float>,
								"unit" : "<unit, e.g. g or cup>",
								"calory": <integer calories>
							},
							…
							]
						},
						…
					],
					"totalCalory": <integer total calories>,
					"totalCarbohydrate": <integer total carbohydrates>,
					"totalProtein": <integer total protein>,
					"totalFat": <integer total fat>
					}

					Do not include any explanatory text or additional fields.
					`,
				},
			},
		},
	})
	if err != nil {
		g.log.WithError(err).Error("failed to generate content")
		return fooddto.FoodInformationResponse{}, err
	}

	raw := response.Text()
	cleaned := strings.Trim(raw, "` \n\r\t")
	cleaned = strings.TrimLeft(cleaned, "json")

	var info fooddto.FoodInformationResponse
	if err := json.Unmarshal([]byte(cleaned), &info); err != nil {
		g.log.WithError(err).
			WithField("raw", raw).
			Error("failed to parse food nutrition JSON")
		return fooddto.FoodInformationResponse{}, fmt.Errorf("invalid JSON format: %w", err)
	}

	return info, nil
}

func (g *Gemini) GenerateDietaryInformation(ctx context.Context, contents []*genai.Content) (fooddto.DietaryInformationChatbotResponse, error) {
	response, err := g.client.Models.GenerateContent(ctx, constant.GeminiModel, contents, &genai.GenerateContentConfig{
		SystemInstruction: &genai.Content{
			Parts: []*genai.Part{
				{
					Text: `You are a professional nutrition assistant.

							You will be provided with:
							- A user's current dietary plan (including meals and snacks), and
							- A list of foods and nutritional values consumed over the past several days.

							Your tasks:
							1. Analyze the given data to estimate the user's average daily caloric intake.
							2. Calculate the total calories consumed per day, including per-meal breakdown: breakfast, snacks, lunch, and dinner.
							3. Provide a target daily calorie intake recommendation that promotes a healthier diet based on the user's consumption patterns.
							4. Include a short explanation and recommendations to help the user meet this daily caloric goal.
							5. If data about breakfast, snack, lunch, or dinner is missing, provide a reasonable estimate based on the user's average daily caloric intake.

							Return only a valid JSON object in the exact format below:

							{
							"totalCalory": <integer total daily calories>,
							"totalBreakfastCalory": <integer breakfast calories>,
							"totalSnackCalory": <integer snack calories>,
							"totalLunchCalory": <integer lunch calories>,
							"totalDinnerCalory": <integer dinner calories>,
							"recomendation": "<brief recommendation and reasoning>"
							}

							Important rules:
							- The "totalCalory" must equal the sum of all individual meal calories.
							- The "recomendation" must be concise and actionable.
							- Do not include any additional fields, explanations, or formatting outside the JSON object.
							- The response must be valid and parsable JSON only.
							`,
				},
			},
		},
	})

	if err != nil {
		g.log.WithError(err).Error("failed to generate content")
		return fooddto.DietaryInformationChatbotResponse{}, err
	}

	raw := response.Text()
	cleaned := strings.Trim(raw, "` \n\r\t")
	cleaned = strings.TrimLeft(cleaned, "json")

	var info fooddto.DietaryInformationChatbotResponse
	if err := json.Unmarshal([]byte(cleaned), &info); err != nil {
		g.log.WithError(err).
			WithField("raw", raw).
			Error("failed to parse food information chatbot response")
		return fooddto.DietaryInformationChatbotResponse{}, fmt.Errorf("invalid JSON format: %w", err)
	}

	return info, nil
}
