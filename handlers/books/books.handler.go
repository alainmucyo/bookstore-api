package books

import (
	"bookstore/core/books"
	"bookstore/handlers/dtos"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Handler struct {
	service *books.Service
}

func New(service *books.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) FindAll(c *gin.Context) {
	allBooks, err := h.service.FindAll()
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"error": "Something went wrong",
		})
		return
	}

	c.JSON(200, gin.H{
		"data": allBooks,
	})

}

func (h *Handler) FindById(c *gin.Context) {
	id := c.Param("id")
	book, err := h.service.FindById(id)
	if err != nil {
		c.AbortWithStatusJSON(404, gin.H{
			"error": "Book not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"data": book,
	})
}

func (h *Handler) Store(c *gin.Context) {
	var book dtos.BookCreateDTO
	if err := c.BindJSON(&book); err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"error": "Invalid JSON object",
		})
		return
	}

	if err := book.Validate(); err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	if _, err := h.service.Store(book.ToEntity()); err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"error": "Something went wrong",
		})
		return
	}

	c.JSON(201, gin.H{
		"data": book,
	})
}

func (h *Handler) Update(c *gin.Context) {
	id := c.Param("id")
	var book dtos.BookUpdateDTO
	if err := c.BindJSON(&book); err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"error": "Invalid JSON object",
		})
		return
	}

	if err := book.Validate(); err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	foundBook, err := h.service.FindById(id)
	if err != nil {
		c.AbortWithStatusJSON(404, gin.H{
			"error": "Book not found",
		})
		return
	}

	if book.Title != "" {
		foundBook.Title = book.Title
	}

	if book.Price != 0 {
		foundBook.Price = book.Price
	}

	if book.ImageUrl != "" {
		foundBook.ImageUrl = book.ImageUrl

	}
	if book.AuthorID != "" {
		authorIdUUID, err := uuid.Parse(book.AuthorID)
		if err != nil {
			c.AbortWithStatusJSON(400, gin.H{
				"error": "Invalid author ID",
			})
			return
		}
		foundBook.AuthorID = authorIdUUID
	}
	if _, err := h.service.Update(id, foundBook); err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"error": "Something went wrong",
		})
		return
	}

	c.JSON(200, gin.H{
		"data": book,
	})
}
