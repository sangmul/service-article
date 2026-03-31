package handler

import (
	"article/internal/domain"
	"article/internal/service"
	"database/sql"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	service service.PostService
}

func NewPostHandler(s service.PostService) *PostHandler {
	return &PostHandler{service: s}
}

func validatePost(post domain.Post) error {
	// Title
	if strings.TrimSpace(post.Title) == "" {
		return errors.New("title is required")
	}
	if len(post.Title) < 20 {
		return errors.New("title minimal 20 karakter")
	}

	// Content
	if strings.TrimSpace(post.Content) == "" {
		return errors.New("content is required")
	}
	if len(post.Content) < 200 {
		return errors.New("content minimal 200 karakter")
	}

	// Category
	if strings.TrimSpace(post.Category) == "" {
		return errors.New("category is required")
	}
	if len(post.Category) < 3 {
		return errors.New("category minimal 3 karakter")
	}

	// Status
	status := strings.ToLower(strings.TrimSpace(post.Status))
	if status == "" {
		return errors.New("status is required")
	}

	allowed := map[string]bool{
		"publish": true,
		"draft":   true,
		"thrash":  true,
	}

	if !allowed[status] {
		return errors.New("status harus publish, draft, atau thrash")
	}

	return nil
}

func (h *PostHandler) Create(c *gin.Context) {
	var post domain.Post

	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 🔥 VALIDATION
	if err := validatePost(post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.Create(post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post created"})
}

// func (h *PostHandler) GetAll(c *gin.Context) {
// 	posts, err := h.service.GetAll()
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, posts)
// }

func (h *PostHandler) GetAll(c *gin.Context) {
    limitParam := c.Query("limit")
    offsetParam := c.Query("offset")

    // default value
    limit := 10
    offset := 0

    if limitParam != "" {
        l, err := strconv.Atoi(limitParam)
        if err == nil {
            limit = l
        }
    }

    if offsetParam != "" {
        o, err := strconv.Atoi(offsetParam)
        if err == nil {
            offset = o
        }
    }

    posts, err := h.service.GetWithPagination(limit, offset)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    var result []gin.H
    for _, p := range posts {
        result = append(result, gin.H{
            "title":    p.Title,
            "content":  p.Content,
            "category": p.Category,
            "status":   p.Status,
        })
    }

    c.JSON(http.StatusOK, result)
}

func (h *PostHandler) GetByID(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	post, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "article not found"})
		return
	}

	// sesuai format yang diminta (tanpa id, created_date, dll)
	c.JSON(http.StatusOK, gin.H{
		"title":    post.Title,
		"content":  post.Content,
		"category": post.Category,
		"status":   post.Status,
	})
}

func (h *PostHandler) Update(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var post domain.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 🔥 VALIDATION (reuse)
	if err := validatePost(post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.service.Update(id, post)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "article not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// response kosong sesuai requirement
	c.JSON(http.StatusOK, gin.H{})
}

func (h *PostHandler) Delete(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	err = h.service.Delete(id)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "article not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// sesuai requirement: response kosong
	c.JSON(http.StatusOK, gin.H{})
}
