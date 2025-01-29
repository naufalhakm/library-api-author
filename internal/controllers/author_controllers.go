package controllers

import (
	"library-api-author/internal/commons/response"
	"library-api-author/internal/models"
	"library-api-author/internal/params"
	"library-api-author/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AuthorController interface {
	CreateAuthor(ctx *gin.Context)
	GetDetailAuthor(ctx *gin.Context)
	UpdateAuthor(ctx *gin.Context)
	DeleteAuthor(ctx *gin.Context)
	GetAllAuthors(ctx *gin.Context)
}

type AuthorControllerImpl struct {
	AuthorService services.AuthorService
}

func NewAuthorController(authorService services.AuthorService) AuthorController {
	return &AuthorControllerImpl{
		AuthorService: authorService,
	}
}

func (controller *AuthorControllerImpl) CreateAuthor(ctx *gin.Context) {
	var req = new(params.AuthorRequest)

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err,
		})
		return
	}

	custErr := controller.AuthorService.CreateAuthor(ctx, req)
	if custErr != nil {
		ctx.AbortWithStatusJSON(custErr.StatusCode, custErr)
		return
	}

	resp := response.CreatedSuccessWithPayload("Success create data author")
	ctx.JSON(resp.StatusCode, resp)
}

func (controller *AuthorControllerImpl) GetDetailAuthor(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err,
		})
		return
	}

	result, custErr := controller.AuthorService.GetDetailAuthor(ctx, uint64(id))

	if custErr != nil {
		ctx.AbortWithStatusJSON(custErr.StatusCode, custErr)
		return
	}

	resp := response.GeneralSuccessCustomMessageAndPayload("Success get detail author", result)
	ctx.JSON(resp.StatusCode, resp)
}

func (controller *AuthorControllerImpl) UpdateAuthor(ctx *gin.Context) {
	var req = new(params.AuthorRequest)

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err,
		})
		return
	}

	id, _ := strconv.Atoi(ctx.Param("id"))

	custErr := controller.AuthorService.UpdateAuthor(ctx, uint64(id), req)
	if custErr != nil {
		ctx.AbortWithStatusJSON(custErr.StatusCode, custErr)
		return
	}

	resp := response.GeneralSuccessCustomMessageAndPayload("Success update data author", nil)
	ctx.JSON(resp.StatusCode, resp)
}

func (controller *AuthorControllerImpl) DeleteAuthor(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err,
		})
		return
	}

	custErr := controller.AuthorService.DeleteAuthor(ctx, uint64(id))
	if custErr != nil {
		ctx.AbortWithStatusJSON(custErr.StatusCode, custErr)
		return
	}

	resp := response.GeneralSuccessCustomMessageAndPayload("Success delete data author", nil)
	ctx.JSON(resp.StatusCode, resp)
}

func (controller *AuthorControllerImpl) GetAllAuthors(ctx *gin.Context) {
	page := ctx.Query("page")
	limit := ctx.Query("limit")

	pageNum := 1
	limitSize := 5

	if page != "" {
		parsedPage, err := strconv.Atoi(page)
		if err == nil && parsedPage > 0 {
			pageNum = parsedPage
		}
	}

	if limit != "" {
		parsedLimit, err := strconv.Atoi(limit)
		if err == nil && parsedLimit > 0 {
			limitSize = parsedLimit
		}
	}

	pagination := models.Pagination{
		Page:     pageNum,
		Offset:   (pageNum - 1) * limitSize,
		PageSize: limitSize,
	}

	result, custErr := controller.AuthorService.GetAllAuthors(ctx, &pagination)

	if custErr != nil {
		ctx.AbortWithStatusJSON(custErr.StatusCode, custErr)
		return
	}

	type Response struct {
		Authors    interface{} `json:"authors"`
		Pagination interface{} `json:"pagination"`
	}

	var responses Response
	responses.Authors = result
	responses.Pagination = pagination

	resp := response.GeneralSuccessCustomMessageAndPayload("Success get data authors", responses)
	ctx.JSON(resp.StatusCode, resp)

}
