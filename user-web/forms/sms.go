package forms

type SendSmsForm struct {
	Phone string `form:"phone" json:"phone" binding:"required,phone"`
	Type  string `form:"Type" json:"Type" binding:"required"`
}
