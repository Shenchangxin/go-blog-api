package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/gin-gonic/gin"
	redis2 "github.com/go-redis/redis/v8"
	"go-blog-api/user-web/forms"
	"go-blog-api/user-web/global"
	"go.uber.org/zap"
	"math/rand"
	"net/http"
	"time"
)

func SendSMS(ctx *gin.Context) {
	sendSmsForm := forms.SendSmsForm{}
	if err := ctx.ShouldBind(&sendSmsForm); err != nil {
		HandleValidatorErr(ctx, err)
		return
	}
	// 生成6位随机Code
	code := GenerateRandomCode(6)
	// 通过accessKey Id和Secret连接服务
	client, err := dysmsapi.NewClientWithAccessKey("cn-hangzhou", global.ServerConfig.AliSmsConfig.ApiKey, global.ServerConfig.AliSmsConfig.ApiSecret)
	if err != nil {
		return
	}
	request := dysmsapi.CreateSendSmsRequest() //创建请求
	request.Scheme = "http"                    //请求协议，可选：https，但会慢一点
	request.PhoneNumbers = sendSmsForm.Phone   //接收短信的手机号码
	request.SignName = "your_signature"        //短信签名名称
	request.TemplateCode = "your_template_id"  //短信模板ID
	Param, err := json.Marshal(map[string]interface{}{
		"code": code, // 验证码参数
	})
	if err != nil {
		return
	}
	request.TemplateParam = string(Param) //将短信模板参数传入短信模板
	_, err = client.SendSms(request)      //调用阿里云API发送信息
	if err != nil {
		zap.S().Errorw("发送验证码失败", err.Error())
		return
	}
	redis := redis2.NewClient(&redis2.Options{
		Addr: fmt.Sprintf("%s:%d", global.ServerConfig.RedisConfig.Host, global.ServerConfig.RedisConfig.Port),
	})
	redis.Set(context.Background(), sendSmsForm.Phone, code, 60*time.Second)
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "验证码发送成功",
	})
	return
}

func GenerateRandomCode(length int) string {
	rand.Seed(time.Now().UnixNano())
	var letters = []rune("0123456789")
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
