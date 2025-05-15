package http

import (
	"github.com/Ndraaa15/foreglyc-server/internal/domain/food/service"
	"github.com/Ndraaa15/foreglyc-server/internal/middleware"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type FoodHandler struct {
	foodService service.IFoodService
	log         *logrus.Logger
	validator   *validator.Validate
}

func New(foodService service.IFoodService, log *logrus.Logger, validator *validator.Validate) *FoodHandler {
	return &FoodHandler{
		foodService: foodService,
		log:         log,
		validator:   validator,
	}
}

func (c *FoodHandler) SetEndpoint(router *fiber.App) {
	v1 := router.Group("/api/v1/foods")
	v1.Post("/generates/informations", middleware.Authentication(), c.GenerateFoodInformation)
	v1.Post("/dietary-plans", middleware.Authentication(), c.CreateDietaryPlan)
	v1.Patch("/dietary-plans/insulines", middleware.Authentication(), c.UpdateInsulineQuestionnaire)
	v1.Post("/monitorings", middleware.Authentication(), c.CreateFoodMonitoring)
	v1.Get("/status/3j/self", middleware.Authentication(), c.GetStatus3J)
	v1.Get("/recomendations/self", middleware.Authentication(), c.GetFoodRecomendation)
	v1.Get("/generates/dietary/informations/self", middleware.Authentication(), c.GenerateDietaryInformation)
	v1.Post("/dietary/informations", middleware.Authentication(), c.CreateDietaryInformation)
}
