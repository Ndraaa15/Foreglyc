package bootstrap

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go"
	_ai "github.com/Ndraaa15/foreglyc-server/config/ai"
	_cache "github.com/Ndraaa15/foreglyc-server/config/cache"
	_database "github.com/Ndraaa15/foreglyc-server/config/database"
	_email "github.com/Ndraaa15/foreglyc-server/config/email"
	"github.com/Ndraaa15/foreglyc-server/config/env"
	_firebase "github.com/Ndraaa15/foreglyc-server/config/firebase"
	_logger "github.com/Ndraaa15/foreglyc-server/config/logger"
	_router "github.com/Ndraaa15/foreglyc-server/config/router"
	_validator "github.com/Ndraaa15/foreglyc-server/config/validator"
	authhandler "github.com/Ndraaa15/foreglyc-server/internal/domain/auth/handler/http"
	authrepository "github.com/Ndraaa15/foreglyc-server/internal/domain/auth/repository"
	authservice "github.com/Ndraaa15/foreglyc-server/internal/domain/auth/service"
	chatbothandler "github.com/Ndraaa15/foreglyc-server/internal/domain/chatbot/handler/http"
	chatbotservice "github.com/Ndraaa15/foreglyc-server/internal/domain/chatbot/service"
	monitoringhandler "github.com/Ndraaa15/foreglyc-server/internal/domain/monitoring/handler/http"
	monitoringrepository "github.com/Ndraaa15/foreglyc-server/internal/domain/monitoring/repository"
	monitoringservice "github.com/Ndraaa15/foreglyc-server/internal/domain/monitoring/service"
	userhandler "github.com/Ndraaa15/foreglyc-server/internal/domain/user/handler/http"
	userrepository "github.com/Ndraaa15/foreglyc-server/internal/domain/user/repository"
	userservice "github.com/Ndraaa15/foreglyc-server/internal/domain/user/service"
	"github.com/Ndraaa15/foreglyc-server/internal/infra/ai"
	"github.com/Ndraaa15/foreglyc-server/internal/infra/cache"
	"github.com/Ndraaa15/foreglyc-server/internal/infra/email"
	"github.com/Ndraaa15/foreglyc-server/internal/infra/storage"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/genai"
	"gopkg.in/gomail.v2"
)

type Bootstrap struct {
	router    *fiber.App
	log       *logrus.Logger
	db        *sqlx.DB
	handlers  []Handler
	dialer    *gomail.Dialer
	message   *gomail.Message
	validator *validator.Validate
	cache     *redis.Client
	firebase  *firebase.App
	ai        *genai.Client
}

func New() *Bootstrap {
	env.Load()

	router := _router.New()
	log := _logger.New()
	db := _database.New()
	message, dialer := _email.New()
	validator := _validator.New()
	cache := _cache.New()
	firebase := _firebase.New()
	ai := _ai.New()

	return &Bootstrap{
		router:    router,
		log:       log,
		db:        db,
		dialer:    dialer,
		message:   message,
		validator: validator,
		cache:     cache,
		firebase:  firebase,
		ai:        ai,
	}
}

type Handler interface {
	SetEndpoint(router *fiber.App)
}

func (b *Bootstrap) DepedencyInjection() {
	firebaseStorageClient, err := b.firebase.Storage(context.Background())
	if err != nil {
		b.log.WithError(err).Fatal("failed to create firebase storage client")
	}

	firebaseStorageService := storage.New(firebaseStorageClient, b.log)
	geminiAiService := ai.New(b.ai, b.log)
	smtpEmailService := email.New(b.message, b.dialer)
	redisCacheRepository := cache.New(b.cache)

	userRepository := userrepository.New(b.db)
	userService := userservice.New(b.log, userRepository)
	userHandler := userhandler.New(userService, b.log, b.validator)

	authRepository := authrepository.New(b.db)
	authService := authservice.New(b.log, authRepository, redisCacheRepository, smtpEmailService)
	authHandler := authhandler.New(authService, b.log, b.validator)

	chatBotService := chatbotservice.New(b.log, geminiAiService, firebaseStorageService)
	chatBotHandler := chatbothandler.New(chatBotService, b.log, b.validator)

	monitoringRepository := monitoringrepository.New(b.db)
	monitoringService := monitoringservice.New(b.log, monitoringRepository, geminiAiService, userService)
	monitoringHandler := monitoringhandler.New(monitoringService, b.log, b.validator)

	b.handlers = []Handler{
		authHandler,
		chatBotHandler,
		monitoringHandler,
		userHandler,
	}
}

func (b *Bootstrap) Run() {
	b.DepedencyInjection()
	b.Health()

	b.router.Use(cors.New(cors.Config{
		AllowOrigins:  viper.GetString("server.cors.allow_origins"),
		AllowMethods:  viper.GetString("server.cors.allow_methods"),
		AllowHeaders:  viper.GetString("server.cors.allow_headers"),
		ExposeHeaders: viper.GetString("server.cors.expose_headers"),
		MaxAge:        viper.GetInt("server.cors.max_age"),
	}))

	for _, handler := range b.handlers {
		handler.SetEndpoint(b.router)
	}

	addr := fmt.Sprintf("%s:%d", viper.GetString("app.address"), viper.GetInt("app.port"))

	if err := b.router.Listen(addr); err != nil {
		b.log.WithError(err).Fatal("failed to start server")
	}
}

func (b *Bootstrap) Shutdown(ctx context.Context) {
	if err := b.router.Shutdown(); err != nil {
		b.log.WithError(err).Error("failed to shutdown server")
	}

	if err := b.db.Close(); err != nil {
		b.log.WithError(err).Error("failed to close database connection")
	}

	if err := b.log.Writer().Close(); err != nil {
		b.log.WithError(err).Error("failed to sync logger")
	}

	if err := b.cache.Close(); err != nil {
		b.log.WithError(err).Error("failed to close cache connection")
	}

	b.log.Info("server shutdown gracefully")
}

func (b *Bootstrap) Health() {
	b.router.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON("OK")
	})
}
