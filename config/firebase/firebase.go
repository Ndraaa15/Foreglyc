package firebase

import (
	"context"

	firebase "firebase.google.com/go"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/api/option"
)

func New() *firebase.App {
	ctx := context.Background()
	client, err := firebase.NewApp(ctx, nil, option.WithCredentialsFile(viper.GetString("firebase.credential_path")))
	if err != nil {
		logrus.WithError(err).Fatal("failed to create firebase client")
	}

	client.Storage(ctx)
	return client
}
