package response

import "go-blog-api/user-web/proto"

type UserResponse struct {
	Id       int32             `json:"id"`
	UserName string            `json:"userName"`
	NickName string            `json:"nickName"`
	Sex      string            `json:"sex"`
	Phone    string            `json:"phone"`
	Role     []*proto.RoleItem `json:"role"`
}
