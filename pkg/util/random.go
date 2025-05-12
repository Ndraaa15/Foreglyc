package util

import (
	"crypto/rand"
	"math/big"

	"github.com/sirupsen/logrus"
)

const numberCharset = "0123456789"

func GenerateCode(length int) string {
	code := make([]byte, length)
	for i := range code {
		n, err := rand.Int(rand.Reader, big.NewInt(9))
		if err != nil {
			logrus.WithError(err).Error("failed to generate random number")
		}

		code[i] = numberCharset[n.Int64()]
	}
	return string(code)
}
