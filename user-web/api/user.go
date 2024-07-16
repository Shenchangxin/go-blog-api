package api

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	redis2 "github.com/go-redis/redis/v8"
	"go-blog-api/user-web/forms"
	"go-blog-api/user-web/global"
	"go-blog-api/user-web/global/response"
	"go-blog-api/user-web/middlewares"
	"go-blog-api/user-web/models"
	"go-blog-api/user-web/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func removeTopStruct(fields map[string]string) map[string]string {
	rsp := map[string]string{}

	for field, err := range fields {
		rsp[field[strings.Index(field, ".")+1:]] = err
	}
	return rsp

}

func HandleValidatorErr(ctx *gin.Context, err error) {
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
	}
	ctx.JSON(http.StatusBadRequest, gin.H{
		"error": removeTopStruct(errs.Translate(global.Trans)),
	})
	return
}

func HandleGrpcErrorToHttp(err error, ctx *gin.Context) {
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
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
			case codes.Unavailable:
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"msg": "用户服务不可用",
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

	userCoon, err := grpc.Dial(fmt.Sprintf("%s:%d", global.ServerConfig.UserSrvInfo.Host, global.ServerConfig.UserSrvInfo.Port), grpc.WithInsecure())
	if err != nil {
		zap.S().Errorf("用户服务连接失败：%s", err.Error())
	}
	userClient := proto.NewUserClient(userCoon)

	pageNo := ctx.DefaultQuery("pageNo", "0")
	pageSize := ctx.DefaultQuery("pageSize", "10")
	pageNoInt, _ := strconv.Atoi(pageNo)
	pageSizeInt, _ := strconv.Atoi(pageSize)

	rsp, err := userClient.GetUserList(context.Background(), &proto.PageInfo{
		PageNo:   uint32(pageNoInt),
		PageSize: uint32(pageSizeInt),
	})
	if err != nil {
		zap.S().Errorf("查询用户列表失败：%s", err.Error())
		HandleGrpcErrorToHttp(err, ctx)
		return
	}

	result := make([]interface{}, 0)

	for _, value := range rsp.Data {
		//data := make(map[string]interface{})

		user := response.UserResponse{
			Id:       value.Id,
			UserName: value.UserName,
			NickName: value.NickName,
			Sex:      value.Sex,
			Phone:    value.Phone,
			Role:     value.Role,
		}

		//data["id"] = value.Id
		//data["username"] = value.UserName
		//data["nickname"] = value.NickName
		//data["sex"] = value.Sex
		//data["phone"] = value.Phone
		//data["role"] = value.Role
		result = append(result, user)
	}
	ctx.JSON(http.StatusOK, result)

}

func Login(ctx *gin.Context) {
	loginForm := forms.PasswordLoginForm{}
	if err := ctx.ShouldBind(&loginForm); err != nil {
		HandleValidatorErr(ctx, err)
		return
	}
	if !store.Verify(loginForm.CaptchaId, loginForm.Captcha, true) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"captcha": "验证码错误",
		})
		return
	}
	//测试github提交记录
	userCoon, err := grpc.Dial(fmt.Sprintf("%s:%d", global.ServerConfig.UserSrvInfo.Host, global.ServerConfig.UserSrvInfo.Port), grpc.WithInsecure())
	if err != nil {
		zap.S().Errorf("用户服务连接失败：%s", err.Error())
	}
	userClient := proto.NewUserClient(userCoon)

	if rsp, err := userClient.GetUserByUserName(context.Background(), &proto.UserNameRequest{
		Username: loginForm.UserName,
	}); err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				ctx.JSON(http.StatusBadRequest, gin.H{
					"username": "用户不存在",
				})
			default:
				ctx.JSON(http.StatusInternalServerError, map[string]string{
					"username": "登录失败",
				})
			}

		}
	} else {
		//校验密码是否正确
		passRsp, passErr := userClient.CheckPassword(context.Background(), &proto.PasswordCheckInfo{
			Password:          loginForm.Password,
			EncryptedPassword: rsp.Sex, //TODO 将用户密码也查出来放到rsp中
		})
		if passErr != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"msg": "登陆失败",
			})
		} else {
			if passRsp.Success {

				j := middlewares.NewJWT()
				claims := models.CustomClaims{
					ID:          uint(rsp.Id),
					NickName:    rsp.NickName,
					AuthorityId: 1,
					StandardClaims: jwt.StandardClaims{
						NotBefore: time.Now().Unix(),
						ExpiresAt: time.Now().Unix() + 60*60*24*30,
						Issuer:    "imooc",
					},
				}
				token, err := j.CreateToken(claims)
				if err != nil {
					ctx.JSON(http.StatusInternalServerError, gin.H{
						"msg": "生成token失败",
					})
					return
				}

				ctx.JSON(http.StatusOK, gin.H{
					"id":         rsp.Id,
					"nick_name":  rsp.NickName,
					"token":      token,
					"expired_at": (time.Now().Unix() + 60*60*24*30) * 1000,
				})
			} else {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"msg": "登录失败",
				})
			}
		}
	}
}

func RegisterUser(ctx *gin.Context) {

	registerUserForm := forms.RegisterUserForm{}
	if err := ctx.ShouldBind(&registerUserForm); err != nil {
		HandleValidatorErr(ctx, err)
		return
	}
	redis := redis2.NewClient(&redis2.Options{
		Addr: fmt.Sprintf("%s:%d", global.ServerConfig.RedisConfig.Host, global.ServerConfig.RedisConfig.Port),
	})
	result, err := redis.Get(context.Background(), registerUserForm.Phone).Result()
	if err == redis2.Nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": "验证码失效",
		})
		return
	}
	if result != registerUserForm.Code {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": "验证码错误",
		})
		return
	}

	userCoon, err := grpc.Dial(fmt.Sprintf("%s:%d", global.ServerConfig.UserSrvInfo.Host, global.ServerConfig.UserSrvInfo.Port), grpc.WithInsecure())
	if err != nil {
		zap.S().Errorf("用户服务连接失败：%s", err.Error())
	}
	userClient := proto.NewUserClient(userCoon)
	userInfoResponse, err := userClient.CreateUser(context.Background(), &proto.CreateUserInfo{
		UserName: registerUserForm.UserName,
		NickName: registerUserForm.NickName,
		Phone:    registerUserForm.Phone,
		Sex:      registerUserForm.Sex,
		Password: registerUserForm.Password,
	})
	if err != nil {
		zap.S().Errorw("创建用户失败", err.Error())
		HandleGrpcErrorToHttp(err, ctx)
		return
	}
	j := middlewares.NewJWT()
	claims := models.CustomClaims{
		ID:          uint(userInfoResponse.Id),
		NickName:    userInfoResponse.NickName,
		AuthorityId: 1,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			ExpiresAt: time.Now().Unix() + 60*60*24*30,
			Issuer:    "imooc",
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "生成token失败",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"id":         userInfoResponse.Id,
		"nick_name":  userInfoResponse.NickName,
		"token":      token,
		"expired_at": (time.Now().Unix() + 60*60*24*30) * 1000,
	})
}
