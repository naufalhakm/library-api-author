package routes

import (
	"fmt"
	"library-api-author/internal/factory"
	"library-api-author/internal/grpc/client"
	"library-api-author/internal/middleware"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(provider *factory.Provider, authClient *client.AuthClient) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger(), CORS())

	router.GET("/", func(ctx *gin.Context) {
		currentYear := time.Now().Year()
		message := fmt.Sprintf("Library API Author %d", currentYear)

		ctx.JSON(http.StatusOK, message)
	})

	api := router.Group("/api")
	{
		v1 := api.Group("v1")
		{
			auth := v1.Use(middleware.CheckAuthIsAdminOrAuthor(authClient))
			auth.GET("/authors", provider.AuthorProvider.GetAllAuthors)
			auth.GET("/authors/:id", provider.AuthorProvider.GetDetailAuthor)

			admin := v1.Use(middleware.CheckAuthIsAdmin(authClient))
			admin.POST("/authors", provider.AuthorProvider.CreateAuthor)
			admin.PUT("/authors/:id", provider.AuthorProvider.UpdateAuthor)
			admin.DELETE("/authors/:id", provider.AuthorProvider.DeleteAuthor)
		}
	}

	return router
}

func CORS() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST, PUT, DELETE")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type, accept, access-control-allow-origin, access-control-allow-headers")
		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(http.StatusNoContent)
		}
		ctx.Next()
	}
}
