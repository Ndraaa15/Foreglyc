package email

import (
	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
)

func New() (*gomail.Message, *gomail.Dialer) {
	return gomail.NewMessage(), gomail.NewDialer(
		viper.GetString("email.host"),
		viper.GetInt("email.port"),
		viper.GetString("email.user"),
		viper.GetString("email.password"),
	)
}
