package ai

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/genai"
)

func New() *genai.Client {
	ctx := context.Background()
	client, err := genai.NewClient(ctx,
		&genai.ClientConfig{
			APIKey:  viper.GetString("gemini.api_key"),
			Backend: genai.BackendGeminiAPI,
		},
	)
	if err != nil {
		logrus.WithError(err).Fatal("failed to create Gemini client")
	}

	return client
}
