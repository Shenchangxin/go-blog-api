package forms

type PasswordLoginForm struct {
	UserName  string `form:"userName" json:"userName" binding:"required"`
	Password  string `form:"password" json:"password" binding:"required,min=3,max=20"`
	Captcha   string `form:"captcha" json:"captcha" binding:"required,min=5,max=5"`
	CaptchaId string `form:"captchaId" json:"captchaId" binding:"required"`
}

type RegisterUserForm struct {
	UserName string `form:"userName" json:"userName" binding:"required"`
	Password string `form:"password" json:"password" binding:"required,min=3,max=20"`
	Phone    string `form:"phone" json:"phone" binding:"required,phone"`
	Sex      string `form:"sex" json:"sex" binding:"required"`
	Code     string `form:"code" json:"code" binding:"required"`
	NickName string `form:"nickName" json:"nickName" binding:"required"`
}
