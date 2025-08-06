package controller

import (
	postErrors "blog/domains/posts/errors"
	"blog/domains/posts/models"
	"blog/domains/posts/service"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PostsController struct {
	service service.PostsService
}

func CreatePostsController(ps service.PostsService) *PostsController {
	return &PostsController{service: ps}
}

func (pc PostsController) CreatePost(ctx *gin.Context) {
	var request models.Post

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body: " + err.Error(),
		})
		return
	}

	post, err := pc.service.CreatePost(&request)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"post": post,
	})
}

func (pc PostsController) UpdatePost(ctx *gin.Context) {
	var request models.Post

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body: " + err.Error(),
		})
		return
	}

	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid UUID: " + err.Error(),
		})
		return
	}

	post, err := pc.service.UpdatePost(id, &request)

	if err != nil {
		if errors.Is(err, postErrors.ErrorPostDoesntExist{ID: id}) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func (pc PostsController) GetPosts(ctx *gin.Context) {
	posts, err := pc.service.GetPosts()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}

func (pc PostsController) GetPost(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid UUID: " + err.Error(),
		})
		return
	}

	post, err := pc.service.GetPost(id)

	if err != nil {
		if errors.Is(err, postErrors.ErrorPostDoesntExist{ID: id}) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func (pc PostsController) DeletePost(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid UUID: " + err.Error(),
		})
		return
	}

	err = pc.service.DeletePost(id)

	if err != nil {
		if errors.Is(err, postErrors.ErrorPostDoesntExist{ID: id}) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Deleted post with id %s", id.String()),
	})
}
