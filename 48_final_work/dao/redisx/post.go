package redisx

import (
	"go_web_demo/48_final_work/models"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

func CreatePost(postID, commID int64) (err error) {
	pipeline := rdb.TxPipeline()
	//两个要同时进行
	// 帖子时间
	pipeline.ZAdd(getKey(KeyPostTime), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})
	// 帖子分数
	pipeline.ZAdd(getKey(KeyPostTime), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})
	key := getKey(KeyCommunity) + strconv.Itoa(int(commID))
	pipeline.SAdd(key, postID)
	_, err = pipeline.Exec()
	return
}

func GetCommunityPostID(page, site int64, key string) ([]string, error) {
	start := (page - 1) * site
	end := start + site - 1
	return rdb.ZRevRange(key, start, end).Result()
}

func GetPostByOrder(p *models.GetPostOder) ([]string, error) {
	key := getKey(KeyPostScore)
	if p.Order == models.OrderTime {
		key = getKey(KeyPostTime)
	}
	return GetCommunityPostID(p.Page, p.Site, key)
}

func GetPostVoteData(ids []string) (voteData []int, err error) {
	voteData = make([]int, 0, len(ids))
	pipeline := rdb.Pipeline()
	for _, id := range ids {
		key := getKey(KeyUserVote + id)
		pipeline.ZCount(key, "1", "1")
	}
	cmders, err := pipeline.Exec()
	if err != nil {
		return nil, err
	}
	for _, cmder := range cmders {
		v := cmder.(*redis.IntCmd).Val()
		voteData = append(voteData, int(v))
	}
	return
}

func GetCommunityPostByOrder2(p *models.CommunityPostOrder) ([]string, error) {
	commkey := getKey(KeyCommunity + strconv.Itoa(int(p.Comm)))
	//bluebell:community:comm_id
	orderKey := getKey(KeyPostScore)
	if p.Order == models.OrderTime {
		orderKey = getKey(KeyPostTime)
	}
	// bluebell:post:score/time
	finkey := orderKey + ":" + strconv.Itoa(int(p.Comm))
	//finkey2 := orderKey + strconv.Itoa(int(p.Comm))
	//fmt.Println("finKey", finkey)
	//fmt.Println("finkey2", finkey2)
	// bluebell:post:score/timeid
	if rdb.Exists(finkey).Val() < 1 {
		rdb.ZInterStore(finkey, redis.ZStore{
			Aggregate: "MAX",
		}, commkey, orderKey)
	}
	return GetCommunityPostID(p.Page, p.Site, finkey)
}
