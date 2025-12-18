package handler

import (
	"time"

	"github.com/faisal-990/age/internal/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type UserHandler struct {
	s        service.UserService
	log      *zap.Logger
	validate *validator.Validate // Validator dependency
}

// New creates a new UserHandler with Service, Logger, and Validator dependencies
func New(svc service.UserService, log *zap.Logger, v *validator.Validate) *UserHandler {
	return &UserHandler{
		s:        svc,
		log:      log,
		validate: v,
	}
}

// --- 1. CREATE USER ---
func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	// Local DTO with Validation Tags
	type Request struct {
		Name string `json:"name" validate:"required,min=3"`
		Dob  string `json:"dob" validate:"required,datetime=2006-01-02"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		h.log.Error("CreateUser: Failed to parse body", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	// VALIDATION STEP
	if err := h.validate.Struct(&req); err != nil {
		h.log.Warn("CreateUser: Validation failed", zap.Error(err))
		// Return the validation error message directly
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Validation failed",
			"details": err.Error(),
		})
	}

	// Manual Date Parsing (Safe to do now because validation passed)
	parsedDob, _ := time.Parse("2006-01-02", req.Dob)

	svcReq := service.CreateUserRequest{
		Name: req.Name,
		Dob:  parsedDob,
	}

	res, err := h.s.CreateUser(c.Context(), svcReq)
	if err != nil {
		h.log.Error("CreateUser: Service failure", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	h.log.Info("User created successfully",
		zap.Int32("user_id", res.ID),
		zap.String("user_name", res.Name),
	)

	return c.Status(fiber.StatusCreated).JSON(res)
}

// --- 2. GET USER ---
func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		h.log.Warn("GetUser: Invalid ID format", zap.String("id_param", c.Params("id")))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid User ID"})
	}

	user, err := h.s.GetUser(c.Context(), int32(id))
	if err != nil {
		h.log.Error("GetUser: Service failure", zap.Int("user_id", id), zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if user == nil {
		h.log.Debug("GetUser: User not found", zap.Int("user_id", id))
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

// --- 3. LIST USERS ---
func (h *UserHandler) ListUsers(c *fiber.Ctx) error {
	users, err := h.s.ListUsers(c.Context())
	if err != nil {
		h.log.Error("ListUsers: Service failure", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if users == nil {
		return c.Status(fiber.StatusOK).JSON([]interface{}{})
	}

	return c.Status(fiber.StatusOK).JSON(users)
}

// --- 4. UPDATE USER ---
func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		h.log.Warn("UpdateUser: Invalid ID format", zap.String("id_param", c.Params("id")))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid User ID"})
	}

	// Local DTO with Validation Tags
	type Request struct {
		Name string `json:"name" validate:"required,min=3"`
		Dob  string `json:"dob" validate:"required,datetime=2006-01-02"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		h.log.Error("UpdateUser: Failed to parse body", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid body"})
	}

	// VALIDATION STEP
	if err := h.validate.Struct(&req); err != nil {
		h.log.Warn("UpdateUser: Validation failed", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Validation failed",
			"details": err.Error(),
		})
	}

	parsedDob, _ := time.Parse("2006-01-02", req.Dob)

	svcReq := service.UpdateUserRequest{
		ID:   int32(id),
		Name: req.Name,
		Dob:  parsedDob,
	}

	res, err := h.s.UpdateUser(c.Context(), &svcReq)
	if err != nil {
		h.log.Error("UpdateUser: Service failure", zap.Int("user_id", id), zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	h.log.Info("User updated successfully", zap.Int("user_id", id))

	return c.Status(fiber.StatusOK).JSON(res)
}

// --- 5. DELETE USER ---
func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		h.log.Warn("DeleteUser: Invalid ID format", zap.String("id_param", c.Params("id")))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid User ID"})
	}

	if err := h.s.DeleteUser(c.Context(), int32(id)); err != nil {
		h.log.Error("DeleteUser: Service failure", zap.Int("user_id", id), zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	h.log.Info("User deleted successfully", zap.Int("user_id", id))

	return c.SendStatus(fiber.StatusNoContent)
}

