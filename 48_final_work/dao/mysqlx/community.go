package mysqlx

import (
	sql2 "database/sql"
	"go_web_demo/48_final_work/models"
)

func GetCommunityList() (community []*models.CommunityList, err error) {
	sql := `select community_id,community_name from community`
	if err = db.Select(&community, sql); err != nil {
		if err == sql2.ErrNoRows {
			err = nil
		}
	}
	return
}

func GetCommunityDetail(id int64) (*models.CommunityDetail, error) {
	community := new(models.CommunityDetail)
	sql := `select community_id,community_name,introduction,create_time,update_time from community where id = ?`
	err := db.Get(community, sql, id)
	return community, err
}
