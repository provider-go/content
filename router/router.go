package router

import (
	"github.com/gin-gonic/gin"
	"github.com/provider-go/content/api"
)

type Group struct {
	Router
}

var GroupApp = new(Group)

type Router struct{}

func (s *Router) InitRouter(Router *gin.RouterGroup) {
	{
		// content_channels 表操作
		Router.POST("channels/add", api.CreateChannel)
		Router.POST("channels/delete", api.DeleteChannel)
		Router.POST("channels/list", api.ListChannel)
		// content_articles 表操作
		Router.POST("articles/add", api.CreateArticle)
		Router.POST("articles/delete", api.DeleteArticle)
		Router.POST("articles/list", api.ListArticle)
		Router.POST("articles/view", api.ViewArticle)
		// content_singles 表操作
		Router.POST("singles/add", api.CreateSingle)
		Router.POST("singles/delete", api.DeleteSingle)
		Router.POST("singles/list", api.ListSingle)
		Router.POST("singles/view", api.ViewSingle)
	}
}
