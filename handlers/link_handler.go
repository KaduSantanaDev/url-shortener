package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"url-shortner/database/repository"
)

type LinkHandler struct {
	repo repository.LinkRepository
}

func NewLinkHandler(repo repository.LinkRepository) *LinkHandler {
	return &LinkHandler{repo}
}

type createLinkRequest struct {
	URL  string `json:"url" binding:"required"`
	Slug string `json:"slug" binding:"required,min=6"`
}

func (handler *LinkHandler) CreateLink(context *gin.Context) {
	var body createLinkRequest

	if err := context.ShouldBindJSON(&body); err != nil {
		context.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	existingLink, _ := handler.repo.FindBySlug(body.Slug)
	if existingLink != nil {
		context.JSON(400, gin.H{
			"error": "link already exists",
		})
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "created successfully",
	})
}

type getLinkRequest struct {
	Slug string `json:"slug" binding:"required,min=6"`
}

func (handler *LinkHandler) GetLink(context *gin.Context) {
	var body getLinkRequest
	if err := context.ShouldBindJSON(&body); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	link, err := handler.repo.FindBySlug(body.Slug)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{})
		return
	}
	_, err = http.Get(link.Url)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Error while trying to access the link URL"})
		return
	}

	context.Redirect(http.StatusFound, link.Url)
}
