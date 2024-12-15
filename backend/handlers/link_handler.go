package handlers

import (
	"fmt"
	"net/http"
	"url-shortner/database/models"
	"url-shortner/database/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LinkHandler struct {
	repo repository.LinkRepository
}

func NewLinkHandler(repo repository.LinkRepository) *LinkHandler {
	return &LinkHandler{repo}
}

type createLinkRequest struct {
	URL  string `json:"url" binding:"required"`
	Slug string `json:"slug" binding:"required,min=3"`
}

func (handler *LinkHandler) CreateLink(context *gin.Context) {
	var body createLinkRequest

	if err := context.ShouldBindJSON(&body); err != nil {
		context.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	fmt.Println(body)

	link := &models.Link{
		Url:      body.URL,
		Slug:     body.Slug,
		ShortUrl: context.Request.Host + "/" + body.Slug,
	}

	err := handler.repo.CreateLink(link)
	if err != gorm.ErrDuplicatedKey {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Esse slug já existe!",
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "created successfully",
		"newURL": link.ShortUrl,
	})
}

type getLinkRequestURI struct {
	Slug string `uri:"slug" binding:"required"`
}

func (handler *LinkHandler) GetLink(context *gin.Context) {
	var uri getLinkRequestURI
	if err := context.ShouldBindUri(&uri); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	link, err := handler.repo.FindBySlug(uri.Slug)
	if link.Slug != uri.Slug || err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"error": "link not found",
		})

		return
	}

	context.Redirect(http.StatusFound, link.Url)
}
