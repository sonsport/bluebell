package logic

import (
	"go_web_demo/48_final_work/dao/mysqlx"
	"go_web_demo/48_final_work/models"
)

func GetCommunityList() (communityList []*models.CommunityList, err error) {
	return mysqlx.GetCommunityList()
	// 错误在controller 处理
	// 数据 在 mysql处理
	// logic 负责处理逻辑
}

func GetCommunityDetail(id int64) (communityDetail *models.CommunityDetail, err error) {
	return mysqlx.GetCommunityDetail(id)
}
