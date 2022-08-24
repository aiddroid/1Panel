package router

import (
	v1 "github.com/1Panel-dev/1Panel/app/api/v1"
	"github.com/1Panel-dev/1Panel/middleware"

	"github.com/gin-gonic/gin"
)

type GroupRouter struct{}

func (s *GroupRouter) InitGroupRouter(Router *gin.RouterGroup) {
	groupRouter := Router.Group("groups").Use(middleware.JwtAuth()).Use(middleware.SessionAuth())
	withRecordRouter := Router.Group("groups").Use(middleware.JwtAuth()).Use(middleware.SessionAuth()).Use(middleware.OperationRecord())
	baseApi := v1.ApiGroupApp.BaseApi
	{
		withRecordRouter.POST("", baseApi.CreateGroup)
		withRecordRouter.DELETE(":id", baseApi.DeleteGroup)
		withRecordRouter.PUT(":id", baseApi.UpdateGroup)
		groupRouter.POST("/search", baseApi.ListGroup)
		groupRouter.GET(":id", baseApi.GetGroupInfo)
	}
}