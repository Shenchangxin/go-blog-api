package forms

type PasswordLoginForm struct {
	UserName string `form:"userName" json:"userName" binding:"required"`
	Password string `form:"password" json:"password" binding:"required,min=3,max=20"`
}
