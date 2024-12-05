package handlers

import (
	"net/http"

	"github.com/0sokrat0/GoGRAFFApi.git/app/internal/auth"
	"github.com/0sokrat0/GoGRAFFApi.git/app/internal/models"
	"github.com/0sokrat0/GoGRAFFApi.git/app/internal/repositories"
	"github.com/0sokrat0/GoGRAFFApi.git/app/pkg/utils"
	"github.com/gin-gonic/gin"
)

// LoginInput описывает входные данные для входа пользователя
type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// RegisterInput описывает входные данные для регистрации пользователя
type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthHandler struct {
	repo *repositories.AuthUserRepository
}

func NewAuthHandler(repo *repositories.AuthUserRepository) *AuthHandler {
	return &AuthHandler{repo: repo}
}

// LoginHandler handles user login
// @Summary User login
// @Description Authenticates the user and returns a JWT token
// @Tags Authentication
// @Accept json
// @Produce json
// @Param input body LoginInput true "User credentials"
// @Success 200 {object} map[string]interface{} "JWT token"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 401 {object} map[string]interface{} "Invalid credentials"
// @Router /auth/login [post]
func (h *AuthHandler) LoginHandler(c *gin.Context) {
	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid input"))
		return
	}

	user, err := h.repo.FindByUsername(input.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, utils.ErrorResponse("User not found"))
		return
	}

	if !auth.CheckPassword(user.Password, input.Password) {
		c.JSON(http.StatusUnauthorized, utils.ErrorResponse("Invalid password"))
		return
	}

	token, err := auth.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to generate token"))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(map[string]interface{}{
		"token": token,
	}))
}

// RegisterHandler handles user registration
// @Summary Register a new user
// @Description Creates a new user account
// @Tags Authentication
// @Accept json
// @Produce json
// @Param input body RegisterInput true "User registration data"
// @Success 201 {string} string "User registered successfully"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 409 {object} map[string]interface{} "Username already exists"
// @Failure 500 {object} map[string]interface{} "Failed to register user"
// @Router /auth/register [post]
func (h *AuthHandler) RegisterHandler(c *gin.Context) {
	var input RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid input"))
		return
	}

	_, err := h.repo.FindByUsername(input.Username)
	if err == nil {
		c.JSON(http.StatusConflict, utils.ErrorResponse("Username already exists"))
		return
	}

	hashedPassword, err := auth.HashPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to hash password"))
		return
	}

	newUser := &models.AuthUser{
		Username: input.Username,
		Password: hashedPassword,
	}

	if err := h.repo.Create(newUser); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to register user"))
		return
	}

	c.JSON(http.StatusCreated, utils.SuccessResponse("User registered successfully"))
}
