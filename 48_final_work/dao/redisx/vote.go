package redisx

import (
	"errors"
	"math"
	"time"

	"github.com/go-redis/redis"
)

const (
	KeyPrefix    = "bluebell:"
	KeyPostTime  = "post:time"
	KeyPostScore = "post:score"
	KeyUserVote  = "user:"
	KeyCommunity = "community:"
	TimeOnWeek   = 60 * 60 * 24 * 7
	ScorePerVote = 432
)

var (
	ErrorTimeLate = errors.New("超时")
	ErrorParaSame = errors.New("你已经通过相同的票了")
)

func getKey(key string) string {
	return KeyPrefix + key
}

func PostVote(userID, postID string, value float64) (err error) {
	/*
				1. 只能一个星期内投票
				2. 到期之后删除

			功能:
		vote = 1 时
		之前没投票，现在赞成	+432
		之前投反对，现在赞成	+432 * 2

		vote = 0 时
		之前投赞成，现在取消	-432
		之前投反对，现在取消	+432

		vote = -1 时
		之前投赞成, 现在反对	-432*2
		之前没投票，现在反对	-432
	*/
	//postIDStr := strconv.Itoa(int(PostID))
	//userIDStr := strconv.Itoa(int(userID))
	// redis读取到帖子发布时间
	postTime := rdb.ZScore(getKey(KeyPostTime), postID).Val()
	if float64(time.Now().Unix())-postTime > TimeOnWeek {
		return ErrorTimeLate
	}

	//查看当前帖子的投票记录
	ov := rdb.ZScore(getKey(KeyUserVote+postID), userID).Val()
	var dir float64
	if value == ov {
		return ErrorParaSame
	}
	if value > ov {
		dir = 1
	} else {
		dir = -1
	}

	pipeline := rdb.TxPipeline()
	diff := math.Abs(ov - value) // 差值
	_, err = rdb.ZIncrBy(getKey(KeyPostScore), dir*diff*ScorePerVote, postID).Result()

	// 记录该用户为帖子投票的数据
	if value == 0 {
		rdb.ZRem(getKey(KeyUserVote+postID), userID)
	} else {
		rdb.ZAdd(getKey(KeyUserVote+postID), redis.Z{
			Score:  value,
			Member: userID,
		})
	}
	_, err = pipeline.Exec()
	return
}
