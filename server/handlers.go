package server

import (
	"62teknologi-senior-backend-test-muhammad_risky_nugraha/internal/contracts"
	"62teknologi-senior-backend-test-muhammad_risky_nugraha/internal/domain"
	"62teknologi-senior-backend-test-muhammad_risky_nugraha/internal/modules/business/usecase"
	"62teknologi-senior-backend-test-muhammad_risky_nugraha/internal/utils/response"

	"github.com/gofiber/fiber/v2"
)

type BusinessHandler struct {
	BusinessUseCase contracts.BusinessUseCase
}

func (s *Server) MapHandlers(app *fiber.App) error {
	handler := &BusinessHandler{}

	app.Get("/business/search", handler.GetBusinessData)
	app.Post("/business", handler.CreateBusinessData)
	app.Delete("/business", handler.DeleteBusinessData)
	app.Put("/business", handler.UpdateBusinessData)

	return nil
}

func (h *BusinessHandler) GetBusinessData(c *fiber.Ctx) error {
	ctx := c.Context()
	service := "62teknologi-senior-backend-test.business.list"

	queryParams := &domain.GetDataQuery{}

	if err := c.QueryParser(queryParams); err != nil {
		return response.ErrResponse(c, service, fiber.StatusBadRequest, "invalid_argument", err)
	}

	data, err := usecase.GetBusinessData(ctx, queryParams)
	if err != nil {
		return response.ErrResponse(c, service, fiber.StatusInternalServerError, "data_not_found", err)
	}

	return response.Success(c, service, "success_get_data", &response.Data{Key: "data", Value: data})
}

func (h *BusinessHandler) CreateBusinessData(c *fiber.Ctx) error {
	ctx := c.Context()
	service := "62teknologi-senior-backend-test.business.create"

	payload := &domain.Business{}

	if err := c.BodyParser(payload); err != nil {
		return response.ErrResponse(c, service, fiber.StatusBadRequest, "invalid_argument", err)
	}

	data, err := usecase.CreateBusinessData(ctx, payload)
	if err != nil {
		return response.ErrResponse(c, service, fiber.StatusInternalServerError, "data_not_found", err)
	}

	return response.Success(c, service, "success_create_data", &response.Data{Key: "data", Value: data})
}

func (h *BusinessHandler) DeleteBusinessData(c *fiber.Ctx) error {
	ctx := c.Context()
	service := "62teknologi-senior-backend-test.business.delete"

	payload := &DeleteBusinessPayload{}

	if err := c.BodyParser(payload); err != nil {
		return response.ErrResponse(c, service, fiber.StatusBadRequest, "invalid_argument", err)
	}

	err := usecase.DeleteBusinessData(ctx, payload.BusinessID)
	if err != nil {
		return response.ErrResponse(c, service, fiber.StatusInternalServerError, "data_not_found", err)
	}

	return response.Success(c, service, "success_delete_data", &response.Data{Key: "data", Value: ""})
}

func (h *BusinessHandler) UpdateBusinessData(c *fiber.Ctx) error {
	ctx := c.Context()
	service := "62teknologi-senior-backend-test.business.delete"

	payload := &domain.Business{}

	if err := c.BodyParser(payload); err != nil {
		return response.ErrResponse(c, service, fiber.StatusBadRequest, "invalid_argument", err)
	}

	err := usecase.UpdateBusinessData(ctx, payload)
	if err != nil {
		return response.ErrResponse(c, service, fiber.StatusInternalServerError, "data_not_found", err)
	}

	return response.Success(c, service, "success_delete_data", &response.Data{Key: "data", Value: ""})
}