package users

import (
	"bookstore/core/users"
	"bookstore/handlers/dtos"
	"bookstore/store/entities"
	jwt2 "bookstore/store/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	service *users.Service
	jwt     *jwt2.JWT
}

func New(service *users.Service, jwt *jwt2.JWT) *Handler {
	return &Handler{service: service, jwt: jwt}
}

func (h *Handler) Check(c *gin.Context) {
	email, exists := c.Get("email")
	if !exists {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid token",
		})
		return
	}

	user, err := h.service.FindByEmail(email.(string))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func (h *Handler) Register(c *gin.Context) {
	var registerDto dtos.RegisterDTO
	if err := c.BindJSON(&registerDto); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON object",
		})
		return
	}
	if err := registerDto.Validate(); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user := entities.User{Email: registerDto.Email, Name: registerDto.Name, Password: registerDto.Password}
	createdUser, err := h.service.Store(user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": createdUser,
	})
}

func (h *Handler) Login(c *gin.Context) {
	var loginDto dtos.LoginDTO
	if err := c.BindJSON(&loginDto); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON object",
		})
		return
	}
	verified, user := h.service.VerifyUser(loginDto.Email, loginDto.Password)
	if !verified {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	token, err := h.jwt.Generate(user.Name, user.Email, user.Id.String())

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to generate token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": token,
	})
}
