package logic

import (
	"go_web_demo/48_final_work/dao/redisx"
	"go_web_demo/48_final_work/models"
	"strconv"
)

func PostVote(userID int64, p *models.VoteData) error {
	return redisx.PostVote(strconv.Itoa(int(userID)), p.PostID, float64(p.Vote))
}
