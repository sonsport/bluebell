package mysqlx

import (
	sql2 "database/sql"
	"fmt"
	"go_web_demo/48_final_work/models"
	"strings"

	"github.com/jmoiron/sqlx"
)

func CreatPost(post *models.Post) (err error) {
	sql := `insert into post(
    post_id,title,content,author_id,community_id)
    values (?,?,?,?,?)
    `
	_, err = db.Exec(sql, post.ID, post.Title, post.Content, post.AuthorID, post.CommunityID)
	return
}

func GetDetailPost(pid int64) (post *models.Post, err error) {
	getPost := new(models.Post)
	sql := `select post_id,title,content,author_id,community_id,status, create_time from post where post_id = ?`
	err = db.Get(getPost, sql, pid)
	return getPost, err
}

func GetUserById(uid int64) (data *models.UserStruct, err error) {
	getData := new(models.UserStruct)
	sql := `select user_id,username from user where user_id = ?`
	err = db.Get(getData, sql, uid)
	if err != nil {
		if err == sql2.ErrNoRows {
			err = nil
		}
	}
	return getData, err
}

func GetPost(site, page int64) (data []*models.Post, err error) {
	getPost := make([]*models.Post, 0, 2)
	sql := `select
    post_id,title,content,author_id,community_id,status, create_time
from post ORDER BY create_time DESC limit ?,?`
	err = db.Select(&getPost, sql, (page-1)*site, site)
	return getPost, err
}

func GetPostByOrder(p []string) (postList []*models.Post, err error) {
	sqlStr := `select post_id,title,content,author_id,community_id,create_time from post where post_id in (?) 
order by FIND_IN_SET(post_id,?)`
	query, args, err := sqlx.In(sqlStr, p, strings.Join(p, ","))
	if err != nil {
		return nil, err
	}
	query = db.Rebind(query)
	err = db.Select(&postList, query, args...)
	return
}

func GetPostByComm(p []string, commID int64) (postList []string, err error) {
	sqlStr := `select post_id from post where post_id in (?) AND community_id = ?
order by FIND_IN_SET(post_id,?)`
	query, args, err := sqlx.In(sqlStr, p, commID, strings.Join(p, ","))
	if err != nil {
		return
	}
	query = db.Rebind(query)
	err = db.Select(&postList, query, args...)
	return
}

func GetTableDetail(username string) (tableList []*models.ApiTableDetail, err error) {
	sqlStr := `select id,create_time,username from table2 where username = ?`
	fmt.Println(username)
	err = db.Select(&tableList, sqlStr, username)
	return
}
