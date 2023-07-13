package models

const (
	OrderScore = "score"
	OrderTime  = "time"
)

type SignUpStruct struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	Repassword string `json:"re_password" binding:"required,eqfield=Password"`
}

type LoginStruct struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type GetPostOder struct {
	Order string `json:"order" form:"order"`
	Site  int64  `json:"site" form:"site"`
	Page  int64  `json:"page" form:"page"`
}

type CommunityPostOrder struct {
	Order string `json:"order" form:"order"`
	Comm  int64  `json:"community_id" form:"community_id" binding:"required"`
	Site  int64  `json:"site" form:"site"`
	Page  int64  `json:"page" form:"page"`
}
