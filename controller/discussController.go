package controller

import (
	"Project/model"
	"Project/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DiscussListResponse struct {
	model.BaseResponse
	Discuss []model.Discuss `json:"data"`
}

func DiscussActionController(ctx *gin.Context) {
	studentId := ctx.PostForm("student_id")
	content := ctx.PostForm("content")
	if studentId == "" || content == "" {
		ctx.JSON(http.StatusBadRequest, model.BaseResponse{
			StatusCode: -1,
			StatusMsg:  "请求参数为空",
		})
		return
	}
	err := service.DiscussActionService(studentId, content)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.BaseResponse{
			StatusCode: -1,
			StatusMsg:  err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, model.BaseResponse{
		StatusCode: 0,
		StatusMsg:  "成功",
	})
	return
}
func DiscussListController(ctx *gin.Context) {
	discusses, err := service.DiscussListService()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.BaseResponse{
			StatusCode: -1,
			StatusMsg:  err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, DiscussListResponse{
		BaseResponse: model.BaseResponse{
			StatusCode: 0,
			StatusMsg:  "成功",
		},
		Discuss: discusses,
	})
	return
}
