package authors

import (
	"bookstore/core/authors"
	"bookstore/handlers/dtos"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	service *authors.Service
}

func New(service *authors.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) FindAll(c *gin.Context) {
	allAuthors, err := h.service.FindAll()
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"error": "Something went wrong",
		})
		return
	}

	c.JSON(200, gin.H{
		"data": allAuthors,
	})

}

func (h *Handler) FindById(c *gin.Context) {
	id := c.Param("id")
	author, err := h.service.FindById(id)
	if err != nil {
		c.AbortWithStatusJSON(404, gin.H{
			"error": "Author not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"data": author,
	})
}

func (h *Handler) Store(c *gin.Context) {
	var author dtos.AuthorCreateDTO
	if err := c.BindJSON(&author); err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"error": "Invalid JSON object",
		})
		return
	}

	if err := author.Validate(); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	createdAuthor, err := h.service.Store(author.ToAuthorEntity())
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"error": "Something went wrong",
		})
		return
	}

	c.JSON(200, gin.H{
		"data": createdAuthor,
	})
}

func (h *Handler) Update(c *gin.Context) {
	id := c.Param("id")
	var author dtos.AuthorUpdateDTO
	if err := c.BindJSON(&author); err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"error": "Invalid JSON object",
		})
		return
	}

	if err := author.Validate(); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	foundAuthor, err := h.service.FindById(id)
	if err != nil {
		c.AbortWithStatusJSON(404, gin.H{
			"error": "Author not found",
		})
		return
	}

	if author.FirstName != "" {
		foundAuthor.FirstName = author.FirstName
	}

	if author.LastName != "" {
		foundAuthor.LastName = author.LastName
	}

	if author.Bio != "" {
		foundAuthor.Bio = author.Bio
	}

	updatedAuthor, err := h.service.Update(id, foundAuthor)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"error": "Something went wrong",
		})
		return
	}

	c.JSON(200, gin.H{
		"data": updatedAuthor,
	})
}

func (h *Handler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.service.Delete(id)
	if err != nil {
		c.AbortWithStatusJSON(404, gin.H{
			"error": "Author not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Author deleted successfully",
	})
}
