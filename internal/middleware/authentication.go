package middleware

import (
	"strings"

	"github.com/Ndraaa15/foreglyc-server/pkg/errx"
	"github.com/Ndraaa15/foreglyc-server/pkg/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func Authentication() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		authorization := c.Get("Authorization")

		if authorization == "" {
			logrus.Warn("missing authorization header")
			return errx.Unauthorized("missing token")
		}

		authorizations := strings.SplitN(authorization, " ", 2)
		if len(authorizations) != 2 || authorizations[0] != "Bearer" {
			logrus.Warn("invalid authorization header format")
			return errx.Unauthorized("invalid token format")
		}

		token := authorizations[1]
		payload, err := jwt.DecodeToken(token)
		if err != nil {
			logrus.WithError(err).Warn("failed to decode token")
			return err
		}

		c.Locals("userId", payload.ID)

		return c.Next()
	}
}
