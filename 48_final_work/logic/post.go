package logic

import (
	"fmt"
	"go_web_demo/48_final_work/dao/mysqlx"
	"go_web_demo/48_final_work/dao/redisx"
	"go_web_demo/48_final_work/models"
	"go_web_demo/48_final_work/pkg/snowflake"

	"go.uber.org/zap"
)

func CreatePost(post *models.Post) (err error) {
	post.ID = snowflake.GetGenID()
	err = mysqlx.CreatPost(post)
	if err != nil {
		return
	}
	err = redisx.CreatePost(post.ID, post.CommunityID)
	return
}

func GetDetailPost(id int64) (data *models.ApiPostDetail, err error) {
	post, err := mysqlx.GetDetailPost(id)
	if err != nil {
		zap.L().Error("mysqlx get detail post failed", zap.Error(err))
		return
	}
	// 根据作者id查询作者信息
	user, err := mysqlx.GetUserById(post.AuthorID)
	if err != nil {
		zap.L().Error("mysqlx get user failed", zap.Error(err))
		return
	}
	// 法2 直接根据request拿到

	// 根据社区id查询到社区详细
	community, err := mysqlx.GetCommunityDetail(post.CommunityID)
	if err != nil {
		zap.L().Error("mysqlx get community detail failed", zap.Error(err))
		return
	}
	data = &models.ApiPostDetail{
		Author:          user.Username,
		Post:            post,
		CommunityDetail: community,
	}
	return
}

func GetPost(site, page int64) (data []*models.ApiPostDetail, err error) {
	posts, err := mysqlx.GetPost(site, page)
	if err != nil {
		return
	}
	for _, post := range posts {
		// 根据作者id查询作者信息
		user, err := mysqlx.GetUserById(post.AuthorID)
		if err != nil {
			zap.L().Error("mysqlx get user failed", zap.Error(err))
			continue
		}
		// 法2 直接根据request拿到

		// 根据社区id查询到社区详细
		community, err := mysqlx.GetCommunityDetail(post.CommunityID)
		if err != nil {
			zap.L().Error("mysqlx get community detail failed", zap.Error(err))
			continue
		}
		tmpData := &models.ApiPostDetail{
			Author:          user.Username,
			CommunityDetail: community,
			Post:            post,
		}
		data = append(data, tmpData)
	}
	return
}

func GetPostByOrder(p *models.GetPostOder) (data []*models.ApiPostDetail, err error) {
	ids, err := redisx.GetPostByOrder(p)
	if err != nil {
		return nil, err
	}
	if len(ids) == 0 {
		return
	}
	// 根据id去数据库查询帖子详细信息
	posts, err := mysqlx.GetPostByOrder(ids)
	fmt.Println(posts)
	if err != nil {
		return
	}
	for _, post := range posts {
		// 根据作者id查询作者信息
		user, err := mysqlx.GetUserById(post.AuthorID)
		if err != nil {
			zap.L().Error("mysqlx get user failed", zap.Error(err))
			continue
		}
		// 法2 直接根据request拿到

		// 根据社区id查询到社区详细
		community, err := mysqlx.GetCommunityDetail(post.CommunityID)
		if err != nil {
			zap.L().Error("mysqlx get community detail failed", zap.Error(err))
			continue
		}
		tmpData := &models.ApiPostDetail{
			Author:          user.Username,
			CommunityDetail: community,
			Post:            post,
		}
		data = append(data, tmpData)
	}
	return
}

func GetPostByOrder2(p *models.GetPostOder) (data []*models.ApiPostDetail, err error) {
	ids, err := redisx.GetPostByOrder(p)
	if err != nil {
		return
	}
	if len(ids) == 0 {
		return
	}
	voteData, err := redisx.GetPostVoteData(ids)
	if err != nil {
		return
	}
	// 根据id去数据库查询帖子详细信息
	posts, err := mysqlx.GetPostByOrder(ids)
	if err != nil {
		return
	}
	for index, post := range posts {
		// 根据作者id查询作者信息
		user, err := mysqlx.GetUserById(post.AuthorID)
		if err != nil {
			zap.L().Error("mysqlx get user failed", zap.Error(err))
			continue
		}
		// 法2 直接根据request拿到

		// 根据社区id查询到社区详细
		community, err := mysqlx.GetCommunityDetail(post.CommunityID)
		if err != nil {
			zap.L().Error("mysqlx get community detail failed", zap.Error(err))
			continue
		}
		tmpData := &models.ApiPostDetail{
			Author:          user.Username,
			VoteData:        voteData[index],
			CommunityDetail: community,
			Post:            post,
		}
		data = append(data, tmpData)
	}
	return
}

func GetCommunityPostOrder(p *models.CommunityPostOrder) (data []*models.ApiPostDetail, err error) {
	ids, err := redisx.GetCommunityPostByOrder2(p)
	if err != nil {
		return
	}
	if len(ids) == 0 {
		return
	}
	//commIDS, err := mysqlx.GetPostByComm(ids, p.Comm)
	//if err != nil {
	//	return
	//}
	// 法2
	voteData, err := redisx.GetPostVoteData(ids)
	if err != nil {
		return
	}
	//根据id去数据库查询帖子详细信息
	posts, err := mysqlx.GetPostByOrder(ids)
	if err != nil {
		return
	}
	for index, post := range posts {
		// 根据作者id查询作者信息
		user, err := mysqlx.GetUserById(post.AuthorID)
		if err != nil {
			zap.L().Error("mysqlx get user failed", zap.Error(err))
			continue
		}
		// 法2 直接根据request拿到

		// 根据社区id查询到社区详细
		community, err := mysqlx.GetCommunityDetail(post.CommunityID)
		if err != nil {
			zap.L().Error("mysqlx get community detail failed", zap.Error(err))
			continue
		}
		tmpData := &models.ApiPostDetail{
			Author:          user.Username,
			VoteData:        voteData[index],
			CommunityDetail: community,
			Post:            post,
		}
		data = append(data, tmpData)
	}
	return
}

func GetTableDetail(username string) (tables []*models.ApiTableDetail, err error) {
	tables, err = mysqlx.GetTableDetail(username)
	if err != nil {
		return
	}
	return
}
