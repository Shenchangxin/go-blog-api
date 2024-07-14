package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-blog-api/user-web/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

func HandleGrpcErrorToHttp(err error, ctx *gin.Context) {
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code {
			case codes.NotFound:
				ctx.JSON(http.StatusNotFound, gin.H{
					"msg": e.Message(),
				})
			case codes.Internal:
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"msg": "内部错误",
				})
			case codes.InvalidArgument:
				ctx.JSON(http.StatusBadRequest, gin.H{
					"msg": "参数错误",
				})
			default:
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"msg": "其他错误",
				})

			}
			return
		}
	}
}

func GetUserList(ctx *gin.Context) {

	userCoon, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		zap.S().Errorf("用户服务连接失败：%s", err.Error())
	}
	userClient := proto.NewUserClient(userCoon)
	rsp, err := userClient.GetUserList(context.Background(), &proto.PageInfo{
		PageNo:   1,
		PageSize: 10,
	})
	if err != nil {
		zap.S().Errorf("查询用户列表失败：%s", err.Error())
		HandleGrpcErrorToHttp(err, ctx)
		return
	}

}
